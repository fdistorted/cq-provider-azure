//go:build integration

package monitor

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMonitorActivityLogs(t *testing.T) {
	client.AzureTestHelper(t, MonitorActivityLogs(),
		client.SnapshotsDirPath)
}
