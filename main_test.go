//go:build unit || all
// +build unit all

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	email string
	name  string
}

func InitData(tests []testCase) {
	for _, test := range tests {
		users.mu.Lock()
		users.storage[test.email] = test.name
		users.mu.Unlock()
	}
}

func TestGetUser(t *testing.T) {

	var tests = []testCase{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	InitData(tests)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res := GetUser(test.email)
			require.Equal(t, test.name, res)
		})
	}
}

func TestAddUser(t *testing.T) {
	var tests = []testCase{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	InitData(tests)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res := GetUser(test.email)
			require.Equal(t, test.name, res)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	var tests = []testCase{
		{"test@gmail.com", "testname"},
		{"test2@gmail.com", ""},
	}
	InitData(tests)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			DeleteUser(test.email)

			users.mu.Lock()
			_, ok := users.storage[test.email]
			users.mu.Unlock()

			require.False(t, ok)
		})
	}
}
