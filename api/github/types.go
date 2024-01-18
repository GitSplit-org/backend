package github

type AddProjectRequest struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Project struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
