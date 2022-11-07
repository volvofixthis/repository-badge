package httpadapter

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volvofixthis/repository-badge/internal/usecase/badge"
)

type BadgeRequest struct {
	ScopeID string `query:"scopeid" validate:"required,uuid4"`
}

func (br *BadgeRequest) GetScopeID() uuid.UUID {
	uuid := uuid.MustParse(br.ScopeID)
	return uuid
}

type BadgeHandler struct {
	bu badge.BadgeUsecaseI
}

func (bh *BadgeHandler) GetBuildBadge(c *gin.Context, in *BadgeRequest) error {
	buf, err := bh.bu.GetBuildBadge(in.GetScopeID())
	if err != nil {
		return err
	}
	c.Data(200, "image/png", buf)
	return nil
}

func GetCoverageBadge(c *gin.Context, in *BadgeRequest) error {
	buf, err := bh.bu.GetCoverageBadge(in.GetScopeID())
	if err != nil {
		return err
	}
	c.Data(200, "image/png", buf)
	return nil
}

func NewBadgeHandler() *BadgeHandler {
	return &BadgeHandler{
		bu: badge.NewBadgeUsecase(),
	}
}
