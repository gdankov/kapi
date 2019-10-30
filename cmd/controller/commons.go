package controller

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfoundry-community/kapi/eirini"
	"github.com/pkg/errors"
)

const (
	EnvVarEiriniCrtPath = "EIRINI_CRT_PATH"
	EnvVarEiriniKeyPath = "EIRINI_KEY_PATH"
	EnvVarEiriniCaPath  = "EIRINI_CA_PATH"
	EnvVarEiriniURL     = "EIRINI_URL"
)

func CreateEiriniClient() (eirini.Client, error) {
	crtPath := os.Getenv(EnvVarEiriniCrtPath)
	keyPath := os.Getenv(EnvVarEiriniKeyPath)
	caPath := os.Getenv(EnvVarEiriniCaPath)
	if crtPath == "" || keyPath == "" || caPath == "" {
		return eirini.Client{}, errors.New("eirini certificates missing")
	}

	httpclient, err := createTLSHTTPClient(
		[]CertPaths{
			{
				Crt: crtPath,
				Key: keyPath,
				Ca:  caPath,
			},
		},
	)
	if err != nil {
		return eirini.Client{}, errors.Wrap(err, "failed to create https client")
	}
	eiriniURL := os.Getenv(EnvVarEiriniURL)
	if eiriniURL == "" {
		return eirini.Client{}, errors.New("eirini url not provided")
	}

	return eirini.Client{
		HTTPClient: httpclient,
		EiriniURL:  eiriniURL,
	}, nil
}

type CertPaths struct {
	Crt, Key, Ca string
}

func createTLSHTTPClient(certPaths []CertPaths) (*http.Client, error) {
	pool := x509.NewCertPool()
	certs := []tls.Certificate{}
	for _, c := range certPaths {
		cert, err := tls.LoadX509KeyPair(c.Crt, c.Key)
		if err != nil {
			return nil, errors.Wrap(err, "could not load cert")
		}
		certs = append(certs, cert)

		cacert, err := ioutil.ReadFile(filepath.Clean(c.Ca))
		if err != nil {
			return nil, err
		}
		if ok := pool.AppendCertsFromPEM(cacert); !ok {
			return nil, errors.New("failed to append cert to cert pool")
		}
	}

	tlsConf := &tls.Config{
		Certificates: certs,
		RootCAs:      pool,
	}

	return &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConf}}, nil
}
