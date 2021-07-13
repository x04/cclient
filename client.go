package cclient

import (
	http "github.com/useflyent/fhttp"

	"golang.org/x/net/proxy"

	"time"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, redirectOption string, timeout int32, proxyUrl ...string) (http.Client, error) {

	if redirectOption == "False" {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			dialer, err := newConnectDialer(proxyUrl[0])
			if err != nil {
				return http.Client{
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					},
					Timeout: time.Duration(timeout) * time.Second,
				}, err
			}
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: newRoundTripper(clientHello, dialer),
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		} else {
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: newRoundTripper(clientHello, proxy.Direct),
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		}
	} else {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			dialer, err := newConnectDialer(proxyUrl[0])
			if err != nil {
				return http.Client{
					Timeout: time.Duration(timeout) * time.Second,
				}, err
			}
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer),
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		} else {
			return http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct),
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		}
	}
}
