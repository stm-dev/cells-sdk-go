// Package cells_sdk provides a ready to use SDK to use the Cells REST API in Go language.
// It also provides some patterns that ease implementation of Go programs that use the SDK.
package cells_sdk

// SdkConfig stores parameters to talk to a running Cells instance REST API via the Go SDK.
type SdkConfig struct {
	// Url stores domain name or IP & port to the server.
	Url string `json:"url"`
	// OIDC ClientKey / ClientSecret
	ClientKey    string `json:"clientKey"`
	ClientSecret string `json:"clientSecret"`
	// Pydio User Authentication
	User     string `json:"user"`
	Password string `json:"password"`

	// OIDC Code Flow
	IdToken        string `json:"idToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenExpiresAt int    `json:"tokenExpiresAt"`

	// SkipVerify tells the transport to ignore expired or self-signed TLS certificates
	SkipVerify bool `json:"skipVerify"`

	// UseTokenCache flags wether we should rely on a local cache to avoid retrieving a new JWT token at each request.
	// It is useful to *not* use the cache when running connection tests for instance.
	UseTokenCache bool `json:"useTokenCache"`

	// Optional list of headers to override in requests, typically User-Agent
	CustomHeaders map[string]string
}

// S3Config stores connection parameters to a running Cells instance S3 gateway via the AWS SDK for Go.
type S3Config struct {
	// Bucket name, usually io.
	Bucket string `json:"bucket"`
	// Endpoint overrides the default URL generated by the S3 SDK from the bucket name.
	Endpoint string `json:"enpoint"`
	// Region param, usually us-east-1
	Region string `json:"region"`
	// ApiKey is used by the Cells SDK to transmit the JWT token.
	ApiKey string `json:"apiKey"`
	// ApiSecret identifies this client.
	ApiSecret string `json:"apiSecret"`
	// Set to true to rather use legacy mode (JWT Auth is transmitted via custom 'X-Pydio-Bearer' header).
	UsePydioSpecificHeader bool `json:"usePydioSpecificHeader"`
	// IsDebug is a convenience legacy flag to add some logging to this S3 client.
	// Should be cleaned as soon as we defined the logging strategy for this repo.
	IsDebug bool `json:"isDebug"`
}

var (
	// DefaultConfig stores a convenience static object that must be configured only once
	// and is globally accessible to easily retrieve an up-and-running connected client.
	DefaultConfig *SdkConfig
	// DefaultS3Config stores a convenience static object that must be configured only once
	// and is globally accessible to easily retrieve an up-and-running connected client.
	DefaultS3Config *S3Config
)
