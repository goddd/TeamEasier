package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goddd/TeamEasier/request"
	"github.com/goddd/TeamEasier/structs"
)

func Task_add(task structs.Todo_item) (string, error) {

	if task.Task.Todo_list_id == 0 {

		return "", errors.New("No ID Passed")
	}

	json, _ := json.Marshal(task)
	jsonStr := string(json)

	task_id := task.Task.Todo_list_id

	callUrl := fmt.Sprintf("tasklists/%v/tasks.json", task_id)

	resp, _ := request.Post(callUrl, jsonStr)

	return resp.String(), nil
}

func Task_get_by_id(task_id int) structs.Todo_item {

	callUrl := fmt.Sprintf("tasks/%v.json", task_id)

	req := request.Req{
		Url: callUrl, Call_type: "GET"}

	resp, _ := request.MakeRequest(req)

	respBody := resp.Body()

	var retObj structs.Todo_item
	json.Unmarshal(respBody, &retObj)

	return retObj
}

func Task_get_dependencies(task_id int) {

	callUrl := fmt.Sprintf("/tasks/%v/dependencies.json", task_id)

	resp, _ := request.Get(callUrl)

	respBody := resp.Body()

	var retObj interface{}

	json.Unmarshal(respBody, &retObj)

	fmt.Println(retObj)
	// return retObj

}

/*
{
  "commentFollowerIds": [
    "203441"
  ],
  "STATUS": "OK",
  "changeFollowerIds": [
    "203441"
  ]
}
*/

func Task_get_followers(task_id int, get_user_data bool) {

	add_string := ""

	if get_user_data {
		add_string = "?returnUserInfo=true"
	}

	callUrl := fmt.Sprintf("/tasks/%v/followers.json"+add_string, task_id)

	resp, _ := request.Get(callUrl)

	show_json(resp.Body())

}

func Task_update_complete(task_id int) {

	callUrl := fmt.Sprintf("tasks/%v/complete.json", task_id)

	resp, _ := request.Put(callUrl, "")

	var retObj interface{}

	json.Unmarshal(resp.Body(), &retObj)

	fmt.Println(retObj)
	// return retObj
}

func Task_update_incomplete(task_id int) {

	callUrl := fmt.Sprintf("tasks/%v/uncomplete.json", task_id)

	resp, _ := request.Put(callUrl, "")

	var retObj interface{}

	json.Unmarshal(resp.Body(), &retObj)

	fmt.Println(string(retObj.([]byte)))
	// return retObj
}

func Task_update_followers(task_id int, comment_followers string, change_followers string) {

	callUrl := fmt.Sprintf("/tasks/%v.json", task_id)

	payload := fmt.Sprintf(`{ "todo-item": { "use-defaults": false, "commentFollowerIds": "%v", "changeFollowerIds": "%v"}}`, comment_followers, change_followers)

	resp, _ := request.Put(callUrl, payload)

	show_json(resp.Body())
}

func Task_update_remove_all_followers(task_id int) {

	Task_update_followers(task_id, "", "")
}

func Task_update(task structs.Todo_item, task_id int) (string, error) {

	json, _ := json.Marshal(task)
	jsonStr := string(json)

	callUrl := fmt.Sprintf("tasks/%v.json", task_id)

	resp, _ := request.Put(callUrl, jsonStr)

	return resp.String(), nil
}

func Task_delete(task_id int) {

	callUrl := fmt.Sprintf("tasks/%v.json", task_id)

	resp, _ := request.Delete(callUrl)

	respBody := resp.Body()

	var retObj interface{}
	json.Unmarshal(respBody, &retObj)

	//	fmt.Println(string(retObj.([]byte)))
	fmt.Println(resp.String())
	// return retObj
}

/*

Filters Available :
-- Key: filter (Optional) Default: anytime
Tasks can be filtered by due dates using the following options:
	all
	anytime
	overdue
	today
	tomorrow
	thisweek
	within7
	within14
	within30
	within365
	nodate
	nostartdate
	completed
Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false.


-- Key: page (Optional) Default: 1
Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the
parameter pageSize.

For example:
	?page=2&pageSize=10 will retrieve results 10-20.


-- Key: pageSize (Optional	Default: 250
The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter.


-- Key: startdate (Optional)
Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD.

For example:
?startdate=20140512&enddate=20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.


-- Key: enddate (Optional)
Must be used in conjunction with startdate, see above.


-- Key: updatedAfterDate (Optional)
Will only return tasks that have been updated after a specified date. Timestamp must be in the following format:
YYYYMMDDHHMMSS.

For example:
?updatedAfterDate=20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC)

-- Key: completedAfterDate (Optional)
Will only return tasks that have been completed after a specified date. Timestamp must be in the following format:
YYYYMMDDHHMMSS

For example:
?completedAfterDate=20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC)


-- Key: completedBeforeDate (Optional)
Will only return tasks that have been completed before a specified date. Timestamp must be in the following format:
YYYYMMDDHHMMSS

For example:
?completedBeforeDate=20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC)


-- Key: showDeleted (Optional)	Default: no
Tasks that have been deleted can be shown by setting this parameter to yes.


-- Key: includeCompletedTasks (Optional) Default: FALSE
Tasks that have been marked as completed can be shown by setting this parameter to true.


-- Key: includeCompletedSubtasks (Optional) Default: FALSE
Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to
include sub-tasks


-- Key: creator-ids (Optional)	 Default:0
For requesting tasks made by a specific person or people.

For example:
	44 would return tasks made by user 44.
	44,45 would return tasks made by users 44 and/or 45 etc.


-- Key: include (Optional)
Extra tasks that can be included with the filter option.
	nodate
	nostartdate
	noduedate
	overdue


-- Key: responsible-party-ids (Optional)
Tasks can be filtered by the person/people a task is assigned to.

For example:
	-1 would return all tasks with an assigned person.
	0 would return all tasks with no assignment.
	32 would return tasks assigned to user 32.
	32,55 would return tasks assigned to users 32 and/or 55 etc.


-- Key: sort (Optional) Default: duedate
Tasks can be sorted by the following options
	duedate
	startdate
	dateadded
	priority
	project
	company


-- Key: getSubTasks (Optional)	Default: yes
Subtasks can be excluded from the results by adding this parameter with no as the value.


-- Key: nestSubTasks (Optional) Default: no
Subtasks can be nested within the parent task object by adding this paramter with yes as the value.


-- Key: getFiles (Optional) Default: FALSE
Files attached to tasks can be returned within the task object by setting this parameter to true.


-- Key: includeToday (Optional) Default: TRUE
When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to
exclude deadlines for today by passing this parameter as false.


-- Key: ignore-start-dates (Optional) Default: FALSE
When using the filter option, you can choose to include start dates matching the filtering critera by passing this
parameter as true. By default, only due dates are checked against the filter.


-- Key: tag-ids (Optional)
A comma separated list of tag ids to filter tasks on


*/

func Tasks_get_all(filter map[string]string) structs.Todo_items {

	fmt.Println(filter)

	var k structs.Todo_items

	resp, _ := request.Get("tasks.json")

	show_json(resp.Body())

	return k
}

// For seeing responses while making the calls
func show_json(byte []byte) {

	buf := new(bytes.Buffer)
	json.Indent(buf, byte, "", "  ")
	fmt.Println(buf)
}
