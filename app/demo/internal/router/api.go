package router

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zchat-team/zim/app/demo/internal/middleware"
	sContact "github.com/zchat-team/zim/app/demo/internal/service/contact"
	sGroup "github.com/zchat-team/zim/app/demo/internal/service/group"
	sPassport "github.com/zchat-team/zim/app/demo/internal/service/passport"
	sUser "github.com/zchat-team/zim/app/demo/internal/service/user"
	"github.com/zchat-team/zim/proto/http/demo/contact"
	"github.com/zchat-team/zim/proto/http/demo/group"
	"github.com/zchat-team/zim/proto/http/demo/passport"
	"github.com/zchat-team/zim/proto/http/demo/user"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

func parseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

func RegisterAPI(r *gin.Engine) {
	Swagger(r)

	skipPath := []string{
		"/api/v1/passport/login",
		"/api/v1/passport/smsLogin",
		"/api/v1/passport/sms",
		"/api/v1/passport/refreshToken",
		"/api/v1/test",
	}

	privKey := config.GetString("auth.privKey")
	priv, err := base64.StdEncoding.DecodeString(privKey)
	_, err = parseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		log.Fatal(err)
	}

	g := r.Group("/api/v1",
		middleware.CheckLogin(skipPath...),
		//middleware.CheckSign(middleware.PrivKey(privKey)),
	)

	passport.RegisterPassportHTTPServer(g, sPassport.GetService())
	contact.RegisterContactHTTPServer(g, sContact.GetService())
	group.RegisterGroupHTTPServer(g, sGroup.GetService())
	user.RegisterUserHTTPServer(g, sUser.GetService())
}
