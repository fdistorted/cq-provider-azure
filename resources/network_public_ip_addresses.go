package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"net"
)

func NetworkPublicIPAddresses() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_public_ip_addresses",
		Description:  "PublicIPAddress public IP address resource.",
		Resolver:     fetchNetworkPublicIpAddresses,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "extended_location_name",
				Description: "Name - The name of the extended location.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "Type - The type of the extended location.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "sku_name",
				Description: "Name - Name of a public IP address SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier - Tier of a public IP address SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "public_ip_allocation_method",
				Description: "PublicIPAllocationMethod - The public IP address allocation method",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAllocationMethod"),
			},
			{
				Name:        "public_ip_address_version",
				Description: "PublicIPAddressVersion - The public IP address version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAddressVersion"),
			},
			{
				Name:        "ip_configuration",
				Description: "IPConfiguration - READ-ONLY; The IP configuration associated with the public IP address.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesIpConfiguration,
			},
			{
				Name:        "dns_settings_domain_name_label",
				Description: "DomainNameLabel - The domain name label",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.DomainNameLabel"),
			},
			{
				Name:        "dns_settings_fqdn",
				Description: "Fqdn - The Fully Qualified Domain Name of the A DNS record associated with the public IP",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.Fqdn"),
			},
			{
				Name:        "dns_settings_reverse_fqdn",
				Description: "ReverseFqdn - The reverse FQDN",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.ReverseFqdn"),
			},
			{
				Name:        "ddos_settings_ddos_custom_policy_id",
				Description: "ID - Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.DdosCustomPolicy.ID"),
			},
			{
				Name:        "ddos_settings_protection_coverage",
				Description: "ProtectionCoverage - The DDoS protection policy customizability of the public IP",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectionCoverage"),
			},
			{
				Name:        "ddos_settings_protected_ip",
				Description: "ProtectedIP - Enables DDoS protection on the public IP.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectedIP"),
			},
			{
				Name:        "ip_tags",
				Description: "IPTags - The list of tags associated with the public IP address.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesIpTags,
			},
			{
				Name:        "ip_address",
				Description: "IPAddress - The IP address associated with the public IP address resource.",
				Type:        schema.TypeCIDR,
				Resolver:    resolveNetworkPublicIPAddressesIpAddress,
			},
			{
				Name:        "public_ip_prefix_id",
				Description: "ID - Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPPrefix.ID"),
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "IdleTimeoutInMinutes - The idle timeout of the public IP address.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IdleTimeoutInMinutes"),
			},
			{
				Name:        "resource_guid",
				Description: "ResourceGUID - READ-ONLY; The resource GUID property of the public IP address resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "ProvisioningState - READ-ONLY; The provisioning state of the public IP address resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "service_public_ip_address",
				Description: "ServicePublicIPAddress - The service public IP address of the public IP address resource.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesServicePublicIpAddress,
			},
			{
				Name:        "nat_gateway",
				Description: "NatGateway - The NatGateway for the Public IP address.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesNatGateway,
			},
			{
				Name:        "migration_phase",
				Description: "MigrationPhase - Migration phase of Public IP Address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.MigrationPhase"),
			},
			{
				Name:        "linked_public_ip_address",
				Description: "LinkedPublicIPAddress - The linked public IP address of the public IP address resource.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesLinkedPublicIpAddress,
			},
			{
				Name:        "etag",
				Description: "Etag - READ-ONLY; A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "zones",
				Description: "Zones - A list of availability zones denoting the IP allocated for the resource needs to come from.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "ID - Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Name - READ-ONLY; Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Type - READ-ONLY; Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Location - Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags - Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkPublicIpAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Network.PublicIPAddresses
	response, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func resolveNetworkPublicIPAddressesIpConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.IPConfiguration)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesIpTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}
	if p.IPTags == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, t := range *p.IPTags {
		j[*t.IPTagType] = *t.Tag
	}
	return resource.Set(c.Name, j)
}
func resolveNetworkPublicIPAddressesIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	i := net.ParseIP(*p.IPAddress)
	_, ip, err := net.ParseCIDR(i.String() + "/32")
	if err != nil {
		i.DefaultMask()
		return err
	}

	return resource.Set(c.Name, ip)
}
func resolveNetworkPublicIPAddressesServicePublicIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.ServicePublicIPAddress == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.ServicePublicIPAddress)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesNatGateway(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.NatGateway == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.NatGateway)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesLinkedPublicIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.LinkedPublicIPAddress == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.LinkedPublicIPAddress)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
