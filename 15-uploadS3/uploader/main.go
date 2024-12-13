package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"sync"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"ASDFASDFASDFASFD",
				"ASDFASdfasregq35ghq3%E",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// ,100 é pra criar buffer para o channel poder suportar ate 100 informacoes/structs antes de descarregar
	// controla a quantidade de uploads ativos, no caso as goroutines criadas pra upload
	uploadControl := make(chan struct{}, 100)

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)

		// faz com que crie até 100 goroutines(uploads), depois ele nao deixa mais fazer upload, ate o channel esvaziar,
		// criando assim um controle de quantidade de goroutines/uploads realizados simultaneamente
		uploadControl <- struct{}{}

		go uploadFile(files[0].Name(), uploadControl)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("open file %s failed.\n", completeFileName)
		<-uploadControl // esvazia o channel
		return
	}
	defer f.Close()

	// subindo arquivo para aws
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("upload file %s failed.\n", completeFileName)
		<-uploadControl // esvazia o channel
		return
	}

	fmt.Printf("File %s uploaded successfully.\n", completeFileName)
	<-uploadControl // esvazia o channel
}
