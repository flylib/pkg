package minio

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"os"
)

const (
	DefaultContentType = "multipart/form-data"
)

var (
	ErrorBucketExist = errors.New("bucket already exist")
)

// {"url":"http://127.0.0.1:9000","accessKey":"MJKJ20AedWTKEfDVXJGh","secretKey":"eVJLUEZv0jNLFZjKBHX62GgiMNCG19YqO5SyWc0F","api":"s3v4","path":"auto"}
type Credentials struct {
	Endpoint  string `json:"endpoint"`
	Url       string `json:"url"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Api       string `json:"api"`
	Path      string `json:"path"`
}

type Service struct {
	conf   Credentials
	client *minio.Client
}

func NewService(credentialFile string) (*Service, error) {
	var (
		service Service
	)
	file, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &service.conf)
	if err != nil {
		return nil, err
	}
	// 解析 URL
	parsedURL, err := url.Parse(service.conf.Url)
	if err != nil {
		return nil, err
	}

	minioClient, err := minio.New(parsedURL.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(service.conf.AccessKey, service.conf.SecretKey, ""),
		Secure: false,
	})
	if !minioClient.IsOnline() {
		return nil, errors.New("minio connection failed")
	}
	service.client = minioClient
	return &service, err
}

func (s *Service) MakeBucket(bucketName string) error {
	ctx := context.Background()
	// Check to see if we already own this bucket (which happens if you run this twice)
	exist, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if exist {
		return ErrorBucketExist
	}
	return s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: ""})
}

// UploadLocalFile 上传本地文件到oss服务
//
// @param bucketName string 桶对象
// @param ossName string 资源对象名(上传后存储的名字)
// @param file string 具体的文件(相对路径，并完整路径名称)
//
// @return
func (s *Service) UploadLocalFile(bucketName, ossName, file string) error {
	// Upload the zip file with FPutObject
	_, err := s.client.FPutObject(context.Background(),
		bucketName,
		ossName,
		file,
		minio.PutObjectOptions{ContentType: DefaultContentType})
	return err
}

// UploadHttpFile http方式存储文件
//
// @param bucketName string 桶对象
// @param ossName string 资源对象名(上传后存储的名字)
// @param fileSize int64 文件大小
// @param reader io.Reader io文件
//
// @return
func (s *Service) UploadHttpFile(bucketName, ossName string, fileSize int64, reader io.Reader) error {
	_, err := s.client.PutObject(
		context.Background(),
		bucketName,
		ossName,
		reader,
		fileSize,
		minio.PutObjectOptions{ContentType: DefaultContentType})
	return err
}

func (s *Service) GetFile(bucketName, objectName string) (*minio.Object, error) {
	return s.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
}
