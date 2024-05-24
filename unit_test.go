//go:build unit || all
// +build unit all

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	var tests = []struct {
		email string
		name  string
	}{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			users.mu.Lock()
			users.storage[test.email] = test.name
			users.mu.Unlock()

			res := GetUser(test.email)

			require.Equal(t, test.name, res)
		})
	}
}

func TestAddUser(t *testing.T) {
	var tests = []struct {
		email string
		name  string
	}{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			users.mu.Lock()
			users.storage[test.email] = test.name
			users.mu.Unlock()

			res := GetUser(test.email)

			require.Equal(t, test.name, res)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	var tests = []struct {
		email string
		name  string
	}{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			users.mu.Lock()
			users.storage[test.email] = test.name
			users.mu.Unlock()

			DeleteUser(test.email)

			users.mu.Lock()
			_, ok := users.storage[test.email]
			users.mu.Unlock()

			require.False(t, ok)
		})
	}
}
