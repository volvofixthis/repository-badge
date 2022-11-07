package service

import (
	"github.com/google/uuid"
	"github.com/volvofixthis/repository-badge/internal/domain/models"
	"github.com/volvofixthis/repository-badge/internal/domain/repository"
)

type ScopeServiceI interface {
	GetScope(scopeID uuid.UUID) (*models.Scope, error)
	SetScope(scopeID uuid.UUID, scope models.Scope) error
}

type ScopeService struct {
	sr repository.ScopeRepositoryI
}

func (ss *ScopeServiceI) GetScope(scopeID uuid.UUID) *models.Scope {
	return ss.sr.GetScope(scopeID)
}

func NewScopeService() ScopeServiceI {
	return &ScopeService{
		sr: repository.NewLRUScopeRepository(),
	}
}
