package repositories

import "fast-search/internal/domain/entities"

type IFQueryRepository interface {
	Save(query *entities.Query) error
}
