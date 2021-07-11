package upstash

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"terraform-provider-upstash/upstash"
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

func Provider() terraform.ResourceProvider {
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
		ResourcesMap: map[string]*schema.Resource{
			"example_item": resourceItem(),
		},
		ConfigureContextFunc: providerConfigure,

	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	email := d.Get("email").(string)
	apiKey := d.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (email != "") && (apiKey != "") {
		c, err := upstash.NewClient(nil, &email, &apiKey)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	c, err := upstash.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}