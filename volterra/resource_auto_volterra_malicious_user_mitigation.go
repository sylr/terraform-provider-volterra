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
	ves_io_schema_malicious_user_mitigation "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/malicious_user_mitigation"
)

// resourceVolterraMaliciousUserMitigation is implementation of Volterra's MaliciousUserMitigation resources
func resourceVolterraMaliciousUserMitigation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraMaliciousUserMitigationCreate,
		Read:   resourceVolterraMaliciousUserMitigationRead,
		Update: resourceVolterraMaliciousUserMitigationUpdate,
		Delete: resourceVolterraMaliciousUserMitigationDelete,

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

			"captcha_challenge_settings": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"javascript_challenge_settings": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"js_script_delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"mitigation_type": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mitigation_action": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"alert_only": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"block_temporarily": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"captcha_challenge": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"javascript_challenge": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"none": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"threat_level": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"high": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"low": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"medium": {

													Type:     schema.TypeBool,
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

			"temporary_blocking_settings": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraMaliciousUserMitigationCreate creates MaliciousUserMitigation resource
func resourceVolterraMaliciousUserMitigationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_malicious_user_mitigation.CreateSpecType{}
	createReq := &ves_io_schema_malicious_user_mitigation.CreateRequest{
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

	if v, ok := d.GetOk("captcha_challenge_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		captchaChallengeSettings := &ves_io_schema_malicious_user_mitigation.CaptchaChallengeSettings{}
		createSpec.CaptchaChallengeSettings = captchaChallengeSettings
		for _, set := range sl {

			captchaChallengeSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := captchaChallengeSettingsMapStrToI["cookie_expiry"]; ok && !isIntfNil(w) {
				captchaChallengeSettings.CookieExpiry = w.(uint32)
			}

			if w, ok := captchaChallengeSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				captchaChallengeSettings.CustomPage = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("javascript_challenge_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		javascriptChallengeSettings := &ves_io_schema_malicious_user_mitigation.JavascriptChallengeSettings{}
		createSpec.JavascriptChallengeSettings = javascriptChallengeSettings
		for _, set := range sl {

			javascriptChallengeSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := javascriptChallengeSettingsMapStrToI["cookie_expiry"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.CookieExpiry = w.(uint32)
			}

			if w, ok := javascriptChallengeSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.CustomPage = w.(string)
			}

			if w, ok := javascriptChallengeSettingsMapStrToI["js_script_delay"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.JsScriptDelay = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("mitigation_type"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		mitigationType := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationType{}
		createSpec.MitigationType = mitigationType
		for _, set := range sl {

			mitigationTypeMapStrToI := set.(map[string]interface{})

			if v, ok := mitigationTypeMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule, len(sl))
				mitigationType.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["mitigation_action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mitigationAction := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction{}
						rules[i].MitigationAction = mitigationAction
						for _, set := range sl {

							mitigationActionMapStrToI := set.(map[string]interface{})

							mitigationActionTypeFound := false

							if v, ok := mitigationActionMapStrToI["alert_only"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_AlertOnly{}
									mitigationActionInt.AlertOnly = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["block_temporarily"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_BlockTemporarily{}
									mitigationActionInt.BlockTemporarily = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["captcha_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_CaptchaChallenge{}
									mitigationActionInt.CaptchaChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["javascript_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_JavascriptChallenge{}
									mitigationActionInt.JavascriptChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["none"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_None{}
									mitigationActionInt.None = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

						}

					}

					if v, ok := rulesMapStrToI["threat_level"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						threatLevel := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel{}
						rules[i].ThreatLevel = threatLevel
						for _, set := range sl {

							threatLevelMapStrToI := set.(map[string]interface{})

							threatLevelTypeFound := false

							if v, ok := threatLevelMapStrToI["high"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_High{}
									threatLevelInt.High = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["low"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Low{}
									threatLevelInt.Low = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["medium"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Medium{}
									threatLevelInt.Medium = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("temporary_blocking_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		temporaryBlockingSettings := &ves_io_schema_malicious_user_mitigation.TemporaryBlockingSettings{}
		createSpec.TemporaryBlockingSettings = temporaryBlockingSettings
		for _, set := range sl {

			temporaryBlockingSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := temporaryBlockingSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				temporaryBlockingSettings.CustomPage = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra MaliciousUserMitigation object with struct: %+v", createReq)

	createMaliciousUserMitigationResp, err := client.CreateObject(context.Background(), ves_io_schema_malicious_user_mitigation.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating MaliciousUserMitigation: %s", err)
	}
	d.SetId(createMaliciousUserMitigationResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraMaliciousUserMitigationRead(d, meta)
}

func resourceVolterraMaliciousUserMitigationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_malicious_user_mitigation.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] MaliciousUserMitigation %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra MaliciousUserMitigation %q: %s", d.Id(), err)
	}
	return setMaliciousUserMitigationFields(client, d, resp)
}

func setMaliciousUserMitigationFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraMaliciousUserMitigationUpdate updates MaliciousUserMitigation resource
func resourceVolterraMaliciousUserMitigationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_malicious_user_mitigation.ReplaceSpecType{}
	updateReq := &ves_io_schema_malicious_user_mitigation.ReplaceRequest{
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

	if v, ok := d.GetOk("captcha_challenge_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		captchaChallengeSettings := &ves_io_schema_malicious_user_mitigation.CaptchaChallengeSettings{}
		updateSpec.CaptchaChallengeSettings = captchaChallengeSettings
		for _, set := range sl {

			captchaChallengeSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := captchaChallengeSettingsMapStrToI["cookie_expiry"]; ok && !isIntfNil(w) {
				captchaChallengeSettings.CookieExpiry = w.(uint32)
			}

			if w, ok := captchaChallengeSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				captchaChallengeSettings.CustomPage = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("javascript_challenge_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		javascriptChallengeSettings := &ves_io_schema_malicious_user_mitigation.JavascriptChallengeSettings{}
		updateSpec.JavascriptChallengeSettings = javascriptChallengeSettings
		for _, set := range sl {

			javascriptChallengeSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := javascriptChallengeSettingsMapStrToI["cookie_expiry"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.CookieExpiry = w.(uint32)
			}

			if w, ok := javascriptChallengeSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.CustomPage = w.(string)
			}

			if w, ok := javascriptChallengeSettingsMapStrToI["js_script_delay"]; ok && !isIntfNil(w) {
				javascriptChallengeSettings.JsScriptDelay = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("mitigation_type"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		mitigationType := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationType{}
		updateSpec.MitigationType = mitigationType
		for _, set := range sl {

			mitigationTypeMapStrToI := set.(map[string]interface{})

			if v, ok := mitigationTypeMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule, len(sl))
				mitigationType.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["mitigation_action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mitigationAction := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction{}
						rules[i].MitigationAction = mitigationAction
						for _, set := range sl {

							mitigationActionMapStrToI := set.(map[string]interface{})

							mitigationActionTypeFound := false

							if v, ok := mitigationActionMapStrToI["alert_only"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_AlertOnly{}
									mitigationActionInt.AlertOnly = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["block_temporarily"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_BlockTemporarily{}
									mitigationActionInt.BlockTemporarily = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["captcha_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_CaptchaChallenge{}
									mitigationActionInt.CaptchaChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["javascript_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_JavascriptChallenge{}
									mitigationActionInt.JavascriptChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["none"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_None{}
									mitigationActionInt.None = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

						}

					}

					if v, ok := rulesMapStrToI["threat_level"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						threatLevel := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel{}
						rules[i].ThreatLevel = threatLevel
						for _, set := range sl {

							threatLevelMapStrToI := set.(map[string]interface{})

							threatLevelTypeFound := false

							if v, ok := threatLevelMapStrToI["high"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_High{}
									threatLevelInt.High = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["low"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Low{}
									threatLevelInt.Low = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["medium"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Medium{}
									threatLevelInt.Medium = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("temporary_blocking_settings"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		temporaryBlockingSettings := &ves_io_schema_malicious_user_mitigation.TemporaryBlockingSettings{}
		updateSpec.TemporaryBlockingSettings = temporaryBlockingSettings
		for _, set := range sl {

			temporaryBlockingSettingsMapStrToI := set.(map[string]interface{})

			if w, ok := temporaryBlockingSettingsMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				temporaryBlockingSettings.CustomPage = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra MaliciousUserMitigation obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_malicious_user_mitigation.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating MaliciousUserMitigation: %s", err)
	}

	return resourceVolterraMaliciousUserMitigationRead(d, meta)
}

func resourceVolterraMaliciousUserMitigationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_malicious_user_mitigation.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] MaliciousUserMitigation %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra MaliciousUserMitigation before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra MaliciousUserMitigation obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_malicious_user_mitigation.ObjectType, namespace, name)
}
