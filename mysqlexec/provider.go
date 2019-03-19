package mysqlexec

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

const(
	Endpoint = "MYSQLEXEC_ENDPOINT"
	Username = "MYSQLEXEC_USERNAME"
	Password = "MYSQLEXEC_PASSWORD"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(Endpoint, ""),
			},

			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(Username, ""),
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(Password, ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mysqlexec_script": resourceMysqlexecScript(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {

		endpoint := d.Get("endpoint").(string)
		if !strings.Contains(endpoint, "/") {
			endpoint += "/"
		}

		client := mySqlDbClient{endpoint: endpoint, username: d.Get("username").(string), password: d.Get("password").(string)}
		return &client, nil
	}
}