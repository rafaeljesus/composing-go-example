package httpclient

var (
	contentType = "application/json"
)

type (
	// HTTPClient is the http wrapper for the application
	HTTPClient struct {
		*HTTPPoster
		*HTTPGetter
	}
)

// New returns a configured HTTPClient
func New(p *HTTPPoster, g *HTTPGetter) *HTTPClient {
	return &HTTPClient{
		HTTPPoster: p,
		HTTPGetter: g,
	}
}
