package texthash

import (
	"fmt"
)

type inMemRepository map[string]*TextHash

func (r inMemRepository) Insert(textHash *TextHash) error {
	if textHash == nil {
		return fmt.Errorf("Não é possível salvar um valor nulo")
	}

	if textHash.Token == "" {
		return fmt.Errorf("É necessário informar o campo 'token'")
	}

	if textHash.Hash == "" {
		return fmt.Errorf("É necessário informar o campo 'hash'")
	}

	if textHash.CreatedAt == "" {
		return fmt.Errorf("É necessário informar o campo 'createdAt'")
	}

	if _, ok := r[textHash.Hash]; ok {
		return fmt.Errorf("Hash %v já existe", textHash.Hash)
	}

	clone := *textHash
	r[textHash.Hash] = &clone

	return nil
}

func (r inMemRepository) FindByHash(hash string) *TextHash {
	item := r[hash]
	clone := *item
	return &clone
}

func (r inMemRepository) FindAll() []*TextHash {
	arr := make([]*TextHash, len(r))

	i := 0
	for _, value := range r {
		clone := *value
		arr[i] = &clone
		i++
	}

	return arr
}

// NewInMemRepository retorna um novo repositório de hashes salvo em memória
func NewInMemRepository() Repository {
	return inMemRepository(make(map[string]*TextHash))
}
