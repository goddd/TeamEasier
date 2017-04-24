package client

// Get all expenses across your projects
func Billing_expenses_get() {
	// GET	/expenses.json
}


// Get all expenses on a single project
func Billing_expenses_get_by_project() {
	// GET	/projects/:id/expenses.json
}

// Get a single expense
func Billing_expenses_get_by_id() {
	// GET	/expenses/:id.json
}

// Create a new expense in a project
func Billing_expenses_add_by_project() {
	// POST	/expenses.json
}

// Update a single expense
func Billing_expenses_update_by_id() {
	// PUT	/expenses/:id.json
}

// Delete a single expense
func Billing_expenses_delete_by_id() {
	// DELETE 	/expenses/:id.json
}

// Add an unbilled expense to an Invoice
func Billing_expenses_add_to_invoice() {
	// PUT	/invoices/:invoiceid/lineitems.json
}

