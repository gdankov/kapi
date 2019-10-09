package apidge

import (
	"net/http"

	"github.com/cloudfoundry-community/kapi/apidge/app"
	"github.com/cloudfoundry-community/kapi/apidge/stage"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/julienschmidt/httprouter"
)

func NewHandler(clientset clientset.Interface, namespace string) http.Handler {
	handler := httprouter.New()

	appHandler := app.NewHandler(clientset, namespace)
	stageHandler := stage.NewHandler(clientset, namespace)

	handler.PUT("/apps/:process_guid", appHandler.Desire)
	handler.POST("/stage/:staging_guid", stageHandler.Stage)

	return handler
}
