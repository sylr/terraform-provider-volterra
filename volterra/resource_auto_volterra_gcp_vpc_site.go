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
	ves_io_schema_views_gcp_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

// resourceVolterraGcpVpcSite is implementation of Volterra's GcpVpcSite resources
func resourceVolterraGcpVpcSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraGcpVpcSiteCreate,
		Read:   resourceVolterraGcpVpcSiteRead,
		Update: resourceVolterraGcpVpcSiteUpdate,
		Delete: resourceVolterraGcpVpcSiteDelete,

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

			"assisted": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"cloud_credentials": {

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

			"gcp_region": {
				Type:     schema.TypeString,
				Required: true,
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

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"gcp_zone_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"inside_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"existing_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"new_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"primary_ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_name": {
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

												"existing_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"new_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"primary_ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_name": {
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

						"gcp_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
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

						"inside_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"existing_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"new_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
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

						"outside_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"existing_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"new_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
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

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"gcp_zone_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"existing_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"new_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"primary_ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_name": {
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

						"gcp_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"local_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"existing_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"new_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
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

			"voltstack_cluster": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"gcp_zone_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"existing_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"new_subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"primary_ipv4": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_name": {
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

						"gcp_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
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

						"site_local_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"existing_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"new_network": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
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

			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraGcpVpcSiteCreate creates GcpVpcSite resource
func resourceVolterraGcpVpcSiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_gcp_vpc_site.CreateSpecType{}
	createReq := &ves_io_schema_views_gcp_vpc_site.CreateRequest{
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

	deploymentTypeFound := false

	if v, ok := d.GetOk("assisted"); ok && !deploymentTypeFound {

		deploymentTypeFound = true

		if v.(bool) {
			deploymentInt := &ves_io_schema_views_gcp_vpc_site.CreateSpecType_Assisted{}
			deploymentInt.Assisted = &ves_io_schema.Empty{}
			createSpec.Deployment = deploymentInt
		}

	}

	if v, ok := d.GetOk("cloud_credentials"); ok && !deploymentTypeFound {

		deploymentTypeFound = true
		deploymentInt := &ves_io_schema_views_gcp_vpc_site.CreateSpecType_CloudCredentials{}
		deploymentInt.CloudCredentials = &ves_io_schema_views.ObjectRefType{}
		createSpec.Deployment = deploymentInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				deploymentInt.CloudCredentials.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				deploymentInt.CloudCredentials.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				deploymentInt.CloudCredentials.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("disk_size"); ok && !isIntfNil(v) {

		createSpec.DiskSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("gcp_region"); ok && !isIntfNil(v) {

		createSpec.GcpRegion =
			v.(string)
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
		siteTypeInt := &ves_io_schema_views_gcp_vpc_site.CreateSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.GCPVPCTwoInterfaceNodeType, len(sl))
				siteTypeInt.IngressEgressGw.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.GCPVPCTwoInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["gcp_zone_name"]; ok && !isIntfNil(w) {
						azNodes[i].GcpZoneName = w.(string)
					}

					if v, ok := azNodesMapStrToI["inside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						insideSubnet := &ves_io_schema_views.GCPVPCSubnetChoiceType{}
						azNodes[i].InsideSubnet = insideSubnet
						for _, set := range sl {

							insideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := insideSubnetMapStrToI["existing_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_ExistingSubnet{}
								choiceInt.ExistingSubnet = &ves_io_schema_views.GCPSubnetType{}
								insideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.ExistingSubnet.SubnetName = v.(string)
									}

								}

							}

							if v, ok := insideSubnetMapStrToI["new_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_NewSubnet{}
								choiceInt.NewSubnet = &ves_io_schema_views.GCPSubnetParamsType{}
								insideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.PrimaryIpv4 = v.(string)
									}

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.SubnetName = v.(string)
									}

								}

							}

						}

					}

					if v, ok := azNodesMapStrToI["outside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						outsideSubnet := &ves_io_schema_views.GCPVPCSubnetChoiceType{}
						azNodes[i].OutsideSubnet = outsideSubnet
						for _, set := range sl {

							outsideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := outsideSubnetMapStrToI["existing_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_ExistingSubnet{}
								choiceInt.ExistingSubnet = &ves_io_schema_views.GCPSubnetType{}
								outsideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.ExistingSubnet.SubnetName = v.(string)
									}

								}

							}

							if v, ok := outsideSubnetMapStrToI["new_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_NewSubnet{}
								choiceInt.NewSubnet = &ves_io_schema_views.GCPSubnetParamsType{}
								outsideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.PrimaryIpv4 = v.(string)
									}

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.SubnetName = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["gcp_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressEgressGw.GcpCertifiedHw = v.(string)
			}

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_GlobalNetworkList{}
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
					globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			if v, ok := cs["inside_network"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				insideNetwork := &ves_io_schema_views.GCPVPCNetworkChoiceType{}
				siteTypeInt.IngressEgressGw.InsideNetwork = insideNetwork
				for _, set := range sl {

					insideNetworkMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := insideNetworkMapStrToI["existing_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_ExistingNetwork{}
						choiceInt.ExistingNetwork = &ves_io_schema_views.GCPVPCNetworkType{}
						insideNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.ExistingNetwork.Name = v.(string)
							}

						}

					}

					if v, ok := insideNetworkMapStrToI["new_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_NewNetwork{}
						choiceInt.NewNetwork = &ves_io_schema_views.GCPVPCNetworkParamsType{}
						insideNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.NewNetwork.Name = v.(string)
							}

						}

					}

				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_InsideStaticRoutes{}
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
					insideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_ActiveNetworkPolicies{}
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
					networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			if v, ok := cs["outside_network"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				outsideNetwork := &ves_io_schema_views.GCPVPCNetworkChoiceType{}
				siteTypeInt.IngressEgressGw.OutsideNetwork = outsideNetwork
				for _, set := range sl {

					outsideNetworkMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := outsideNetworkMapStrToI["existing_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_ExistingNetwork{}
						choiceInt.ExistingNetwork = &ves_io_schema_views.GCPVPCNetworkType{}
						outsideNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.ExistingNetwork.Name = v.(string)
							}

						}

					}

					if v, ok := outsideNetworkMapStrToI["new_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_NewNetwork{}
						choiceInt.NewNetwork = &ves_io_schema_views.GCPVPCNetworkParamsType{}
						outsideNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.NewNetwork.Name = v.(string)
							}

						}

					}

				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_OutsideStaticRoutes{}
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
				servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_ActiveForwardProxyPolicies{}
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
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_gcp_vpc_site.CreateSpecType_IngressGw{}
		siteTypeInt.IngressGw = &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.GCPVPCOneInterfaceNodeType, len(sl))
				siteTypeInt.IngressGw.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.GCPVPCOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["gcp_zone_name"]; ok && !isIntfNil(w) {
						azNodes[i].GcpZoneName = w.(string)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.GCPVPCSubnetChoiceType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["existing_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_ExistingSubnet{}
								choiceInt.ExistingSubnet = &ves_io_schema_views.GCPSubnetType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.ExistingSubnet.SubnetName = v.(string)
									}

								}

							}

							if v, ok := localSubnetMapStrToI["new_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_NewSubnet{}
								choiceInt.NewSubnet = &ves_io_schema_views.GCPSubnetParamsType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.PrimaryIpv4 = v.(string)
									}

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.SubnetName = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["gcp_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressGw.GcpCertifiedHw = v.(string)
			}

			if v, ok := cs["local_network"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				localNetwork := &ves_io_schema_views.GCPVPCNetworkChoiceType{}
				siteTypeInt.IngressGw.LocalNetwork = localNetwork
				for _, set := range sl {

					localNetworkMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := localNetworkMapStrToI["existing_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_ExistingNetwork{}
						choiceInt.ExistingNetwork = &ves_io_schema_views.GCPVPCNetworkType{}
						localNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.ExistingNetwork.Name = v.(string)
							}

						}

					}

					if v, ok := localNetworkMapStrToI["new_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_NewNetwork{}
						choiceInt.NewNetwork = &ves_io_schema_views.GCPVPCNetworkParamsType{}
						localNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.NewNetwork.Name = v.(string)
							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("voltstack_cluster"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_gcp_vpc_site.CreateSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.GCPVPCOneInterfaceNodeType, len(sl))
				siteTypeInt.VoltstackCluster.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.GCPVPCOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["gcp_zone_name"]; ok && !isIntfNil(w) {
						azNodes[i].GcpZoneName = w.(string)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.GCPVPCSubnetChoiceType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["existing_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_ExistingSubnet{}
								choiceInt.ExistingSubnet = &ves_io_schema_views.GCPSubnetType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.ExistingSubnet.SubnetName = v.(string)
									}

								}

							}

							if v, ok := localSubnetMapStrToI["new_subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.GCPVPCSubnetChoiceType_NewSubnet{}
								choiceInt.NewSubnet = &ves_io_schema_views.GCPSubnetParamsType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.PrimaryIpv4 = v.(string)
									}

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.NewSubnet.SubnetName = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["gcp_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.VoltstackCluster.GcpCertifiedHw = v.(string)
			}

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_GlobalNetworkList{}
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
					globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_ActiveNetworkPolicies{}
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
					networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_OutsideStaticRoutes{}
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
				servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_ActiveForwardProxyPolicies{}
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
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["site_local_network"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				siteLocalNetwork := &ves_io_schema_views.GCPVPCNetworkChoiceType{}
				siteTypeInt.VoltstackCluster.SiteLocalNetwork = siteLocalNetwork
				for _, set := range sl {

					siteLocalNetworkMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := siteLocalNetworkMapStrToI["existing_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_ExistingNetwork{}
						choiceInt.ExistingNetwork = &ves_io_schema_views.GCPVPCNetworkType{}
						siteLocalNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.ExistingNetwork.Name = v.(string)
							}

						}

					}

					if v, ok := siteLocalNetworkMapStrToI["new_network"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.GCPVPCNetworkChoiceType_NewNetwork{}
						choiceInt.NewNetwork = &ves_io_schema_views.GCPVPCNetworkParamsType{}
						siteLocalNetwork.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								choiceInt.NewNetwork.Name = v.(string)
							}

						}

					}

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

	log.Printf("[DEBUG] Creating Volterra GcpVpcSite object with struct: %+v", createReq)

	createGcpVpcSiteResp, err := client.CreateObject(context.Background(), ves_io_schema_views_gcp_vpc_site.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating GcpVpcSite: %s", err)
	}
	d.SetId(createGcpVpcSiteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraGcpVpcSiteRead(d, meta)
}

func resourceVolterraGcpVpcSiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_gcp_vpc_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] GcpVpcSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra GcpVpcSite %q: %s", d.Id(), err)
	}
	return setGcpVpcSiteFields(client, d, resp)
}

func setGcpVpcSiteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraGcpVpcSiteUpdate updates GcpVpcSite resource
func resourceVolterraGcpVpcSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_gcp_vpc_site.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_gcp_vpc_site.ReplaceRequest{
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
		siteTypeInt := &ves_io_schema_views_gcp_vpc_site.ReplaceSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_GlobalNetworkList{}
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
					globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_InsideStaticRoutes{}
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
					insideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_ActiveNetworkPolicies{}
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
					networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_OutsideStaticRoutes{}
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
				servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_ActiveForwardProxyPolicies{}
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
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCIngressEgressGwReplaceType_NoForwardProxyPolicy{}
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
		siteTypeInt := &ves_io_schema_views_gcp_vpc_site.ReplaceSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_GlobalNetworkList{}
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
					globalNetworkChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_ActiveNetworkPolicies{}
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
					networkPolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_OutsideStaticRoutes{}
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
				servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_ActiveForwardProxyPolicies{}
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
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_gcp_vpc_site.GCPVPCVoltstackClusterReplaceType_NoForwardProxyPolicy{}
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

	log.Printf("[DEBUG] Updating Volterra GcpVpcSite obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_gcp_vpc_site.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating GcpVpcSite: %s", err)
	}

	return resourceVolterraGcpVpcSiteRead(d, meta)
}

func resourceVolterraGcpVpcSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_gcp_vpc_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] GcpVpcSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra GcpVpcSite before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra GcpVpcSite obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_gcp_vpc_site.ObjectType, namespace, name)
}
