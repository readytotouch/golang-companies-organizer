package domain

type LinkedInProfile struct {
	ID    int64
	Alias string
	Name  string
}

type Article struct {
	Title string
	URL   string
}

type Company struct {
	Name             string
	URL              string
	LinkedInProfiles []LinkedInProfile
	GitHubProfile    string
	Articles         []Article
	MainLanguage     bool // Golang is the main language
}
