package client

// List Files on a Task
func Files_get_by_task () {

 // GET  /tasks/{task_id}/files.json

 } 

// List Files on a Project
func Files_get_by_project () {

 // GET  /projects/{project_id}/files.json

 } 

// Get a Single File
func Files_get_by_id() {

 // GET  /files/{file_id}.json

 } 

// Add a File to a Project
func Files_add_to_project () {

 // POST  /projects/{project_id}/files.json

 } 

// Add a new File Version to a File
func Files_create_new_version () {

 // POST  /files/{file_id}.json

 } 

// Delete a File from a Project
func Files_delete_by_project() {

 // DELETE  /files/{file_id}.json

 } 

// Get a short URL for sharing a file
func Files_create_short_url() {

 // GET  /files/{file_id}/sharedlink.json
 // GET  /files/{file_id}/sharedlink.json?version=2
 } 

// Copy or Move a file to another project
func Files_() {

   // ???

 } 

// Add or update a Project logo
func Files_ () {

  // ???
 // PUT  to /projects/{project_id}.json

 } 


// Files Uploading
func Files_uploading_ () {

  // ???
 // POST  /pendingfiles.json

 } 


