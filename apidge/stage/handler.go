package stage

import (
	"encoding/json"
	"net/http"

	cfmodels "code.cloudfoundry.org/eirini/models/cf"
	v1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/julienschmidt/httprouter"
)

func NewHandler(clientset clientset.Interface, namespace string) *StageBridge {
	return &StageBridge{
		Creator: CRDCreator{
			Clientset: clientset,
			Namespace: namespace,
		},
	}
}

type StageBridge struct {
	Creator CRDCreator
}

func (s *StageBridge) Stage(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var stagingRequest v1alpha1.StagingSpec
	if err := json.NewDecoder(req.Body).Decode(&stagingRequest); err != nil {
		writeErrorResponse(resp, http.StatusBadRequest, err)
		return
	}

	if err := s.Creator.Create(stagingRequest); err != nil {
		writeErrorResponse(resp, http.StatusInternalServerError, err)
		return
	}

	resp.WriteHeader(http.StatusAccepted)
}

func writeErrorResponse(resp http.ResponseWriter, status int, err error) {
	resp.WriteHeader(status)
	encodingErr := json.NewEncoder(resp).Encode(&cfmodels.StagingError{Message: err.Error()})
	if encodingErr != nil {
		panic(encodingErr)
	}
}
