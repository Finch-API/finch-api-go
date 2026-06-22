package main

import (
	"context"
	"fmt"

	finch "github.com/Finch-API/finch-api-go/v2"
	"github.com/Finch-API/finch-api-go/v2/option"
)

func main() {
	client := finch.NewClient(option.WithClientID("foo-client-id"), option.WithClientSecret("foo-client-secret"))

	url, err := client.GetAuthURL("products", "https://example.com/redirect", false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("auth url: %s\n", url)

	accessTokenResponse, err := client.AccessTokens.New(context.TODO(), finch.AccessTokenNewParams{
		Code:        finch.F("my-code"),
		RedirectUri: finch.F("https://example.com/redirect"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("access token: %s\n", accessTokenResponse.AccessToken)

}
