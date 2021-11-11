package v1alpha1

import (
	"github.com/netgroup-polito/CrownLabs/operators/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *Tenant) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.Tenant)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.CreateSandbox = src.Spec.CreateSandbox
	dst.Spec.Email = src.Spec.Email
	dst.Spec.FirstName = src.Spec.FirstName
	dst.Spec.LastName = src.Spec.LastName
	dst.Spec.PublicKeys = src.Spec.PublicKeys
	dst.Spec.Workspaces = make([]v1alpha2.TenantWorkspaceEntry, len(src.Spec.Workspaces))
	for i, w := range src.Spec.Workspaces {
		dst.Spec.Workspaces[i].Name = w.WorkspaceRef.Name
		dst.Spec.Workspaces[i].Role = v1alpha2.WorkspaceUserRole(w.Role)
	}

	dst.Status.PersonalNamespace = src.Status.PersonalNamespace
	dst.Status.SandboxNamespace = src.Status.SandboxNamespace
	dst.Status.FailingWorkspaces = src.Status.FailingWorkspaces
	dst.Status.Subscriptions = src.Status.Subscriptions
	dst.Status.Ready = src.Status.Ready

	return nil
}

func (dst *Tenant) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.Tenant)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.CreateSandbox = src.Spec.CreateSandbox
	dst.Spec.Email = src.Spec.Email
	dst.Spec.FirstName = src.Spec.FirstName
	dst.Spec.LastName = src.Spec.LastName
	dst.Spec.PublicKeys = src.Spec.PublicKeys
	dst.Spec.Workspaces = make([]TenantWorkspaceEntry, len(src.Spec.Workspaces))
	for i, w := range src.Spec.Workspaces {
		dst.Spec.Workspaces[i].WorkspaceRef.Name = w.Name
		dst.Spec.Workspaces[i].Role = WorkspaceUserRole(w.Role)
	}

	dst.Status.PersonalNamespace = src.Status.PersonalNamespace
	dst.Status.SandboxNamespace = src.Status.SandboxNamespace
	dst.Status.FailingWorkspaces = src.Status.FailingWorkspaces
	dst.Status.Subscriptions = src.Status.Subscriptions
	dst.Status.Ready = src.Status.Ready

	return nil
}
