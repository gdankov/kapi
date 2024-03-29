module github.com/cloudfoundry-community/kapi

go 1.13

require (
	code.cloudfoundry.org/eirini v0.0.0-20191003133837-d1f227bf250a
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.8.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	k8s.io/api v0.0.0-20190703205437-39734b2a72fe
	k8s.io/apimachinery v0.0.0-20190703205208-4cfb76a8bf76
	k8s.io/client-go v0.0.0-20190704045512-07281898b0f0
	k8s.io/code-generator v0.0.0-20190703204957-583809a49343
	k8s.io/klog v0.3.3
	k8s.io/sample-controller v0.0.0-20190704050429-35c85454ecd6
)

replace k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190703204957-583809a49343
