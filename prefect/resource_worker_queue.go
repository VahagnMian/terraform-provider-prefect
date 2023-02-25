package prefect

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ph "terraform-provider-prefect/prefect-client-go"
	"time"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
				Optional: true,
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
	}
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	url := m.(string)

	var diags diag.Diagnostics

	workQueue := ph.WorkQueue{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		IsPaused:         d.Get("is_paused").(bool),
		ConcurrencyLimit: d.Get("concurrency_limit").(int),
	}

	createdWorkQueue, err := workQueue.CreateWorkQueue(workQueue, url)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdWorkQueue.ID)
	resourceOrderRead(ctx, d, m)

	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	wqID := d.Id()

	var diags = diag.Diagnostics{}

	url := m.(string)

	workQueue, err := (&ph.WorkQueue{}).GetWorkQueue(wqID, url)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the attributes of the resource data
	if err := d.Set("name", workQueue.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", workQueue.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("is_paused", workQueue.IsPaused); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("concurrency_limit", workQueue.ConcurrencyLimit); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(workQueue.ID)

	return diags
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	url := m.(string)

	workQueue := ph.WorkQueue{
		//ID:          d.Get("id").(string),
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		IsPaused:         d.Get("is_paused").(bool),
		ConcurrencyLimit: d.Get("concurrency_limit").(int),
	}

	workQueue.UpdateWorkQueue(d.Get("id").(string), workQueue, url)

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceOrderRead(ctx, d, m)
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	url := m.(string)
	wqID := d.Id()
	workQueue, err := (&ph.WorkQueue{}).GetWorkQueue(wqID, url)

	var diags diag.Diagnostics

	workQueue.DeleteWorkQueue(wqID, url)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
