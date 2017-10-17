// Package vcodeHMAC creates an authorization header for accessing Veracode APIs using an API ID and Key.
package vcodeHMAC

// GenerateAuthHeader takes the location of your credentials file, the HTTP Method, and URL and returns the header value to be used for Authorization
func GenerateAuthHeader(credsFile, httpMethod, url string) string {
	credentials := getCredentials(credsFile)
	apiKeyID := credentials[0]
	apiKeySecret := credentials[1]

	host := getHost(url)
	path := getPathParams(url)

	headerValue := generateHeader(host, path, httpMethod, apiKeyID, apiKeySecret, defaultAuthScheme)
	return headerValue
}
