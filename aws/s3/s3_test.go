package s3_test

import (
	"bytes"
	"fmt"
	"github.com/nicktaobo/go_tool/aws/s3"
	"os"
	"testing"
)

const ()

var client *s3.Client

func init() {
	client = s3.NewS3(&s3.Conf{AccessKey: "xxx", AccessSecret: "xxx", Region: "xxx"})
}

func TestUpload(t *testing.T) {
	f := "/Users/hank/Pictures/bing/' '.jpg"
	bs, _ := os.ReadFile(f)
	err := client.UploadFile("xxx", "a/00027c7c193111eebf32063de4e620c1.jpg", bytes.NewReader(bs))
	fmt.Println(err)
}

func TestDelete(t *testing.T) {
	err := client.DeleteFile("xxx", "a/00027c7c193111eebf32063de4e620ce.jpg")
	fmt.Println(err)
}
