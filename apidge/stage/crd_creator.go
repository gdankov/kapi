package stage

import (
	v1alpha1 "github.com/cloudfoundry-community/kapi/pkg/apis/kapi/v1alpha1"
	clientset "github.com/cloudfoundry-community/kapi/pkg/generated/clientset/versioned"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CRDCreator struct {
	Clientset clientset.Interface
	Namespace string
}

func (c CRDCreator) Create(spec v1alpha1.StagingSpec) error {
	spec.State = v1alpha1.NotStartedState

	staging := &v1alpha1.Staging{
		ObjectMeta: metav1.ObjectMeta{
			Name:      spec.AppGUID,
			Namespace: c.Namespace,
		},
		Spec: spec,
	}

	_, err := c.Clientset.SamplecontrollerV1alpha1().Stagings(c.Namespace).Create(staging)
	if err != nil {
		return errors.Wrap(err, "failed to create staging crd")
	}

	return nil
}
