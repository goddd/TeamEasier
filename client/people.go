package client

// Add a new user
func People_create_user() {

 // POST  /people.json

 } 

// Edit user
func People_update_user() {

 // PUT  /people/{id}.json

 } 

// Delete user
func People_delete_user() {

 // DELETE  /people/{id}.json

 } 

// Get Current User Details
func People_get_self_details() {

 // GET  /me.json

 } 

// Current User Summary Stats
func People_get_user_summary() {

 // GET  /stats.json

 } 

// ???
// Get people
func People_get_all() {

 // GET  /people.json

 } 

// Get all People (within a Project)
func People_get_by_project() {

 // GET  /projects/{project_id}/people.json

 } 

// Get People (within a Company)
func People_get_by_company() {

 // GET  /companies/{company_id}/people.json

 } 

// Retrieve a Specific Person
func People_get_user() {

 // GET  /people/{person_id}.json

 } 

// Retrieve a API Keys for all people on account
func People_get_all_users_api_keys() {

 // GET  /people/APIKeys.json

 } 

// Unassign a user from all Tasks
func People_update_user_unassign_all_tasks() {

 // PUT  /people/{person_id}.json

 } 


