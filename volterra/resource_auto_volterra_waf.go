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
	ves_io_schema_waf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf"
)

// resourceVolterraWaf is implementation of Volterra's Waf resources
func resourceVolterraWaf() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraWafCreate,
		Read:   resourceVolterraWafRead,
		Update: resourceVolterraWafUpdate,
		Delete: resourceVolterraWafDelete,

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

			"app_profile": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cms": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"language": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"webserver": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"disabled_detection_tags": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detection_tag_type": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraWafCreate creates Waf resource
func resourceVolterraWafCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_waf.CreateSpecType{}
	createReq := &ves_io_schema_waf.CreateRequest{
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

	if v, ok := d.GetOk("app_profile"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		appProfile := &ves_io_schema_waf.AppProfile{}
		createSpec.AppProfile = appProfile
		for _, set := range sl {

			appProfileMapStrToI := set.(map[string]interface{})

			if v, ok := appProfileMapStrToI["cms"]; ok && !isIntfNil(v) {

				cmsList := []ves_io_schema_waf.ContentManagementSystemType{}
				for _, j := range v.([]interface{}) {
					cmsList = append(cmsList, ves_io_schema_waf.ContentManagementSystemType(ves_io_schema_waf.ContentManagementSystemType_value[j.(string)]))
				}
				appProfile.Cms = cmsList

			}

			if v, ok := appProfileMapStrToI["language"]; ok && !isIntfNil(v) {

				languageList := []ves_io_schema_waf.LanguageType{}
				for _, j := range v.([]interface{}) {
					languageList = append(languageList, ves_io_schema_waf.LanguageType(ves_io_schema_waf.LanguageType_value[j.(string)]))
				}
				appProfile.Language = languageList

			}

			if v, ok := appProfileMapStrToI["webserver"]; ok && !isIntfNil(v) {

				webserverList := []ves_io_schema_waf.WebServerType{}
				for _, j := range v.([]interface{}) {
					webserverList = append(webserverList, ves_io_schema_waf.WebServerType(ves_io_schema_waf.WebServerType_value[j.(string)]))
				}
				appProfile.Webserver = webserverList

			}

		}

	}

	if v, ok := d.GetOk("disabled_detection_tags"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		disabledDetectionTags := &ves_io_schema_waf.DisabledDetectionTags{}
		createSpec.DisabledDetectionTags = disabledDetectionTags
		for _, set := range sl {

			disabledDetectionTagsMapStrToI := set.(map[string]interface{})

			if v, ok := disabledDetectionTagsMapStrToI["detection_tag_type"]; ok && !isIntfNil(v) {

				detection_tag_typeList := []ves_io_schema_waf.DetectionTagType{}
				for _, j := range v.([]interface{}) {
					detection_tag_typeList = append(detection_tag_typeList, ves_io_schema_waf.DetectionTagType(ves_io_schema_waf.DetectionTagType_value[j.(string)]))
				}
				disabledDetectionTags.DetectionTagType = detection_tag_typeList

			}

		}

	}

	if v, ok := d.GetOk("mode"); ok && !isIntfNil(v) {

		createSpec.Mode = ves_io_schema.WafModeType(ves_io_schema.WafModeType_value[v.(string)])

	}

	log.Printf("[DEBUG] Creating Volterra Waf object with struct: %+v", createReq)

	createWafResp, err := client.CreateObject(context.Background(), ves_io_schema_waf.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Waf: %s", err)
	}
	d.SetId(createWafResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraWafRead(d, meta)
}

func resourceVolterraWafRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_waf.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Waf %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Waf %q: %s", d.Id(), err)
	}
	return setWafFields(client, d, resp)
}

func setWafFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraWafUpdate updates Waf resource
func resourceVolterraWafUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_waf.ReplaceSpecType{}
	updateReq := &ves_io_schema_waf.ReplaceRequest{
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

	if v, ok := d.GetOk("app_profile"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		appProfile := &ves_io_schema_waf.AppProfile{}
		updateSpec.AppProfile = appProfile
		for _, set := range sl {

			appProfileMapStrToI := set.(map[string]interface{})

			if v, ok := appProfileMapStrToI["cms"]; ok && !isIntfNil(v) {

				cmsList := []ves_io_schema_waf.ContentManagementSystemType{}
				for _, j := range v.([]interface{}) {
					cmsList = append(cmsList, ves_io_schema_waf.ContentManagementSystemType(ves_io_schema_waf.ContentManagementSystemType_value[j.(string)]))
				}
				appProfile.Cms = cmsList

			}

			if v, ok := appProfileMapStrToI["language"]; ok && !isIntfNil(v) {

				languageList := []ves_io_schema_waf.LanguageType{}
				for _, j := range v.([]interface{}) {
					languageList = append(languageList, ves_io_schema_waf.LanguageType(ves_io_schema_waf.LanguageType_value[j.(string)]))
				}
				appProfile.Language = languageList

			}

			if v, ok := appProfileMapStrToI["webserver"]; ok && !isIntfNil(v) {

				webserverList := []ves_io_schema_waf.WebServerType{}
				for _, j := range v.([]interface{}) {
					webserverList = append(webserverList, ves_io_schema_waf.WebServerType(ves_io_schema_waf.WebServerType_value[j.(string)]))
				}
				appProfile.Webserver = webserverList

			}

		}

	}

	if v, ok := d.GetOk("disabled_detection_tags"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		disabledDetectionTags := &ves_io_schema_waf.DisabledDetectionTags{}
		updateSpec.DisabledDetectionTags = disabledDetectionTags
		for _, set := range sl {

			disabledDetectionTagsMapStrToI := set.(map[string]interface{})

			if v, ok := disabledDetectionTagsMapStrToI["detection_tag_type"]; ok && !isIntfNil(v) {

				detection_tag_typeList := []ves_io_schema_waf.DetectionTagType{}
				for _, j := range v.([]interface{}) {
					detection_tag_typeList = append(detection_tag_typeList, ves_io_schema_waf.DetectionTagType(ves_io_schema_waf.DetectionTagType_value[j.(string)]))
				}
				disabledDetectionTags.DetectionTagType = detection_tag_typeList

			}

		}

	}

	if v, ok := d.GetOk("mode"); ok && !isIntfNil(v) {

		updateSpec.Mode = ves_io_schema.WafModeType(ves_io_schema.WafModeType_value[v.(string)])

	}

	log.Printf("[DEBUG] Updating Volterra Waf obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_waf.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Waf: %s", err)
	}

	return resourceVolterraWafRead(d, meta)
}

func resourceVolterraWafDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_waf.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Waf %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Waf before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Waf obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_waf.ObjectType, namespace, name)
}
