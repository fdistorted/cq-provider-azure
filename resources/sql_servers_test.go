package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockSqlServerClient(ctrl)
	databaseSvc := mocks.NewMockSqlDatabaseClient(ctrl)
	firewallSvc := mocks.NewMockSQLFirewallClient(ctrl)
	adminsSvc := mocks.NewMockSQLServerAdminClient(ctrl)
	databaseBlobSvc := mocks.NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl)
	serverBlobSvc := mocks.NewMockSQLServerBlobAuditingPolicies(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Databases:                    databaseSvc,
			DatabaseBlobAuditingPolicies: databaseBlobSvc,
			Firewall:                     firewallSvc,
			ServerAdmins:                 adminsSvc,
			ServerBlobAuditingPolicies:   serverBlobSvc,
			Servers:                      serverSvc,
		},
	}
	server := sql.Server{}
	if err := faker.FakeData(&server); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	server.ID = &fakeResourceGroup
	serverSvc.EXPECT().List(gomock.Any()).Return(
		sql.NewServerListResultPage(
			sql.ServerListResult{Value: &[]sql.Server{server}},
			func(context.Context, sql.ServerListResult) (sql.ServerListResult, error) {
				return sql.ServerListResult{}, nil
			},
		),
		nil,
	)

	database := sql.Database{}
	if err := faker.FakeData(&database); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	database.ID = &fakeResourceGroup
	databaseSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewDatabaseListResultPage(
			sql.DatabaseListResult{Value: &[]sql.Database{database}},
			func(context.Context, sql.DatabaseListResult) (sql.DatabaseListResult, error) {
				return sql.DatabaseListResult{}, nil
			},
		), nil,
	)

	var databaseBlobPolicy sql.DatabaseBlobAuditingPolicy
	if err := faker.FakeData(&databaseBlobPolicy); err != nil {
		t.Fatal(err)
	}
	databaseBlobSvc.EXPECT().ListByDatabase(gomock.Any(), "test", *server.Name, *database.Name).Return(
		sql.NewDatabaseBlobAuditingPolicyListResultPage(
			sql.DatabaseBlobAuditingPolicyListResult{Value: &[]sql.DatabaseBlobAuditingPolicy{databaseBlobPolicy}},
			func(context.Context, sql.DatabaseBlobAuditingPolicyListResult) (sql.DatabaseBlobAuditingPolicyListResult, error) {
				return sql.DatabaseBlobAuditingPolicyListResult{}, nil
			},
		), nil,
	)

	var rule sql.FirewallRule
	if err := faker.FakeData(&rule); err != nil {
		t.Fatal(err)
	}
	firewallSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.FirewallRuleListResult{Value: &[]sql.FirewallRule{rule}}, nil,
	)

	var admin sql.ServerAzureADAdministrator
	if err := faker.FakeData(&admin); err != nil {
		t.Fatal(err)
	}
	adminPage := sql.NewAdministratorListResultPage(
		sql.AdministratorListResult{Value: &[]sql.ServerAzureADAdministrator{admin}},
		func(context.Context, sql.AdministratorListResult) (sql.AdministratorListResult, error) {
			return sql.AdministratorListResult{}, nil
		},
	)
	adminsSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(adminPage, nil)

	var serverBlobPolicy sql.ServerBlobAuditingPolicy
	if err := faker.FakeData(&serverBlobPolicy); err != nil {
		t.Fatal(err)
	}
	serverBlobSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.NewServerBlobAuditingPolicyListResultPage(
			sql.ServerBlobAuditingPolicyListResult{Value: &[]sql.ServerBlobAuditingPolicy{serverBlobPolicy}},
			func(context.Context, sql.ServerBlobAuditingPolicyListResult) (sql.ServerBlobAuditingPolicyListResult, error) {
				return sql.ServerBlobAuditingPolicyListResult{}, nil
			},
		), nil,
	)
	return s
}

func TestSQLServers(t *testing.T) {
	azureTestHelper(t, resources.SQLServers(), buildSQLServerMock)
}
