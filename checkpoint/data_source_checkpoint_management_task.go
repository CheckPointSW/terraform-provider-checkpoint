package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTaskRead,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tasks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task status.",
						},
						"progress_percentage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
						},
						"suppressed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
						},
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Asynchronous task unique identifier. Use show-task command to check the progress of the task.",
						},
						"task_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments string.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementTaskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("task_id"); ok {
		payload["task-id"] = v.(*schema.Set).List()
	}

	showTaskRes, err := client.ApiCall("show-task", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTaskRes.Success {
		return fmt.Errorf(showTaskRes.ErrorMsg)
	}

	task := showTaskRes.GetData()

	log.Println("Read Task - Show JSON = ", task)

	d.SetId("show-task-" + acctest.RandString(10))

	if task["tasks"] != nil {
		tasksList := task["tasks"].([]interface{})

		if len(tasksList) > 0 {

			var tasksListToReturn []map[string]interface{}

			for i := range tasksList {
				tasksMap := tasksList[i].(map[string]interface{})

				tasksMapToAdd := make(map[string]interface{})

				if v, _ := tasksMap["comments"]; v != nil {
					tasksMapToAdd["comments"] = v
				}
				if v, _ := tasksMap["task-name"]; v != nil {
					tasksMapToAdd["task_name"] = v
				}
				if v, _ := tasksMap["task-id"]; v != nil {
					tasksMapToAdd["task_id"] = v
				}
				if v, _ := tasksMap["status"]; v != nil {
					tasksMapToAdd["status"] = v
				}
				if v, _ := tasksMap["progress-percentage"]; v != nil {
					tasksMapToAdd["progress_percentage"] = v
				}
				if v, _ := tasksMap["suppressed"]; v != nil {
					tasksMapToAdd["suppressed"] = v
				}
				tasksListToReturn = append(tasksListToReturn, tasksMapToAdd)
			}
			_ = d.Set("tasks", tasksListToReturn)
		} else {
			_ = d.Set("tasks", tasksList)
		}
	} else {
		_ = d.Set("tasks", nil)
	}

	return nil
}
