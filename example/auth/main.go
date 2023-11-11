package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/redis/go-redis/v9"
	hertzlaravel "github.com/rogerogers/hertz-laravel"
	"github.com/rogerogers/hertz-laravel/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{})
	db, err := gorm.Open(sqlite.Open("laravel"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	appKey, err := utils.B64Decode("aj5F8ziDtRkdSShsM30vsWNNK1UqJcDDbuUsBlGkz2I=")
	if err != nil {
		panic("appKey error")
	}
	h := server.Default()
	h.Use(
		hertzlaravel.Auth(hertzlaravel.WithAppKey(appKey),
			hertzlaravel.WithRedisClient(redisClient),
			hertzlaravel.WithCachePrefix("laravel_database_laravel_cache_:"),
			hertzlaravel.WithDb(db),
		),
	)
	h.Handle(http.MethodGet, "/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, map[string]string{"hello": "hertz-laravel"})
	})
	h.Spin()
}
