package vcodeHMAC

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"
	"strconv"
	"encoding/hex"
)

const DEFAULT_AUTH_SCHEME = "VERACODE-HMAC-SHA-256"

func generateHeader(host, path, method, api_key_id, api_key_secret, auth_scheme string) string {
	signing_data := formatSigningData(api_key_id, host, path, method)
	timestamp := getCurrentTimestamp()
	nonce := generateNonce()
	auth_scheme = DEFAULT_AUTH_SCHEME
	signature := createSignature(auth_scheme, api_key_secret, signing_data, timestamp, nonce)
	return formatHeader(auth_scheme, api_key_id, timestamp, nonce, signature)
}

func createSignature(auth_scheme string, api_key_secret string, signing_data string, timestamp int64, nonce string) string {
	if auth_scheme == DEFAULT_AUTH_SCHEME{
		signature := hmacSig(api_key_secret, signing_data, timestamp, nonce)
		return signature
	} else {
	log.Fatal("Unsupported auth scheme")
		return "error"
	}
}

func hmacSig(api_key_secret string, signing_data string, timestamp int64, nonce string) string {

	time_string := strconv.Itoa(int(timestamp))
	api_key_sec_dec, _ := hex.DecodeString(api_key_secret)
	nonce_dec, _ := hex.DecodeString(nonce)

	h := hmac.New(sha256.New, api_key_sec_dec)
	h.Write(nonce_dec)
	key_nonce := h.Sum(nil)

	h = hmac.New(sha256.New, key_nonce)
	h.Write([]byte(time_string))
	key_date := h.Sum(nil)

	h = hmac.New(sha256.New, key_date)
	h.Write([]byte("vcode_request_version_1"))
	signature_key := h.Sum(nil)

	h = hmac.New(sha256.New, signature_key)
	h.Write([]byte(signing_data))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}