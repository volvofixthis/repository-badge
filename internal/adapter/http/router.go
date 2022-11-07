package httpadapter

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
)

func NewRouter() *gin.Engine {
	tonic.SetErrorHook(jujerr.ErrHook)
	engine := gin.New()
	badgeHandler := NewBadgeHandler()
	engine.GET("/buildbadge", tonic.Handler(badgeHandler.GetBuildBadge, 200))
	engine.GET("/coveragebadge", tonic.Handler(badgeHandler.GetCoverageBadge, 200))
	return engine
}
