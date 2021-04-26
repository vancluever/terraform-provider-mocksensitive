package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMockSensitiveResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
			d.SetId("foo")
			return nil
		},
		ReadContext:   func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		UpdateContext: func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		DeleteContext: func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },

		Schema: map[string]*schema.Schema{
			"sensitive_string": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sensitive_number": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sensitive_boolean": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sensitive_map": {
				Type:      schema.TypeMap,
				Required:  true,
				Sensitive: true,
				Elem:      schema.TypeString,
			},
			"sensitive_object": {
				Type:      schema.TypeList,
				Required:  true,
				Sensitive: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nested": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"nonsensitive_object": {
				Type:      schema.TypeList,
				Required:  true,
				Sensitive: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sensitive_nested": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
