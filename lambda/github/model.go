package github

type Repo struct {
	Name             *string          `json:"name,omitempty"`
	FullName         *string          `json:"full_name,omitempty"`
	Description      *string          `json:"description,omitempty"`
	GitURL           *string          `json:"git_url,omitempty"`
}
