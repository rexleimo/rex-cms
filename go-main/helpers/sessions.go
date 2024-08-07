package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type SessionHelper struct {
	Ctx sessions.Session
}

func RegisterSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("rex"))
	r.Use(sessions.Sessions("rex_session", store))
}

func (h *SessionHelper) New(c *gin.Context) {
	h.Ctx = sessions.Default(c)
}

func (h *SessionHelper) Get(key string) interface{} {
	return h.Ctx.Get(key)
}

func (h *SessionHelper) Set(key string, value interface{}) {
	h.Ctx.Set(key, value)
	h.Ctx.Save()
}
