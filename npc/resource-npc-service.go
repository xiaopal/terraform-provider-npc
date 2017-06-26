package npc

import (
	"log"
	"time"
	
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNpcService() *schema.Resource {
	return &schema.Resource{
		Create: resourceNpcServiceCreate,
		Read:   resourceNpcServiceRead,
		Update: resourceNpcServiceUpdate,
		Delete: resourceNpcServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Default:  "default",
				Optional: true,
				ForceNew: true,
			},
			"stateful": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
				ForceNew: true,
			},
			"spec": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lan_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceNpcServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Read instance: id=%s, openapi_token=%s", d.Id(), meta.(*schema.ResourceData).Get("openapi_token"))
	return nil
}
func resourceNpcServiceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(resource.PrefixedUniqueId("npc-"))
	log.Printf("[DEBUG] Create the instance: id=%s, openapi_token=%s", d.Id(), meta.(*schema.ResourceData).Get("openapi_token"))
	return resourceNpcServiceUpdate(d, meta)
}
func resourceNpcServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	meta.(*schema.ResourceData).Set("openapi_token", "Token2")
	log.Printf("[DEBUG] Update the instance: id=%s, openapi_token=%s", d.Id(), meta.(*schema.ResourceData).Get("openapi_token"))

	return resourceNpcServiceRead(d, meta)
}
func resourceNpcServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Delete the instance: id=%s, openapi_token=%s", d.Id(), meta.(*schema.ResourceData).Get("openapi_token"))

	d.SetId("")
	return nil
}
