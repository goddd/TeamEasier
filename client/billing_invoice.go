package client

// Get all invoices across your projects
func Billing_invoices_get(){

	// GET	/invoices.json

}

// Get all invoices on a single project
func Billing_invoices_get_by_project(){

	// GET	/projects/:id/invoices.json

}

// Get a single invoice
func Billing_invoices_get_single(){

	// GET	/invoices/:id.json

}

// * Create a new invoice in a project
func Billing_invoices_add_to_project(){

	// POST	/invoices.json
}


// Update a specific invoice
func Billing_invoices_update_by_id(){

	// Invoice	PUT	/invoices/:id.json

}

// Delete a specific invoice
func Billing_invoices_delete_by_id(){

	// DELETE 	/invoices/:id.json

}

// Mark a specific invoice as complete
func Billing_invoices_update_complete(){

	// PUT	/invoices/:id/complete.json
}

// Mark a specific invoice as not complete
func Billing_invoices_update_incomplete(){

	// PUT	/invoices/:id/uncomplete.json
}

// Add a time entry to an Invoice
func Billing_invoices_add_time_entry() {
	// PUT	/invoices/:invoiceid/lineitems.json
}

// Get a list of valid currency codes
func Billing_invoices_get_currency_codes(){

	// GET	/currencycodes.json

}