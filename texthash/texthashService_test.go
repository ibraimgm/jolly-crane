package texthash_test

import (
	"testing"

	"github.com/ibraimgm/jolly-crane/texthash"
)

func TestCreate(t *testing.T) {
	service := texthash.NewService(texthash.NewInMemRepository())

	tests := []struct {
		input    texthash.TextHash
		expected string
	}{
		{input: texthash.TextHash{Token: "foo"}, expected: "76d3bc41c9f588f7fcd0d5bf4718f8f84b1c41b20882703100b9eb9413807c01"},
		{input: texthash.TextHash{Token: "bar"}, expected: "cceefd7e0545bcf8b6d19f3b5750c8a3ee8350418877bc6fb12e32de28137355"},
		{input: texthash.TextHash{Token: "baz"}, expected: "9713fc828dd6313c2975127f77e1681499b9d80c0bef9645837ed6555f24fb76"},
	}

	for i, test := range tests {
		created, err := service.Create(&test.input)

		if err != nil {
			t.Errorf("Case %v, shoul not have returned an error. Error: %v", i, err)
		}

		if created == nil {
			t.Errorf("Case %v, should have returned the saved object.", i)
		}

		if created.Hash != test.expected {
			t.Errorf("Case %v, expected hash '%v', received '%v'.", i, test.expected, created.Hash)
		}

		if created.CreatedAt == "" {
			t.Errorf("Case %v, createdAt should not be empty.", i)
		}
	}
}

func TestCreateInvalid(t *testing.T) {
	service := texthash.NewService(texthash.NewInMemRepository())

	if _, err := service.Create(&texthash.TextHash{Token: ""}); err == nil {
		t.Errorf("Should have returned an error.")
	}

	if _, err := service.Create(nil); err == nil {
		t.Errorf("Should have returned an error.")
	}
}
