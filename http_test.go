//go:build integration
// +build integration

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProcess(t *testing.T) {
	AddUser("myemail", "myname")

	req := httptest.NewRequest(http.MethodGet, "/?email=myemail", nil)
	w := httptest.NewRecorder()

	process(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := UserResponse{
		Email: "myemail",
		Name:  "myname",
	}
	expectedJson, err := json.Marshal(expected)

	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)

	assert.Equal(t, expectedJson, data)
}

func TestPostProcess(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/?email=test@test.com&name=test", nil)
	w := httptest.NewRecorder()
	process(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := []byte("test has been added with next email:  test@test.com\n")

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data), string(expected))
	assert.Equal(t, expected, data)
}

func TestDeleteProcess(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/?email=test@test.com", nil)
	w := httptest.NewRecorder()
	process(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := []byte("test@test.com has been deleted")

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data), string(expected))
	assert.Equal(t, expected, data)
}
