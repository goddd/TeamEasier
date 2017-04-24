// Valid tag colors :
// #d84640,#f78234,#f4bd38,#b1da34,#53c944,#37ced0,#2f8de4,#9b7cdb,#f47fbe,#a6a6a6,#4d4d4d,#9e6957

package structs

type Account_Details struct {
	Requirehttps               bool   `json:"requirehttps,omitempty"`
	Time_tracking_enabled      bool   `json:"time-tracking-enabled,omitempty"`
	Name                       string `json:"name,omitempty"`
	Datesignedup               string `json:"datesignedup,omitempty"`
	Companyname                string `json:"companyname,omitempty"`
	Ssl_enabled                bool   `json:"ssl-enabled,omitempty"`
	Created_at                 string `json:"created-at,omitempty"`
	CacheUUID                  string `json:"cacheUUID,omitempty"`
	Account_holder_id          string `json:"account-holder-id,omitempty"`
	Logo                       string `json:"logo,omitempty"`
	Id                         string `json:"id,omitempty"`
	URL                        string `json:"URL,omitempty"`
	Email_notification_enabled bool   `json:"email-notification-enabled,omitempty"`
	Companyid                  string `json:"companyid,omitempty"`
	Lang                       string `json:"lang,omitempty"`
	Code                       string `json:"code,omitempty"`
}

type Account struct {
	Details Account_Details `json:"account",omitempty`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type CategoryDetails struct {
	Id             int    `json:"id,omitempty"`
	Elements_count int    `json:"elements_count,omitempty"`
	Name           string `json:"name,omitempty"`
	Project_id     int    `json:"project-id,omitempty"`
	Type           string `json:"type,omitempty"`
	Parent_id      string `json:"parent-id,omitempty"`
}

type Category struct {
	Details CategoryDetails `json:"category,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type CommentDetails struct {
	Id                int    `json:"id,omitempty"`
	Commentable_type  string `json:"commentable_type,omitempty"`
	Body              string `json:"body,omitempty"`
	Author_id         string `json:"author_id,omitempty"`
	Author_firstname  string `json:"author-firstname,omitempty"`
	Author_lastname   string `json:"author-lastname,omitempty"`
	Author_avatar_url string `json:"author-avatar-url,omitempty"`
	Commentable_id    int    `json:"commentable-id,omitempty"`
	Attachments_count int    `json:"attachments-count,omitempty"`
	Emailed_from      string `json:"emailed-from,omitempty"`
	Datetime          string `json:"datetime,omitempty"`
	Private           bool   `json:"private,omitempty"`
}

type Comment struct {
	Details CommentDetails `json:"comment,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type CompanyDetails struct {
	Id               int    `json:"id,omitempty"`
	Can_see_private  bool   `json:"can_see_private,omitempty"`
	Company_name_url string `json:"company_name_url,omitempty"`
	Name             string `json:"name,omitempty"`
	Address_one      string `json:"address_one,omitempty"`
	Address_two      string `json:"address_two,omitempty"`
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	Zip              string `json:"zip,omitempty"`
	Country          string `json:"country,omitempty"`
	Website          string `json:"website,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Fax              string `json:"fax,omitempty"`
}

type Company struct {
	Details CompanyDetails `json:"company,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type PostDetails struct {
	Id                int    `json:"id,omitempty"`
	Title             string `json:"title,omitempty"`
	Author_id         int    `json:"author-id,omitempty"`
	Author_first_name string `json:"author-first-name,omitempty"`
	Author_last_name  string `json:"author-last-name,omitempty"`
	Author_avatar_url string `json:"author-avatar-url,omitempty"`
	Body              string `json:"body,omitempty"`
	Category_id       int    `json:"category-id,omitempty"`
	Posted_on         string `json:"posted-on,omitempty"`
	Project_id        int    `json:"project-id,omitempty"`
	Comments_count    int    `json:"comments-count,omitempty"`
	Attachments_count int    `json:"attachments-count,omitempty"`
	Display_body      string `json:"display-body,omitempty"`
	Private           bool   `json:"private,omitempty"`
}

type Post struct {
	Details PostDetails `json:"post,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type MilestoneDetails struct {
	Id                      int    `json:"id,omitempty"`
	Title                   string `json:"title,omitempty"`
	Project_id              int    `json:"project-id,omitempty"`
	Created_on              string `json:"created-on,omitempty"`
	Creator_id              int    `json:"creator-id,omitempty"`
	Comments_count          int    `json:"comments-count,omitempty"`
	Completed_on            string `json:"completed-on,omitempty"`
	Deadline                string `json:"deadline,omitempty"`
	Completed               bool   `json:"completed,omitempty"`
	Completer_id            int    `json:"completer-id,omitempty"`
	Responsible_party_ids   string `json:"responsible-party-ids,omitempty"`
	Responsible_party_names string `json:"responsible-party-names,omitempty"`
	Reminder                bool   `json:"reminder,omitempty"`
	Private                 bool   `json:"private,omitempty"`
}

type Milestone struct {
	Details MilestoneDetails `json:"milestone,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type NotebookDetails struct {
	Id                       int    `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	Content                  string `json:"content,omitempty"`
	Private                  int    `json:"private,omitempty"`
	Category_id              int    `json:"category-id,omitempty"`
	Category_name            string `json:"category-name,omitempty"`
	Created_date             string `json:"created-date,omitempty"`
	Created_by_userId        int    `json:"created-by-userId,omitempty"`
	Created_by_userfirstname string `json:"created-by-userfirstname,omitempty"`
	Created_by_userlastname  string `json:"created-by-userlastname,omitempty"`
	Locked                   int    `json:"locked,omitempty"`
	Locked_date              string `json:"locked-date,omitempty"`
	Locked_by_userid         int    `json:"locked-by-userid,omitempty"`
	Locked_by_userfirstname  string `json:"locked-by-userfirstname,omitempty"`
	Locked_by_userlastname   string `json:"locked-by-userlastname,omitempty"`
	Project_id               int    `json:"project-id,omitempty"`
}

type Notebook struct {
	Details NotebookDetails `json:"notebook,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type PermissionsDetails struct {
	View_messages_and_files                 string `json:"view-messages-and-files,omitempty"`
	View_tasks_and_milestones               string `json:"view-tasks-and-milestones,omitempty"`
	View_time                               string `json:"view-time,omitempty"`
	View_notebooks                          string `json:"view-notebooks,omitempty"`
	View_risk_register                      string `json:"view-risk-register,omitempty"`
	View_invoices                           string `json:"view-invoices,omitempty"`
	View_links                              string `json:"view-links,omitempty"`
	Add_tasks                               string `json:"add-tasks,omitempty"`
	Add_milestones                          string `json:"add-milestones,omitempty"`
	Add_taskLists                           string `json:"add-taskLists,omitempty"`
	Add_messages                            string `json:"add-messages,omitempty"`
	Add_files                               string `json:"add-files,omitempty"`
	Add_time                                string `json:"add-time,omitempty"`
	Add_notebooks                           string `json:"add-notebooks,omitempty"`
	Add_links                               string `json:"add-links,omitempty"`
	Set_privacy                             string `json:"set-privacy,omitempty"`
	Can_be_assigned_to_tasks_and_milestones string `json:"can-be-assigned-to-tasks-and-milestones,omitempty"`
	Project_administrator                   string `json:"project-administrator,omitempty"`
	Add_people_to_project                   string `json:"add-people-to-project,omitempty"`
}

type Permissions struct {
	Details PermissionsDetails `json:"permissions,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type PersonDetails struct {
	Id                         int    `json:"id,omitempty"`
	User_type                  string `json:"user-type,omitempty"`
	First_name                 string `json:"first-name,omitempty"`
	Last_name                  string `json:"last-name,omitempty"`
	Title                      string `json:"title,omitempty"`
	Email_address              string `json:"email-address,omitempty"`
	Im_handle                  string `json:"im-handle,omitempty"`
	Im_service                 string `json:"im-service,omitempty"`
	Notes                      string `json:"notes,omitempty"`
	Phone_number_office        string `json:"phone-number-office,omitempty"`
	Phone_number_office_ext    string `json:"phone-number-office-ext,omitempty"`
	Phone_number_mobile        string `json:"phone-number-mobile,omitempty"`
	Phone_number_home          string `json:"phone-number-home,omitempty"`
	Phone_number_fax           string `json:"phone-number-fax,omitempty"`
	Last_login                 string `json:"last-login,omitempty"`
	Company_id                 int    `json:"company-id,omitempty"`
	Company_name               string `json:"company-name,omitempty"`
	In_owner_company           string `json:"in-owner-company,omitempty"`
	Created_at                 string `json:"created-at,omitempty"`
	Last_changed_on            string `json:"last-changed-on,omitempty"`
	Avatar_url                 string `json:"avatar-url,omitempty"`
	User_name                  string `json:"user-name,omitempty"`
	Deleted                    bool   `json:"deleted,omitempty"`
	Has_access_to_new_projects bool   `json:"has-access-to-new-projects,omitempty"`
	Administrator              bool   `json:"administrator,omitempty"`
}

type Person struct {
	Details PersonDetails `json:"person,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type ProjectDetails struct {
	Name                   string `json:"name,omitempty"`
	Description            string `json:"description,omitempty"`
	CompanyId              int    `json:"companyId,omitempty"`
	NewCompany             string `json:"newCompany,omitempty"`
	Category_id            int    `json:"category-id,omitempty"`
	Start_date             string `json:"start-date,omitempty"`
	End_date               string `json:"end-date,omitempty"`
	Use_tasks              string `json:"use-tasks,omitempty"`
	Use_milestones         string `json:"use-milestones,omitempty"`
	Use_messages           string `json:"use-messages,omitempty"`
	Use_files              string `json:"use-files,omitempty"`
	Use_time               string `json:"use-time,omitempty"`
	Use_notebook           string `json:"use-notebook,omitempty"`
	Use_riskregister       string `json:"use-riskregister,omitempty"`
	Use_links              string `json:"use-links,omitempty"`
	Use_billing            string `json:"use-billing,omitempty"`
	Start_page             string `json:"start-page,omitempty"`
	Harvest_timers_enabled bool   `json:"harvest-timers-enabled,omitempty"`
	DefaultPrivacy         string `json:"defaultPrivacy,omitempty"`
}

type Project struct {
	Details ProjectDetails `json:"project,omitempty"`
}

// ------ ------ ------ ------ ------ ------ ------ ------ ------

type LinkDetails struct {
	Id                       int    `json:"id,omitempty"`
	Project_id               int    `json:"project-id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	Created_by_userfirstname string `json:"created-by-userfirstname,omitempty"`
	Height                   int    `json:"height,omitempty"`
	Private                  string `json:"private,omitempty"`
	Width                    int    `json:"width,omitempty"`
	Created_by_userId        int    `json:"created-by-userId,omitempty"`
	Created_by_userlastname  string `json:"created-by-userlastname,omitempty"`
	Category_id              int    `json:"category-id,omitempty"`
	Project_name             string `json:"project-name,omitempty"`
	Open_in_new_window       bool   `json:"open-in-new-window,omitempty"`
	Provider                 string `json:"provider,omitempty"`
	Created_date             string `json:"created-date,omitempty"`
	Category_name            string `json:"category-name,omitempty"`
	Code                     string `json:"code,omitempty"`
}

type Link struct {
	Details LinkDetails `json:"link,omitempty"`
}
