package mysqlexec

import (
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"os"
)

func resourceMysqlexecScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceMysqlexecScriptUpdate,
		Read:   resourceMysqlexecScriptRead,
		Update: resourceMysqlexecScriptUpdate,
		Delete: resourceMysqlexecScriptDelete,

		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
			"file_path": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

// Because FirewallRule can be created during current terraform apply Update method will always be executed
func resourceMysqlexecScriptRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMysqlexecScriptUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*mySqlDbClient)
	query := d.Get("query").(string)
	filePath := d.Get("file_path").(string)

	if query != "" {
		return execQuery(client, query)
	} else if filePath != "" {
		f, err := os.Open(filePath)
		defer f.Close()
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(f)
		return execQuery(client, string(content))
	}

	return resourceMysqlexecScriptRead(d, m)
}

func resourceMysqlexecScriptDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func execQuery(client *mySqlDbClient, query string) error {
	err := client.open()
	if err != nil {
		return err
	}
	defer client.close()
	return client.exec(query)
}