package shopeepay

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func Stringify(m interface{}) string {
	json, _ := json.Marshal(m)
	return string(json)
}

func Sha1(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	bs := h.Sum(nil)

	result := fmt.Sprintf("%x", bs)

	return result
}

func CreateSignatureHeaders(req []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	// Cannot return error
	if _, err := mac.Write(req); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
