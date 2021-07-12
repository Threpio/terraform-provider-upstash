package upstash

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-upstash/upstash/client"
)

//{"database_id":"92fb6455-692b-4b8f-b588-a2e85588079a","database_name":"test-tha2","region":"eu-west-1","tls":false,"api_enabled":true}]
func dataSourceDatabases() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDatabasesRead,
		Schema: map[string]*schema.Schema{
			"databases": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tls": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"api_enabled": &schema.Schema{
						Type:     schema.TypeBool,
						Computed: true,
						},
					},
				},
			},
		},
	}
}


func dataSourceDatabasesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// I actually don't know what this line does properly
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	databaseId := strconv.Itoa(d.Get("database_id").(int))

	databases, err := c.GetDatabases()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("databases", databases); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(databaseId)

	// TODO: Handle when no resources are returned (AKA resources manually deleted)

	return diags
}
