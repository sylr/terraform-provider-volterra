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
	ves_io_schema_network_firewall "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_firewall"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

// resourceVolterraAwsTgwSite is implementation of Volterra's AwsTgwSite resources
func resourceVolterraAwsTgwSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAwsTgwSiteCreate,
		Read:   resourceVolterraAwsTgwSiteRead,
		Update: resourceVolterraAwsTgwSiteUpdate,
		Delete: resourceVolterraAwsTgwSiteDelete,

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

			"aws_parameters": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"aws_region": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aws_az_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disk_size": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"inside_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"existing_subnet_id": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"outside_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"existing_subnet_id": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"ipv6": {
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
							},
						},

						"assisted": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"aws_cred": {

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

						"disk_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"instance_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"nodes_per_az": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"new_vpc": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allocate_ipv6": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name_tag": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"primary_ipv4": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"vpc_id": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"ssh_key": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"existing_tgw": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"tgw_asn": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"tgw_id": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"volterra_site_asn": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"new_tgw": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"system_generated": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"user_assigned": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"tgw_asn": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"volterra_site_asn": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tgw_security": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"active_network_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {

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

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_forward_proxy_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {

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

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"vn_config": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"global_network_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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
														},
													},
												},

												"slo_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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
														},
													},
												},

												"disable_forward_proxy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_forward_proxy": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"max_connect_attempts": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"white_listed_ports": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {

																Type: schema.TypeList,

																Optional: true,
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
								},
							},
						},

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_static_routes": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {

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

																		"nexthop_address": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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
													},
												},

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"no_inside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {

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

																		"nexthop_address": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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
													},
												},

												"simple_static_route": {

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
				},
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vpc_attachments": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"vpc_list": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"labels": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"vpc_id": {
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

// resourceVolterraAwsTgwSiteCreate creates AwsTgwSite resource
func resourceVolterraAwsTgwSiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_aws_tgw_site.CreateSpecType{}
	createReq := &ves_io_schema_views_aws_tgw_site.CreateRequest{
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

	if v, ok := d.GetOk("aws_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		awsParameters := &ves_io_schema_views_aws_tgw_site.ServicesVPCType{}
		createSpec.AwsParameters = awsParameters
		for _, set := range sl {

			awsParametersMapStrToI := set.(map[string]interface{})

			if w, ok := awsParametersMapStrToI["aws_certified_hw"]; ok && !isIntfNil(w) {
				awsParameters.AwsCertifiedHw = w.(string)
			}

			if w, ok := awsParametersMapStrToI["aws_region"]; ok && !isIntfNil(w) {
				awsParameters.AwsRegion = w.(string)
			}

			if v, ok := awsParametersMapStrToI["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AWSVPCTwoInterfaceNodeType, len(sl))
				awsParameters.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AWSVPCTwoInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["aws_az_name"]; ok && !isIntfNil(w) {
						azNodes[i].AwsAzName = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(string)
					}

					if v, ok := azNodesMapStrToI["inside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						insideSubnet := &ves_io_schema_views.CloudSubnetType{}
						azNodes[i].InsideSubnet = insideSubnet
						for _, set := range sl {

							insideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := insideSubnetMapStrToI["existing_subnet_id"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_ExistingSubnetId{}

								insideSubnet.Choice = choiceInt

								choiceInt.ExistingSubnetId = v.(string)

							}

							if v, ok := insideSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								insideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

					if v, ok := azNodesMapStrToI["outside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						outsideSubnet := &ves_io_schema_views.CloudSubnetType{}
						azNodes[i].OutsideSubnet = outsideSubnet
						for _, set := range sl {

							outsideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := outsideSubnetMapStrToI["existing_subnet_id"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_ExistingSubnetId{}

								outsideSubnet.Choice = choiceInt

								choiceInt.ExistingSubnetId = v.(string)

							}

							if v, ok := outsideSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								outsideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

				}

			}

			deploymentTypeFound := false

			if v, ok := awsParametersMapStrToI["assisted"]; ok && !isIntfNil(v) && !deploymentTypeFound {

				deploymentTypeFound = true

				if v.(bool) {
					deploymentInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_Assisted{}
					deploymentInt.Assisted = &ves_io_schema.Empty{}
					awsParameters.Deployment = deploymentInt
				}

			}

			if v, ok := awsParametersMapStrToI["aws_cred"]; ok && !isIntfNil(v) && !deploymentTypeFound {

				deploymentTypeFound = true
				deploymentInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_AwsCred{}
				deploymentInt.AwsCred = &ves_io_schema_views.ObjectRefType{}
				awsParameters.Deployment = deploymentInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						deploymentInt.AwsCred.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						deploymentInt.AwsCred.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						deploymentInt.AwsCred.Tenant = v.(string)
					}

				}

			}

			if w, ok := awsParametersMapStrToI["disk_size"]; ok && !isIntfNil(w) {
				awsParameters.DiskSize = w.(uint32)
			}

			if w, ok := awsParametersMapStrToI["instance_type"]; ok && !isIntfNil(w) {
				awsParameters.InstanceType = w.(string)
			}

			if w, ok := awsParametersMapStrToI["nodes_per_az"]; ok && !isIntfNil(w) {
				awsParameters.NodesPerAz = w.(uint32)
			}

			serviceVpcChoiceTypeFound := false

			if v, ok := awsParametersMapStrToI["new_vpc"]; ok && !isIntfNil(v) && !serviceVpcChoiceTypeFound {

				serviceVpcChoiceTypeFound = true
				serviceVpcChoiceInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_NewVpc{}
				serviceVpcChoiceInt.NewVpc = &ves_io_schema_views.AWSVPCParamsType{}
				awsParameters.ServiceVpcChoice = serviceVpcChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["allocate_ipv6"]; ok && !isIntfNil(v) {

						serviceVpcChoiceInt.NewVpc.AllocateIpv6 = v.(bool)
					}

					if v, ok := cs["name_tag"]; ok && !isIntfNil(v) {

						serviceVpcChoiceInt.NewVpc.NameTag = v.(string)
					}

					if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

						serviceVpcChoiceInt.NewVpc.PrimaryIpv4 = v.(string)
					}

				}

			}

			if v, ok := awsParametersMapStrToI["vpc_id"]; ok && !isIntfNil(v) && !serviceVpcChoiceTypeFound {

				serviceVpcChoiceTypeFound = true
				serviceVpcChoiceInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_VpcId{}

				awsParameters.ServiceVpcChoice = serviceVpcChoiceInt

				serviceVpcChoiceInt.VpcId = v.(string)

			}

			if w, ok := awsParametersMapStrToI["ssh_key"]; ok && !isIntfNil(w) {
				awsParameters.SshKey = w.(string)
			}

			tgwChoiceTypeFound := false

			if v, ok := awsParametersMapStrToI["existing_tgw"]; ok && !isIntfNil(v) && !tgwChoiceTypeFound {

				tgwChoiceTypeFound = true
				tgwChoiceInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_ExistingTgw{}
				tgwChoiceInt.ExistingTgw = &ves_io_schema_views_aws_tgw_site.ExistingTGWType{}
				awsParameters.TgwChoice = tgwChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["tgw_asn"]; ok && !isIntfNil(v) {

						tgwChoiceInt.ExistingTgw.TgwAsn = uint32(v.(int))
					}

					if v, ok := cs["tgw_id"]; ok && !isIntfNil(v) {

						tgwChoiceInt.ExistingTgw.TgwId = v.(string)
					}

					if v, ok := cs["volterra_site_asn"]; ok && !isIntfNil(v) {

						tgwChoiceInt.ExistingTgw.VolterraSiteAsn = uint32(v.(int))
					}

				}

			}

			if v, ok := awsParametersMapStrToI["new_tgw"]; ok && !isIntfNil(v) && !tgwChoiceTypeFound {

				tgwChoiceTypeFound = true
				tgwChoiceInt := &ves_io_schema_views_aws_tgw_site.ServicesVPCType_NewTgw{}
				tgwChoiceInt.NewTgw = &ves_io_schema_views_aws_tgw_site.TGWParamsType{}
				awsParameters.TgwChoice = tgwChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					asnChoiceTypeFound := false

					if v, ok := cs["system_generated"]; ok && !isIntfNil(v) && !asnChoiceTypeFound {

						asnChoiceTypeFound = true

						if v.(bool) {
							asnChoiceInt := &ves_io_schema_views_aws_tgw_site.TGWParamsType_SystemGenerated{}
							asnChoiceInt.SystemGenerated = &ves_io_schema.Empty{}
							tgwChoiceInt.NewTgw.AsnChoice = asnChoiceInt
						}

					}

					if v, ok := cs["user_assigned"]; ok && !isIntfNil(v) && !asnChoiceTypeFound {

						asnChoiceTypeFound = true
						asnChoiceInt := &ves_io_schema_views_aws_tgw_site.TGWParamsType_UserAssigned{}
						asnChoiceInt.UserAssigned = &ves_io_schema_views_aws_tgw_site.TGWAssignedASNType{}
						tgwChoiceInt.NewTgw.AsnChoice = asnChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["tgw_asn"]; ok && !isIntfNil(v) {

								asnChoiceInt.UserAssigned.TgwAsn = uint32(v.(int))
							}

							if v, ok := cs["volterra_site_asn"]; ok && !isIntfNil(v) {

								asnChoiceInt.UserAssigned.VolterraSiteAsn = uint32(v.(int))
							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		createSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("tgw_security"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tgwSecurity := &ves_io_schema_views_aws_tgw_site.SecurityConfigType{}
		createSpec.TgwSecurity = tgwSecurity
		for _, set := range sl {

			tgwSecurityMapStrToI := set.(map[string]interface{})

			networkPolicyChoiceTypeFound := false

			if v, ok := tgwSecurityMapStrToI["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				tgwSecurity.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := tgwSecurityMapStrToI["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					tgwSecurity.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := tgwSecurityMapStrToI["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := tgwSecurityMapStrToI["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := tgwSecurityMapStrToI["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("vn_config"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vnConfig := &ves_io_schema_views_aws_tgw_site.VnConfiguration{}
		createSpec.VnConfig = vnConfig
		for _, set := range sl {

			vnConfigMapStrToI := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				vnConfig.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := vnConfigMapStrToI["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					vnConfig.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				vnConfig.InsideStaticRouteChoice = insideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						insideStaticRouteChoiceInt.InsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			if v, ok := vnConfigMapStrToI["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					vnConfig.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					vnConfig.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := vnConfigMapStrToI["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				vnConfig.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		createSpec.VolterraSoftwareVersion =
			v.(string)
	}

	if v, ok := d.GetOk("vpc_attachments"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vpcAttachments := &ves_io_schema_views_aws_tgw_site.VPCAttachmentListType{}
		createSpec.VpcAttachments = vpcAttachments
		for _, set := range sl {

			vpcAttachmentsMapStrToI := set.(map[string]interface{})

			if v, ok := vpcAttachmentsMapStrToI["vpc_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				vpcList := make([]*ves_io_schema_views_aws_tgw_site.VPCAttachmentType, len(sl))
				vpcAttachments.VpcList = vpcList
				for i, set := range sl {
					vpcList[i] = &ves_io_schema_views_aws_tgw_site.VPCAttachmentType{}

					vpcListMapStrToI := set.(map[string]interface{})

					if w, ok := vpcListMapStrToI["labels"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						vpcList[i].Labels = ms
					}

					if w, ok := vpcListMapStrToI["vpc_id"]; ok && !isIntfNil(w) {
						vpcList[i].VpcId = w.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AwsTgwSite object with struct: %+v", createReq)

	createAwsTgwSiteResp, err := client.CreateObject(context.Background(), ves_io_schema_views_aws_tgw_site.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AwsTgwSite: %s", err)
	}
	d.SetId(createAwsTgwSiteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAwsTgwSiteRead(d, meta)
}

func resourceVolterraAwsTgwSiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_aws_tgw_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AwsTgwSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AwsTgwSite %q: %s", d.Id(), err)
	}
	return setAwsTgwSiteFields(client, d, resp)
}

func setAwsTgwSiteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAwsTgwSiteUpdate updates AwsTgwSite resource
func resourceVolterraAwsTgwSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_aws_tgw_site.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_aws_tgw_site.ReplaceRequest{
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

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		updateSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("tgw_security"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tgwSecurity := &ves_io_schema_views_aws_tgw_site.SecurityConfigType{}
		updateSpec.TgwSecurity = tgwSecurity
		for _, set := range sl {

			tgwSecurityMapStrToI := set.(map[string]interface{})

			networkPolicyChoiceTypeFound := false

			if v, ok := tgwSecurityMapStrToI["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				tgwSecurity.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := tgwSecurityMapStrToI["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					tgwSecurity.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := tgwSecurityMapStrToI["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := tgwSecurityMapStrToI["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := tgwSecurityMapStrToI["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_tgw_site.SecurityConfigType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					tgwSecurity.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("vn_config"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vnConfig := &ves_io_schema_views_aws_tgw_site.VnConfiguration{}
		updateSpec.VnConfig = vnConfig
		for _, set := range sl {

			vnConfigMapStrToI := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				vnConfig.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := vnConfigMapStrToI["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					vnConfig.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				vnConfig.InsideStaticRouteChoice = insideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						insideStaticRouteChoiceInt.InsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			if v, ok := vnConfigMapStrToI["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					vnConfig.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := vnConfigMapStrToI["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					vnConfig.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := vnConfigMapStrToI["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_tgw_site.VnConfiguration_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				vnConfig.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareVersion =
			v.(string)
	}

	if v, ok := d.GetOk("vpc_attachments"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vpcAttachments := &ves_io_schema_views_aws_tgw_site.VPCAttachmentListType{}
		updateSpec.VpcAttachments = vpcAttachments
		for _, set := range sl {

			vpcAttachmentsMapStrToI := set.(map[string]interface{})

			if v, ok := vpcAttachmentsMapStrToI["vpc_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				vpcList := make([]*ves_io_schema_views_aws_tgw_site.VPCAttachmentType, len(sl))
				vpcAttachments.VpcList = vpcList
				for i, set := range sl {
					vpcList[i] = &ves_io_schema_views_aws_tgw_site.VPCAttachmentType{}

					vpcListMapStrToI := set.(map[string]interface{})

					if w, ok := vpcListMapStrToI["labels"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						vpcList[i].Labels = ms
					}

					if w, ok := vpcListMapStrToI["vpc_id"]; ok && !isIntfNil(w) {
						vpcList[i].VpcId = w.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AwsTgwSite obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_aws_tgw_site.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AwsTgwSite: %s", err)
	}

	return resourceVolterraAwsTgwSiteRead(d, meta)
}

func resourceVolterraAwsTgwSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_aws_tgw_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AwsTgwSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AwsTgwSite before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AwsTgwSite obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_aws_tgw_site.ObjectType, namespace, name)
}