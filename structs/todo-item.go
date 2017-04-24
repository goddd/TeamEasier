package structs

import "time"

type Tag struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type ParentTask struct {
	Content string `json:"content,omitempty"`
	Id      string `json:"id,omitempty"`
}

type Item struct {
	CanComplete               bool         `json:"canComplete,omitempty"`
	Project_id                int          `json:"project-id,omitempty"`
	Creator_lastname          string       `json:"creator-lastname,omitempty"`
	Has_reminders             bool         `json:"has-reminders,omitempty"`
	Todo_list_name            string       `json:"todo-list-name,omitempty"`
	Has_unread_comments       bool         `json:"has-unread-comments,omitempty"`
	Has_tickets               bool         `json:"hasTickets,omitempty"`
	Order                     int          `json:"order,omitempty"`
	Comments_count            int          `json:"comments-count,omitempty"`
	Private                   int          `json:"private,omitempty"`
	Grant_access_to           string       `json:"grant-access-to,omitempty"`
	Status                    string       `json:"status,omitempty"`
	Todo_list_id              int          `json:"todo-list-id,omitempty"`
	Predecessors              []string     `json:"predecessors,omitempty"`
	Created_on                time.Time    `json:"created_on,omitempty"` // 2014_04_01T15 52 15Z
	CanEdit                   bool         `json:"canEdit,omitempty"`
	Content                   string       `json:"content,omitempty"`
	Has_predecessors          int          `json:"has-predecessors,omitempty"`
	Company_name              string       `json:"company-name,omitempty"`
	Id                        int          `json:"id,omitempty"`
	Creator_firstname         string       `json:"creator-firstname,omitempty"`
	Last_changed_on           time.Time    `json:"last-changed-on,omitempty"`
	Has_dependencies          int          `json:"has-dependencies,omitempty"`
	Completed                 bool         `json:"completed,omitempty"`
	Position                  int          `json:"position,omitempty"`
	Attachments_count         int          `json:"attachments-count,omitempty"`
	Estimated_minutes         int          `json:"estimated-minutes,omitempty"`
	Description               string       `json:"description,omitempty"`
	Priority                  string       `json:"priority,omitempty"`
	Progress                  int          `json:"progress,omitempty"`
	Harvest_enabled           bool         `json:"harvest-enabled,omitempty"`
	ViewEstimatedTime         bool         `json:"viewEstimatedTime,omitempty"`
	ParentTaskId              string       `json:"parentTaskId,omitempty"`
	Company_id                int          `json:"company-id,omitempty"`
	Tasklist_lockdownId       string       `json:"tasklist-lockdownId,omitempty"`
	Creator_avatar_url        string       `json:"creator-avatar-url,omitempty"`
	CanLogTime                bool         `json:"canLogTime,omitempty"`
	Creator_id                int          `json:"creator-id,omitempty"`
	Project_name              string       `json:"project-name,omitempty"`
	Parent_task               []ParentTask `json:"parent-task,omitempty"`
	Attachments               []string     `json:"attachments,omitempty"`
	Responsible_party_ids     string       `json:"responsible-party-ids,omitempty"`
	Responsible_party_names   string       `json:"responsible-party-names,omitempty"`
	Responsible_party_summary string       `json:"responsible-party-summary,omitempty"`
	Tasklist_private          bool         `json:"tasklist-private,omitempty"`
	TimeIsLogged              string       `json:"timeIsLogged,omitempty"`
	LockdownId                string       `json:"lockdownId,omitempty"`
	Tags                      []Tag        `json:"tags,omitempty"`
	User_following_comments   bool         `json:"userFollowingComments,omitempty"`
	User_following_changes    bool         `json:"userFollowingChanges,omitempty"`

	//	Start_date  20140402
	// 	Due_date_base  20140405
	//	Due_date  20140405

}

type Todo_item struct {
	Task Item `json:"todo-item"`
}

type Todo_items struct {
	Tasks []Item `json:"todo-items"`
}
