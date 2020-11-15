//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.volterra.us/stdlib/sets"

	ves_io_schema_namespace "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"

	ves_io_schema_advertise_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/advertise_policy"
	ves_io_schema_alert_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_policy"
	ves_io_schema_alert_policy_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_policy_set"
	ves_io_schema_alert_receiver "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_receiver"
	ves_io_schema_app_type "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_type"
	ves_io_schema_bgp "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/bgp"
	ves_io_schema_bgp_asn_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/bgp_asn_set"
	ves_io_schema_cloud_credentials "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
	ves_io_schema_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cluster"
	ves_io_schema_discovery "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/discovery"
	ves_io_schema_endpoint "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/endpoint"
	ves_io_schema_fast_acl "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl"
	ves_io_schema_fast_acl_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl_rule"
	ves_io_schema_fast_acl_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl_set"
	ves_io_schema_fleet "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fleet"
	ves_io_schema_healthcheck "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/healthcheck"
	ves_io_schema_ip_prefix_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ip_prefix_set"
	ves_io_schema_malicious_user_mitigation "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/malicious_user_mitigation"
	ves_io_schema_network_connector "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_connector"
	ves_io_schema_network_firewall "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_firewall"
	ves_io_schema_network_interface "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_interface"
	ves_io_schema_network_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy"
	ves_io_schema_network_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_rule"
	ves_io_schema_network_policy_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_set"
	ves_io_schema_policer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policer"
	ves_io_schema_protocol_policer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/protocol_policer"
	ves_io_schema_rate_limiter "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/rate_limiter"
	ves_io_schema_role "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/role"
	ves_io_schema_route "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/route"
	ves_io_schema_secret_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/secret_policy"
	ves_io_schema_secret_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/secret_policy_rule"
	ves_io_schema_service_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
	ves_io_schema_service_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy_rule"
	ves_io_schema_service_policy_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy_set"
	ves_io_schema_token "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/token"
	ves_io_schema_user "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user"
	ves_io_schema_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
	ves_io_schema_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	ves_io_schema_azure_vnet_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
	ves_io_schema_forward_proxy_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/forward_proxy_policy"
	ves_io_schema_gcp_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
	ves_io_schema_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_network_policy_view "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/network_policy_view"
	ves_io_schema_origin_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/origin_pool"
	ves_io_schema_tcp_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/tcp_loadbalancer"
	ves_io_schema_virtual_host "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	ves_io_schema_virtual_k8s "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_k8s"
	ves_io_schema_virtual_network "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_network"
	ves_io_schema_virtual_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_site"
	ves_io_schema_waf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf"
	ves_io_schema_waf_rule_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rule_list"
	ves_io_schema_waf_rules "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rules"
)

func getVolterraResourceMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"volterra_namespace":                 resourceVolterraNamespace(),
		"volterra_advertise_policy":          resourceVolterraAdvertisePolicy(),
		"volterra_alert_policy":              resourceVolterraAlertPolicy(),
		"volterra_alert_policy_set":          resourceVolterraAlertPolicySet(),
		"volterra_alert_receiver":            resourceVolterraAlertReceiver(),
		"volterra_app_type":                  resourceVolterraAppType(),
		"volterra_bgp":                       resourceVolterraBgp(),
		"volterra_bgp_asn_set":               resourceVolterraBgpAsnSet(),
		"volterra_cloud_credentials":         resourceVolterraCloudCredentials(),
		"volterra_cluster":                   resourceVolterraCluster(),
		"volterra_discovery":                 resourceVolterraDiscovery(),
		"volterra_endpoint":                  resourceVolterraEndpoint(),
		"volterra_fast_acl":                  resourceVolterraFastAcl(),
		"volterra_fast_acl_rule":             resourceVolterraFastAclRule(),
		"volterra_fast_acl_set":              resourceVolterraFastAclSet(),
		"volterra_fleet":                     resourceVolterraFleet(),
		"volterra_healthcheck":               resourceVolterraHealthcheck(),
		"volterra_ip_prefix_set":             resourceVolterraIpPrefixSet(),
		"volterra_malicious_user_mitigation": resourceVolterraMaliciousUserMitigation(),
		"volterra_network_connector":         resourceVolterraNetworkConnector(),
		"volterra_network_firewall":          resourceVolterraNetworkFirewall(),
		"volterra_network_interface":         resourceVolterraNetworkInterface(),
		"volterra_network_policy":            resourceVolterraNetworkPolicy(),
		"volterra_network_policy_rule":       resourceVolterraNetworkPolicyRule(),
		"volterra_network_policy_set":        resourceVolterraNetworkPolicySet(),
		"volterra_policer":                   resourceVolterraPolicer(),
		"volterra_protocol_policer":          resourceVolterraProtocolPolicer(),
		"volterra_rate_limiter":              resourceVolterraRateLimiter(),
		"volterra_role":                      resourceVolterraRole(),
		"volterra_route":                     resourceVolterraRoute(),
		"volterra_secret_policy":             resourceVolterraSecretPolicy(),
		"volterra_secret_policy_rule":        resourceVolterraSecretPolicyRule(),
		"volterra_service_policy":            resourceVolterraServicePolicy(),
		"volterra_service_policy_rule":       resourceVolterraServicePolicyRule(),
		"volterra_service_policy_set":        resourceVolterraServicePolicySet(),
		"volterra_token":                     resourceVolterraToken(),
		"volterra_user":                      resourceVolterraUser(),
		"volterra_aws_tgw_site":              resourceVolterraAwsTgwSite(),
		"volterra_aws_vpc_site":              resourceVolterraAwsVpcSite(),
		"volterra_azure_vnet_site":           resourceVolterraAzureVnetSite(),
		"volterra_forward_proxy_policy":      resourceVolterraForwardProxyPolicy(),
		"volterra_gcp_vpc_site":              resourceVolterraGcpVpcSite(),
		"volterra_http_loadbalancer":         resourceVolterraHttpLoadbalancer(),
		"volterra_network_policy_view":       resourceVolterraNetworkPolicyView(),
		"volterra_origin_pool":               resourceVolterraOriginPool(),
		"volterra_tcp_loadbalancer":          resourceVolterraTcpLoadbalancer(),
		"volterra_virtual_host":              resourceVolterraVirtualHost(),
		"volterra_virtual_k8s":               resourceVolterraVirtualK8S(),
		"volterra_virtual_network":           resourceVolterraVirtualNetwork(),
		"volterra_virtual_site":              resourceVolterraVirtualSite(),
		"volterra_waf":                       resourceVolterraWaf(),
		"volterra_waf_rule_list":             resourceVolterraWafRuleList(),
		"volterra_waf_rules":                 resourceVolterraWafRules(),
	}
}

func getAllAkarObjectTypes() sets.String {

	return sets.NewString(
		ves_io_schema_namespace.ObjectType,
		ves_io_schema_advertise_policy.ObjectType,
		ves_io_schema_alert_policy.ObjectType,
		ves_io_schema_alert_policy_set.ObjectType,
		ves_io_schema_alert_receiver.ObjectType,
		ves_io_schema_app_type.ObjectType,
		ves_io_schema_bgp.ObjectType,
		ves_io_schema_bgp_asn_set.ObjectType,
		ves_io_schema_cloud_credentials.ObjectType,
		ves_io_schema_cluster.ObjectType,
		ves_io_schema_discovery.ObjectType,
		ves_io_schema_endpoint.ObjectType,
		ves_io_schema_fast_acl.ObjectType,
		ves_io_schema_fast_acl_rule.ObjectType,
		ves_io_schema_fast_acl_set.ObjectType,
		ves_io_schema_fleet.ObjectType,
		ves_io_schema_healthcheck.ObjectType,
		ves_io_schema_ip_prefix_set.ObjectType,
		ves_io_schema_malicious_user_mitigation.ObjectType,
		ves_io_schema_network_connector.ObjectType,
		ves_io_schema_network_firewall.ObjectType,
		ves_io_schema_network_interface.ObjectType,
		ves_io_schema_network_policy.ObjectType,
		ves_io_schema_network_policy_rule.ObjectType,
		ves_io_schema_network_policy_set.ObjectType,
		ves_io_schema_policer.ObjectType,
		ves_io_schema_protocol_policer.ObjectType,
		ves_io_schema_rate_limiter.ObjectType,
		ves_io_schema_role.ObjectType,
		ves_io_schema_route.ObjectType,
		ves_io_schema_secret_policy.ObjectType,
		ves_io_schema_secret_policy_rule.ObjectType,
		ves_io_schema_service_policy.ObjectType,
		ves_io_schema_service_policy_rule.ObjectType,
		ves_io_schema_service_policy_set.ObjectType,
		ves_io_schema_token.ObjectType,
		ves_io_schema_user.ObjectType,
		ves_io_schema_aws_tgw_site.ObjectType,
		ves_io_schema_aws_vpc_site.ObjectType,
		ves_io_schema_azure_vnet_site.ObjectType,
		ves_io_schema_forward_proxy_policy.ObjectType,
		ves_io_schema_gcp_vpc_site.ObjectType,
		ves_io_schema_http_loadbalancer.ObjectType,
		ves_io_schema_network_policy_view.ObjectType,
		ves_io_schema_origin_pool.ObjectType,
		ves_io_schema_tcp_loadbalancer.ObjectType,
		ves_io_schema_virtual_host.ObjectType,
		ves_io_schema_virtual_k8s.ObjectType,
		ves_io_schema_virtual_network.ObjectType,
		ves_io_schema_virtual_site.ObjectType,
		ves_io_schema_waf.ObjectType,
		ves_io_schema_waf_rule_list.ObjectType,
		ves_io_schema_waf_rules.ObjectType,
	)
}
