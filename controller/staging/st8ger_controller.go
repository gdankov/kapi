package staging

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	kapiv1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	samplescheme "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned/scheme"
	informers "github.com/cloudfoundry-community/kapi/pkg/generated/informers/externalversions/kapi/v1alpha1"
	listers "github.com/cloudfoundry-community/kapi/pkg/generated/listers/kapi/v1alpha1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

const controllerAgentName = "stager-controller"

type StagerController struct {
	kapiClientset clientset.Interface

	stagersLister listers.StagingLister
	stagersSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	kapiclientset clientset.Interface,
	stagerInformer informers.StagingInformer) *StagerController {

	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &StagerController{
		kapiClientset: kapiclientset,
		stagersLister: stagerInformer.Lister(),
		stagersSynced: stagerInformer.Informer().HasSynced,
		workqueue:     workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "CFStagers"),
		recorder:      recorder,
	}

	klog.Info("Setting up event handlers")
	stagerInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueFoo,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueFoo(new)
		},
	})
	return controller
}

func (c *StagerController) enqueueFoo(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

func (c *StagerController) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.stagersSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch two workers to process Foo resources
	go wait.Until(c.runWorker, time.Second, stopCh)

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *StagerController) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *StagerController) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *StagerController) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	staging, err := c.stagersLister.Stagings(namespace).Get(name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("foo '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}
	fmt.Println("Staging state: ", staging.Spec.State)
	if staging.Spec.State != kapiv1alpha1.NotStartedState {
		fmt.Println("Staging job already started")
		return nil
	}
	fmt.Println("Starting staging task")

	spec := staging.Spec
	b, err := json.Marshal(spec)
	if err != nil {
		return errors.Wrap(err, "failed to marshal spec")
	}

	fmt.Printf("STAGING SPEC LOOKS LIKE THIS: %+v", spec)
	url := fmt.Sprintf("https://eirini-opi.scf.svc.cluster.local:8085/stage/%s", spec.AppGUID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "failed to create the post request")
	}

	client, err := createTLSHTTPClient(
		[]CertPaths{
			{
				Crt: "/workspace/jobs/st8ger_controller/eirini.crt",
				Key: "/workspace/jobs/st8ger_controller/eirini.key",
				Ca:  "/workspace/jobs/st8ger_controller/ca.crt",
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create https client")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to execute the post request")
	}
	fmt.Printf("WE GOT A RESPONSE: %+v", resp)

	staging.Spec.State = kapiv1alpha1.StartedState

	_, err = c.kapiClientset.SamplecontrollerV1alpha1().Stagings(namespace).Update(staging)
	if err != nil {
		fmt.Println("Failed to update status", err)
	}

	return err
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
