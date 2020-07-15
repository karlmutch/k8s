package v1beta1

import "github.com/karlmutch.k8s"

func init() {
	k8s.Register("storage.k8s.io", "v1beta1", "storageclasses", false, &StorageClass{})

	k8s.RegisterList("storage.k8s.io", "v1beta1", "storageclasses", false, &StorageClassList{})
}
