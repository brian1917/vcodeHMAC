package vcodeHMAC

import (
	"fmt"
	"strings"
)

func formatSigningData(apiKeyID string, host string, url string, method string) string {
	apiKeyIDLower := strings.ToLower(apiKeyID)
	hostName := strings.ToLower(host)
	method = strings.ToUpper(method)

	return fmt.Sprintf("id=%v&host=%v&url=%v&method=%v", apiKeyIDLower, hostName, url, method)
}

func formatHeader(authScheme string, apiKeyID string, timestamp int64, nonce string, signature string) string {
	return fmt.Sprintf("%v id=%v,ts=%v,nonce=%v,sig=%v", authScheme, apiKeyID, timestamp, nonce, signature)
}
