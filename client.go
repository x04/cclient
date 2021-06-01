package cclient

import (
	"net/http"

	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, redirectOption string, proxyUrl ...string) (http.Client, error) {

	if redirectOption == "False" {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			dialer, err := newConnectDialer(proxyUrl[0])
			if err != nil {
				return http.Client{
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					},
				}, err
			}
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: newRoundTripper(clientHello, dialer),
			}, nil
		} else {
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: newRoundTripper(clientHello, proxy.Direct),
			}, nil
		}
	} else {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			dialer, err := newConnectDialer(proxyUrl[0])
			if err != nil {
				return http.Client{}, err
			}
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer),
			}, nil
		} else {
			return http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct),
			}, nil
		}
	}
}
