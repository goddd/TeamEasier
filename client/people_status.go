package client

// Create Status
func People_status_create_my_status () {

 // POST  /me/status.json

 } 

// Create Status
func People_status_create_user_status() {

 // POST  /people/{person_id}/status.json

 } 

// Update Status
func People_status_update_my_status() {

 // PUT  /me/status/{status_id}.json

 } 

// Update Status
func People_status_update_user_status() {

 // PUT  /people/status/{status_id}.json

 } 

// Update Status
func People_status_update_() {

 // PUT  /people/{person_id}/status/{status_id}.json

 } 

// Delete Status
func People_status_delete_my_status () {

 // DELETE  /me/status/{status_id}.json

 } 

// Delete Status
func People_status_delete_user_status () {

 // DELETE  /people/status/{status_id}.json

 } 

// ???  Why 2 delete status?

/*
// Delete Status
func People_status_delete () {

 // DELETE  /people/{person_id}/status/{status_id}.json

 } 
*/

// Retrieve a Persons Status
func People_status_get_my_status() {

 // GET  /me/status.json

 }

// Retrieve a Persons Status
func People_status_get_user_status() {

 // GET  /people/{user_id}/

 } 

// Retrieve Everybodys Status
func People_status_get_all_users() {

 // GET  /people/status.json

 } 


