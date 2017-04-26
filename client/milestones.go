package client

// List All Milestones
func Milestones_get_all() {

 // GET  /milestones.json

 } 

// List Milestones on a Project
func Milestones_get_by_project() {

 // GET  /projects/{project_id}/milestones.json

 } 

// Get a Single Milestone
func Milestones_get_by_id() {

 // GET  /milestones/{milestone_id}.json

 } 

// Complete
func Milestones_update_completed() {

 // PUT  /milestones/{id}/complete.json

 } 

// Incomplete
func Milestones_update_incomplete() {

 // PUT  /milestones/{id}/uncomplete.json

 } 

// Create a Single Milestone
func Milestones_create() {

 // POST  /projects/{project_id}/milestones.json

 } 

// Update
func Milestones_update() {

 // PUT  /milestones/{milestone_id}.json

 } 

// Delete
func Milestones_delete() {

 // DELETE  /milestones/{id}.json

 } 


