package ports

import "go-hex-example/internal/core/domain"

type GameRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
