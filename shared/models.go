package shared

type CreateDeployedAppInput struct {
	RepoName string `json:"repo_name"`
	Port     int    `json:"port"`
}
