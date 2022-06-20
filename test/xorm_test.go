package test

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"testing"
	"wangjiandev/cloud-disk/core/model"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*model.UserBasic, 0)

	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())

}

func TestMd5(t *testing.T) {
	p := fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	fmt.Println(p)
}
