package domain

type LinkedInProfile struct {
	ID    int
	Alias string
	Name  string
}

type GitHubProfile struct {
	Login           string
	RepositoryCount int
}

type Company struct {
	Name             string
	URL              string
	LinkedInProfiles []LinkedInProfile
	GitHubProfile    GitHubProfile
	MainLanguage     bool // Golang is the main language
}
