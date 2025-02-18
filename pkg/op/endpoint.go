package op

import "strings"

type Endpoint struct {
	path string
	url  string
}

func NewEndpoint(path string) Endpoint {
	return Endpoint{path: path}
}

func NewEndpointWithURL(path, url string) Endpoint {
	return Endpoint{path: path, url: url}
}

func NewEndpointWithIssuersPath(issuer, path string) Endpoint {
	issuerPath := getIssuerPath(issuer)
	if len(issuerPath) > 0 {
		issuerPath = issuerPath + "/"
	}

	return Endpoint{path: issuerPath + path}
}

func (e Endpoint) Relative() string {
	return relativeEndpoint(e.path)
}

func (e Endpoint) Absolute(host string) string {
	if e.url != "" {
		return e.url
	}
	return absoluteEndpoint(host, e.path)
}

func (e Endpoint) Validate() error {
	return nil //TODO:
}

func absoluteEndpoint(host, endpoint string) string {
	return strings.TrimSuffix(host, "/") + relativeEndpoint(endpoint)
}

func relativeEndpoint(endpoint string) string {
	return "/" + strings.TrimPrefix(endpoint, "/")
}
