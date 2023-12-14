package main

import (
	"context"
	"fmt"

	finch "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func main() {
	client := finch.NewClient(option.WithClientID("foo-client-id"), option.WithClientSecret("foo-client-secret"))

	url, err := client.GetAuthURL("products", "https://example.com/redirect", false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("auth url: %s\n", url)

	accessToken, err := client.GetAccessToken(context.TODO(), "my-code", "https://example.com/redirect")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("access token: %s\n", accessToken)

}
