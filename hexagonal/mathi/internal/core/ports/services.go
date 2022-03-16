package ports

import "mathi/internal/core/domain"

type GamesService interface {
	et(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}
