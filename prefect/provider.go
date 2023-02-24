package prefect

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PREFECT_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"prefect_work_queue": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"prefect_work_queues": dataSourceWorkQueues(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)

	var diags diag.Diagnostics

	return url, diags
}
