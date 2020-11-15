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
	ves_io_schema_virtual_k8s "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_k8s"
)

// resourceVolterraVirtualK8S is implementation of Volterra's VirtualK8S resources
func resourceVolterraVirtualK8S() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraVirtualK8SCreate,
		Read:   resourceVolterraVirtualK8SRead,
		Update: resourceVolterraVirtualK8SUpdate,
		Delete: resourceVolterraVirtualK8SDelete,

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

			"vsite_refs": {

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
	}
}

// resourceVolterraVirtualK8SCreate creates VirtualK8S resource
func resourceVolterraVirtualK8SCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_virtual_k8s.CreateSpecType{}
	createReq := &ves_io_schema_virtual_k8s.CreateRequest{
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

	if v, ok := d.GetOk("vsite_refs"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		vsiteRefsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.VsiteRefs = vsiteRefsInt
		for i, ps := range sl {

			vrMapToStrVal := ps.(map[string]interface{})
			vsiteRefsInt[i] = &ves_io_schema.ObjectRefType{}

			vsiteRefsInt[i].Kind = "virtual_site"

			if v, ok := vrMapToStrVal["name"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Name = v.(string)
			}

			if v, ok := vrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Namespace = v.(string)
			}

			if v, ok := vrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Tenant = v.(string)
			}

			if v, ok := vrMapToStrVal["uid"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra VirtualK8S object with struct: %+v", createReq)

	createVirtualK8SResp, err := client.CreateObject(context.Background(), ves_io_schema_virtual_k8s.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating VirtualK8S: %s", err)
	}
	d.SetId(createVirtualK8SResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraVirtualK8SRead(d, meta)
}

func resourceVolterraVirtualK8SRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_virtual_k8s.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualK8S %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualK8S %q: %s", d.Id(), err)
	}
	return setVirtualK8SFields(client, d, resp)
}

func setVirtualK8SFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraVirtualK8SUpdate updates VirtualK8S resource
func resourceVolterraVirtualK8SUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_virtual_k8s.ReplaceSpecType{}
	updateReq := &ves_io_schema_virtual_k8s.ReplaceRequest{
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

	if v, ok := d.GetOk("vsite_refs"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		vsiteRefsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.VsiteRefs = vsiteRefsInt
		for i, ps := range sl {

			vrMapToStrVal := ps.(map[string]interface{})
			vsiteRefsInt[i] = &ves_io_schema.ObjectRefType{}

			vsiteRefsInt[i].Kind = "virtual_site"

			if v, ok := vrMapToStrVal["name"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Name = v.(string)
			}

			if v, ok := vrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Namespace = v.(string)
			}

			if v, ok := vrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Tenant = v.(string)
			}

			if v, ok := vrMapToStrVal["uid"]; ok && !isIntfNil(v) {
				vsiteRefsInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra VirtualK8S obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_virtual_k8s.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating VirtualK8S: %s", err)
	}

	return resourceVolterraVirtualK8SRead(d, meta)
}

func resourceVolterraVirtualK8SDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_virtual_k8s.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualK8S %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualK8S before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra VirtualK8S obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_virtual_k8s.ObjectType, namespace, name)
}