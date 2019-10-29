package bridge

import (
	"encoding/json"
	"net/http"

	v1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi.cf.org/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/julienschmidt/httprouter"
)

func NewStagingBridge(clientset clientset.Interface, namespace string) *Staging {
	return &Staging{
		Creator: CRDCreator{
			Clientset: clientset,
			Namespace: namespace,
		},
	}
}

type Staging struct {
	Creator CRDCreator
}

func (s *Staging) Stage(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var stagingRequest v1alpha1.StagingSpec
	if err := json.NewDecoder(req.Body).Decode(&stagingRequest); err != nil {
		writeErrorResponse(resp, http.StatusBadRequest, err)
		return
	}

	if err := s.Creator.CreateStaging(stagingRequest); err != nil {
		writeErrorResponse(resp, http.StatusInternalServerError, err)
		return
	}

	resp.WriteHeader(http.StatusAccepted)
}

func writeErrorResponse(resp http.ResponseWriter, status int, err error) {
	resp.WriteHeader(status)
	encodingErr := json.NewEncoder(resp).Encode(&StagingError{Message: err.Error()})
	if encodingErr != nil {
		panic(encodingErr)
	}
}

type StagingError struct {
	Message string `json:"message"`
}
