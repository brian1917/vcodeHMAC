package vcodeHMAC

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
)

const defaultAuthScheme = "VERACODE-HMAC-SHA-256"

func generateHeader(host, path, method, apiKeyID, apiKeySecret, authScheme string) (string, error) {
	signingData := formatSigningData(apiKeyID, host, path, method)
	timestamp := getCurrentTimestamp()
	nonce := generateNonce()
	authScheme = defaultAuthScheme
	signature, err := createSignature(authScheme, apiKeySecret, signingData, timestamp, nonce)
	if err != nil {
		return "", err
	}
	return formatHeader(authScheme, apiKeyID, timestamp, nonce, signature), nil
}

func createSignature(authScheme string, apiKeySecret string, signingData string, timestamp int64, nonce string) (string, error) {
	if authScheme == defaultAuthScheme {
		signature := hmacSig(apiKeySecret, signingData, timestamp, nonce)
		return signature, nil
	}
	return "", errors.New("unsupported auth scheme")
}

func hmacSig(apiKeySecret string, signingData string, timestamp int64, nonce string) string {

	timeString := strconv.FormatInt(timestamp, 10)
	apiKeySecDecoded, _ := hex.DecodeString(apiKeySecret)
	nonceDecoded, _ := hex.DecodeString(nonce)

	h := hmac.New(sha256.New, apiKeySecDecoded)
	h.Write(nonceDecoded)
	keyNonce := h.Sum(nil)

	h = hmac.New(sha256.New, keyNonce)
	h.Write([]byte(timeString))
	keyDate := h.Sum(nil)

	h = hmac.New(sha256.New, keyDate)
	h.Write([]byte("vcode_request_version_1"))
	signatureKey := h.Sum(nil)

	h = hmac.New(sha256.New, signatureKey)
	h.Write([]byte(signingData))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}
