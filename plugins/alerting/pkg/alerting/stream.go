package alerting

import (
	"github.com/rancher/opni/pkg/capabilities/wellknown"
	streamext "github.com/rancher/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/rancher/opni/plugins/alerting/pkg/apis/node"
)

func (p *Plugin) StreamServers() []streamext.Server {
	return []streamext.Server{
		{
			Desc:              &node.NodeAlertingCapability_ServiceDesc,
			Impl:              &p.node,
			RequireCapability: wellknown.CapabilityAlerting,
		},
	}
}
