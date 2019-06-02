package texthash

// TextHash é um conjunto que armazena os dados de um token, seu hash e a data de criação
type TextHash struct {
	Token     string `json:"token,omitempty"`
	Hash      string `json:"hash,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// Service é um serviço que expõe as operações da API de TextHash
type Service interface {
	Create(textHash *TextHash) (*TextHash, error)
	FindByHash(hash string) *TextHash
	FindAll() []*TextHash
}

// Repository é um repositório de dados que salva as informações
// dos TextHash
type Repository interface {
	Insert(textHash *TextHash) error
	FindByHash(hash string) *TextHash
	FindAll() []*TextHash
}
