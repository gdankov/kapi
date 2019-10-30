package lrp

import (
	"fmt"
	"time"

	"github.com/cloudfoundry-community/kapi/eirini"
	kapiv1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi.cf.org/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	samplescheme "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned/scheme"
	informers "github.com/cloudfoundry-community/kapi/pkg/generated/informers/externalversions/kapi.cf.org/v1alpha1"
	listers "github.com/cloudfoundry-community/kapi/pkg/generated/listers/kapi.cf.org/v1alpha1"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

type LRPController struct {
	clientset clientset.Interface
	workqueue workqueue.RateLimitingInterface

	lrpLister  listers.LRPLister
	lrpsSynced cache.InformerSynced

	eiriniClient eirini.Client
}

// NewController returns a new sample controller
func NewController(
	clientset clientset.Interface,
	lrpInformer informers.LRPInformer,
	eiriniClient eirini.Client) *LRPController {

	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)

	controller := &LRPController{
		clientset:    clientset,
		lrpLister:    lrpInformer.Lister(),
		lrpsSynced:   lrpInformer.Informer().HasSynced,
		workqueue:    workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "LRPs"),
		eiriniClient: eiriniClient,
	}

	klog.Info("Setting up event handlers")
	lrpInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueLRP,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueLRP(new)
		},
	})
	return controller
}

func (c *LRPController) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.lrpsSynced); !ok {
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

func (c *LRPController) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *LRPController) processNextWorkItem() bool {
	key, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(keyObj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(keyObj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = keyObj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(keyObj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", keyObj))
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
		c.workqueue.Forget(keyObj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(key)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *LRPController) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	return c.handleResource(name, namespace)
}

// The business logic
func (c *LRPController) handleResource(name, namespace string) error {
	lrp, err := c.lrpLister.LRPs(namespace).Get(name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("LRP '%s/%s' in work queue no longer exists", name, namespace))
			return nil
		}
		return err
	}

	if lrp.Status.State != kapiv1alpha1.NotStartedState {
		return nil
	}

	if err := c.eiriniClient.Desire(lrp.Spec); err != nil {
		return errors.Wrap(err, "eirini failed to stage the app")
	}

	lrp.Status.State = kapiv1alpha1.StartedState

	_, err = c.clientset.KapiV1alpha1().LRPs(namespace).UpdateStatus(lrp)
	return errors.Wrap(err, "failed to update lrp status")
}

func (c *LRPController) enqueueLRP(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}
