package bridge

import (
	v1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi.cf.org/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CRDCreator struct {
	Clientset clientset.Interface
	Namespace string
}

func (c CRDCreator) CreateStaging(spec v1alpha1.StagingSpec) error {
	staging := &v1alpha1.Staging{
		ObjectMeta: metav1.ObjectMeta{
			Name:      spec.AppGUID,
			Namespace: c.Namespace,
		},
		Spec:   spec,
		Status: v1alpha1.StagingStatus{State: v1alpha1.NotStartedState},
	}

	if _, err := c.Clientset.KapiV1alpha1().Stagings(c.Namespace).Create(staging); err != nil {
		return errors.Wrap(err, "failed to create staging crd")
	}

	return nil
}

func (c CRDCreator) CreateLRP(spec v1alpha1.LRPSpec) error {
	staging := &v1alpha1.LRP{
		ObjectMeta: metav1.ObjectMeta{
			Name:      spec.ProcessGUID,
			Namespace: c.Namespace,
		},
		Spec:   spec,
		Status: v1alpha1.LRPStatus{State: v1alpha1.NotStartedState},
	}

	if _, err := c.Clientset.KapiV1alpha1().LRPs(c.Namespace).Create(staging); err != nil {
		return errors.Wrap(err, "failed to create app crd")
	}

	return nil
}
