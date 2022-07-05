package signature

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

var aesKeySizes = []int{16, 24, 32}

func TestEncryptDecrypt(t *testing.T) {
	plainText := "helloworld,this is golang language. welcome"
	for _, keySize := range aesKeySizes {
		key := make([]byte, keySize)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		cipherText, err := Encrypt(string(key), plainText)
		require.NoError(t, err)

		got, err := Decrypt(string(key), cipherText)
		require.NoError(t, err)
		require.Equal(t, []byte(plainText), got)
	}
}

var pri = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA2G4SFmN+v/KO1ppBGSW4FJ4tXRzSUjTCjIQlUCAUE1xBT6+V
ZmMt28vxromutTune5d054tcH2Vq9QMebWwmcHBCthuVaOQ1Ij/CgKsD/UUhEBLA
PBocFOtPadTR4m95gQYioOGPQZb90VkGgO5y7OGzu5KG/wZsi8QQSNmHeTnr4uZq
CAFjlbSgcMFqfqgh1E4dAIglQl54MCOXiI9xnGTbKddaBl3IlpVlqfmYwmrWJbku
8r/WlEBtjTteOMpwDKKxSX3rGtvg9Jgok0EnG0LEwsCfmuHfNFfplzzqj8X1XaZA
f8hqu2XpgIqQjPHWzBG9PIYd3DxYRSGlUhuvOQIDAQABAoIBAQC1kQHTjnyTAyYZ
mybptd8MTPa5mqhHFsPvphy3b3HoHAkelKq9To72Sc3jItZSbE1BPfpxFVSfcjGc
gpVQLt7AjS0qIVHiwTBiHyNJVi7ulsP5/AERasYMNqxUmJnLYMGKIF+EoDXSTJ16
tzjhiSkY2PAzd+WQpQ8C4eTXeMZSR15LXr+K2hy7ysbAq6pHNuUxFZOEEGDsOb0T
TreA4xyIgcej/Vv2m/buR4tr652xLreloJ5oUpMXDAfVmgZqHIr8fwDOu3v8vDfG
J/mT++Y+BpqwSQXDu9i40juT4gsAgF7U3RxzAMe1HV1DG7zu+YZaK4iUA7cWnGL/
2tkOp1pBAoGBAO2T/5JBk1+G/TDJGPlG6EuIHI/gFBYOfwtz2cS992Bj6MOv142x
b2y9c5dp0Kk2iDOa+a/Ds43GfCFw9lXiN8JHvSx0Q6qDXAXQJy4xmJwB6ntl9hKL
stbG7lCN3JNj4maxytiDNxAWk/+7t98lnA0OqyuLWChH/zEnHOBdIYZ1AoGBAOk2
RlHHnTHdsqzI2uIu7y4KNgl/H964FDOp7NWkkXxteweGfe00d9dNc7XV9wNqYIkn
kaRa1KeF0ptCRkfGLbUyFdz6DDemmnM5ASLtnjb60ItpZtLOZgyFkuy6MbCwk3RZ
N+kFDhdmH5Y3EFdlLooBtU14k7FmoPmdiPvRrlU1AoGAY2diLrLLU9PqSihKD7rQ
ZRINSVGrddMY6xTNEBmf0K/c60u+t+V+xpO6MqcujC5p7JWyVQ1gKjjbJS7bkvG0
/NABYgE/cq/FqBUA374WqWfP0VPHEtlquZzAh+njWbQYPXm0csTsHAomYIENnQti
cMArdGu4NhpxtwIzfdjZtyUCgYEAyBqa4cbeeZAZpKo/Lb5J2f5G+YULqoXWR7Ix
Feu8LcCexQlAec0AW0wI0ehCp7qaFHVQQW7ycr+fwzptpV5Fj+jm25Ht875PXjh2
YirzC4fQcx7AbHdPFsVyGQ92XX5VN4rqL1X4DlnBFpouul6GPUJT96JTT++YhjYG
+NOku1UCgYAS4UE6A4WKsTDUtHrP8/+fk0OFdQrbR5IVYvIjywkGJEcVOUv3FlOu
vrCHHg0CkJHobWcaCMCCBMF21aKZZlNMJjtTlqjvlFyP4WVG+BCFWI213iYYgV4T
YbZQlr1h33HSOIQ6qihuOSRC0+0Vo108ZwKOFIJ8p+5rZkVpo6bgQA==
-----END RSA PRIVATE KEY-----
`

var pub = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2G4SFmN+v/KO1ppBGSW4
FJ4tXRzSUjTCjIQlUCAUE1xBT6+VZmMt28vxromutTune5d054tcH2Vq9QMebWwm
cHBCthuVaOQ1Ij/CgKsD/UUhEBLAPBocFOtPadTR4m95gQYioOGPQZb90VkGgO5y
7OGzu5KG/wZsi8QQSNmHeTnr4uZqCAFjlbSgcMFqfqgh1E4dAIglQl54MCOXiI9x
nGTbKddaBl3IlpVlqfmYwmrWJbku8r/WlEBtjTteOMpwDKKxSX3rGtvg9Jgok0En
G0LEwsCfmuHfNFfplzzqj8X1XaZAf8hqu2XpgIqQjPHWzBG9PIYd3DxYRSGlUhuv
OQIDAQAB
-----END PUBLIC KEY-----
`

func TestRsaEncryptDecrypt(t *testing.T) {
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pri))
	require.NoError(t, err)
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pub))
	require.NoError(t, err)

	want := "helloworld,this is golang language. welcome"

	mid, err := RsaEncrypt(pubKey, want)
	require.NoError(t, err)

	got, err := RsaDecrypt(priKey, mid)
	require.NoError(t, err)

	require.Equal(t, want, got)
}

func TestPCKS(t *testing.T) {
	want := []byte("1234567890abcdef")

	b := PCKSPadding(want, 16)

	got, err := PCKSUnPadding(b, 16)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
