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
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_secret_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/secret_policy_rule"
)

// resourceVolterraSecretPolicyRule is implementation of Volterra's SecretPolicyRule resources
func resourceVolterraSecretPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSecretPolicyRuleCreate,
		Read:   resourceVolterraSecretPolicyRuleRead,
		Update: resourceVolterraSecretPolicyRuleUpdate,
		Delete: resourceVolterraSecretPolicyRuleDelete,

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

			"action": {
				Type:     schema.TypeString,
				Required: true,
			},

			"client_name": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"client_name_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"exact_values": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"regex_values": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"client_selector": {

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

			"label_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"keys": {

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
	}
}

// resourceVolterraSecretPolicyRuleCreate creates SecretPolicyRule resource
func resourceVolterraSecretPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_secret_policy_rule.CreateSpecType{}
	createReq := &ves_io_schema_secret_policy_rule.CreateRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		createSpec.Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

	}

	clientChoiceTypeFound := false

	if v, ok := d.GetOk("client_name"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.CreateSpecType_ClientName{}

		createSpec.ClientChoice = clientChoiceInt

		clientChoiceInt.ClientName = v.(string)

	}

	if v, ok := d.GetOk("client_name_matcher"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.CreateSpecType_ClientNameMatcher{}
		clientChoiceInt.ClientNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		createSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("client_selector"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.CreateSpecType_ClientSelector{}
		clientChoiceInt.ClientSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientSelector.Expressions = ls

			}

		}

	}

	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		createSpec.LabelMatcher = labelMatcher
		for _, set := range sl {

			labelMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				labelMatcher.Keys = ls
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra SecretPolicyRule object with struct: %+v", createReq)

	createSecretPolicyRuleResp, err := client.CreateObject(context.Background(), ves_io_schema_secret_policy_rule.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating SecretPolicyRule: %s", err)
	}
	d.SetId(createSecretPolicyRuleResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraSecretPolicyRuleRead(d, meta)
}

func resourceVolterraSecretPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_secret_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] SecretPolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra SecretPolicyRule %q: %s", d.Id(), err)
	}
	return setSecretPolicyRuleFields(client, d, resp)
}

func setSecretPolicyRuleFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraSecretPolicyRuleUpdate updates SecretPolicyRule resource
func resourceVolterraSecretPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_secret_policy_rule.ReplaceSpecType{}
	updateReq := &ves_io_schema_secret_policy_rule.ReplaceRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		updateSpec.Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

	}

	clientChoiceTypeFound := false

	if v, ok := d.GetOk("client_name"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.ReplaceSpecType_ClientName{}

		updateSpec.ClientChoice = clientChoiceInt

		clientChoiceInt.ClientName = v.(string)

	}

	if v, ok := d.GetOk("client_name_matcher"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.ReplaceSpecType_ClientNameMatcher{}
		clientChoiceInt.ClientNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		updateSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("client_selector"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_secret_policy_rule.ReplaceSpecType_ClientSelector{}
		clientChoiceInt.ClientSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientSelector.Expressions = ls

			}

		}

	}

	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		updateSpec.LabelMatcher = labelMatcher
		for _, set := range sl {

			labelMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				labelMatcher.Keys = ls
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra SecretPolicyRule obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_secret_policy_rule.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating SecretPolicyRule: %s", err)
	}

	return resourceVolterraSecretPolicyRuleRead(d, meta)
}

func resourceVolterraSecretPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_secret_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] SecretPolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra SecretPolicyRule before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra SecretPolicyRule obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_secret_policy_rule.ObjectType, namespace, name)
}