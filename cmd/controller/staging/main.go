package main

import (
	"flag"
	"time"

	commons "github.com/cloudfoundry-community/kapi/cmd/controller"
	"github.com/cloudfoundry-community/kapi/controller/staging"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	informers "github.com/cloudfoundry-community/kapi/pkg/generated/informers/externalversions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/sample-controller/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	clientset, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	eiriniClient, err := commons.CreateEiriniClient()
	if err != nil {
		klog.Fatalf("Error creating Eirini client: %s", err.Error())
	}

	controller := staging.NewController(
		clientset,
		informerFactory.Kapi().V1alpha1().Stagings(),
		eiriniClient,
	)

	informerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
