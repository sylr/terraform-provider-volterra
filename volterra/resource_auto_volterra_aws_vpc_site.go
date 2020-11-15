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
	ves_io_schema_views_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
)

// resourceVolterraAwsVpcSite is implementation of Volterra's AwsVpcSite resources
func resourceVolterraAwsVpcSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAwsVpcSiteCreate,
		Read:   resourceVolterraAwsVpcSiteRead,
		Update: resourceVolterraAwsVpcSiteUpdate,
		Delete: resourceVolterraAwsVpcSiteDelete,

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

			"aws_region": {
				Type:     schema.TypeString,
				Required: true,
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
				Required: true,
			},

			"nodes_per_az": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ingress_egress_gw": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_certified_hw": {
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

			"ingress_gw": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_certified_hw": {
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
										Type:     schema.TypeInt,
										Optional: true,
									},

									"local_subnet": {

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
					},
				},
			},

			"voltstack_cluster": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_certified_hw": {
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
										Type:     schema.TypeInt,
										Optional: true,
									},

									"local_subnet": {

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

			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vpc": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
					},
				},
			},
		},
	}
}

// resourceVolterraAwsVpcSiteCreate creates AwsVpcSite resource
func resourceVolterraAwsVpcSiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_aws_vpc_site.CreateSpecType{}
	createReq := &ves_io_schema_views_aws_vpc_site.CreateRequest{
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

	if v, ok := d.GetOk("aws_region"); ok && !isIntfNil(v) {

		createSpec.AwsRegion =
			v.(string)
	}

	deploymentTypeFound := false

	if v, ok := d.GetOk("assisted"); ok && !deploymentTypeFound {

		deploymentTypeFound = true

		if v.(bool) {
			deploymentInt := &ves_io_schema_views_aws_vpc_site.CreateSpecType_Assisted{}
			deploymentInt.Assisted = &ves_io_schema.Empty{}
			createSpec.Deployment = deploymentInt
		}

	}

	if v, ok := d.GetOk("aws_cred"); ok && !deploymentTypeFound {

		deploymentTypeFound = true
		deploymentInt := &ves_io_schema_views_aws_vpc_site.CreateSpecType_AwsCred{}
		deploymentInt.AwsCred = &ves_io_schema_views.ObjectRefType{}
		createSpec.Deployment = deploymentInt

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

	if v, ok := d.GetOk("disk_size"); ok && !isIntfNil(v) {

		createSpec.DiskSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("instance_type"); ok && !isIntfNil(v) {

		createSpec.InstanceType =
			v.(string)
	}

	if v, ok := d.GetOk("nodes_per_az"); ok && !isIntfNil(v) {

		createSpec.NodesPerAz =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		createSpec.OperatingSystemVersion =
			v.(string)
	}

	siteTypeTypeFound := false

	if v, ok := d.GetOk("ingress_egress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_aws_vpc_site.CreateSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["aws_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressEgressGw.AwsCertifiedHw = v.(string)
			}

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AWSVPCTwoInterfaceNodeType, len(sl))
				siteTypeInt.IngressEgressGw.AzNodes = azNodes
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

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt

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

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt

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

			if v, ok := cs["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt

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

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

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

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt

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

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_aws_vpc_site.CreateSpecType_IngressGw{}
		siteTypeInt.IngressGw = &ves_io_schema_views_aws_vpc_site.AWSVPCIngressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["aws_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressGw.AwsCertifiedHw = v.(string)
			}

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AWSVPCOneInterfaceNodeType, len(sl))
				siteTypeInt.IngressGw.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AWSVPCOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["aws_az_name"]; ok && !isIntfNil(w) {
						azNodes[i].AwsAzName = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(uint32)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.CloudSubnetType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["existing_subnet_id"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_ExistingSubnetId{}

								localSubnet.Choice = choiceInt

								choiceInt.ExistingSubnetId = v.(string)

							}

							if v, ok := localSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								localSubnet.Choice = choiceInt

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

		}

	}

	if v, ok := d.GetOk("voltstack_cluster"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_aws_vpc_site.CreateSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["aws_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.VoltstackCluster.AwsCertifiedHw = v.(string)
			}

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AWSVPCOneInterfaceNodeType, len(sl))
				siteTypeInt.VoltstackCluster.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AWSVPCOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["aws_az_name"]; ok && !isIntfNil(w) {
						azNodes[i].AwsAzName = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(uint32)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.CloudSubnetType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["existing_subnet_id"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_ExistingSubnetId{}

								localSubnet.Choice = choiceInt

								choiceInt.ExistingSubnetId = v.(string)

							}

							if v, ok := localSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.CloudSubnetType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								localSubnet.Choice = choiceInt

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

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt

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

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt

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

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

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

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt

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

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ssh_key"); ok && !isIntfNil(v) {

		createSpec.SshKey =
			v.(string)
	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		createSpec.VolterraSoftwareVersion =
			v.(string)
	}

	if v, ok := d.GetOk("vpc"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vpc := &ves_io_schema_views.AWSVPCchoiceType{}
		createSpec.Vpc = vpc
		for _, set := range sl {

			vpcMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := vpcMapStrToI["new_vpc"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views.AWSVPCchoiceType_NewVpc{}
				choiceInt.NewVpc = &ves_io_schema_views.AWSVPCParamsType{}
				vpc.Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["allocate_ipv6"]; ok && !isIntfNil(v) {

						choiceInt.NewVpc.AllocateIpv6 = v.(bool)
					}

					if v, ok := cs["name_tag"]; ok && !isIntfNil(v) {

						choiceInt.NewVpc.NameTag = v.(string)
					}

					if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

						choiceInt.NewVpc.PrimaryIpv4 = v.(string)
					}

				}

			}

			if v, ok := vpcMapStrToI["vpc_id"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views.AWSVPCchoiceType_VpcId{}

				vpc.Choice = choiceInt

				choiceInt.VpcId = v.(string)

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AwsVpcSite object with struct: %+v", createReq)

	createAwsVpcSiteResp, err := client.CreateObject(context.Background(), ves_io_schema_views_aws_vpc_site.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AwsVpcSite: %s", err)
	}
	d.SetId(createAwsVpcSiteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAwsVpcSiteRead(d, meta)
}

func resourceVolterraAwsVpcSiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_aws_vpc_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AwsVpcSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AwsVpcSite %q: %s", d.Id(), err)
	}
	return setAwsVpcSiteFields(client, d, resp)
}

func setAwsVpcSiteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAwsVpcSiteUpdate updates AwsVpcSite resource
func resourceVolterraAwsVpcSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_aws_vpc_site.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_aws_vpc_site.ReplaceRequest{
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

	if v, ok := d.GetOk("nodes_per_az"); ok && !isIntfNil(v) {

		updateSpec.NodesPerAz =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		updateSpec.OperatingSystemVersion =
			v.(string)
	}

	siteTypeTypeFound := false

	if v, ok := d.GetOk("ingress_egress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_aws_vpc_site.ReplaceSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt

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

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt

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

			if v, ok := cs["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt

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

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

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

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt

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

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCIngressEgressGwReplaceType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		_ = v
	}

	if v, ok := d.GetOk("voltstack_cluster"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_aws_vpc_site.ReplaceSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt

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

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt

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

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

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

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt

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

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_aws_vpc_site.AWSVPCVoltstackClusterReplaceType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareVersion =
			v.(string)
	}

	log.Printf("[DEBUG] Updating Volterra AwsVpcSite obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_aws_vpc_site.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AwsVpcSite: %s", err)
	}

	return resourceVolterraAwsVpcSiteRead(d, meta)
}

func resourceVolterraAwsVpcSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_aws_vpc_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AwsVpcSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AwsVpcSite before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AwsVpcSite obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_aws_vpc_site.ObjectType, namespace, name)
}
