package models


// User represents a user in the system
type Company struct {
	ID                   int    `json:"id"`
	CompanyName             string `json:"name"`
	CompanyEmail              string `json:"email"`
	About string `json:"about"`
	Password             string `json:"password"`
	MobileNumber         string `json:"mobile_number"`
	CompanyLogo      string `json:"logo"`
	CompanyGst      string `json:"gstin"`
	CompanyLinkedin      string `json:"linkedin"`
	CompanyWebsite		string `json:"webiste_link"`
	CompanyAddress		string `json:"address"`
	CompanyIndustry		string `json:"industry"`
	CompanyStatus		string `json:"status"`
	NumberOfEmployee	string `json:"number_of_employee"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}


// Company struct should not have the password field
type CompanyModel struct {
    ID              int       `json:"id"`
    CompanyName     string    `json:"name"`
    CompanyEmail    string    `json:"email"`
    About           string    `json:"about"`
    MobileNumber    string    `json:"mobile_number"`
    CompanyLogo     string    `json:"logo"`
    CompanyGst      string    `json:"gstin"`
    CompanyLinkedin string    `json:"linkedin"`
    CompanyWebsite  string    `json:"website_link"`
    CompanyAddress  string    `json:"address"`
    CompanyIndustry string    `json:"industry"`
    CompanyStatus   string    `json:"status"`
    NumberOfEmployee string   `json:"number_of_employee"`
CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}
