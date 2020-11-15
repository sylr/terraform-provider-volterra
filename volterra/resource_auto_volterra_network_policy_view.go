//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_network_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_rule"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_network_policy_view "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/network_policy_view"
)

// resourceVolterraNetworkPolicyView is implementation of Volterra's NetworkPolicyView resources
func resourceVolterraNetworkPolicyView() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkPolicyViewCreate,
		Read:   resourceVolterraNetworkPolicyViewRead,
		Update: resourceVolterraNetworkPolicyViewUpdate,
		Delete: resourceVolterraNetworkPolicyViewDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"egress_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"keys": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"any": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ip_prefix_set": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tenant": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"label_selector": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expressions": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"namespace": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"outside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"prefix_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"prefixes": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"rule_description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"rule_name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"all_tcp_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"all_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"all_udp_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"applications": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"applications": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"protocol_port_range": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"port_ranges": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"endpoint": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"any": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"label_selector": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expressions": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"namespace": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"outside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"prefix_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"prefixes": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},

			"ingress_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"keys": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"any": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ip_prefix_set": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tenant": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"label_selector": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expressions": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"namespace": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"outside_endpoints": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"prefix_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"prefixes": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"rule_description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"rule_name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"all_tcp_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"all_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"all_udp_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"applications": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"applications": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"protocol_port_range": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"port_ranges": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraNetworkPolicyViewCreate creates NetworkPolicyView resource
func resourceVolterraNetworkPolicyViewCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_network_policy_view.CreateSpecType{}
	createReq := &ves_io_schema_views_network_policy_view.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("egress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		egressRules := make([]*ves_io_schema_views_network_policy_view.NetworkPolicyRuleType, len(sl))
		createSpec.EgressRules = egressRules
		for i, set := range sl {
			egressRules[i] = &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType{}

			egressRulesMapStrToI := set.(map[string]interface{})

			if v, ok := egressRulesMapStrToI["action"]; ok && !isIntfNil(v) {

				egressRules[i].Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

			}

			if w, ok := egressRulesMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				egressRules[i].Keys = ls
			}

			otherEndpointTypeFound := false

			if v, ok := egressRulesMapStrToI["any"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Any{}
					otherEndpointInt.Any = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_InsideEndpoints{}
					otherEndpointInt.InsideEndpoints = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_IpPrefixSet{}
				otherEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						otherEndpointInt.IpPrefixSet.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "ip_prefix_set"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := egressRulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_LabelSelector{}
				otherEndpointInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := egressRulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Namespace{}

				egressRules[i].OtherEndpoint = otherEndpointInt

				otherEndpointInt.Namespace = v.(string)

			}

			if v, ok := egressRulesMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_OutsideEndpoints{}
					otherEndpointInt.OutsideEndpoints = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_PrefixList{}
				otherEndpointInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.PrefixList.Prefixes = ls

					}

				}

			}

			if w, ok := egressRulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
				egressRules[i].RuleDescription = w.(string)
			}

			if w, ok := egressRulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
				egressRules[i].RuleName = w.(string)
			}

			trafficChoiceTypeFound := false

			if v, ok := egressRulesMapStrToI["all_tcp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTcpTraffic{}
					trafficChoiceInt.AllTcpTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["all_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTraffic{}
					trafficChoiceInt.AllTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["all_udp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllUdpTraffic{}
					trafficChoiceInt.AllUdpTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["applications"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Applications{}
				trafficChoiceInt.Applications = &ves_io_schema_views_network_policy_view.ApplicationsType{}
				egressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["applications"]; ok && !isIntfNil(v) {

						applicationsList := []ves_io_schema_views_network_policy_view.ApplicationEnumType{}
						for _, j := range v.([]interface{}) {
							applicationsList = append(applicationsList, ves_io_schema_views_network_policy_view.ApplicationEnumType(ves_io_schema_views_network_policy_view.ApplicationEnumType_value[j.(string)]))
						}
						trafficChoiceInt.Applications.Applications = applicationsList

					}

				}

			}

			if v, ok := egressRulesMapStrToI["protocol_port_range"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_ProtocolPortRange{}
				trafficChoiceInt.ProtocolPortRange = &ves_io_schema_views_network_policy_view.ProtocolPortType{}
				egressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["port_ranges"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						trafficChoiceInt.ProtocolPortRange.PortRanges = ls

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						trafficChoiceInt.ProtocolPortRange.Protocol = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("endpoint"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		endpoint := &ves_io_schema_views_network_policy_view.EndpointChoiceType{}
		createSpec.Endpoint = endpoint
		for _, set := range sl {

			endpointMapStrToI := set.(map[string]interface{})

			endpointChoiceTypeFound := false

			if v, ok := endpointMapStrToI["any"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Any{}
					endpointChoiceInt.Any = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_InsideEndpoints{}
					endpointChoiceInt.InsideEndpoints = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["interface"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Interface{}
				endpointChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Tenant = v.(string)
					}

				}

			}

			if v, ok := endpointMapStrToI["label_selector"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_LabelSelector{}
				endpointChoiceInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						endpointChoiceInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := endpointMapStrToI["namespace"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Namespace{}

				endpoint.EndpointChoice = endpointChoiceInt

				endpointChoiceInt.Namespace = v.(string)

			}

			if v, ok := endpointMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_OutsideEndpoints{}
					endpointChoiceInt.OutsideEndpoints = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_PrefixList{}
				endpointChoiceInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						endpointChoiceInt.PrefixList.Prefixes = ls

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ingressRules := make([]*ves_io_schema_views_network_policy_view.NetworkPolicyRuleType, len(sl))
		createSpec.IngressRules = ingressRules
		for i, set := range sl {
			ingressRules[i] = &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType{}

			ingressRulesMapStrToI := set.(map[string]interface{})

			if v, ok := ingressRulesMapStrToI["action"]; ok && !isIntfNil(v) {

				ingressRules[i].Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

			}

			if w, ok := ingressRulesMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				ingressRules[i].Keys = ls
			}

			otherEndpointTypeFound := false

			if v, ok := ingressRulesMapStrToI["any"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Any{}
					otherEndpointInt.Any = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_InsideEndpoints{}
					otherEndpointInt.InsideEndpoints = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_IpPrefixSet{}
				otherEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						otherEndpointInt.IpPrefixSet.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "ip_prefix_set"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_LabelSelector{}
				otherEndpointInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Namespace{}

				ingressRules[i].OtherEndpoint = otherEndpointInt

				otherEndpointInt.Namespace = v.(string)

			}

			if v, ok := ingressRulesMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_OutsideEndpoints{}
					otherEndpointInt.OutsideEndpoints = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_PrefixList{}
				otherEndpointInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.PrefixList.Prefixes = ls

					}

				}

			}

			if w, ok := ingressRulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
				ingressRules[i].RuleDescription = w.(string)
			}

			if w, ok := ingressRulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
				ingressRules[i].RuleName = w.(string)
			}

			trafficChoiceTypeFound := false

			if v, ok := ingressRulesMapStrToI["all_tcp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTcpTraffic{}
					trafficChoiceInt.AllTcpTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["all_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTraffic{}
					trafficChoiceInt.AllTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["all_udp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllUdpTraffic{}
					trafficChoiceInt.AllUdpTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["applications"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Applications{}
				trafficChoiceInt.Applications = &ves_io_schema_views_network_policy_view.ApplicationsType{}
				ingressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["applications"]; ok && !isIntfNil(v) {

						applicationsList := []ves_io_schema_views_network_policy_view.ApplicationEnumType{}
						for _, j := range v.([]interface{}) {
							applicationsList = append(applicationsList, ves_io_schema_views_network_policy_view.ApplicationEnumType(ves_io_schema_views_network_policy_view.ApplicationEnumType_value[j.(string)]))
						}
						trafficChoiceInt.Applications.Applications = applicationsList

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["protocol_port_range"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_ProtocolPortRange{}
				trafficChoiceInt.ProtocolPortRange = &ves_io_schema_views_network_policy_view.ProtocolPortType{}
				ingressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["port_ranges"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						trafficChoiceInt.ProtocolPortRange.PortRanges = ls

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						trafficChoiceInt.ProtocolPortRange.Protocol = v.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkPolicyView object with struct: %+v", createReq)

	createNetworkPolicyViewResp, err := client.CreateObject(context.Background(), ves_io_schema_views_network_policy_view.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkPolicyView: %s", err)
	}
	d.SetId(createNetworkPolicyViewResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkPolicyViewRead(d, meta)
}

func resourceVolterraNetworkPolicyViewRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_network_policy_view.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicyView %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicyView %q: %s", d.Id(), err)
	}
	return setNetworkPolicyViewFields(client, d, resp)
}

func setNetworkPolicyViewFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkPolicyViewUpdate updates NetworkPolicyView resource
func resourceVolterraNetworkPolicyViewUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_network_policy_view.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_network_policy_view.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("egress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		egressRules := make([]*ves_io_schema_views_network_policy_view.NetworkPolicyRuleType, len(sl))
		updateSpec.EgressRules = egressRules
		for i, set := range sl {
			egressRules[i] = &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType{}

			egressRulesMapStrToI := set.(map[string]interface{})

			if v, ok := egressRulesMapStrToI["action"]; ok && !isIntfNil(v) {

				egressRules[i].Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

			}

			if w, ok := egressRulesMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				egressRules[i].Keys = ls
			}

			otherEndpointTypeFound := false

			if v, ok := egressRulesMapStrToI["any"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Any{}
					otherEndpointInt.Any = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_InsideEndpoints{}
					otherEndpointInt.InsideEndpoints = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_IpPrefixSet{}
				otherEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						otherEndpointInt.IpPrefixSet.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "ip_prefix_set"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := egressRulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_LabelSelector{}
				otherEndpointInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := egressRulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Namespace{}

				egressRules[i].OtherEndpoint = otherEndpointInt

				otherEndpointInt.Namespace = v.(string)

			}

			if v, ok := egressRulesMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_OutsideEndpoints{}
					otherEndpointInt.OutsideEndpoints = &ves_io_schema.Empty{}
					egressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := egressRulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_PrefixList{}
				otherEndpointInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				egressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.PrefixList.Prefixes = ls

					}

				}

			}

			if w, ok := egressRulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
				egressRules[i].RuleDescription = w.(string)
			}

			if w, ok := egressRulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
				egressRules[i].RuleName = w.(string)
			}

			trafficChoiceTypeFound := false

			if v, ok := egressRulesMapStrToI["all_tcp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTcpTraffic{}
					trafficChoiceInt.AllTcpTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["all_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTraffic{}
					trafficChoiceInt.AllTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["all_udp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllUdpTraffic{}
					trafficChoiceInt.AllUdpTraffic = &ves_io_schema.Empty{}
					egressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := egressRulesMapStrToI["applications"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Applications{}
				trafficChoiceInt.Applications = &ves_io_schema_views_network_policy_view.ApplicationsType{}
				egressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["applications"]; ok && !isIntfNil(v) {

						applicationsList := []ves_io_schema_views_network_policy_view.ApplicationEnumType{}
						for _, j := range v.([]interface{}) {
							applicationsList = append(applicationsList, ves_io_schema_views_network_policy_view.ApplicationEnumType(ves_io_schema_views_network_policy_view.ApplicationEnumType_value[j.(string)]))
						}
						trafficChoiceInt.Applications.Applications = applicationsList

					}

				}

			}

			if v, ok := egressRulesMapStrToI["protocol_port_range"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_ProtocolPortRange{}
				trafficChoiceInt.ProtocolPortRange = &ves_io_schema_views_network_policy_view.ProtocolPortType{}
				egressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["port_ranges"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						trafficChoiceInt.ProtocolPortRange.PortRanges = ls

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						trafficChoiceInt.ProtocolPortRange.Protocol = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("endpoint"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		endpoint := &ves_io_schema_views_network_policy_view.EndpointChoiceType{}
		updateSpec.Endpoint = endpoint
		for _, set := range sl {

			endpointMapStrToI := set.(map[string]interface{})

			endpointChoiceTypeFound := false

			if v, ok := endpointMapStrToI["any"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Any{}
					endpointChoiceInt.Any = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_InsideEndpoints{}
					endpointChoiceInt.InsideEndpoints = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["interface"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Interface{}
				endpointChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						endpointChoiceInt.Interface.Tenant = v.(string)
					}

				}

			}

			if v, ok := endpointMapStrToI["label_selector"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_LabelSelector{}
				endpointChoiceInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						endpointChoiceInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := endpointMapStrToI["namespace"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_Namespace{}

				endpoint.EndpointChoice = endpointChoiceInt

				endpointChoiceInt.Namespace = v.(string)

			}

			if v, ok := endpointMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true

				if v.(bool) {
					endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_OutsideEndpoints{}
					endpointChoiceInt.OutsideEndpoints = &ves_io_schema.Empty{}
					endpoint.EndpointChoice = endpointChoiceInt
				}

			}

			if v, ok := endpointMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !endpointChoiceTypeFound {

				endpointChoiceTypeFound = true
				endpointChoiceInt := &ves_io_schema_views_network_policy_view.EndpointChoiceType_PrefixList{}
				endpointChoiceInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				endpoint.EndpointChoice = endpointChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						endpointChoiceInt.PrefixList.Prefixes = ls

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ingressRules := make([]*ves_io_schema_views_network_policy_view.NetworkPolicyRuleType, len(sl))
		updateSpec.IngressRules = ingressRules
		for i, set := range sl {
			ingressRules[i] = &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType{}

			ingressRulesMapStrToI := set.(map[string]interface{})

			if v, ok := ingressRulesMapStrToI["action"]; ok && !isIntfNil(v) {

				ingressRules[i].Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

			}

			if w, ok := ingressRulesMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				ingressRules[i].Keys = ls
			}

			otherEndpointTypeFound := false

			if v, ok := ingressRulesMapStrToI["any"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Any{}
					otherEndpointInt.Any = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["inside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_InsideEndpoints{}
					otherEndpointInt.InsideEndpoints = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_IpPrefixSet{}
				otherEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						otherEndpointInt.IpPrefixSet.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "ip_prefix_set"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_LabelSelector{}
				otherEndpointInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.LabelSelector.Expressions = ls

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Namespace{}

				ingressRules[i].OtherEndpoint = otherEndpointInt

				otherEndpointInt.Namespace = v.(string)

			}

			if v, ok := ingressRulesMapStrToI["outside_endpoints"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true

				if v.(bool) {
					otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_OutsideEndpoints{}
					otherEndpointInt.OutsideEndpoints = &ves_io_schema.Empty{}
					ingressRules[i].OtherEndpoint = otherEndpointInt
				}

			}

			if v, ok := ingressRulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !otherEndpointTypeFound {

				otherEndpointTypeFound = true
				otherEndpointInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_PrefixList{}
				otherEndpointInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
				ingressRules[i].OtherEndpoint = otherEndpointInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						otherEndpointInt.PrefixList.Prefixes = ls

					}

				}

			}

			if w, ok := ingressRulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
				ingressRules[i].RuleDescription = w.(string)
			}

			if w, ok := ingressRulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
				ingressRules[i].RuleName = w.(string)
			}

			trafficChoiceTypeFound := false

			if v, ok := ingressRulesMapStrToI["all_tcp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTcpTraffic{}
					trafficChoiceInt.AllTcpTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["all_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllTraffic{}
					trafficChoiceInt.AllTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["all_udp_traffic"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true

				if v.(bool) {
					trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_AllUdpTraffic{}
					trafficChoiceInt.AllUdpTraffic = &ves_io_schema.Empty{}
					ingressRules[i].TrafficChoice = trafficChoiceInt
				}

			}

			if v, ok := ingressRulesMapStrToI["applications"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_Applications{}
				trafficChoiceInt.Applications = &ves_io_schema_views_network_policy_view.ApplicationsType{}
				ingressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["applications"]; ok && !isIntfNil(v) {

						applicationsList := []ves_io_schema_views_network_policy_view.ApplicationEnumType{}
						for _, j := range v.([]interface{}) {
							applicationsList = append(applicationsList, ves_io_schema_views_network_policy_view.ApplicationEnumType(ves_io_schema_views_network_policy_view.ApplicationEnumType_value[j.(string)]))
						}
						trafficChoiceInt.Applications.Applications = applicationsList

					}

				}

			}

			if v, ok := ingressRulesMapStrToI["protocol_port_range"]; ok && !isIntfNil(v) && !trafficChoiceTypeFound {

				trafficChoiceTypeFound = true
				trafficChoiceInt := &ves_io_schema_views_network_policy_view.NetworkPolicyRuleType_ProtocolPortRange{}
				trafficChoiceInt.ProtocolPortRange = &ves_io_schema_views_network_policy_view.ProtocolPortType{}
				ingressRules[i].TrafficChoice = trafficChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["port_ranges"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						trafficChoiceInt.ProtocolPortRange.PortRanges = ls

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						trafficChoiceInt.ProtocolPortRange.Protocol = v.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkPolicyView obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_network_policy_view.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkPolicyView: %s", err)
	}

	return resourceVolterraNetworkPolicyViewRead(d, meta)
}

func resourceVolterraNetworkPolicyViewDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_network_policy_view.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicyView %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicyView before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkPolicyView obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_network_policy_view.ObjectType, namespace, name)
}