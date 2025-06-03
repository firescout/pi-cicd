package handler

// const (
// 	REPO_CLONE_DIR = "/tmp/repos"
// 	PATH_SEPARATOR = string(os.PathSeparator)
// )

type HTTPHandler struct {
	repoCloneDir  string
	pathSeparator string
	repos         Repos
}

type StdResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Repo struct {
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	AfterScript []AfterScript `json:"after_script"`
	// Path        string        `json:"path"`
}
type AfterScript struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Repos struct {
	System    string `json:"system"`
	ClonePath string `json:"clone_path"`
	Repos     []Repo `json:"repos"`
}
