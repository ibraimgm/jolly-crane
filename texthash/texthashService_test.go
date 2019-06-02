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

		if created.Token != "" {
			t.Errorf("Case %v, token should be null", i)
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

func TestFind(t *testing.T) {
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
		if _, err := service.Create(&test.input); err != nil {
			t.Errorf("Case %v, should not have returned an error on save. Error: %v.", i, err)
		}

		found := service.FindByHash(test.expected)

		if found == nil {
			t.Errorf("Case %v, should have returned a item.", i)
		}

		if found.Hash != test.expected {
			t.Errorf("Case %v, hashes should be equal. Expected: '%v', received '%v'", i, test.expected, found.Hash)
		}

		if found.Token != test.input.Token {
			t.Errorf("Case %v, tokens should be equal. Expected: '%v', received '%v'", i, test.input.Token, found.Token)
		}
	}
}

func TestFindAll(t *testing.T) {
	service := texthash.NewService(texthash.NewInMemRepository())

	tests := []struct {
		input texthash.TextHash
		hash  string
	}{
		{input: texthash.TextHash{Token: "Texto a ser cifrado"}, hash: "369ee900da8fd705ea41965e3df5df6cb7cc87a682bd29cf6c5c99253e9f87d5"},
		{input: texthash.TextHash{Token: "Texto a ser cifrado 2"}, hash: "a37f2f5b614918c29b2d89a75810568f4926febd91be04190680fb0d9d52bb49"},
		{input: texthash.TextHash{Token: "Texto a ser cifrado 3"}, hash: "aae889e4b2d49258c278632345044426982d8adcd701443ab9b8ede1f23a9032"},
		{input: texthash.TextHash{Token: "Texto a ser cifrado 4"}, hash: "dc43782555f0d52629dfd1fce8617ff6b53ec4937755ba73398feb2dd27722b0"},
	}

	for i, test := range tests {
		created, err := service.Create(&test.input)

		if err != nil {
			t.Errorf("Case %v, shoul not have returned an error. Error: %v", i, err)
		}

		if created == nil {
			t.Errorf("Case %v, should have returned the saved object.", i)
		}

		if created.Hash != test.hash {
			t.Errorf("Case %v, expected hash '%v', received '%v'.", i, test.hash, created.Hash)
		}
	}

	all := service.FindAll()
	countTests := len(tests)
	countAll := len(all)

	if countAll != countTests {
		t.Errorf("Shoul return the same number of records that were saved. Expected: %v, received %v.", countTests, countAll)
	}
}
