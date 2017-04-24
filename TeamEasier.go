package main

import (
	"github.com/goddd/TeamEasier/client"
)

func main() {


	//k  := structs.Tag{
	//	Id : 1,
	//	Color: "#000000",
	//	Name: "Not Done3",
	//}

	//myItem := structs.Todo_item{ Task : structs.Item{
	//
	//	Content: "This is the fourth task you are supposed to be doing!!",
	//	Description: "These are the finer details of the task you only get to know when this is called",
	//	Tags: []structs.Tag{k},
	//	Todo_list_id: 955796,
	//}}

	//client.Add_task(myItem)


	client.Task_get_dependencies(9927032)

	//client.Task_delete(9927032)

	// client.Task_get_followers(9880223, true);
	// client.Tasks_get_all(map[string]string{})


	// client.Task_update_followers(9880223, "203441", "203441")
	// client.Task_update_remove_all_followers(9880223)

	// fmt.Println(client.Get_tasks(955796, map[string]string{ "hey" : "hello", "what" : "when" }))

}
