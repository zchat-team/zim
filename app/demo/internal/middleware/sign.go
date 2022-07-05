package middleware

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zmicro-team/zmicro/core/errors"
	"github.com/zmicro-team/zmicro/core/log"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"

	"github.com/zchat-team/zim/pkg/signature"
)

// 客户端签名加密过程
// 随机生成一个randomKey
// 如果body不为空并且body不为空需要加密 body=Base64(AesCBCEncrypt(randomKey,body))
// 拼接str = timestamp+method+url+body
// sign = Base64(HMAC(randomKey,str))
// secret = Base64(RsaEncrypt(randomKey, pubkey))

// 服务端验签解密过程则是上述的逆过程

func CheckSign(opts ...Option) gin.HandlerFunc {
	options := Options{
		skip: func(c *gin.Context) bool { return false },
	}
	for _, o := range opts {
		o(&options)
	}
	return func(c *gin.Context) {
		if options.skip(c) {
			c.Next()
			return
		}
		method := strings.ToUpper(c.Request.Method)
		target := c.Request.RequestURI
		timestamp := c.GetHeader("Timestamp")
		secret := c.GetHeader("Secret")
		if timestamp == "" || secret == "" {
			zgin.Error(c, errors.ErrBadRequest("timestamp, secret 不能为空"))
			return
		}

		// 根据secret解出randomKey
		randomKey, err := signature.RsaDecrypt(options.privKey, secret)
		if err != nil {
			zgin.Error(c, errors.ErrBadRequest(err.Error()))
			return
		}

		// body处理
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			zgin.Error(c, errors.ErrBadRequest(err.Error()))
			return
		}
		// 是否加密body
		var origBody []byte
		cipherBody := string(body)
		encrypt := c.GetHeader("Encrypt")
		if encrypt == "1" && len(body) > 0 {
			origBody, err = signature.Decrypt(randomKey, cipherBody)
		} else {
			origBody, err = base64.StdEncoding.DecodeString(cipherBody)
		}
		if err != nil {
			zgin.Error(c, errors.ErrBadRequest(err.Error()))
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(origBody))

		str := timestamp + method + target + cipherBody
		calcSign := signature.Signature(randomKey, str)
		sign := c.GetHeader("Sign")

		log.Debugf("客户端上传的sign:%s", sign)
		log.Debugf("计算得到的sign:%s", calcSign)
		if calcSign != sign {
			zgin.Error(c, errors.ErrBadRequest("签名错误"))
			return
		}
		c.Next()
	}
}

type Option func(*Options)

type Options struct {
	privKey *rsa.PrivateKey
	skip    func(*gin.Context) bool
}

func PrivKey(privKey *rsa.PrivateKey) Option {
	return func(o *Options) {
		o.privKey = privKey
	}
}

func MustPrivKeyFromFile(privKeyFile string) Option {
	keyData, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		panic(err)
	}
	return PrivKey(key)
}

func WithSkip(skip func(c *gin.Context) bool) Option {
	return func(o *Options) {
		if skip != nil {
			o.skip = skip
		}
	}
}
