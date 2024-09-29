package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	s := &Server{
		PackSizes: []int{3, 2, 1},
	}
	server := httptest.NewServer(http.HandlerFunc(s.PacketHandler))
	defer server.Close()
	resp, err := http.Post(server.URL+"/packets", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(`items=9`)))
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := `{"3":3}`
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if !strings.EqualFold(expected, strings.TrimSpace(string(body))) {
		t.Errorf("expected response body to be %v; got %v;", expected, string(body))
	}
}
