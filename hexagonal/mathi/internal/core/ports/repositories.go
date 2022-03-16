package pors

import "mathi/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
