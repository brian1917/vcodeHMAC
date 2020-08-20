// Package vcodeHMAC creates an authorization header for accessing Veracode APIs using an API ID and Key.
package vcodeHMAC

// GenerateAuthHeader takes the location of your credentials file, the HTTP Method, and URL and returns the header value to be used for Authorization
func GenerateAuthHeader(credsFile, httpMethod, url string) (string, error) {
	credentials, err := getCredentials(credsFile)
	if err != nil {
		return "", err
	}

	host, err := getHost(url)
	if err != nil {
		return "", err
	}
	params, err := getPathParams(url)
	if err != nil {
		return "", err
	}

	headerValue, err := GenerateHeader(host, params, httpMethod, credentials[0], credentials[1], defaultAuthScheme)
	if err != nil {
		return "", err
	}
	return headerValue, nil
}
