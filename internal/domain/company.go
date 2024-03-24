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

type GlassdoorProfile struct {
	OverviewURL string
	ReviewsURL  string
}

type Company struct {
	Name             string
	URL              string
	LinkedInProfiles []LinkedInProfile
	GitHubProfile    GitHubProfile
	GlassdoorProfile GlassdoorProfile
	OttaProfileSlug  string
	MainLanguage     bool // Golang is the main language
}
