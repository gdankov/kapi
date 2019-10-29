module github.com/cloudfoundry-community/kapi

go 1.13

require (
	code.cloudfoundry.org/eirini v0.0.0-20191028145513-876c52b5ef67
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.8.1
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/apimachinery v0.0.0-20191025225532-af6325b3a843
	k8s.io/client-go v0.0.0-20191026065934-0bdba2f91880
	k8s.io/code-generator v0.0.0-20191017183038-0b22993d207c
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20191026065934-0bdba2f91880
