package projects

type AddProjectRequest struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Twitter     string `json:"twitter"`
	Instagram   string `json:"instagram"`
	Linkedin    string `json:"linkedin"`
	Owner       string `json:"owner"`
	Image       string `json:"image"`
}
