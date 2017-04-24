package client

// List All Notebooks
func Notebooks_get () {

 // GET  /notebooks.json

 } 


// List Notebooks on a Project
func Notebooks_get () {

 // GET  /projects/{project_id}/notebooks.json

 } 


// List Notebooks in a specific category
func Notebooks_get () {

 // GET  /notebookCategories/{id}/notebooks.json

 } 


// Get a Single Notebook
func Notebooks_get () {

 // GET  /notebooks/{notebook_id}.json

 } 


// Create a Single Notebook
func Notebooks_create () {

 // POST  /projects/{project_id}/notebooks.json

 } 


// Update a Single Notebook
func Notebooks_update () {

 // PUT  /notebooks/{notebook_id}.json

 } 


// Lock a Single Notebook For Editing
func Notebooks_update () {

 // PUT  /notebooks/{id}/lock.json

 } 


// Unlock a Single Notebook
func Notebooks_update () {

 // PUT  /notebooks/{id}/unlock.json

 } 


// Delete a Single Notebook
func Notebooks_delete () {

 // DELETE  /notebooks/{id}.json

 } 


// Copy a Notebook to another Project
func Notebooks_update () {

 // PUT  /notebooks/{notebook_id}/copy.json

 } 


// Move a Notebook to another Project
func Notebooks_update () {

 // PUT  /notebooks/{notebook_id}/move.json

 } 


