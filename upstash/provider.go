package upstash

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Environment Variables
const (
	UPSTASH_EMAIL_ADDRESS = "UPSTASH_EMAIL_ADDRESS"
	UPSTASH_API_KEY       = "UPSTASH_API_KEY"
)

// Provider keys
const (
	email   = "email"
	api_key = "api_key"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			email: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(UPSTASH_EMAIL_ADDRESS, nil),
				Description: "Your Upstash email address",
			},
			api_key: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(UPSTASH_API_KEY, nil),
				Description: "Your Upstash API_KEY associated with your defined email address",
			},
		},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	email := d.Get("email").(string)
	apiKey := d.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (email != "") && (apiKey != "") {
		c, err := NewClient(nil, &email, &apiKey)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	c, err := NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
