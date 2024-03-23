package db

import "github.com/readytotouch/gocompanies/internal/domain"

func Companies() []domain.Company {
	return []domain.Company{
		{
			Name: "Google",
			URL:  "https://www.google.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    1441,
					Alias: "google",
					Name:  "Google",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "google",
				RepositoryCount: 157,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Google-EI_IE9079.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Google-Reviews-E9079.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "VictoriaMetrics",
			URL:  "https://victoriametrics.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    30169914,
					Alias: "victoriametrics",
					Name:  "VictoriaMetrics",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "VictoriaMetrics",
				RepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+VictoriaMetrics",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+VictoriaMetrics+reviews",
			},
			MainLanguage: true,
		},
		{
			Name: "DocHQ",
			URL:  "https://dochq.co.uk/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    14136494,
					Alias: "dochq",
					Name:  "DocHQ",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "dochq",
				RepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+DocHQ",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+DocHQ+reviews",
			},
			MainLanguage: true,
		},
		{
			Name: "Uber",
			URL:  "https://www.uber.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    1815218,
					Alias: "uber-com",
					Name:  "Uber",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "uber",
				RepositoryCount: 30,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Uber-EI_IE575263.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Uber-Reviews-E575263.htm",
			},
			MainLanguage: true,
		},
		{
			Name: "Careem",
			URL:  "https://www.careem.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    2852511,
					Alias: "careem",
					Name:  "Careem",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "careem",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Careem-EI_IE1438731.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Careem-Reviews-E1438731.htm",
			},
			MainLanguage: true,
		},
		{
			Name: "Cloudflare",
			URL:  "https://www.cloudflare.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    407222,
					Alias: "cloudflare",
					Name:  "Cloudflare",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cloudflare",
				RepositoryCount: 87,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudflare-EI_IE430862.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudflare-Reviews-E430862.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "Bitly",
			URL:  "https://bitly.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    552285,
					Alias: "bitly",
					Name:  "Bitly",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "bitly",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bitly-EI_IE800313.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bitly-Reviews-E800313.htm",
			},
			MainLanguage: true, // https://bitly.com/blog/why-we-write-everything-in-go/
		},
		{
			Name: "Cockroach Labs",
			URL:  "https://www.cockroachlabs.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    9309408,
					Alias: "cockroach-labs",
					Name:  "Cockroach Labs",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cockroachdb",
				RepositoryCount: 92,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cockroach-Labs-EI_IE1168502.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cockroach-Labs-Reviews-E1168502.htm",
			},
			MainLanguage: true,
		},
		{
			Name: "Reddit",
			URL:  "https://www.reddit.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    150573,
					Alias: "reddit-com",
					Name:  "Reddit, Inc.",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "reddit",
				RepositoryCount: 20,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Reddit-EI_IE796358.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Reddit-Reviews-E796358.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "Dailymotion",
			URL:  "https://www.dailymotion.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    24411,
					Alias: "dailymotion",
					Name:  "Dailymotion",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "dailymotion",
				RepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dailymotion-EI_IE372880.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dailymotion-Reviews-E372880.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "Stream",
			URL:  "https://getstream.io/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    5338728,
					Alias: "getstream",
					Name:  "Stream",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "GetStream",
				RepositoryCount: 32,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stream-CO-EI_IE1703813.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stream-CO-Reviews-E1703813.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "Workato",
			URL:  "https://www.workato.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    3675685,
					Alias: "workato",
					Name:  "Workato",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "workato",
				RepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Workato-EI_IE1333040.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Workato-Reviews-E1333040.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "FerretDB",
			URL:  "https://www.ferretdb.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    80672744,
					Alias: "ferretdb",
					Name:  "FerretDB",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "FerretDB",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+FerretDB",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+FerretDB+reviews",
			},
			MainLanguage: true,
		},
		{
			Name: "Form3",
			URL:  "https://www.form3.tech/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15156804,
					Alias: "form3-financial-cloud",
					Name:  "Form3",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "form3tech-oss",
				RepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Form3-EI_IE2008415.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Form3-Reviews-E2008415.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "Docker",
			URL:  "https://www.docker.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    1301808,
					Alias: "docker",
					Name:  "Docker, Inc",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "docker",
				RepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Docker-EI_IE1089506.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Docker-Reviews-E1089506.htm",
			},
			MainLanguage: true,
		},
		{
			Name: "Assertive Yield",
			URL:  "https://www.assertiveyield.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    76806664,
					Alias: "assertive-yield",
					Name:  "Assertive Yield B.V.",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Assertive-Yield",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: `https://www.google.com/search?q=Glassdoor+"Assertive+Yield"`,
				ReviewsURL:  `https://www.google.com/search?q=Glassdoor+%22Assertive+Yield%22+reviews`,
			},
			MainLanguage: true,
		},
		{
			Name: "Splunk",
			URL:  "https://www.splunk.com/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    20226,
					Alias: "splunk",
					Name:  "Splunk",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "splunk",
				RepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splunk-EI_IE117313.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splunk-Reviews-E117313.htm",
			},
			MainLanguage: false,
		},
		{
			Name: "90POE",
			URL:  "https://www.90poe.io/",
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    18466590,
					Alias: "90poe",
					Name:  "Ninety Percent of Everything (90POE)",
				},
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "90poe",
				RepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-90-POE-EI_IE3898368.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/90-POE-Reviews-E3898368.htm",
			},
			MainLanguage: true,
		},
		//{
		//	Name: "",
		//	URL:  "",
		//	LinkedInProfiles: []domain.LinkedInProfile{
		//		{
		//			ID:    0,
		//			Alias: "",
		//			Name:  "",
		//		},
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:           "",
		//		RepositoryCount: 0,
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//	},
		//	MainLanguage: false,
		//},
	}
}
