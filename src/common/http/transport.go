// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	tracelib "github.com/goharbor/harbor/src/lib/trace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	// InsecureTransport used to get the insecure http Transport
	InsecureTransport = iota
	// SecureTransport used to get the external secure http Transport
	SecureTransport
)

var (
	secureHTTPTransport   http.RoundTripper
	insecureHTTPTransport http.RoundTripper
)

func init() {
	insecureHTTPTransport = NewTransport(WithInsecureSkipVerify(true))
	if InternalTLSEnabled() {
		secureHTTPTransport = NewTransport(WithInternalTLSConfig())
	} else {
		secureHTTPTransport = NewTransport()
	}
	if tracelib.Enabled() {
		insecureHTTPTransport = otelhttp.NewTransport(insecureHTTPTransport)
		secureHTTPTransport = otelhttp.NewTransport(secureHTTPTransport)
	}
}

// Use this instead of Default Transport in library because it sets ForceAttemptHTTP2 to true
// And that options introduced in go 1.13 will cause the https requests hang forever in replication environment
func newDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig:       &tls.Config{},
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func WithInternalTLSConfig() func(*http.Transport) {
	return func(tr *http.Transport) {
		tlsConfig, err := GetInternalTLSConfig()
		if err != nil {
			panic(err)
		}
		tr.TLSClientConfig = tlsConfig
	}
}
func WithInsecureSkipVerify(skipVerify bool) func(*http.Transport) {
	return func(tr *http.Transport) {
		tr.TLSClientConfig.InsecureSkipVerify = skipVerify
	}
}

func WithMaxIdleConns(maxIdleConns int) func(*http.Transport) {
	return func(tr *http.Transport) {
		tr.MaxIdleConns = maxIdleConns
	}
}

func WithIdleconnectionTimeout(idleConnectionTimeout time.Duration) func(*http.Transport) {
	return func(tr *http.Transport) {
		tr.IdleConnTimeout = idleConnectionTimeout
	}
}

func NewTransport(opts ...func(*http.Transport)) http.RoundTripper {
	tr := newDefaultTransport()
	for _, opt := range opts {
		opt(tr)
	}
	return tr
}

type TransportConfig struct {
	Insecure bool
}
type TransportOption func(*TransportConfig)

func WithInsecure(skipVerify bool) TransportOption {
	return func(cfg *TransportConfig) {
		cfg.Insecure = skipVerify
	}
}

// GetHTTPTransport returns HttpTransport based on insecure configuration
func GetHTTPTransport(opts ...TransportOption) http.RoundTripper {
	cfg := &TransportConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.Insecure {
		return insecureHTTPTransport
	}
	return secureHTTPTransport
}
