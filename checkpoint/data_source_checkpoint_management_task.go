package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
						"previous_task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the previous task in the execution sequence, on which this task depends.",
						},
						"progress_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"revert_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"start_time": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
						"last_update_time": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
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
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response message in JSON format",
			},
		},
	}
}

func dataSourceManagementTaskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("task_id"); ok {
		payload["task-id"] = v.(*schema.Set).List()
		payload["details-level"] = "full"
	}

	showTaskRes, err := client.ApiCall("show-task", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showTaskRes.Success {
		return fmt.Errorf("%s", showTaskRes.ErrorMsg)
	}

	task := showTaskRes.GetData()

	log.Println("Read Task - Show JSON = ", task)

	if task["tasks"] != nil {
		tasksList := task["tasks"].([]interface{})

		if len(tasksList) > 0 {

			var tasksListToReturn []map[string]interface{}

			for i := range tasksList {
				tasksMap := tasksList[i].(map[string]interface{})

				tasksMapToAdd := make(map[string]interface{})

				if v := tasksMap["start-time"]; v != nil {
					startTimeShow := v.(map[string]interface{})
					startTimeState := make(map[string]interface{})
					if v := startTimeShow["iso-8601"]; v != nil {
						startTimeState["iso_8601"] = v
					}
					if v := startTimeShow["posix"]; v != nil {
						startTimeState["posix"] = v
					}
					tasksMapToAdd["start_time"] = []interface{}{startTimeState}
				}
				if v := tasksMap["last-update-time"]; v != nil {
					lastUpdateTimeShow := v.(map[string]interface{})
					lastUpdateTimeState := make(map[string]interface{})
					if v := lastUpdateTimeShow["iso-8601"]; v != nil {
						lastUpdateTimeState["iso_8601"] = v
					}
					if v := lastUpdateTimeShow["posix"]; v != nil {
						lastUpdateTimeState["posix"] = v
					}
					tasksMapToAdd["last_update_time"] = []interface{}{lastUpdateTimeState}
				}
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
				if v := tasksMap["previous-task-id"]; v != nil {
					tasksMapToAdd["previous_task_id"] = v
				}
				if v := tasksMap["progress-description"]; v != nil {
					tasksMapToAdd["progress_description"] = v
				}
				if v := tasksMap["revert-status"]; v != nil {
					tasksMapToAdd["revert_status"] = v
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

	jsonResponse, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if jsonResponse != nil {
		_ = d.Set("response", string(jsonResponse))
	}

	d.SetId("show-task-" + acctest.RandString(10))

	return nil
}
