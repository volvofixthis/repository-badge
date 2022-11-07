package repository

import (
	"errors"

	"github.com/dboslee/lru"
	"github.com/google/uuid"
	"github.com/volvofixthis/repository-badge/internal/domain/models"
)

type ScopeRepositoryI interface {
	GetScope(scopeID uuid.UUID) (*models.Scope, error)
	SetScope(scopeID uuid.UUID, scope models.Scope) error
}

type LRUScopeRepository struct {
	cache *lru.Cache[uuid.UUID, models.Scope]
}

func (sr *LRUScopeRepository) GetScope(scopeID uuid.UUID) (*models.Scope, error) {
	value, ok := sr.cache.Get(scopeID)
	if !ok {
		return nil, errors.New("No such scope")
	}
	return &value
}

func (sr *LRUScopeRepository) SetScope(scopeID uuid.UUID, scope models.Scope) error {
	sr.cache.Set(scopeID)
	return nil
}

func NewLRUScopeRepository() *ScopeRepositoryI {
	return &LRUScopeRepository{
		cache: lru.New[string, models.Scope](),
	}
}
