package bridge

import (
	"bytes"
	"encoding/json"
	"net/http"

	v1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi.cf.org/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/julienschmidt/httprouter"
)

func NewLRPBridge(clientset clientset.Interface, namespace string) *LRP {
	return &LRP{
		Creator: CRDCreator{
			Clientset: clientset,
			Namespace: namespace,
		},
	}
}

type LRP struct {
	Creator CRDCreator
}

func (a *LRP) Desire(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request v1alpha1.LRPSpec
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := a.Creator.CreateLRP(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
