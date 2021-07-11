package upstash

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
							Type: schema.TypeString,
							Computed: true,
						},
						"region": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDatabasesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databases", "https://api.upstash.com/v1"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	databases := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&databases)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("databases", database); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	// TODO: Handle when no resources are returned (AKA resources manually deleted)

	return diags
}