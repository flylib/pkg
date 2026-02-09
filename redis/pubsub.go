package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type ClientType int8

const (
	Publisher ClientType = iota
	Subscriber
)

type PubSubQueue struct {
	client     *redis.Client
	channel    string
	clientType ClientType

	// retry options
	retryMaxTimes   int
	retryIntervalMs int
}

func (c *Cli) NewPublisher(channel string) *PubSubQueue {
	return &PubSubQueue{
		client:     c.Client,
		channel:    channel,
		clientType: Publisher,
	}
}

func (c *Cli) NewSubscriber(channel string, retryMaxTimes int, retryIntervalMs int) *PubSubQueue {
	return &PubSubQueue{
		client:          redis.NewClient(c.opt),
		channel:         channel,
		clientType:      Subscriber,
		retryMaxTimes:   retryMaxTimes,
		retryIntervalMs: retryIntervalMs,
	}
}

func (psq *PubSubQueue) Publish(ctx context.Context, message string) error {
	return psq.client.Publish(ctx, psq.channel, message).Err()
}

func (psq *PubSubQueue) Subscribe(ctx context.Context, handler func(string) error) error {
	if psq.clientType != Subscriber {
		return errors.New("cannot subscribe using publisher client")
	}

	retryCount := 0

	for {
		// 检查是否已取消
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		pubsub := psq.client.Subscribe(ctx, psq.channel)
		ch := pubsub.Channel()

		log.Printf("✅ subscribed to channel: %s", psq.channel)

		// 重置重试计数（成功订阅）
		retryCount = 0

		err := psq.consumeMessages(ctx, pubsub, ch, handler)

		// 确保关闭订阅
		if closeErr := pubsub.Close(); closeErr != nil {
			log.Printf("⚠️ error closing pubsub: %v", closeErr)
		}

		// 如果是正常取消，直接返回
		if err == context.Canceled || err == context.DeadlineExceeded {
			return err
		}

		// 处理重连逻辑
		if !psq.shouldRetry(retryCount) {
			return fmt.Errorf("max retry attempts (%d) exceeded", psq.retryMaxTimes)
		}

		retryCount++
		retryInterval := time.Duration(psq.retryIntervalMs) * time.Millisecond
		log.Printf("⚠️ pubsub disconnected, retrying in %v (attempt %d/%d)...",
			retryInterval, retryCount, psq.retryMaxTimes)

		time.Sleep(retryInterval)
	}
}

func (psq *PubSubQueue) consumeMessages(
	ctx context.Context,
	pubsub *redis.PubSub,
	ch <-chan *redis.Message,
	handler func(string) error,
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-ch:
			if !ok {
				// 通道关闭，连接断开
				return errors.New("pubsub channel closed")
			}

			if err := handler(msg.Payload); err != nil {
				log.Printf("❌ handler error: %v", err)
				// 可以根据需求决定是否继续处理后续消息
			}
		}
	}
}

func (psq *PubSubQueue) shouldRetry(currentRetry int) bool {
	// 如果未设置重试配置，不重试
	if psq.retryMaxTimes <= 0 {
		return false
	}
	return currentRetry < psq.retryMaxTimes
}

// Close 关闭订阅者的客户端连接
func (psq *PubSubQueue) Close() error {
	if psq.clientType == Subscriber && psq.client != nil {
		return psq.client.Close()
	}
	return nil
}
