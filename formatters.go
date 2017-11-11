package vcodeHMAC

import (
	"fmt"
	"strings"
)

func formatSigningData(apiKeyID string, host string, url string, method string) string {
	apiKeyIDLower := strings.ToLower(apiKeyID)
	hostName := strings.ToLower(host)
	method = strings.ToUpper(method)

	return fmt.Sprintf("id=%s&host=%s&url=%s&method=%s", apiKeyIDLower, hostName, url, method)
}

func formatHeader(authScheme string, apiKeyID string, timestamp int64, nonce string, signature string) string {
	return fmt.Sprintf("%s id=%s,ts=%d,nonce=%s,sig=%s", authScheme, apiKeyID, timestamp, nonce, signature)
}
