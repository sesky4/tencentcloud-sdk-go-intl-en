package profile

// HttpProfile encapsulates the configuration for HTTP requests.
// It defines settings such as the request method, timeout, scheme, and endpoints.
type HttpProfile struct {
	// ReqMethod specifies the HTTP request method (e.g., "GET", "POST").
	// Default: "POST".
	ReqMethod string

	// ReqTimeout defines the request timeout in seconds.
	// This determines how long the client will wait for a response before timing out.
	// Default: 60 seconds.
	ReqTimeout int

	// Deprecated: Use Scheme instead.
	Protocol string
	// Scheme specifies the protocol scheme to use for the request (e.g., "HTTP", "HTTPS").
	// Default: "HTTPS".
	Scheme string

	// RootDomain specifies the root domain for the API request.
	// This is often used as a base for constructing the full API endpoint.
	// Default: "".
	RootDomain string

	// Endpoint specifies the specific API endpoint to which the request is sent.
	// This can be a full URL or a domain name, depending on the service.
	// Default: "".
	Endpoint string

	ApigwEndpoint string

	// Proxy specifies the URL of a proxy server to use for the requests.
	// If set, all requests will be routed through this proxy.
	// Default: "".
	Proxy string

	// DomainFailover enables transparent retries against backup TLDs when the
	// primary domain fails for transport-level reasons (DNS miss, TLS handshake,
	// connection refused, timeout, non-2xx HTTP status, malformed JSON body).
	// When the request host belongs to the tencentcloudapi.{com,cn,com.cn}
	// family, the SDK will rotate within that family. Region-pinned hosts
	// (containing an ap-/na-/eu-/sa-/af-/me- label) opt out of TLD rotation
	// to avoid silently changing the resolved region.
	// Default: true.
	DomainFailover bool
}

// NewHttpProfile creates and initializes a new HttpProfile with default values.
// This function provides a convenient way to obtain a HttpProfile instance
// with sensible defaults, which can then be customized as needed.
func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:      "POST",
		ReqTimeout:     60,
		Scheme:         "HTTPS",
		RootDomain:     "",
		Endpoint:       "",
		ApigwEndpoint:  "",
		DomainFailover: true,
	}
}
