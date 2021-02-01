package updater

import (
	"context"
	"fmt"

	nais_v1alpha1 "github.com/nais/liberator/pkg/apis/nais.io/v1alpha1"
	naiserator_scheme "github.com/nais/naiserator/pkg/naiserator/scheme"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateOrUpdate(ctx context.Context, cli client.Client, scheme *runtime.Scheme, resource runtime.Object) func() error {
	return func() error {
		log.Infof("CreateOrUpdate %s", naiserator_scheme.TypeName(resource))
		existing, err := scheme.New(resource.GetObjectKind().GroupVersionKind())
		if err != nil {
			return fmt.Errorf("internal error: %w", err)
		}
		objectKey, err := client.ObjectKeyFromObject(resource)
		if err != nil {
			return fmt.Errorf("unable to derive object key: %w", err)
		}
		err = cli.Get(ctx, objectKey, existing)

		if errors.IsNotFound(err) {
			err = cli.Create(ctx, resource)
		} else if err == nil {
			err = CopyMeta(existing, resource)
			if err != nil {
				return err
			}
			err = CopyImmutable(existing, resource)
			if err != nil {
				return err
			}
			err = cli.Update(ctx, resource)
		}

		if err != nil {
			return err
		}
		return err
	}
}

func CreateOrRecreate(ctx context.Context, cli client.Client, resource runtime.Object) func() error {
	return func() error {
		log.Infof("CreateOrRecreate %s", naiserator_scheme.TypeName(resource))
		err := cli.Delete(ctx, resource)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		return cli.Create(ctx, resource)
	}
}

func CreateIfNotExists(ctx context.Context, cli client.Client, resource runtime.Object) func() error {
	return func() error {
		log.Infof("CreateIfNotExists %s", naiserator_scheme.TypeName(resource))
		err := cli.Create(ctx, resource)
		if err != nil && errors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}
}

func DeleteIfExists(ctx context.Context, cli client.Client, resource runtime.Object) func() error {
	return func() error {
		log.Infof("DeleteIfExists %s", naiserator_scheme.TypeName(resource))
		err := cli.Delete(ctx, resource)
		if err != nil && errors.IsNotFound(err) {
			return nil
		}
		return err
	}
}

// Find all Kubernetes resource matching label selector 'app=NAME' for all specified types
func FindAll(ctx context.Context, cli client.Client, scheme *runtime.Scheme, types []runtime.Object, app *nais_v1alpha1.Application) ([]runtime.Object, error) {
	// Set up label selector 'app=NAME'
	labelSelector := labels.NewSelector()
	labelreq, err := labels.NewRequirement("app", selection.Equals, []string{app.Name})
	if err != nil {
		return nil, err
	}
	labelSelector.Add(*labelreq)
	listopt := &client.ListOptions{
		LabelSelector: labelSelector,
	}

	resources := make([]runtime.Object, 0)

	for _, obj := range types {
		err = cli.List(ctx, obj, listopt)
		if err != nil {
			return nil, fmt.Errorf("list %T: %w", obj, err)
		}

		_ = meta.EachListItem(obj, func(item runtime.Object) error {
			resources = append(resources, item)
			return nil
		})
	}

	return withOwnerReference(app, resources), nil
}

// CopyMeta copies resource metadata from one resource to another.
// used when updating existing resources in the cluster.
func CopyMeta(src, dst runtime.Object) error {
	srcacc, err := meta.Accessor(src)
	if err != nil {
		return err
	}

	dstacc, err := meta.Accessor(dst)
	if err != nil {
		return err
	}

	dstacc.SetResourceVersion(srcacc.GetResourceVersion())
	dstacc.SetUID(srcacc.GetUID())
	dstacc.SetSelfLink(srcacc.GetSelfLink())

	return err
}

func CopyImmutable(src, dst runtime.Object) error {
	switch srcTyped := src.(type) {
	case *corev1.Service:
		// ClusterIP must be retained as the field is immutable.
		dstTyped, ok := dst.(*corev1.Service)
		if !ok {
			return fmt.Errorf("source and destination types differ (%T != %T)", src, dst)
		}
		dstTyped.Spec.ClusterIP = srcTyped.Spec.ClusterIP
	}
	return nil
}

func withOwnerReference(app *nais_v1alpha1.Application, resources []runtime.Object) []runtime.Object {
	owned := make([]runtime.Object, 0, len(resources))

	hasOwnerReference := func(r runtime.Object) (bool, error) {
		m, err := meta.Accessor(r)
		if err != nil {
			return false, err
		}
		for _, ref := range m.GetOwnerReferences() {
			if ref.UID == app.UID {
				return true, nil
			}
		}
		return false, nil
	}

	for _, resource := range resources {
		ok, err := hasOwnerReference(resource)
		if err == nil && ok {
			owned = append(owned, resource)
		}
	}

	return owned
}
