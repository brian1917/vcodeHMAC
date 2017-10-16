// vcodeHMAC creates an authorization header for accessing Veracode APIs using an API ID and Key.
package vcodeHMAC

// GenerateAuthHeader takes the location of your credentials file, the HTTP Method, and URL and returns the header value to be used for Authorization
func GenerateAuthHeader(credsFile, http_method, url string) string {
	credentials := getCredentials(credsFile)
	api_key_id_val := credentials[0]
	api_key_secret_val := credentials[1]

	host := getHost(url)
	path := getPathParams(url)

	header_value := generateHeader(host, path, http_method, api_key_id_val, api_key_secret_val, DEFAULT_AUTH_SCHEME)
	return header_value
}