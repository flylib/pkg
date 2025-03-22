package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
	"testing"
)

func TestGetFile(t *testing.T) {
	service, err := NewService("./credentials.json")
	if err != nil {
		t.Fatal(err)
	}
	err = service.UploadLocalFile("image", "test.png", "./baber_logo.jpg")
	if err != nil {
		t.Fatal(err)
	}
	//file, err := service.GetFile("image", "image/shop/queuing_shop.png")
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//stat, err := file.Stat()
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//t.Log("TestUploadFile:", stat.Size)
}

func TestUploadFile(t *testing.T) {
	// FileUploader.go MinIO example

	ctx := context.Background()
	endpoint := "192.168.1.50:9090"
	accessKeyID := "5srZg1yepoRFRENJqqZb"
	secretAccessKey := "nltqPIeESgt4yoZRbTn8CI94pGTECQtfonxdDAwE"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		t.Fatal(err)
	}
	object, err := minioClient.GetObject(ctx, "image", "shop/images.png", minio.GetObjectOptions{})
	if err != nil {
		t.Fatal(err)
	}

	defer object.Close()

	// 创建一个足够大的切片以读取对象数据
	//data := make([]byte, 101024) // 例如，创建一个 1024 字节大小的切片
	//
	//n, err := object.Read(data)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log("TestUploadFile:", n)

	// 将对象内容写入本地文件
	localFile, err := os.Create("images.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	// 复制对象内容到本地文件
	if _, err := io.Copy(localFile, object); err != nil {
		log.Fatalln(err)
	}

	t.Log("File downloaded successfully")

}
