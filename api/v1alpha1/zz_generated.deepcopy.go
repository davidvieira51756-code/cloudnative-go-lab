package v1alpha1

import runtime "k8s.io/apimachinery/pkg/runtime"

func (in *AppDeployment) DeepCopyInto(out *AppDeployment) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

func (in *AppDeployment) DeepCopy() *AppDeployment {
	if in == nil {
		return nil
	}

	out := new(AppDeployment)
	in.DeepCopyInto(out)

	return out
}

func (in *AppDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

func (in *AppDeploymentList) DeepCopyInto(out *AppDeploymentList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)

	if in.Items != nil {
		out.Items = make([]AppDeployment, len(in.Items))

		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *AppDeploymentList) DeepCopy() *AppDeploymentList {
	if in == nil {
		return nil
	}

	out := new(AppDeploymentList)
	in.DeepCopyInto(out)

	return out
}

func (in *AppDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}
