package client

// List All Notebooks
func Notebooks_get_all() {

 // GET  /notebooks.json

}

// List Notebooks on a Project
func Notebooks_get_by_project_id() {

 // GET  /projects/{project_id}/notebooks.json

}

// List Notebooks in a specific category
func Notebooks_get_by_category() {

 // GET  /notebookCategories/{id}/notebooks.json

 } 

// Get a Single Notebook
func Notebooks_get_by_id() {

 // GET  /notebooks/{notebook_id}.json

 } 

// Create a Single Notebook
func Notebooks_create_by_project() {

 // POST  /projects/{project_id}/notebooks.json

 } 

// Update a Single Notebook
func Notebooks_update() {

 // PUT  /notebooks/{notebook_id}.json

 } 

// Lock a Single Notebook For Editing
func Notebooks_update_lock_for_editing() {

 // PUT  /notebooks/{id}/lock.json

 } 

// Unlock a Single Notebook
func Notebooks_update_unlock_for_editing() {

 // PUT  /notebooks/{id}/unlock.json

 } 

// Delete a Single Notebook
func Notebooks_delete_by_id() {

 // DELETE  /notebooks/{id}.json

 } 

// Copy a Notebook to another Project
func Notebooks_copy_to_other_project() {

 // PUT  /notebooks/{notebook_id}/copy.json

 } 

// Move a Notebook to another Project
func Notebooks_move_to_other_project () {

 // PUT  /notebooks/{notebook_id}/move.json

 } 


