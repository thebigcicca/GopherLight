package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

// CORSOptions holds the configuration settings for Cross-Origin Resource Sharing (CORS).
// These options determine which origins, headers, and methods are allowed when accessing
// the API from a different domain. They can also control caching and credentials for CORS requests.

type CORSOptions struct {
	// AllowOrigin specifies the allowed origin for cross-origin requests.
	// Use "*" to allow any origin or specify a specific origin like "http://example.com".
	AllowOrigin string

	// AllowMethods lists the HTTP methods that are allowed for cross-origin requests.
	// Common methods include "GET", "POST", "PUT", etc.
	AllowMethods []string

	// AllowHeaders lists the headers that can be used during a request.
	// This includes headers such as "Content-Type", "Authorization", etc.
	AllowHeaders []string

	// AllowCredentials indicates whether credentials (cookies, HTTP authentication, etc.)
	// are allowed in cross-origin requests. Set to true to allow credentials.
	AllowCredentials bool

	// ExposeHeaders specifies the headers that can be exposed to the browser.
	// This allows the client to read these headers from the response.
	ExposeHeaders []string

	// MaxAge defines how long (in seconds) the results of a preflight request can be cached.
	// A longer MaxAge reduces the number of preflight requests. Set to 0 to disable caching.
	MaxAge int
}

var DefaultCORSOptions = CORSOptions{
	AllowOrigin:      "*",
	AllowHeaders:     []string{"Content-Type", "Authorization"},
	ExposeHeaders:    []string{},
	AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPatch, http.MethodPost},
	AllowCredentials: false,
	MaxAge:           600,
}

// CORSMiddleware creates a middleware function that applies the given CORS options.
func CORSMiddleware(opts CORSOptions) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			originHeader := r.Header.Get("Origin")

			if (originHeader != opts.AllowOrigin) && opts.AllowOrigin != "*" {
				http.Error(w, "CORS Origin not allowed", http.StatusForbidden)
				return
			}

			if opts.AllowOrigin == "*" && opts.AllowCredentials {
				http.Error(w, "Could not set Access-Control-Allow-Credentials to true when Access-Control-Allow-Origin is *", http.StatusForbidden)
				return
			}

			if opts.AllowOrigin != "" {
				w.Header().Add("Access-Control-Allow-Origin", opts.AllowOrigin)
			}
			if len(opts.AllowMethods) != 0 {
				w.Header().Add("Access-Control-Allow-Methods", strings.Join(opts.AllowMethods, ","))
			}
			if opts.AllowCredentials {
				w.Header().Add("Access-Control-Allow-Credentials", "true")
			}

			if len(opts.AllowHeaders) != 0 {
				w.Header().Add("Access-Control-Allow-Headers", strings.Join(opts.AllowHeaders, ","))
			}
			if opts.MaxAge != 0 {
				w.Header().Add("Access-Control-Max-Age", strconv.Itoa(opts.MaxAge))
			}
			if len(opts.ExposeHeaders) != 0 {
				w.Header().Add("Access-Control-Expose-Headers", strings.Join(opts.ExposeHeaders, ","))
			}

			next(w, r)
		}
	}
}
