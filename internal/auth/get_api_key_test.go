package auth

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestAPI(t *testing.T) {

	jsonBody := []byte(`{"name": "John Doe", "age": 30}`)
	req, err := http.NewRequest(http.MethodPost, "https://example.com/api/users", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Add("Authorization", "ApiKey testingforkeycheck")

	got, _ := GetAPIKey(req.Header)
	want := "testingforkeycheck"

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	req.Header.Del("Authorization")

	_, err = GetAPIKey(req.Header)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

}
