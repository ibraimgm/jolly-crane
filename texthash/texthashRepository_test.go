package texthash_test

import (
	"testing"

	"github.com/ibraimgm/jolly-crane/texthash"
)

func TestInsertAndFind(t *testing.T) {
	repo := texthash.NewInMemRepository()

	toInsert := []texthash.TextHash{
		{Token: "foo", Hash: "001", CreatedAt: "xxxx"},
		{Token: "bar", Hash: "002", CreatedAt: "xxxx"},
		{Token: "baz", Hash: "003", CreatedAt: "xxxx"},
	}

	for i, item := range toInsert {
		err := repo.Insert(&item)

		if err != nil {
			t.Errorf("Case %v, insert should not return error. Received: %v", i, err)
		}

		if saved := repo.FindByHash(item.Hash); saved == nil {
			t.Errorf("Case %v, item '%v' should be in the repository, but received nil.", i, item.Hash)
		}
	}

	all := repo.FindAll()
	if len(all) != len(toInsert) {
		t.Errorf("Should have %v itens in repository, but %v were found.", len(toInsert), len(all))
	}
}

func TestInsertEmptyValues(t *testing.T) {
	repo := texthash.NewInMemRepository()

	toInsert := []texthash.TextHash{
		{Token: "foo", Hash: "001"},
		{Token: "bar", CreatedAt: "xxxx"},
		{Hash: "003", CreatedAt: "xxxx"},
	}

	for i, item := range toInsert {
		err := repo.Insert(&item)

		if err == nil {
			t.Errorf("Case %v, should have returned an error. Item: %v.", i, item)
		}
	}
}

func TestFindByHashReturnsClone(t *testing.T) {
	repo := texthash.NewInMemRepository()

	original := texthash.TextHash{Token: "A", Hash: "1", CreatedAt: "Z"}
	repo.Insert(&original)

	found := repo.FindByHash(original.Hash)
	if found == nil {
		t.Errorf("Should have found the item with hash %v.", original.Hash)
	}

	found.Token = "X"
	if original.Token == found.Token {
		t.Errorf("Changes on the found object shoul not change the original.")
	}

	original.Token = "X"
	found = repo.FindByHash(original.Hash)
	if original.Token == found.Token {
		t.Errorf("Changes on the original object shoul not change the stored one.")
	}
}

func TestFindAllReturnsClose(t *testing.T) {
	repo := texthash.NewInMemRepository()

	toInsert := []texthash.TextHash{
		{Token: "foo", Hash: "001", CreatedAt: "xxxx"},
		{Token: "bar", Hash: "002", CreatedAt: "xxxx"},
		{Token: "baz", Hash: "003", CreatedAt: "xxxx"},
	}

	originalItens := make(map[string]*texthash.TextHash)

	for _, item := range toInsert {
		clone := item
		originalItens[item.Hash] = &clone
		repo.Insert(&item)
	}

	all := repo.FindAll()
	for i, found := range all {
		original := originalItens[found.Hash]

		if found == nil {
			t.Errorf("Case %v, should have found the item with hash %v.", i, original.Hash)
		}

		found.Token = "X"
		if original.Token == found.Token {
			t.Errorf("Case %v, changes on the found object shoul not change the original.", i)
		}

		original.Token = "X"
		found = repo.FindByHash(original.Hash)
		if original.Token == found.Token {
			t.Errorf("Case %v, changes on the original object shoul not change the stored one.", i)
		}
	}

}
