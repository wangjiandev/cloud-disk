package test

import (
	"fmt"
	"testing"
	"wangjiandev/cloud-disk/core/helper"
)

func TestRand(t *testing.T) {
	code := helper.RandCode()
	fmt.Println(code)
}
