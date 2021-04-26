package main

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"mocksensitive_resource": resourceMockSensitiveResource(),
		},
	}
}
