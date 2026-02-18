package main

import (
	"context"
	"fmt"
	"os"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func main() {
	client := finchgo.NewClient(
		option.WithAccessToken(os.Getenv("FINCH_ACCESS_TOKEN")),
	)

	resp, err := client.HRIS.Company.GetWithAssistedSupport(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Handle the union response
	switch u := resp.AsUnion().(type) {
	case finchgo.Company:
		// 200 OK - Company data is ready
		fmt.Printf("Company ID: %s\n", u.ID)
		fmt.Printf("Legal Name: %s\n", u.LegalName)
		fmt.Printf("EIN: %s\n", u.Ein)
		fmt.Printf("Primary Email: %s\n", u.PrimaryEmail)

	case finchgo.CompanyDataSyncInProgress:
		// 202 Accepted - Data sync in progress
		fmt.Printf("Status: %s (code %v)\n", u.Name, u.Code)
		fmt.Printf("Message: %s\n", u.Message)
		fmt.Printf("Finch Code: %s\n", u.FinchCode)
		fmt.Println("\nPlease retry this request later. Data is still syncing from the provider.")

	default:
		fmt.Fprintf(os.Stderr, "Unexpected response type: %T\n", u)
		os.Exit(1)
	}
}
