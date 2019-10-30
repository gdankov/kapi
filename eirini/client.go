package eirini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	kapiv1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi.cf.org/v1alpha1"
	"github.com/pkg/errors"
)

type Client struct {
	HTTPClient *http.Client
	EiriniURL  string
}

func (c Client) Stage(spec kapiv1alpha1.StagingSpec) error {
	b, err := json.Marshal(spec)
	if err != nil {
		return errors.Wrap(err, "failed to marshal staging spec")
	}

	url := fmt.Sprintf("%s/stage/%s", c.EiriniURL, spec.AppGUID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "failed to create the post request")
	}

	_, err = c.HTTPClient.Do(req)
	return errors.Wrap(err, "failed to execute stage request")
}

func (c Client) Desire(spec kapiv1alpha1.LRPSpec) error {
	b, err := json.Marshal(spec)
	if err != nil {
		return errors.Wrap(err, "failed to marshal spec")
	}

	url := fmt.Sprintf("%s/apps/%s", c.EiriniURL, spec.ProcessGUID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "failed to create the post request")
	}

	_, err = c.HTTPClient.Do(req)
	return errors.Wrap(err, "failed to execute desire request")
}
