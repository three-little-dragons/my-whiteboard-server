package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/three-little-dragons/my-whiteboard-server/internal/app/service"
	"github.com/three-little-dragons/my-whiteboard-server/internal/pkg/com"
	"github.com/three-little-dragons/my-whiteboard-server/internal/pkg/validate"
)

func setupRouter() *gin.Engine {
	//gin.SetMode(config.Val.Gin.LogLevel)
	r := gin.Default()

	store := cookie.NewStore([]byte(uuid.NewString()))
	r.Use(sessions.Sessions("session", store))

	initRouter(r)

	return r
}

func initRouter(r *gin.Engine) {
	key := "user_id"
	r.Use(func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userId := session.Get(key)
		if userId == nil {
			userId = service.GenerateId()
			session.Set(key, userId)
			err := session.Save()
			if err != nil {
				panic(err)
			}
		}
		ctx.Set(key, userId)
	})
	getUserId := func(c *gin.Context) int64 { return c.GetInt64(key) }

	api := r.Group("/api")
	{
		api.GET("/new_paint", func(c *gin.Context) {
			vo := validate.Struct(c, &struct {
				Nickname string `form:"nickname" valid:"required"`
			}{})
			if vo != nil {
				com.Success(c, service.NewPaint(getUserId(c), vo.Nickname))
			}
		})
	}
}
