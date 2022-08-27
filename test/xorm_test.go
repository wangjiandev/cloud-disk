package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"wangjiandev/cloud-disk/core/model"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("postgres", "postgresql://wangjian:123456@127.0.0.1:5432/cloud-disk?sslmode=disable")
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
