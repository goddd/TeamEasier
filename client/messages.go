package client

// Create a message
func Messages_create() {

 // POST  /projects/{project_id}/posts.json

 } 

// Retrieve a Single Message
func Messages_get_by_id() {

 // GET  /posts/{id}.json

 } 

// Retrieve Latest Messages
func Messages_get_latest() {

 // GET  /projects/{project_id}/posts.json

 } 

// Retrieve Messages by Category
func Messages_get_by_category() {

 // GET  /projects/{project_id}/cat/{category_id}/posts.json

 } 

// Update message
func Messages_update() {

 // PUT  /posts/{id}.json

 }

// Get archived messages
func Messages_get_archived() {

 // GET  /projects/{project_id}/posts/archive.json

 } 

// Get archived messages by category
func Messages_get_archived_by_category() {

 // GET  /projects/{project_id}/cat/{category_id}/posts/archive.json

 } 


// Archive a message
func Messages_update_archive() {

 // PUT  /messages/{id}/archive.json

 }

// Un-archive a message
func Messages_update_unarchive() {

 // PUT  /messages/{id}/unarchive.json

 } 

// Destroy message
func Messages_delete () {

 // DELETE  /posts/{id}.json

 } 


// Mark Message Read
func Messages_update_as_read() {

 // PUT  /messages/{id}/markread.json

 } 


