package texthash

import (
	"fmt"
	"time"

	"golang.org/x/crypto/sha3"
)

type serviceImpl struct {
	repo Repository
}

func (s *serviceImpl) Create(textHash *TextHash) (*TextHash, error) {
	if textHash == nil || textHash.Token == "" {
		return nil, fmt.Errorf("É necessário informar um objeto json com o campo 'token' para criar um hash")
	}

	hash := sha3.Sum256([]byte(textHash.Token))
	textHash.Hash = fmt.Sprintf("%x", hash)
	textHash.CreatedAt = fmt.Sprintf("%v", time.Now().UTC())

	if err := s.repo.Insert(textHash); err != nil {
		return nil, err
	}

	saved := s.repo.FindByHash(textHash.Hash)
	saved.Token = ""
	return saved, nil
}

func (s *serviceImpl) FindByHash(hash string) *TextHash {
	return s.repo.FindByHash(hash)
}

func (s *serviceImpl) FindAll() []*TextHash {
	return s.repo.FindAll()
}

// NewService retorna um novo serviço para a API de hashes, usando o
// repositório de dados passado como parâmetro
func NewService(repository Repository) Service {
	return Service(&serviceImpl{repository})
}
