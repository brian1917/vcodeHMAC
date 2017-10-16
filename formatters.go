package vcodeHMAC

import (
	"strings"
	"fmt"
)

func formatSigningData(api_key_id string, host string, url string, method string) string  {
	api_key_id_valabc := strings.ToLower(api_key_id)
	host_abc := strings.ToLower(host)
	method = strings.ToUpper(method)

	return fmt.Sprintf("id=%v&host=%v&url=%v&method=%v", api_key_id_valabc, host_abc,url, method)
}

func formatHeader(auth_scheme string, api_key_id string, timestamp int64, nonce string, signature string) string{
	return fmt.Sprintf("%v id=%v,ts=%v,nonce=%v,sig=%v", auth_scheme, api_key_id, timestamp, nonce, signature)
}