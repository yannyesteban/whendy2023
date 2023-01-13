package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"crypto/hmac"
	"crypto/sha256"
)

var (
	sep    = []byte(".")
	pad    = []byte("=")
	padStr = string(pad)
)

type JWT struct {
	Payload interface{}
	Key     string
}

func (jwt *JWT) VerifyHeader(req *http.Request) ([]byte, error) {

	auth := req.Header.Get("Authorization")

	if auth != "" {

		value := strings.Split(auth, " ")
		token := value[len(value)-1]

		return jwt.Verify(token)
	}

	return nil, fmt.Errorf("token not found")
}

func (jwt *JWT) Generate(payload interface{}) string {
	header := Header{"HS256", "JWT"}
	jwt.Payload = payload

	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)

	bytH := Base64Encode([]byte(h))
	bytP := Base64Encode([]byte(p))

	message := append(bytH, sep...)
	message = append(message, bytP...)

	signature := jwt.hash(message)

	message = append(message, sep...)
	message = append(message, signature...)

	token := string(message)

	return token

}

func (jwt *JWT) Verify(t string) ([]byte, error) {

	tk := []byte(t)

	token := bytes.Split(tk, sep)

	//header, _ := Base64Decode(token[0])
	payload, _ := Base64Decode(token[1])
	signature := token[2]

	expectedSignature := jwt.hash(bytes.Join([][]byte{token[0], token[1]}, sep))

	if !hmac.Equal(signature, expectedSignature) {
		return nil, fmt.Errorf("invalid token %s", t)
	}

	return payload, nil
}

func (jwt *JWT) hash(message []byte) []byte {

	key := []byte(jwt.Key)
	ha := hmac.New(sha256.New, key)
	ha.Write(message)
	return Base64Encode(ha.Sum(nil))
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func Base64Encode(src []byte) []byte {
	buf := make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(buf, src)

	return bytes.TrimRight(buf, padStr) // JWT: no trailing '='.
}

// Base64Decode decodes "src" to jwt base64 url format.
// We could use the base64.RawURLEncoding but the below is a bit faster.
func Base64Decode(src []byte) ([]byte, error) {
	if n := len(src) % 4; n > 0 {
		// JWT: Because of no trailing '=' let's suffix it
		// with the correct number of those '=' before decoding.
		src = append(src, bytes.Repeat(pad, 4-n)...)
	}

	buf := make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(buf, src)
	return buf[:n], err
}
