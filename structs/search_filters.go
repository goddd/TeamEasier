package structs

// Valid search filters

var Search_filters = map[string][]string{

	// Tasks
	// Reference : http://developer.teamwork.com/todolistitems
	"tasks_tasks": {"filter", "page", "pageSize", "startdate", "enddate", "updatedAfterDate",
		"completedAfterDate", "completedBeforeDate", "showDeleted", "includeCompletedTasks",
		"includeCompletedSubtasks", "creator-ids", "include", "responsible-party-ids", "sort",
		"getSubTasks", "nestSubTasks", "getFiles", "includeToday", "ignore-start-dates", "tag-ids"},

	"tasks_task": {"getFiles", "nestSubTasks", "includeCompletedSubtasks"},

	"tasks_completed_tasks": {"page", "pageSize", "startdate", "enddate", "includeArchivedProjects"},

	// Activity
	// Reference : http://developer.teamwork.com/activity
	"activity_latest_activity": {"maxItems", "onlyStarred"},

	// Invoices
	// Reference : http://developer.teamwork.com/billing
	"invoices_get_invoices_all_projects": {"type", "projectStatus", "page"},
	"invoices_get_invoices_one_project":  {"type", "page"},

	// Boards
	// Refrence : http://developer.teamwork.com/boards
	"boards_list_columns": {"page", "pageSize", "showDeleted", "deletedAfterDate", "updatedAfterDate"},
	"boards_list_cards":   {"page", "pageSize", "showDeleted", "searchTerm", "responsible-party-ids", "deletedAfterDate", "updatedAfterDate"},

	// Companies
	// Reference : http://developer.teamwork.com/companies
	"companies_get_companies": {"page", "pageSize"},

	// Message Replies
	// Reference : http://developer.teamwork.com/messagereplies
	"message_replies_get_replies": {"page", "pageSize"},

	// Milesotones
	// Reference : http://developer.teamwork.com/milestones
	"milestones_list_all": {"page", "pageSize", "find"},

	// Notebooks
	// Refrenece : http://developer.teamwork.com/notebooks
	"notebooks_get": {"includeContent"},

	// People
	// Reference : http://developer.teamwork.com/people
	"people_stats": {"getPermissions", "onlymyprojects", "onlyMyEvents", "eventsInNext"},
	"people_get":   {"emailaddress", "fullprofile", "returnProjectIds", "page", "pageSize"},

	// Projects
	// Reference : http://developer.teamwork.com/projectsapi
	"projects_get": {"status", "updatedAfterDate", "updatedAfterTime", "orderby", "createdAfterDate",
		"createdAfterTime", "catId", "includePeople", "page"},

	"projects_get_one":   {"includePeople"},
	"projects_get_rates": {"page", "pageSize"},

	// Search Projects
	// Reference : http://developer.teamwork.com/searchProjects
	"projects_search": {"searchFor", "searchTerm", "projectId", "sortOrder", "includArchivedProjects",
		"includeCompletedItems", "pageSize"},

	// Timetracking
	// Reference : http://developer.teamwork.com/timetracking
	"timetracking_get_timeentries": {"page", "fromdate", "fromtime", "todate", "totime", "sortby", "sortorder",
		"userId", "billableType", "invoicedType", "projectType", "showDeleted", "tagIds"},

	"timetracking_totals": {"userId", "fromDate", "toDate", "fromTime", "fromDate", "toTime", "toDate",
		"projectType", "page", "pageSize"},

	// Tasklists
	// Reference : http://developer.teamwork.com/tasklists
	"tasklists_get": {"responsible-party-id", "getOverdueCount", "status", "getCompletedCount", "showMilestones",
		"page"},

	// Tasks
	// Reference : http://developer.teamwork.com/todolistitems
	"tasks_get": {"filter", "page", "pageSize", "startdate", "enddate", "updatedAfterDate", "completedAfterDate",
		"completedBeforeDate", "showDeleted", "includeCompletedTasks", "includeCompletedSubtasks",
		"creator-ids", "include", "responsible-party-ids", "sort", "getSubTasks", "nestSubTasks",
		"getFiles", "includeToday", "ignore-start-dates", "tag-ids"},

	"tasks_get_one": {"getFiles", "nestSubTasks", "includeCompletedSubtasks"},

	"tasks_get_completed": {"page", "pageSize", "startdate", "enddate", "includeArchivedProjects"},

	"tasks_get_followers": {"returnUserInfo"},

	// Workload
	// Reference : http://developer.teamwork.com/workload
	"workload_get": {"startDate", "endDate", "includeTasksWithoutDates", "distributeEstimatedTimeToAssignees",
		"prorataEstimatedTimeToAssignees", "page", "pageSize", "sortBy", "prorataEstimatedTime"},
}
