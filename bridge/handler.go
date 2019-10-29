package bridge

import (
	"net/http"

	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/julienschmidt/httprouter"
)

func NewHandler(clientset clientset.Interface, namespace string) http.Handler {
	handler := httprouter.New()

	stager := NewStagingBridge(clientset, namespace)
	lrpDesirer := NewLRPBridge(clientset, namespace)

	handler.POST("/stage/:staging_guid", stager.Stage)
	handler.PUT("/apps/:process_guid", lrpDesirer.Desire)

	return handler
}
