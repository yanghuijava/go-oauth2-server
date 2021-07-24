package test

import (
	"fmt"
	"go-oauth2-server/service"
	"testing"
)

func TestClientIdSecret(t *testing.T) {
	client := &service.DefaultClientProduce{}
	ci, cs := client.ClientIdSecret()
	fmt.Printf("clientId=%s,clientSecret=%s\n", ci, cs)
}
