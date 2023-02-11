package prefect

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strconv"
	"time"
)

func dataSourceWorkQueues() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkQueuesRead,
		Schema: map[string]*schema.Schema{
			"work_queues": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_paused": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"concurrency_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_polled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWorkQueuesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	data := []byte("{\n    \"limit\": 10\n}")
	payload := bytes.NewReader(data)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/work_queues/filter", "http://localhost:4200"), payload)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	workQueues := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&workQueues)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("work_queues", workQueues); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
