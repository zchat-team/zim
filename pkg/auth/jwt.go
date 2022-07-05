package auth

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generate(privateKey string, acc *Account, d time.Duration) (string, error) {
	var priv []byte
	if strings.HasPrefix(privateKey, "-----BEGIN RSA PRIVATE KEY-----") ||
		strings.HasPrefix(privateKey, "-----BEGIN PRIVATE KEY-----") {
		priv = []byte(privateKey)
	} else {
		var err error
		priv, err = base64.StdEncoding.DecodeString(privateKey)
		if err != nil {
			return "", err
		}
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return "", ErrEncodingToken
	}

	expiresAt := jwt.NewNumericDate(time.Now().Add(d))
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaims{
		Type: acc.Type, Scopes: acc.Scopes, Metadata: acc.Metadata,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   acc.ID,
			Issuer:    acc.Issuer,
			ExpiresAt: expiresAt,
		},
	})
	token, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func inspect(publicKey string, t string) (*Account, error) {
	if len(strings.Split(t, ".")) != 3 {
		return nil, ErrInvalidToken
	}

	var pub []byte
	if strings.HasPrefix(publicKey, "-----BEGIN RSA PUBLIC KEY-----") ||
		strings.HasPrefix(publicKey, "-----BEGIN PUBLIC KEY-----") ||
		strings.HasPrefix(publicKey, "-----BEGIN CERTIFICATE-----") {
		pub = []byte(publicKey)
	} else {
		var err error
		pub, err = base64.StdEncoding.DecodeString(publicKey)
		if err != nil {
			return nil, err
		}
	}

	res, err := jwt.ParseWithClaims(t, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(pub)
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if !res.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := res.Claims.(*authClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &Account{
		ID:       claims.Subject,
		Issuer:   claims.Issuer,
		Type:     claims.Type,
		Scopes:   claims.Scopes,
		Metadata: claims.Metadata,
	}, nil
}
