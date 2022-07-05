package auth

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/zmicro-team/zmicro/core/log"
)

type Auth struct {
	privateKey string
	publicKey  string
	client     *redis.Client
}

func NewAuth(privateKey, publicKey string, client *redis.Client) *Auth {
	a := &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
		client:     client,
	}
	return a
}

type Account struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Issuer       string            `json:"issuer"`
	Metadata     map[string]string `json:"metadata"`
	Scopes       []string          `json:"scopes"`
	RefreshToken string            `json:"refresh_token"`
}

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Created      time.Time `json:"created"`
	Expires      time.Time `json:"expires"`
}

type authClaims struct {
	Type     string            `json:"type"`
	Scopes   []string          `json:"scopes"`
	Metadata map[string]string `json:"metadata"`
	Name     string            `json:"name"`

	jwt.RegisteredClaims
}

const (
	storePrefixRefreshToken = "refresh"
	joinKey                 = ":"
)

func (a *Auth) GenerateToken(id string, opts ...GenerateOption) (*Token, error) {
	options := NewGenerateOptions(opts...)
	acc := &Account{
		ID:       id,
		Type:     options.Type,
		Issuer:   options.Issuer,
		Metadata: options.Metadata,
		Scopes:   options.Scopes,
	}

	if acc.Metadata == nil {
		acc.Metadata = map[string]string{}
	}

	tok, err := generate(a.privateKey, acc, time.Hour)
	if err != nil {
		return nil, err
	}

	var refreshToken string
	key := strings.Join([]string{storePrefixRefreshToken, id}, joinKey)
	val, err := a.client.Get(context.Background(), key).Bytes()
	if err != nil {
		if err == redis.Nil {
			// 不存在刷新令牌，写入一个刷新令牌
			refreshToken = uuid.New().String()
			acc.RefreshToken = refreshToken
			b, err := json.Marshal(acc)
			if err != nil {
				return nil, ErrEncodingToken
			}
			if err = a.client.Set(
				context.Background(),
				key,
				b,
				time.Hour*24*180,
			).Err(); err != nil {
				log.Errorf("GenerateToken error = %v", err)
				return nil, err
			}
		} else {
			log.Errorf("GenerateToken error = %v", err)
			return nil, err
		}
	} else {
		if err := json.Unmarshal(val, &acc); err != nil {
			return nil, err
		}
		refreshToken = acc.RefreshToken
		// 重新登录,刷新令牌将被续期
		if err := a.client.Set(
			context.Background(),
			key,
			val,
			time.Hour*24*180,
		).Err(); err != nil {
			log.Errorf("GenerateToken error = %v", err)
			return nil, err
		}
	}

	token := &Token{
		AccessToken:  tok,
		RefreshToken: refreshToken,
		Created:      time.Now(),
		Expires:      time.Now().Add(time.Hour),
	}

	return token, nil
}

func (a *Auth) Inspect(token string) (*Account, error) {
	return inspect(a.publicKey, token)
}

// 后续考虑 refresh token中也保存用户信息，以简化metadata处理
func (a *Auth) RefreshToken(id, refreshToken string, opts ...RefreshTokenOption) (*Token, error) {
	options := NewRefreshTokenOptions(opts...)

	key := strings.Join([]string{storePrefixRefreshToken, id}, joinKey)
	val, err := a.client.Get(context.Background(), key).Bytes()

	if err != nil {
		if err == redis.Nil {
			err = ErrInvalidToken
		}

		return nil, err
	}

	var acc Account

	if err := json.Unmarshal(val, &acc); err != nil {
		return nil, err
	}

	if refreshToken != acc.RefreshToken {
		return nil, ErrInvalidToken
	}

	if acc.Metadata == nil {
		acc.Metadata = map[string]string{}
	}

	for k, v := range options.Metadata {
		acc.Metadata[k] = v
	}

	tok, err := generate(a.privateKey, &acc, time.Hour)
	if err != nil {
		return nil, err
	}

	token := &Token{
		AccessToken:  tok,
		RefreshToken: refreshToken,
		Created:      time.Now(),
		Expires:      time.Now().Add(time.Hour),
	}

	return token, nil
}
