package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	p := fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	fmt.Println(p)
}
