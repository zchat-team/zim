package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// pair lookup token pair
type pair struct {
	key   string
	value string
	// headerName is a string in the header.
	// Possible value is "Bearer"
	headerName string
}

type Lookup struct {
	// lookup pair slice from lookup parse
	lookups []pair
}

// NewLookup new lookup
// lookup is a string in the form of "<source>:<name>[:<headerName>]" that is used
// to extract token from the request.
// use like "header:<name>[:<headerName>],query:<name>,cookie:<name>,param:<name>"
// Optional, Default value "header:Authorization:Bearer".
// Possible values:
// - "header:<name>:<headerName>", <headerName> is a special string in the header, Possible value is "Bearer"
// - "query:<name>"
// - "cookie:<name>"
// - "param:<name>"
func NewLookup(lookup string) *Lookup {
	if lookup == "" {
		lookup = "header:Authorization:Bearer"
	}
	methods := strings.Split(lookup, ",")
	lookups := make([]pair, 0, len(methods))
	for _, method := range methods {
		parts := strings.Split(strings.TrimSpace(method), ":")
		headerName := ""
		if !(len(parts) == 2 || len(parts) == 3) {
			continue
		}
		if len(parts) == 3 && parts[0] == "header" {
			headerName = strings.TrimSpace(parts[2])
		}
		lookups = append(lookups, pair{
			strings.TrimSpace(parts[0]),
			strings.TrimSpace(parts[1]),
			headerName,
		})
	}
	return &Lookup{lookups}
}

// GetToken 获取token, 从Request中获取,由 Lookup 定义
func (sf *Lookup) GetToken(c *gin.Context) (string, error) {
	var token string
	var err error

	for _, lookup := range sf.lookups {
		if len(token) > 0 {
			break
		}
		switch lookup.key {
		case "header":
			token, err = FromHeader(c, lookup.value, lookup.headerName)
		case "query":
			token, err = FromQuery(c, lookup.value)
		case "cookie":
			token, err = FromCookie(c, lookup.value)
		case "param":
			token, err = FromParam(c, lookup.value)
		}
	}
	if err != nil {
		return "", err
	}
	if len(token) == 0 {
		return "", ErrMissingToken
	}
	return token, nil
}

// FromHeader get token from header
// key is header key, like "Authorization"
// headerName is a string in the header, like "Bearer", if it is empty, it will return value.
func FromHeader(c *gin.Context, key, headerName string) (string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", ErrMissingToken
	}
	if headerName == "" {
		return strings.TrimSpace(authHeader), nil
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == headerName) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimSpace(parts[1]), nil
}

// FromQuery get token from query
// key is query key
func FromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)
	if token == "" {
		return "", ErrMissingToken
	}
	return strings.TrimSpace(token), nil
}

// FromCookie get token from Cookie
// key is cookie key
func FromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)
	if cookie == "" {
		return "", ErrMissingToken
	}
	return strings.TrimSpace(cookie), nil
}

// FromParam get token from param
// key is param key
func FromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)
	if token == "" {
		return "", ErrMissingToken
	}
	return strings.TrimSpace(token), nil
}
