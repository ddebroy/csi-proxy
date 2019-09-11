// Code generated by csi-proxy-gen. DO NOT EDIT.

package dummy

import (
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal/v1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal/v1alpha2"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
)

const name = "dummy"

// ensure the server defines all the required methods
var _ internal.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*server.VersionedAPI {
	v1alpha1Server := v1alpha1.NewVersionedServer(s)
	v1alpha2Server := v1alpha2.NewVersionedServer(s)
	v1Server := v1.NewVersionedServer(s)

	return []*server.VersionedAPI{
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha1"),
			Registrant: v1alpha1Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha2"),
			Registrant: v1alpha2Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1"),
			Registrant: v1Server.Register,
		},
	}
}
