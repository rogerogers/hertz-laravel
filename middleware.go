package hertz_laravel

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/elliotchance/phpserialize"
	"github.com/redis/go-redis/v9"

	"github.com/rogerogers/hertz-laravel/utils"
)

type (
	cookieValue struct {
		Iv    string `json:"iv"`
		Mac   string `json:"mac"`
		Value string `json:"value"`
		Tag   string `json:"tag"`
	}
	userModel struct {
		Id            int
		Password      string
		RememberToken string
	}
	LaravelKey   string
	sessionModel map[any]any
)

func Auth(options ...Option) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		cfg := defaultAuthConfig()
		if len(options) > 0 {
			for _, o := range options {
				o.apply(cfg)
			}
		}

		if utils.InArray(c.FullPath(), cfg.ignorePaths) {
			c.Next(ctx)
			return
		}
		userId, err := getUserId(c, cfg)
		if err != nil || userId == 0 {
			cfg.UnAuthHandler(ctx, c)
			return
		}
		ctx = context.WithValue(ctx, LaravelKey("userId"), userId)

		c.Next(ctx)
	}
}

func getUserId(c *app.RequestContext, cfg *authConfig) (int, error) {
	laravelSession := c.Cookie(cfg.sessionCookieName)
	userId, err := getUserIdFromLaravelSession(laravelSession, cfg)
	if err != nil {
		rememberWeb := c.Cookie(cfg.rememberCookieName)
		userId, err = getUserIdFromRememberWeb(rememberWeb, cfg)
		if err != nil {
			return 0, nil
		}
	}
	return userId, nil
}

func parseCookie(cookie []byte, cfg *authConfig) (string, error) {
	decoded, err := utils.B64SafeDecode(cookie)
	if err != nil {
		return "", err
	}
	cv := &cookieValue{}
	err = json.Unmarshal(decoded, cv)
	if err != nil {
		return "", err
	}
	iv, err := utils.B64Decode(cv.Iv)
	if err != nil {
		return "", err
	}
	decryptValue, err := utils.Aes256CbcDecrypt(cv.Value, cfg.appKey, iv)
	if err != nil {
		return "", err
	}

	return decryptValue, nil
}

func getUserIdFromLaravelSession(cookie []byte, cfg *authConfig) (int, error) {
	decryptValue, err := parseCookie(cookie, cfg)
	if err != nil {
		return 0, err
	}
	ssArr := strings.Split(decryptValue, "|")
	ssId := ssArr[0]
	if !cfg.disableEncryptCookies && len(ssArr) == 2 && !utils.InArray(cfg.sessionCookieName, cfg.exceptEncryptCookies) {
		ssId = ssArr[1]
	}
	redisRes := cfg.redisClient.Get(context.Background(), strings.Join([]string{cfg.cachePrefix, ssId}, ""))
	if redisRes.Err() == redis.Nil {
		return 0, errors.New("401")
	}

	var payloadByte string
	redisResByte, err := redisRes.Bytes()
	if err != nil {
		return 0, errors.New("401")
	}
	err = phpserialize.Unmarshal(redisResByte, &payloadByte)

	if err != nil {
		return 0, errors.New("401")
	}

	var payload sessionModel

	switch cfg.serialization {
	case PhpSerialize:
		err = phpserialize.Unmarshal([]byte(payloadByte), &payload)
		if err != nil {
			return 0, errors.New("401")
		}
	case JsonSerialize:
		err = json.Unmarshal([]byte(payloadByte), &payload)
		if err != nil {
			return 0, errors.New("401")
		}
	}

	userid, ok := payload["login_web_59ba36addc2b2f9401580f014c7f58ea4e30989d"]
	if !ok {
		return 0, errors.New("401")
	}

	return int(userid.(int64)), nil
}

func getUserIdFromRememberWeb(cookie []byte, cfg *authConfig) (int, error) {
	decryptValue, err := parseCookie(cookie, cfg)
	if err != nil {
		return 0, err
	}
	ssArr := strings.Split(decryptValue, "|")
	userid, rememberToken, hashedPass := ssArr[0], ssArr[1], ssArr[2]
	if !cfg.disableEncryptCookies && len(ssArr) == 4 && !utils.InArray(cfg.rememberCookieName, cfg.exceptEncryptCookies) {
		userid, rememberToken, hashedPass = ssArr[1], ssArr[2], ssArr[3]
	}

	var user userModel

	err = cfg.db.Table(cfg.tableName).First(&user, userid).Error
	if err != nil {
		return 0, err
	}
	if strconv.Itoa(user.Id) == userid && user.Password == hashedPass && user.RememberToken == rememberToken {
		return user.Id, nil
	}

	return 0, errors.New("401")
}
