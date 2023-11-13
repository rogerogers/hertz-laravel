package hertz_laravel

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Option interface {
	apply(cfg *authConfig)
}

type option func(cfg *authConfig)

func (fn option) apply(cfg *authConfig) {
	fn(cfg)
}

type serialization string

const (
	PhpSerialize  serialization = "php"
	JsonSerialize serialization = "json"
)

type unAuthHandler func(ctx context.Context, c *app.RequestContext)

type authConfig struct {
	appKey                []byte
	sessionCookieName     string
	rememberCookieName    string
	ignorePaths           []string
	disableEncryptCookies bool
	exceptEncryptCookies  []string
	db                    *gorm.DB
	tableName             string
	redisClient           redis.UniversalClient
	serialization         serialization
	cachePrefix           string
	UnAuthHandler         unAuthHandler
}

func defaultAuthConfig() *authConfig {
	return &authConfig{
		sessionCookieName:  "laravel_session",
		rememberCookieName: "remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d",
		ignorePaths:        []string{"/login", "/api/login"},
		serialization:      PhpSerialize,
		tableName:          "users",
		UnAuthHandler: func(ctx context.Context, c *app.RequestContext) {
			c.AbortWithStatus(401)
		},
	}
}

func WithAppKey(appKey []byte) Option {
	return option(func(cfg *authConfig) {
		cfg.appKey = appKey
	})
}

func WithSessionCookieName(cm string) Option {
	return option(func(cfg *authConfig) {
		cfg.sessionCookieName = cm
	})
}

func WithRememberCookieName(cm string) Option {
	return option(func(cfg *authConfig) {
		cfg.rememberCookieName = cm
	})
}

func WithIgnorePaths(paths []string) Option {
	return option(func(cfg *authConfig) {
		cfg.ignorePaths = paths
	})
}

func WithDisableEncryptCookies(disable bool) Option {
	return option(func(cfg *authConfig) {
		cfg.disableEncryptCookies = disable
	})
}

func WithExceptEncryptCookies(except []string) Option {
	return option(func(cfg *authConfig) {
		cfg.exceptEncryptCookies = except
	})
}

func WithDb(db *gorm.DB) Option {
	return option(func(cfg *authConfig) {
		cfg.db = db
	})
}

func WithTableName(tableName string) Option {
	return option(func(cfg *authConfig) {
		cfg.tableName = tableName
	})
}

func WithRedisClient(client redis.UniversalClient) Option {
	return option(func(cfg *authConfig) {
		cfg.redisClient = client
	})
}

func WithSerialization(serialization serialization) Option {
	return option(func(cfg *authConfig) {
		cfg.serialization = serialization
	})
}

func WithCachePrefix(prefix string) Option {
	return option(func(cfg *authConfig) {
		cfg.cachePrefix = prefix
	})
}

func WithUnAuthHandler(handler unAuthHandler) Option {
	return option(func(cfg *authConfig) {
		cfg.UnAuthHandler = handler
	})
}
