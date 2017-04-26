package client

// Create a new Column
func Boards_create_column () {

 // POST /projects/{project_id}/boards/columns.json

 } 


// List Columns
func Boards_get_by_project_id() {

 // GET /projects/{project_id}/boards/columns.json

 } 


// Delete a Column
func Boards_delete_by_column_id () {

 // DELETE  /boards/columns/{column_id}.json

 } 


// Add a Task from the Backlog to a Column
func Boards_add_task_to_column() {

 // POST /boards/columns/{column_id}/cards.json

 } 


// List Cards in a Column
func Boards_get_cards_by_column_id () {

 // GET /boards/columns/{column_id}/cards.json

 } 


// Move a Card
func Boards_update_move_cards () {

 // PUT /boards/columns/cards/{card_id}/move.json

 } 


// Edit a Card
func Boards_update_card() {

 // PUT /boards/columns/cards/{card_id}.json

 } 


// Remove a Card
func Boards_delete_card() {

 // DELETE  /boards/columns/cards/{card_id}.json

 } 



