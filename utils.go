package vcodeHMAC

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"time"
)

func getCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func generateNonce() string {
	token := make([]byte, 16)
	rand.Read(token)
	return hex.EncodeToString(token)
}

func getHost(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	host := u.Host

	return host
}

func getPathParams(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	path := u.Path
	params := u.RawQuery

	if len(params) > 0 {
		return fmt.Sprintf("%v?%v", path, params)
	}
	return fmt.Sprintf("%v", path)
}

func getScheme(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	return u.Scheme
}
