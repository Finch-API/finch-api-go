package finchgo_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func TestCompanyGetWithAssistedSupport_200Response(t *testing.T) {
	// Mock server that returns 200 with Company data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/employer/company" {
			t.Errorf("Expected path /employer/company, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := map[string]interface{}{
			"id":                   "12345678-1234-1234-1234-123456789012",
			"legal_name":           "Test Company LLC",
			"ein":                  "12-3456789",
			"primary_email":        "test@example.com",
			"primary_phone_number": "+11234567890",
			"entity": map[string]interface{}{
				"type":    "llc",
				"subtype": nil,
			},
			"locations":   []interface{}{},
			"accounts":    []interface{}{},
			"departments": []interface{}{},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := finchgo.NewClient(
		option.WithAccessToken("test-token"),
		option.WithBaseURL(server.URL),
	)

	resp, err := client.HRIS.Company.GetWithAssistedSupport(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that we got a Company response
	switch u := resp.AsUnion().(type) {
	case finchgo.Company:
		if u.ID != "12345678-1234-1234-1234-123456789012" {
			t.Errorf("Expected ID 12345678-1234-1234-1234-123456789012, got %s", u.ID)
		}
		if u.LegalName != "Test Company LLC" {
			t.Errorf("Expected LegalName 'Test Company LLC', got %s", u.LegalName)
		}
	case finchgo.CompanyDataSyncInProgress:
		t.Fatalf("Expected Company, got CompanyDataSyncInProgress")
	default:
		t.Fatalf("Unexpected union type: %T", u)
	}
}

func TestCompanyGetWithAssistedSupport_202Response(t *testing.T) {
	// Mock server that returns 202 with sync status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/employer/company" {
			t.Errorf("Expected path /employer/company, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // SDK doesn't check HTTP status, relies on body
		resp := map[string]interface{}{
			"code":       202,
			"finch_code": "data_sync_in_progress",
			"message":    "The company data is being fetched. Please check back later.",
			"name":       "accepted",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := finchgo.NewClient(
		option.WithAccessToken("test-token"),
		option.WithBaseURL(server.URL),
	)

	resp, err := client.HRIS.Company.GetWithAssistedSupport(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that we got a CompanyDataSyncInProgress response
	switch u := resp.AsUnion().(type) {
	case finchgo.Company:
		t.Fatalf("Expected CompanyDataSyncInProgress, got Company")
	case finchgo.CompanyDataSyncInProgress:
		if u.Code != finchgo.CompanyDataSyncInProgressCode202 {
			t.Errorf("Expected Code 202, got %v", u.Code)
		}
		if u.FinchCode != finchgo.CompanyDataSyncInProgressFinchCodeDataSyncInProgress {
			t.Errorf("Expected FinchCode 'data_sync_in_progress', got %s", u.FinchCode)
		}
		if u.Name != finchgo.CompanyDataSyncInProgressNameAccepted {
			t.Errorf("Expected Name 'accepted', got %s", u.Name)
		}
	default:
		t.Fatalf("Unexpected union type: %T", u)
	}
}
