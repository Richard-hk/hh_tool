package tool

import (
	"fmt"
	"testing"
)

func TestGetSiteDocPost(t *testing.T) {
	res := GetBitDomainPost("https://register-api.did.id/v1/account/search", "rookie.bit")
	fmt.Println(res)
}
