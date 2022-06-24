package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetRegStr(t *testing.T) {
	reg := regexp.MustCompile(`([0-9.]+):([0-9]+)`)
	fmt.Println(reg.FindAllString(("proxy/vmess/encoding: failed to read request header > read tcp 104.224.187.69:24328->117.155.240.186:5536: i/o timeout"), -1))
}
