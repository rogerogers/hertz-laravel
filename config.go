package hertz_laravel

import (
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
}

func defaultAuthConfig() *authConfig {
	return &authConfig{
		sessionCookieName:  "laravel_session",
		rememberCookieName: "remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d",
		ignorePaths:        []string{"/login", "/api/login"},
		serialization:      PhpSerialize,
		tableName:          "users",
	}
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

func WithAppKey(appKey []byte) Option {
	return option(func(cfg *authConfig) {
		cfg.appKey = appKey
	})
}

func WithRedisClient(client redis.UniversalClient) Option {
	return option(func(cfg *authConfig) {
		cfg.redisClient = client
	})
}

func WithCachePrefix(prefix string) Option {
	return option(func(cfg *authConfig) {
		cfg.cachePrefix = prefix
	})
}

func WithDb(db *gorm.DB) Option {
	return option(func(cfg *authConfig) {
		cfg.db = db
	})
}

func WithSerialization(serialization serialization) Option {
	return option(func(cfg *authConfig) {
		cfg.serialization = serialization
	})
}
