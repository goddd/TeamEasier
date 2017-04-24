package client

// Retrieve All Time Entries across all projects
func Time tracking_get () {

 // GET  /time_entries.json

 } 


// Retrieve All Time Entries for a Project
func Time tracking_get () {

 // GET  /projects/{project_id}/time_entries.json

 } 


// Retrieve all To-do Item Times
func Time tracking_get () {

 // GET  /todo_items/{todo_item_id}/time_entries.json

 } 


// Create a Time-Entry
func Time tracking_create () {

 // POST  /projects/{project_id}/time_entries.json

 } 


// Create a Time-Entry (for a task/todo item)
func Time tracking_create () {

 // POST  /tasks/{taskid}/time_entries.json 

 } 


// Retrieve Single Time-Entry
func Time tracking_get () {

 // GET  /time_entries/{id}.json

 } 


// Update an Entry
func Time tracking_update () {

 // PUT  /time_entries/{id}.json

 } 


// Delete Time-Entry
func Time tracking_delete () {

 // DELETE  /time_entries/{id}.json

 } 


// Time Totals Across Account
func Time tracking_get () {

 // GET  /time/total.json

 } 


// Project Times Totals
func Time tracking_get () {

 // GET  /projects/{id}/time/total.json

 } 


// Tasklist Time Totals
func Time tracking_get () {

 // GET  /tasklists/{id}/time/total.json

 } 


// Task Times Totals
func Time tracking_get () {

 // GET  /tasks/{id}/time/total.json

 } 


// Time Totals per Project
func Time tracking_get () {

 // GET  /projects/time/total.json

 } 


