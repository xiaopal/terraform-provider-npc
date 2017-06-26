package npc

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NPC_ACCESS_KEY",
					"ACCESS_KEY",
				}, nil),
				Description: "Access Key",
			},

			"access_secret": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NPC_ACCESS_SECRET",
					"ACCESS_SECRET",
				}, nil),
				Description: "Access Secret",
			},
			"access_token": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NPC_ACCESS_TOKEN",
					"ACCESS_TOKEN",
				}, nil),
				Description: "OpenAPI token ",
			},
			"openapi_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NPC_OPENAPI_ENDPOINT",
					"OPENAPI_ENDPOINT",
				}, "https://open.c.163.com/"),
				Description: "Override the default OpenAPI endpoint URL",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"npc_service": resourceNpcService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		ApiEndpoint: d.Get("openapi_endpoint").(string),
		Credentials: &Credentials{
			Key:    d.Get("access_key").(string),
			Secret: d.Get("access_secret").(string),
			Token:  d.Get("access_token").(string),
		},
	}
	return config, nil
}
