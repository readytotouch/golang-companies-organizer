package db

import "github.com/readytotouch/golang-companies-organizer/internal/domain"

func GoCompanies() []domain.Company {
	return []domain.Company{
		{
			Name: "Google",
			URL:  "https://www.google.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1441,
				Alias: "google",
				Name:  "Google",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "google",
				RepositoryCount: 157,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Google-EI_IE9079.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Google-Reviews-E9079.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    false,
		},
		{
			Name: "VictoriaMetrics",
			URL:  "https://victoriametrics.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    30169914,
				Alias: "victoriametrics",
				Name:  "VictoriaMetrics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "VictoriaMetrics",
				RepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+VictoriaMetrics",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+VictoriaMetrics+reviews",
			},
			OttaProfileSlug: "",
			MainLanguage:    true,
		},
		{
			Name: "DocHQ",
			URL:  "https://dochq.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14136494,
				Alias: "dochq",
				Name:  "DocHQ",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "dochq",
				RepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+DocHQ",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+DocHQ+reviews",
			},
			OttaProfileSlug: "",
			MainLanguage:    true,
		},
		{
			Name: "Uber",
			URL:  "https://www.uber.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1815218,
				Alias: "uber-com",
				Name:  "Uber",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "uber",
				RepositoryCount: 30,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Uber-EI_IE575263.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Uber-Reviews-E575263.htm",
			},
			OttaProfileSlug: "Uber",
			MainLanguage:    true,
		},
		{
			Name: "Careem",
			URL:  "https://www.careem.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2852511,
				Alias: "careem",
				Name:  "Careem",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "careem",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Careem-EI_IE1438731.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Careem-Reviews-E1438731.htm",
			},
			OttaProfileSlug: "Careem",
			MainLanguage:    true,
		},
		{
			Name: "Cloudflare",
			URL:  "https://www.cloudflare.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    407222,
				Alias: "cloudflare",
				Name:  "Cloudflare",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cloudflare",
				RepositoryCount: 87,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudflare-EI_IE430862.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudflare-Reviews-E430862.htm",
			},
			OttaProfileSlug:   "Cloudflare-1",
			YouTubeChannelURL: "https://www.youtube.com/@cloudflare",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/RWRkVWNJ", "https://app.otta.com/jobs/B_RaZ6l5", "https://app.otta.com/jobs/e4CavNqx"},
		},
		{
			Name: "Bitly",
			URL:  "https://bitly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    552285,
				Alias: "bitly",
				Name:  "Bitly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "bitly",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bitly-EI_IE800313.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bitly-Reviews-E800313.htm",
			},
			OttaProfileSlug: "Bitly",
			MainLanguage:    true, // https://bitly.com/blog/why-we-write-everything-in-go/
		},
		{
			Name: "Cockroach Labs",
			URL:  "https://www.cockroachlabs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9309408,
				Alias: "cockroach-labs",
				Name:  "Cockroach Labs",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cockroachdb",
				RepositoryCount: 92,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cockroach-Labs-EI_IE1168502.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cockroach-Labs-Reviews-E1168502.htm",
			},
			OttaProfileSlug:   "Cockroach-Labs",
			YouTubeChannelURL: "https://www.youtube.com/@cockroachdb",
			MainLanguage:      true,
			Vacancies:         []string{"https://app.otta.com/jobs/uPMrLAMV"},
			Skip:              false,
		},
		{
			Name: "Reddit",
			URL:  "https://www.reddit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    150573,
				Alias: "reddit-com",
				Name:  "Reddit, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "reddit",
				RepositoryCount: 20,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Reddit-EI_IE796358.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Reddit-Reviews-E796358.htm",
			},
			OttaProfileSlug:   "Reddit-1",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/l-8sLbJL"},
		},
		{
			Name: "Dailymotion",
			URL:  "https://www.dailymotion.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    24411,
				Alias: "dailymotion",
				Name:  "Dailymotion",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "dailymotion",
				RepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dailymotion-EI_IE372880.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dailymotion-Reviews-E372880.htm",
			},
			OttaProfileSlug: "Dailymotion",
			MainLanguage:    false,
		},
		{
			Name: "Stream",
			URL:  "https://getstream.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5338728,
				Alias: "getstream",
				Name:  "Stream",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "GetStream",
				RepositoryCount: 32,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stream-CO-EI_IE1703813.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stream-CO-Reviews-E1703813.htm",
			},
			OttaProfileSlug:   "Stream",
			YouTubeChannelURL: "https://www.youtube.com/@streamdevelopers",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3873637399/"},
		},
		{
			Name: "Workato",
			URL:  "https://www.workato.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3675685,
				Alias: "workato",
				Name:  "Workato",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "workato",
				RepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Workato-EI_IE1333040.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Workato-Reviews-E1333040.htm",
			},
			OttaProfileSlug: "Workato",
			MainLanguage:    false,
		},
		{
			Name: "FerretDB",
			URL:  "https://www.ferretdb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80672744,
				Alias: "ferretdb",
				Name:  "FerretDB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "FerretDB",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.google.com/search?q=Glassdoor+FerretDB",
				ReviewsURL:  "https://www.google.com/search?q=Glassdoor+FerretDB+reviews",
			},
			OttaProfileSlug: "FerretDB",
			MainLanguage:    true,
		},
		{
			Name: "Form3",
			URL:  "https://www.form3.tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15156804,
				Alias: "form3-financial-cloud",
				Name:  "Form3",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "form3tech-oss",
				RepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Form3-EI_IE2008415.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Form3-Reviews-E2008415.htm",
			},
			OttaProfileSlug: "Form3",
			MainLanguage:    false,
			Vacancies: []string{
				"https://relocate.me/the-united-kingdom/london/form3/senior-software-engineer-go-7646",
			},
		},
		{
			Name: "Docker",
			URL:  "https://www.docker.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1301808,
				Alias: "docker",
				Name:  "Docker, Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "docker",
				RepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Docker-EI_IE1089506.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Docker-Reviews-E1089506.htm",
			},
			OttaProfileSlug: "Docker-1",
			MainLanguage:    true,
		},
		{
			Name: "Assertive Yield",
			URL:  "https://www.assertiveyield.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    76806664,
				Alias: "assertive-yield",
				Name:  "Assertive Yield B.V.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Assertive-Yield",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: `https://www.google.com/search?q=Glassdoor+"Assertive+Yield"`,
				ReviewsURL:  `https://www.google.com/search?q=Glassdoor+%22Assertive+Yield%22+reviews`,
			},
			OttaProfileSlug: "",
			MainLanguage:    true,
		},
		{
			Name: "Splunk",
			URL:  "https://www.splunk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20226,
				Alias: "splunk",
				Name:  "Splunk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "splunk",
				RepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splunk-EI_IE117313.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splunk-Reviews-E117313.htm",
			},
			OttaProfileSlug:   "Splunk",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://justjoin.it/offers/splunk-senior-software-engineer-backend-go-olkusz"},
		},
		{
			Name: "90POE",
			URL:  "https://www.90poe.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18466590,
				Alias: "90poe",
				Name:  "Ninety Percent of Everything (90POE)",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "90poe",
				RepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-90-POE-EI_IE3898368.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/90-POE-Reviews-E3898368.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    true,
		},
		{
			Name: "HelloFresh",
			URL:  "https://www.hellofresh.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2454643,
				Alias: "hellofresh",
				Name:  "HelloFresh",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "hellofresh",
				RepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HelloFresh-EI_IE998728.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HelloFresh-Reviews-E998728.htm",
			},
			OttaProfileSlug: "HelloFresh",
			MainLanguage:    false,
		},
		{
			Name: "AUTODOC",
			URL:  "https://autodoc.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7703911,
				Alias: "autodoc",
				Name:  "AUTODOC",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // unknown
				RepositoryCount: 0,  // unknown
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AUTODOC-EI_IE2179604.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AUTODOC-Reviews-E2179604.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    false,
		},
		{
			Name: "Gymondo",
			URL:  "https://www.gymondo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5233814,
				Alias: "gymondo-gmbh",
				Name:  "Gymondo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Gymondo-git",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gymondo-EI_IE1344198.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gymondo-Reviews-E1344198.htm",
			},
			OttaProfileSlug: "Gymondo",
			MainLanguage:    false,
		},
		{
			Name: "Delivery Hero",
			URL:  "https://www.deliveryhero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2393200,
				Alias: "delivery-hero-se",
				Name:  "Delivery Hero",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "deliveryhero",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Delivery-Hero-EI_IE504556.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Delivery-Hero-Reviews-E504556.htm",
			},
			OttaProfileSlug:   "Delivery-Hero",
			YouTubeChannelURL: "https://www.youtube.com/@DeliveryHeroDH",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3908507037/"},
			Skip:              false,
		},
		{
			Name: "HashiCorp",
			URL:  "https://www.hashicorp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2830763,
				Alias: "hashicorp",
				Name:  "HashiCorp",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "hashicorp",
				RepositoryCount: 296,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HashiCorp-EI_IE1359860.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HashiCorp-Reviews-E1359860.htm",
			},
			OttaProfileSlug: "HashiCorp",
			MainLanguage:    true,
		},
		{
			Name: "Grafana Labs",
			URL:  "https://grafana.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11062162,
				Alias: "grafana-labs",
				Name:  "Grafana Labs",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "grafana",
				RepositoryCount: 233,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grafana-Labs-EI_IE2300269.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grafana-Labs-Reviews-E2300269.htm",
			},
			OttaProfileSlug:   "Grafana-Labs",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCYCwgQAMm9sTJv0rgwQLCxw",
			MainLanguage:      true,
			Vacancies:         nil,
		},
		{
			Name: "Weaviate",
			URL:  "https://weaviate.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11702022,
				Alias: "weaviate-io",
				Name:  "Weaviate",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "weaviate",
				RepositoryCount: 13,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Weaviate-EI_IE7479527.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Weaviate-Reviews-E7479527.htm",
			},
			OttaProfileSlug: "Weaviate",
			MainLanguage:    true,
		},
		{
			Name: "Palantir",
			URL:  "https://www.palantir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20708,
				Alias: "palantir-technologies",
				Name:  "Palantir Technologies",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "palantir",
				RepositoryCount: 57,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Palantir-Technologies-EI_IE236375.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Palantir-Technologies-Reviews-E236375.htm",
			},
			OttaProfileSlug: "Palantir",
			MainLanguage:    false,
		},
		{
			Name: "Okta",
			URL:  "https://www.okta.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    926041,
				Alias: "okta-inc-",
				Name:  "Okta",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "okta",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Okta-EI_IE444756.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Okta-Reviews-E444756.htm",
			},
			OttaProfileSlug:   "Okta",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/FtbF4gpB", "https://app.otta.com/jobs/QcMdOhR_"},
			Skip:              false,
		},
		{
			Name: "1Password",
			URL:  "https://1password.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18648301,
				Alias: "1password",
				Name:  "1Password",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "1Password",
				RepositoryCount: 19,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1Password-EI_IE2984143.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1Password-Reviews-E2984143.htm",
			},
			OttaProfileSlug:   "1Password",
			YouTubeChannelURL: "https://www.youtube.com/@1PasswordVideos",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3905310871/"},
			Skip:              false,
		},
		{
			Name: "Fubo",
			URL:  "https://www.fubo.tv/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5316737,
				Alias: "fubotv",
				Name:  "Fubo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "fubotv",
				RepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FuboTV-EI_IE1006878.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FuboTV-Reviews-E1006878.htm",
			},
			OttaProfileSlug: "fubo",
			MainLanguage:    false,
		},
		{
			Name: "Yassir",
			URL:  "https://yassir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19069709,
				Alias: "yassir",
				Name:  "Yassir",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "YAtechnologies",
				RepositoryCount: 0, // 0
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-YASSIR-EI_IE2601333.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/YASSIR-Reviews-E2601333.htm",
			},
			OttaProfileSlug: "Yassir",
			MainLanguage:    false,
		},
		{
			Name: "DoorDash",
			URL:  "https://doordash.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3205573,
				Alias: "doordash",
				Name:  "DoorDash",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "doordash",
				RepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DoorDash-EI_IE813073.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DoorDash-Reviews-E813073.htm",
			},
			OttaProfileSlug: "DoorDash",
			MainLanguage:    false,
		},
		{
			Name: "BeReal.",
			URL:  "https://bereal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    34731272,
				Alias: "bereal-app",
				Name:  "BeReal.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "BeReal-App",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BeReal-EI_IE7468524.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BeReal-Reviews-E7468524.htm",
			},
			OttaProfileSlug: "BeReal",
			MainLanguage:    false,
		},
		{
			Name: "Vio.com",
			URL:  "https://www.vio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1192098,
				Alias: "viodotcom",
				Name:  "Vio.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "viodotcom",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vio-com-EI_IE940798.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vio-com-Reviews-E940798.htm",
			},
			OttaProfileSlug: "Vio-com",
			MainLanguage:    false, // https://www.linkedin.com/jobs/view/3819771736/
		},
		{
			Name: "Monzo",
			URL:  "https://monzo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9471107,
				Alias: "monzo-bank",
				Name:  "Monzo Bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "monzo",
				RepositoryCount: 76,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Monzo-Bank-EI_IE1557148.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Monzo-Bank-Reviews-E1557148.htm",
			},
			OttaProfileSlug:   "Monzo",
			YouTubeChannelURL: "",
			MainLanguage:      true,
			Vacancies:         []string{"https://app.otta.com/jobs/ZwnXtENr"},
			Skip:              false,
		},
		{
			Name: "Mastercard",
			URL:  "https://www.mastercard.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3015,
				Alias: "mastercard",
				Name:  "Mastercard",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Mastercard",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mastercard-EI_IE3677.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mastercard-Reviews-E3677.htm",
			},
			OttaProfileSlug:   "Mastercard",
			YouTubeChannelURL: "https://www.youtube.com/@MasterCard",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3842471708/", "https://www.linkedin.com/jobs/view/3928338037/"},
		},
		{
			Name: "Cynergy Bank",
			URL:  "https://www.cynergybank.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18921842,
				Alias: "cynergy-bank",
				Name:  "Cynergy Bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynergy-Bank-EI_IE769090.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.co.uk/Reviews/Cynergy-Bank-Reviews-E769090.htm",
			},
			OttaProfileSlug:   "", // NOP
			YouTubeChannelURL: "https://www.youtube.com/@cynergybank",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3917097524/", "https://www.linkedin.com/jobs/view/3868356139/"},
		},
		{
			Name: "Canonical",
			URL:  "https://canonical.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    234280,
				Alias: "canonical",
				Name:  "Canonical",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "canonical",
				RepositoryCount: 83,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Canonical-EI_IE230560.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Canonical-Reviews-E230560.htm",
			},
			OttaProfileSlug: "canonical",
			MainLanguage:    false,
			Skip:            true, // too many vacancies on LinkedIn
		},
		{
			Name: "Vodeno",
			URL:  "https://vodeno.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19016391,
				Alias: "vodeno",
				Name:  "Vodeno",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "vodeno",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vodeno-EI_IE2877672.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vodeno-Reviews-E2877672.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://justjoin.it/offers/vodeno-go-developer-kielce-358668"},
		},
		{
			Name: "Utility Warehouse",
			URL:  "https://uw.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    457903,
				Alias: "utilitywarehouse",
				Name:  "Utility Warehouse",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "utilitywarehouse",
				RepositoryCount: 85,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Utility-Warehouse-EI_IE512935.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Utility-Warehouse-Reviews-E512935.htm",
			},
			OttaProfileSlug: "Utility-Warehouse",
			MainLanguage:    false,
		},
		{
			Name: "Codenotary",
			URL:  "https://codenotary.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    35523736,
				Alias: "codenotary",
				Name:  "Codenotary",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "codenotary",
				RepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CodeNotary-EI_IE3677292.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CodeNotary-Reviews-E3677292.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    true,
		},
		{
			Name: "Audigent",
			URL:  "https://audigent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10642467,
				Alias: "audigent",
				Name:  "Audigent",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "AuDigent",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Audigent-EI_IE5815298.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Audigent-Reviews-E5815298.htm",
			},
			OttaProfileSlug: "Audigent",
			MainLanguage:    false,
		},
		{
			Name: "runZero",
			URL:  "https://www.runzero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    33274038,
				Alias: "runzero",
				Name:  "runZero",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "runZeroInc",
				RepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-runZero-EI_IE7717209.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/runZero-Reviews-E7717209.htm",
			},
			OttaProfileSlug: "runZero",
			MainLanguage:    false,
		},
		{
			Name: "Pinterest",
			URL:  "https://www.pinterest.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1124131,
				Alias: "pinterest",
				Name:  "Pinterest",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "pinterest",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinterest-EI_IE503467.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinterest-Reviews-E503467.htm",
			},
			OttaProfileSlug: "Pinterest",
			MainLanguage:    false,
			Vacancies:       []string{"https://app.otta.com/jobs/JwXeQ-gm"},
		},
		{
			Name: "Tyk",
			URL:  "https://tyk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10775050,
				Alias: "tyk",
				Name:  "Tyk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "TykTechnologies",
				RepositoryCount: 79,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tyk-EI_IE2321465.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tyk-Reviews-E2321465.htm",
			},
			OttaProfileSlug: "Tyk",
			MainLanguage:    true,
		},
		{
			Name: "Sixt",
			URL:  "https://www.sixt.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    17120,
				Alias: "sixt",
				Name:  "SIXT",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Sixt",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sixt-EI_IE10875.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sixt-Reviews-E10875.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    false,
		},
		{
			Name: "OpenTag",
			URL:  "https://theopentag.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20565935,
				Alias: "theopentag",
				Name:  "OpenTag",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "OpenTagOS",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenTag-EI_IE3310869.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenTag-Reviews-E3310869.htm",
			},
			OttaProfileSlug: "",
			MainLanguage:    false,
		},
		{
			Name: "Oxla",
			URL:  "https://oxla.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    79378182,
				Alias: "oxla",
				Name:  "Oxla",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://justjoin.it/offers/oxla-golang-developer-gdansk-362959"},
		},
		{
			Name: "Lightspeed",
			URL:  "https://www.lightspeedhq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1557218,
				Alias: "lightspeedcommerce",
				Name:  "Lightspeed Commerce",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "lightspeed",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lightspeed-EI_IE648762.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lightspeed-Reviews-E648762.htm",
			},
			OttaProfileSlug:   "Lightspeed",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCqOEKwLpolZBcj4LfU3R0Fg",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/RnXM1YTv"},
		},
		{
			Name: "Squarespace",
			URL:  "https://www.squarespace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    265314,
				Alias: "squarespace",
				Name:  "Squarespace",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "squarespace",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Squarespace-EI_IE466343.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Squarespace-Reviews-E466343.htm",
			},
			OttaProfileSlug:   "Squarespace-1",
			YouTubeChannelURL: "https://www.youtube.com/squarespace",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/qAnITcbo"},
		},
		{
			Name: "Curve",
			URL:  "https://curve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10023464,
				Alias: "curve-ltd",
				Name:  "Curve",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Curve-EI_IE1739754.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Curve-Reviews-E1739754.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3872933701/"},
		},
		{
			Name: "Tradevest",
			URL:  "https://www.tradevest.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    92827682,
				Alias: "tradevestgmbh",
				Name:  "Tradevest",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://justjoin.it/offers/tv-development-gmbh-senior-backend-developer"},
		},
		{
			Name: "Samsung",
			URL:  "https://www.samsung.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{1753, 13749561, 3270132, 77752302, 3812597, 3220515, 10633911, 1447520, 3238801, 2711168, 10261221, 87464967, 5089912, 9500165, 10218505, 895705, 78467539, 10332849, 10660176, 79815984, 3290134, 76958044, 85881048, 27127559, 9278177, 37470202, 22316801, 81590082, 11229641, 15213487, 68478415, 5552815, 14472582},
				Alias: "samsung-electronics",
				Name:  "Samsung Electronics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Samsung",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Samsung-Electronics-EI_IE3363.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Samsung-Electronics-Reviews-E3363.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Samsung",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3919334242/", "https://www.linkedin.com/jobs/view/3882780776/", "https://www.linkedin.com/jobs/view/3864617797/"},
		},
		{
			Name: "Woolsocks",
			URL:  "https://woolsocks.eu/",
			LinkedInProfile: domain.LinkedInProfile{

				ID:    79728837,
				Alias: "woolsocks",
				Name:  "Woolsocks",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "amsterdam-platform-creation",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3869628047/"},
		},
		{
			Name: "Applied Systems Canada",
			URL:  "https://www.appliedsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    908801,
				Alias: "applied-systems-canada",
				Name:  "Applied Systems Canada",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Applied-Systems-EI_IE8534.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Applied-Systems-Reviews-E8534.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3879152828/"},
		},
		{
			Name: "Autodesk",
			URL:  "https://www.autodesk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1879,
				Alias: "autodesk",
				Name:  "Autodesk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Autodesk",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Autodesk-EI_IE1155.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Autodesk-Reviews-E1155.htm",
			},
			OttaProfileSlug:   "Autodesk",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3843045776/"},
		},
		{
			Name: "Vonage",
			URL:  "https://www.vonage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{5028, 1345545, 66428, 76778, 2102},
				Alias: "vonage",
				Name:  "Vonage",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Vonage",
				RepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vonage-EI_IE23019.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vonage-Reviews-E23019.htm",
			},
			OttaProfileSlug:   "Vonage",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3818119706/"},
		},
		{
			Name: "OpenWeb",
			URL:  "https://www.openweb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2506872,
				Alias: "openwebhq",
				Name:  "OpenWeb",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenWeb-EI_IE1675932.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenWeb-Reviews-E1675932.htm",
			},
			OttaProfileSlug:   "OpenWeb",
			YouTubeChannelURL: "https://www.youtube.com/@openwebhq",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/x-Xm2wSF"},
		},
		{
			Name: "Timescale",
			URL:  "https://www.timescale.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11065434,
				Alias: "timescaledb",
				Name:  "Timescale",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "timescale",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Timescale-EI_IE2214356.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Timescale-Reviews-E2214356.htm",
			},
			OttaProfileSlug:   "Timescale",
			YouTubeChannelURL: "https://www.youtube.com/@TimescaleDB",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/qZT84OgV"},
		},
		{
			Name: "Arenko",
			URL:  "https://arenko.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10288973,
				Alias: "arenko-cleantech",
				Name:  "Arenko",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Arenko-Group-EI_IE4554199.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Arenko-Group-Reviews-E4554199.htm",
			},
			OttaProfileSlug:   "Arenko-Group",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3889419131/", "https://app.otta.com/jobs/hKvg-QNr"},
		},
		{
			Name: "Kong",
			URL:  "https://konghq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    278819,
				Alias: "konghq",
				Name:  "Kong Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Kong",
				RepositoryCount: 51,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kong-EI_IE801963.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kong-Reviews-E801963.htm",
			},
			OttaProfileSlug:   "Kong",
			YouTubeChannelURL: "https://www.youtube.com/@KongInc",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/7-RRNkSk"},
		},
		{
			Name: "Xata",
			URL:  "https://xata.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69560619,
				Alias: "xataio",
				Name:  "Xata.io",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "xataio",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Xata",
			YouTubeChannelURL: "https://www.youtube.com/@xataio",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3887215661/"},
		},
		{
			Name: "Dojo",
			URL:  "https://dojo.careers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    42391390,
				Alias: "dojo-tech",
				Name:  "Dojo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dojo-EI_IE687810.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dojo-Reviews-E687810.htm",
			},
			OttaProfileSlug:   "Dojo",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3926951138/", "https://www.linkedin.com/jobs/view/3887771635/", "https://app.otta.com/jobs/kdaq4_rK"},
		},
		{
			Name: "ScyllaDB",
			URL:  "https://www.scylladb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10201068,
				Alias: "scylladb",
				Name:  "ScyllaDB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "scylladb",
				RepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ScyllaDB-EI_IE1622223.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ScyllaDB-Reviews-E1622223.htm",
			},
			OttaProfileSlug:   "ScyllaDB",
			YouTubeChannelURL: "https://www.youtube.com/@ScyllaDB",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/Xj1J-uAC"},
		},
		{
			Name: "Unnax",
			URL:  "https://www.unnax.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10295665,
				Alias: "unnax-emi",
				Name:  "Unnax",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "unnax",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unnax-EI_IE2108310.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unnax-Reviews-E2108310.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@unnax-emi",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3890431726/"},
		},
		{
			Name: "AB Tasty",
			URL:  "https://www.abtasty.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1885711,
				Alias: "ab-tasty",
				Name:  "AB Tasty",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AB-Tasty-EI_IE1309242.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AB-Tasty-Reviews-E1309242.htm",
			},
			OttaProfileSlug:   "AB-Tasty",
			YouTubeChannelURL: "https://www.youtube.com/@abtasty",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3872953963/", "https://app.otta.com/jobs/eYKE7Ta_"},
		},
		{
			Name: "Firebolt",
			URL:  "https://www.firebolt.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    40719957,
				Alias: "Firebolt",
				Name:  "Firebolt",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "firebolt-db",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firebolt-EI_IE5001853.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firebolt-Reviews-E5001853.htm",
			},
			OttaProfileSlug:   "Firebolt",
			YouTubeChannelURL: "https://www.youtube.com/@FireboltHQ",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3855486951/"},
		},
		{
			Name: "IBM",
			URL:  "https://www.ibm.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1009,
				Alias: "ibm",
				Name:  "IBM",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "IBM",
				RepositoryCount: 199,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IBM-EI_IE354.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IBM-Reviews-E354.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@IBMTechnology",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3890489682/"},
		},
		{
			Name: "Nine",
			URL:  "https://www.nineforbrands.com.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14933,
				Alias: "nine-entertainment-co.",
				Name:  "Nine",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nine-Entertainment-EI_IE229827.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nine-Entertainment-Reviews-E229827.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3890491488/"},
		},
		{
			Name: "Isovalent",
			URL:  "https://isovalent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    34714411,
				Alias: "isovalent",
				Name:  "Isovalent",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "isovalent",
				RepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Isovalent-EI_IE3180689.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Isovalent-Reviews-E3180689.htm",
			},
			OttaProfileSlug:   "Isovalent",
			YouTubeChannelURL: "https://www.youtube.com/@isovalent",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3893426061/"},
		},
		{
			Name: "ABC Fitness",
			URL:  "https://abcfitness.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    110372,
				Alias: "abc-fitness",
				Name:  "ABC Fitness",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ABC-Fitness-EI_IE28305.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ABC-Fitness-Reviews-E28305.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ABC-Fitness",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3890939506/"},
		},
		{
			Name: "Device42",
			URL:  "https://www.device42.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2345405,
				Alias: "device42",
				Name:  "Device42",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "device42",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Device42-EI_IE1705087.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Device42-Reviews-E1705087.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Device42",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3887639999/"},
		},
		{
			Name: "Acronis",
			URL:  "https://www.acronis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13179,
				Alias: "acronis",
				Name:  "Acronis",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "acronis",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Acronis-EI_IE152824.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Acronis-Reviews-E152824.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Acronis",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3888974596/"},
		},
		{
			Name: "Gcore",
			URL:  "https://gcore.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10351246,
				Alias: "g-core",
				Name:  "Gcore",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Gcore",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gcore-EI_IE2156026.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gcore-Reviews-E2156026.htm",
			},
			OttaProfileSlug:   "Gcore",
			YouTubeChannelURL: "https://www.youtube.com/@GCoreOfficial",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3888070191/"},
		},
		{
			Name: "Zep AI",
			URL:  "https://www.getzep.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    97181599,
				Alias: "zep-ai",
				Name:  "Zep AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3885860612/"},
		},
		{
			Name: "Gelato",
			URL:  "https://www.gelato.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5037871,
				Alias: "gelato",
				Name:  "Gelato",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "gelatoas",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gelato-EI_IE1297272.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gelato-Reviews-E1297272.htm",
			},
			OttaProfileSlug:   "Gelato",
			YouTubeChannelURL: "https://www.youtube.com/@GelatoConnects",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3866739991/"},
		},
		{
			Name: "SumUp",
			URL:  "https://www.sumup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2619512,
				Alias: "sumup",
				Name:  "SumUp",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "sumup",
				RepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SumUp-EI_IE673829.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SumUp-Reviews-E673829.htm",
			},
			OttaProfileSlug:   "sumup",
			YouTubeChannelURL: "https://www.youtube.com/@SumUpGlobal",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3848461191/"},
		},
		{
			Name: "Level Home",
			URL:  "https://level.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    28779237,
				Alias: "levelhome",
				Name:  "Level Home Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "LevelHome",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Level-Home-EI_IE3556695.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Level-Home-Reviews-E3556695.htm",
			},
			OttaProfileSlug:   "Level-Home",
			YouTubeChannelURL: "https://www.youtube.com/@LevelHome",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3897014183/"},
		},
		{
			Name: "Nord Security",
			URL:  "https://nordsecurity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64277258,
				Alias: "nordsecurity",
				Name:  "Nord Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "NordSecurity",
				RepositoryCount: 7, // Rust 22 repositories
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nord-Security-EI_IE4015819.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nord-Security-Reviews-E4015819.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@nordsecurity",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3890883832/", "https://www.linkedin.com/jobs/view/3892388207/"},
		},
		{
			Name: "SonicWall",
			URL:  "https://www.sonicwall.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4926,
				Alias: "sonicwall",
				Name:  "SonicWall",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "sonicwall",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SonicWall-EI_IE9777.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SonicWall-Reviews-E9777.htm",
			},
			OttaProfileSlug:   "SonicWall",
			YouTubeChannelURL: "https://www.youtube.com/@SonicWallInc",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3875719837/"},
		},
		{
			Name: "Pindrop",
			URL:  "https://www.pindrop.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2326557,
				Alias: "pindrop",
				Name:  "Pindrop",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pindrop-EI_IE709157.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pindrop-Reviews-E709157.htm",
			},
			OttaProfileSlug:   "Pindrop",
			YouTubeChannelURL: "https://www.youtube.com/@Pindropsecurity",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3901803458/"},
		},
		{
			Name: "Namecheap",
			URL:  "https://www.namecheap.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    486932,
				Alias: "namecheap-inc",
				Name:  "Namecheap, Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "namecheap",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Namecheap-EI_IE994113.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Namecheap-Reviews-E994113.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@namecheap",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3902607066/"},
		},
		{
			Name: "Seedtag",
			URL:  "https://www.seedtag.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5171806,
				Alias: "seedtag",
				Name:  "Seedtag",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "seedtag",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Seedtag-EI_IE1421858.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Seedtag-Reviews-E1421858.htm",
			},
			OttaProfileSlug:   "Seedtag",
			YouTubeChannelURL: "https://www.youtube.com/@seedtag.advertising",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3853123918/"},
		},
		{
			Name: "Flix",
			URL:  "https://www.flixbus.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2726149,
				Alias: "flixbus",
				Name:  "Flix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "flix-tech",
				RepositoryCount: 17,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flix-EI_IE976463.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flix-Reviews-E976463.htm",
			},
			OttaProfileSlug:   "FlixBus-1",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3901538380/"},
		},
		{
			Name: "Press Ganey",
			URL:  "https://www.pressganey.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18007,
				Alias: "press-ganey-associates",
				Name:  "Press Ganey",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Press-Ganey-EI_IE35100.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Press-Ganey-Reviews-E35100.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PressGaney",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3898197086/"},
		},
		{
			Name: "Atmail",
			URL:  "https://www.atmail.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    918978,
				Alias: "atmail",
				Name:  "Atmail",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3843643222/"},
		},
		{
			Name: "Dusty Robotics",
			URL:  "https://www.dustyrobotics.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    33298433,
				Alias: "dusty-robotics",
				Name:  "Dusty Robotics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dusty-Robotics-EI_IE3518259.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dusty-Robotics-Reviews-E3518259.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dustyrobotics",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3804188613/", "https://www.linkedin.com/jobs/view/3832850412/"},
		},
		{
			Name: "Cimri",
			URL:  "https://www.cimri.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    411498,
				Alias: "cimri",
				Name:  "Cimri",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cimri-EI_IE2401296.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cimri-Reviews-E2401296.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cimri",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3863565726/"},
		},
		{
			Name: "Quadcode",
			URL:  "https://quadcode.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    42345997,
				Alias: "quadcodecareer",
				Name:  "Quadcode",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quadcode-EI_IE3293771.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quadcode-Reviews-E3293771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3907613945/"},
		},
		{
			Name: "FinCompare",
			URL:  "https://fincompare.de/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10976436,
				Alias: "fincompare",
				Name:  "FinCompare",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "fincompare",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FinCompare-EI_IE1644778.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FinCompare-Reviews-E1644778.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3846776554/"},
		},
		{
			Name: "Mellifera",
			URL:  "https://mellifera.team/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    98533069,
				Alias: "mellifera-operations-limited",
				Name:  "Mellifera Operations Limited",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912533167/", "https://www.linkedin.com/jobs/view/3907553857/"},
		},
		{
			Name: "BNP Paribas - Securities Services",
			URL:  "https://securities.cib.bnpparibas/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3625182,
				Alias: "bnpparibassecuritiesservices",
				Name:  "BNP Paribas - Securities Services",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BNP-Paribas-EI_IE10342.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BNP-Paribas-Reviews-E10342.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@labanquedunmondequichange",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3909846279/"},
		},
		{
			Name: "Apifonica",
			URL:  "https://www.apifonica.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10943475,
				Alias: "apifonica",
				Name:  "Apifonica",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "apifonica",
				RepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apifonica-EI_IE1805118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apifonica-Reviews-E1805118.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3909698518/"},
		},
		{
			Name: "Cybus",
			URL:  "https://www.cybus.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7798667,
				Alias: "cybus",
				Name:  "Cybus",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cybusio",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cybus-EI_IE2928520.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cybus-Reviews-E2928520.htm",
			},
			OttaProfileSlug:   "Cybus",
			YouTubeChannelURL: "https://www.youtube.com/@cybus_io",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3905530455/"},
		},
		{
			Name: "Flink",
			URL:  "https://www.goflink.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71241902,
				Alias: "goflink",
				Name:  "Flink",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flink-EI_IE4921496.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flink-Reviews-E4921496.htm",
			},
			OttaProfileSlug:   "Flink",
			YouTubeChannelURL: "https://www.youtube.com/@flink7309",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3906306472/"},
		},
		{
			Name: "Greenbone",
			URL:  "https://www.greenbone.net/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    697428,
				Alias: "greenbone-networks-gmbh",
				Name:  "Greenbone AG",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "greenbone",
				RepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Greenbone-Networks-EI_IE379229.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Greenbone-Networks-Reviews-E379229.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@greenbone",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3892843872/"},
		},
		{
			Name: "Microsoft",
			URL:  "https://www.microsoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1035,
				Alias: "microsoft",
				Name:  "Microsoft",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "microsoft",
				RepositoryCount: 78,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Microsoft-EI_IE1651.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Microsoft-Reviews-E1651.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Microsoft",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3908309859/", "https://www.linkedin.com/jobs/view/3905989512/"},
		},
		{
			Name: "OLX",
			URL:  "https://www.olxgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    167557,
				Alias: "olx-group",
				Name:  "OLX",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OLX-Group-EI_IE517166.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OLX-Group-Reviews-E517166.htm",
			},
			OttaProfileSlug:   "OLX-Group",
			YouTubeChannelURL: "https://www.youtube.com/@OLXGroup",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3904643356/"},
		},
		{
			Name: "Snyk",
			URL:  "https://snyk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10043614,
				Alias: "snyk",
				Name:  "Snyk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "snyk",
				RepositoryCount: 22,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snyk-EI_IE2094989.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snyk-Reviews-E2094989.htm",
			},
			OttaProfileSlug:   "Snyk",
			YouTubeChannelURL: "https://www.youtube.com/@Snyksec",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3903469567/", "https://app.otta.com/jobs/AOeyI9AM"},
		},
		{
			Name: "Siemens",
			URL:  "https://www.siemens.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{1043, 157241},
				Alias: "siemens",
				Name:  "Siemens",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "siemens",
				RepositoryCount: 19,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Siemens-EI_IE3510.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Siemens-Reviews-E3510.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@SiemensKnowledgeHub",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3901324430/"},
		},
		{
			Name: "Sinch",
			URL:  "https://www.sinch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3726743,
				Alias: "sinch",
				Name:  "Sinch",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "sinch",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sinch-EI_IE788805.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sinch-Reviews-E788805.htm",
			},
			OttaProfileSlug:   "sinch",
			YouTubeChannelURL: "https://www.youtube.com/@WeAreSinch",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3903797937/"},
		},
		{
			Name: "FOX Tech",
			URL:  "https://tech.fox.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14850572,
				Alias: "foxtechteam",
				Name:  "FOX Tech",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FOX-Broadcasting-EI_IE13272.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FOX-Broadcasting-Reviews-E13272.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3903338524/"},
		},
		{
			Name: "SailPoint",
			URL:  "https://www.sailpoint.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    47456,
				Alias: "sailpoint-technologies",
				Name:  "SailPoint",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "sailpoint-oss",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SailPoint-Technologies-EI_IE449696.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SailPoint-Technologies-Reviews-E449696.htm",
			},
			OttaProfileSlug:   "Sailpoint",
			YouTubeChannelURL: "https://www.youtube.com/@SailPointTechnologies",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3903153316/"},
		},
		{
			Name: "Proofpoint",
			URL:  "https://www.proofpoint.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11681,
				Alias: "proofpoint",
				Name:  "Proofpoint",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "proofpoint",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proofpoint-EI_IE39140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proofpoint-Reviews-E39140.htm",
			},
			OttaProfileSlug:   "Proofpoint",
			YouTubeChannelURL: "https://www.youtube.com/@proofpoint",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3797062001/"},
		},
		{
			Name: "Medium",
			URL:  "https://medium.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3039001,
				Alias: "medium-com",
				Name:  "Medium",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Medium",
				RepositoryCount: 20,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Medium-EI_IE784883.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Medium-Reviews-E784883.htm",
			},
			OttaProfileSlug:   "Medium",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3910516815/"},
		},
		{
			Name: "Asset Reality",
			URL:  "https://www.assetreality.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    42805677,
				Alias: "asset-reality",
				Name:  "Asset Reality",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Asset-Reality",
			YouTubeChannelURL: "https://www.youtube.com/@assetreality",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3897435235/"},
		},
		{
			Name: "Limango",
			URL:  "https://www.limango.pl/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2966982,
				Alias: "limango-sp-z-o-o-",
				Name:  "Limango",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Limango-Polska-EI_IE2884426.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Limango-Polska-Reviews-E2884426.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@limangoPolska",
			MainLanguage:      false,
			Vacancies:         []string{"https://justjoin.it/offers/limango-polska-mid-senior-backend-developer-golang"},
		},
		{
			Name: "RxBenefits",
			URL:  "https://www.rxbenefits.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2873210,
				Alias: "rxbenefits-inc-",
				Name:  "RxBenefits, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RxBenefits-EI_IE1175839.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RxBenefits-Reviews-E1175839.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3910221910/"},
		},
		{
			Name: "SmithRx",
			URL:  "https://www.smithrx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10886362,
				Alias: "smithrx",
				Name:  "SmithRx",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SmithRx-EI_IE1901555.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SmithRx-Reviews-E1901555.htm",
			},
			OttaProfileSlug:   "SmithRx",
			YouTubeChannelURL: "https://www.youtube.com/@SmithRxPBM",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3910420916/"},
		},
		{
			Name: "k-ID",
			URL:  "https://www.k-id.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    45973117,
				Alias: "k-id",
				Name:  "k-ID",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@k-IDofficial",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912660665/"},
		},
		{
			Name: "CAFU",
			URL:  "https://www.cafu.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13892019,
				Alias: "mycafu",
				Name:  "CAFU",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CAFU-EI_IE3713615.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CAFU-Reviews-E3713615.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3910059830/"},
		},
		{
			Name: "Rollee",
			URL:  "https://www.getrollee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    76353840,
				Alias: "rollee",
				Name:  "Rollee",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "Rollee",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3908431251/", "https://app.otta.com/jobs/4ajPfsAU"},
		},
		{
			Name: "Net2Phone",
			URL:  "https://www.net2phone.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2965,
				Alias: "net2phone",
				Name:  "Net2Phone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "net2phone",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Net2Phone-EI_IE9342.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Net2Phone-Reviews-E9342.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914839595/"},
		},
		{
			Name: "Ola Chat",
			URL:  "https://olachat.sg/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68332088,
				Alias: "ola-chat",
				Name:  "Ola Chat",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "olachat",
				RepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912130901/"},
		},
		{
			Name: "Veracity Software Inc",
			URL:  "https://veracity-us.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11137552,
				Alias: "veracitysoftwareinc",
				Name:  "Veracity Software Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Veracity-Software-EI_IE1357198.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Veracity-Software-Reviews-E1357198.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3913559527/"},
		},
		{
			Name: "Salesforge",
			URL:  "https://www.salesforge.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    91706116,
				Alias: "salesforgeai",
				Name:  "Salesforge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@salesforge",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912551615/"},
		},
		{
			Name: "Treecard",
			URL:  "https://www.treecard.org/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68821773,
				Alias: "treecardapp",
				Name:  "Treecard",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "TreeCard",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Treecard-EI_IE5675051.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Treecard-Reviews-E5675051.htm",
			},
			OttaProfileSlug:   "Treecard",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3909975571/"},
		},
		{
			Name: "Openprovider",
			URL:  "https://www.openprovider.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    550698,
				Alias: "openprovider",
				Name:  "Openprovider",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "openprovider",
				RepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Openprovider-EI_IE1267124.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Openprovider-Reviews-E1267124.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@openprovider453",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914419419/"},
		},
		{
			Name: "fiskaly",
			URL:  "https://www.fiskaly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18929063,
				Alias: "fiskaly",
				Name:  "fiskaly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "fiskaly",
				RepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-fiskaly-EI_IE3059515.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/fiskaly-Reviews-E3059515.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fiskaly",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914242491/"},
		},
		{
			Name: "Vay",
			URL:  "https://vay.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12584218,
				Alias: "vaytechnology",
				Name:  "Vay",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Reemote",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Vay",
			YouTubeChannelURL: "https://www.youtube.com/@vay_io",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912106948/"},
		},
		{
			Name: "Fortinet",
			URL:  "https://www.fortinet.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6460,
				Alias: "fortinet",
				Name:  "Fortinet",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "fortinet",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fortinet-EI_IE23128.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fortinet-Reviews-E23128.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fortinet",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3907970778/"},
		},
		{
			Name: "GitHub",
			URL:  "https://github.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1418841,
				Alias: "github",
				Name:  "GitHub",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "github",
				RepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitHub-EI_IE671945.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitHub-Reviews-E671945.htm",
			},
			OttaProfileSlug:   "Github",
			YouTubeChannelURL: "https://www.youtube.com/@GitHub",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914880703/"},
		},
		{
			Name: "Voltus",
			URL:  "https://www.voltus.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10866628,
				Alias: "voltus-inc.",
				Name:  "Voltus",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "voltusdev",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voltus-EI_IE2090197.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voltus-Reviews-E2090197.htm",
			},
			OttaProfileSlug:   "Voltus",
			YouTubeChannelURL: "https://www.youtube.com/@Voltusinc",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/pB_hGf0W"},
		},
		{
			Name: "Grammarly",
			URL:  "https://www.grammarly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1253671,
				Alias: "grammarly",
				Name:  "Grammarly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "grammarly",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grammarly-EI_IE636873.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grammarly-Reviews-E636873.htm",
			},
			OttaProfileSlug:   "Grammarly",
			YouTubeChannelURL: "https://www.youtube.com/@grammarly",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/jonw8YN_"},
		},
		{
			Name: "Stonebranch",
			URL:  "https://www.stonebranch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71261,
				Alias: "stonebranch",
				Name:  "Stonebranch",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "stonebranch-marketplace",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stonebranch-EI_IE408996.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stonebranch-Reviews-E408996.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Stonebranch",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914103960/"},
		},
		{
			Name: "Rapid7",
			URL:  "https://www.rapid7.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39624,
				Alias: "rapid7",
				Name:  "Rapid7",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "rapid7",
				RepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rapid7-EI_IE243542.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rapid7-Reviews-E243542.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@GoRapid7",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3915341496/"},
		},
		{
			Name: "Toggle AI",
			URL:  "https://toggle.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    28508827,
				Alias: "toggle-ai",
				Name:  "Toggle AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Toggle-EI_IE3924898.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Toggle-Reviews-E3924898.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ToggleAI-Investing",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3915194291/"},
		},
		{
			Name: "hearX Group",
			URL:  "https://www.hearxgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18002825,
				Alias: "hearx-group",
				Name:  "hearX Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "hearxgroup",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-hearX-Group-EI_IE5800566.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/hearX-Group-Reviews-E5800566.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hearxgroupptyltd8061",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3913102407/"},
		},
		{
			Name: "MarkiTech.AI",
			URL:  "https://markitech.ca/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9294422,
				Alias: "markitech-ai",
				Name:  "MarkiTech.AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Markitech-EI_IE4190913.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Markitech-Reviews-E4190913.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@markitech-digitaltransform6173",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912880597/"},
		},
		{
			Name: "Lantronix",
			URL:  "https://www.lantronix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12612,
				Alias: "lantronix",
				Name:  "Lantronix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Lantronix",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lantronix-EI_IE5498.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lantronix-Reviews-E5498.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@LantronixInc",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3915346698/"},
		},
		{
			Name: "Citi",
			URL:  "https://www.citi.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11448,
				Alias: "citi",
				Name:  "Citi",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Citi",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/citi",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Citi-Reviews-E8843.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Citi",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912849197/"},
		},
		{
			Name: "INFOLOB",
			URL:  "https://www.infolob.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    404211,
				Alias: "infolob-global",
				Name:  "INFOLOB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infolob-Solutions-EI_IE423113.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infolob-Solutions-Reviews-E423113.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914415181/"},
		},
		{
			Name: "Argela Technologies",
			URL:  "https://www.argela.com.tr/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    26805,
				Alias: "argela-technologies",
				Name:  "Argela Technologies",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Argela-EI_IE389528.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Argela-Reviews-E389528.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@argelatechnologies",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914290466/"},
		},
		{
			Name: "Top Doctors",
			URL:  "https://topdoctors.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3035481,
				Alias: "top-doctors-europe",
				Name:  "Top Doctors",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Top-Doctors-EI_IE1712721.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Top-Doctors-Reviews-E1712721.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TopDoctorsUK",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912190311/"},
		},
		{
			Name: "Recurly",
			URL:  "https://recurly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    810383,
				Alias: "recurly",
				Name:  "Recurly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "recurly",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Recurly-EI_IE692611.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Recurly-Reviews-E692611.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Recurly",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912991626/"},
		},
		{
			Name: "Cynet Systems",
			URL:  "https://www.cynetsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5003556,
				Alias: "cynet-systems",
				Name:  "Cynet Systems",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynet-Systems-EI_IE654628.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cynet-Systems-Reviews-E654628.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cynetsystems2026",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912977786/"},
		},
		{
			Name: "Odyssey Information Services",
			URL:  "https://www.odysseyis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    56445,
				Alias: "odyssey-information-services",
				Name:  "Odyssey Information Services",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Odyssey-Information-Services-EI_IE558201.11,39.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Odyssey-Information-Services-Reviews-E558201.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@odysseyinformationservices1641",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3905563294/"},
		},
		{
			Name: "Infomatics Corp",
			URL:  "https://infomaticscorp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    821065,
				Alias: "infomatics-corp",
				Name:  "Infomatics Corp",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infomatics-EI_IE925978.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infomatics-Reviews-E925978.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912956713/"},
		},
		{
			Name: "Mindera",
			URL:  "https://mindera.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80044268,
				Alias: "mindera-world",
				Name:  "Mindera",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Mindera",
				RepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mindera-EI_IE1139926.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mindera-Reviews-E1139926.htm",
			},
			OttaProfileSlug:   "Mindera",
			YouTubeChannelURL: "https://www.youtube.com/@MinderaSoftware",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3915092213/"},
		},
		{
			Name: "Sytac",
			URL:  "https://sytac.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    626751,
				Alias: "sytac",
				Name:  "Sytac",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "sytac",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sytac-EI_IE1255983.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sytac-Reviews-E1255983.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@sytac",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3912222258/"},
		},
		{
			Name: "Qumulus Cloud Platform",
			URL:  "https://www.qumulus.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    81905634,
				Alias: "qumuluscloudplatform",
				Name:  "Qumulus Cloud Platform",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "QumulusTechnology",
				RepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914358799/"},
		},
		{
			Name: "Saxon AI",
			URL:  "https://saxon.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    224935,
				Alias: "saxonai",
				Name:  "Saxon AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@saxonai",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914185384/"},
		},
		{
			Name: "Dyninno Group",
			URL:  "https://dyninno.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9456141,
				Alias: "dyninno-group",
				Name:  "Dyninno Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DYNINNO-Group-EI_IE2562842.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DYNINNO-Group-Reviews-E2562842.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dyninnogroup2702",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914135515/"},
		},
		{
			Name: "Holland & Barrett",
			URL:  "https://www.hollandandbarrett.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    815488,
				Alias: "holland-&-barrett",
				Name:  "Holland & Barrett",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Holland-and-Barrett-EI_IE36174.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Holland-and-Barrett-Reviews-E36174.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hollandandbarrett",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3911059005/"},
		},
		{
			Name: "Group Avows",
			URL:  "https://avowstech.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3651016,
				Alias: "group-avows",
				Name:  "Group Avows",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AVOWS-Technologies-EI_IE870406.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AVOWS-Technologies-Reviews-E870406.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@avowsgroupofficial9907",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3911007016/"},
		},
		{
			Name: "Cognizant",
			URL:  "https://www.cognizant.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1680,
				Alias: "cognizant",
				Name:  "Cognizant",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognizant-Technology-Solutions-EI_IE8014.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognizant-Technology-Solutions-Reviews-E8014.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cognizant",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3914839625/"},
		},
		{
			Name: "DigitalOcean",
			URL:  "https://www.digitalocean.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2601253,
				Alias: "digitalocean",
				Name:  "DigitalOcean",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "digitalocean",
				RepositoryCount: 116,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DigitalOcean-EI_IE823482.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DigitalOcean-Reviews-E823482.htm",
			},
			OttaProfileSlug:   "DigitalOcean",
			YouTubeChannelURL: "https://www.youtube.com/@DigitalOcean",
			MainLanguage:      true,
			Vacancies:         []string{"https://app.otta.com/jobs/bwHc4GXE"},
		},
		{
			Name: "Snap",
			URL:  "https://snap.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15191764,
				Alias: "snap-inc-co",
				Name:  "Snap Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Snapchat",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snap-EI_IE671946.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snap-Reviews-E671946.htm",
			},
			OttaProfileSlug:   "Snap",
			YouTubeChannelURL: "https://www.youtube.com/@snapinc",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/hr1D1_FZ"},
		},
		{
			Name: "Nuro",
			URL:  "https://www.nuro.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12957486,
				Alias: "nuro-inc.",
				Name:  "Nuro",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nuro-EI_IE1550693.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nuro-Reviews-E1550693.htm",
			},
			OttaProfileSlug:   "Nuro",
			YouTubeChannelURL: "https://www.youtube.com/@NuroTeam",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/P5a_50Xb"},
		},
		{
			Name: "CloudWalk",
			URL:  "https://cloudwalk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3523168,
				Alias: "cloudwalk-inc",
				Name:  "CloudWalk, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "cloudwalk",
				RepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudWalk-Inc-EI_IE2722308.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudWalk-Inc-Reviews-E2722308.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cloudwalk_shorts",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3917078597/"},
		},
		{
			Name: "Transition Technologies PSC",
			URL:  "https://ttpsc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    17880075,
				Alias: "transition-technologies-psc",
				Name:  "Transition Technologies PSC",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Transition-Technologies-PSC-EI_IE1875542.11,38.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Transition-Technologies-PSC-Reviews-E1875542.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TransitionTechnologiesPSC",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3886949034/"},
		},
		{
			Name: "Kroger",
			URL:  "https://www.kroger.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4914,
				Alias: "kroger",
				Name:  "Kroger",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "Kroger-Technology",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kroger-EI_IE386.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kroger-Reviews-E386.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@KrogerCo",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3917025739/"},
		},
		{
			Name: "Precisely",
			URL:  "https://www.precisely.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64863146,
				Alias: "preciselydata",
				Name:  "Precisely",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "PreciselyData",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Precisely-EI_IE3372755.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Precisely-Reviews-E3372755.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PreciselyData",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3917000992/"},
		},
		{
			Name: "R Systems",
			URL:  "https://www.rsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    165636,
				Alias: "r-systems",
				Name:  "R Systems",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-R-Systems-EI_IE32444.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/R-Systems-Reviews-E32444.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@RSystems_inc",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3916555239/"},
		},
		{
			Name: "Hays",
			URL:  "https://www.haysplc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3486,
				Alias: "hays",
				Name:  "Hays",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hays-EI_IE10166.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hays-Reviews-E10166.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@HaysTV",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3913846166/"},
		},
		{
			Name: "Consort Group",
			URL:  "https://consort-group.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    46088,
				Alias: "consortgroup",
				Name:  "Consort Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Groupe-Consort-EI_IE915503.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Groupe-Consort-Reviews-E915503.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@consortgroup",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3913831485/"},
		},
		{
			Name: "Ascendion",
			URL:  "https://ascendion.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    86694680,
				Alias: "ascendion",
				Name:  "Ascendion",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "", // NOP
				RepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ascendion-EI_IE7774544.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ascendion-Reviews-E7774544.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ascendioninc",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3916325061/"},
		},
		{
			Name: "American Express",
			URL:  "https://www.americanexpress.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1277,
				Alias: "american-express",
				Name:  "American Express",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "americanexpress",
				RepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-American-Express-EI_IE35.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/American-Express-Reviews-E35.htm",
			},
			OttaProfileSlug:   "American-Express",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3926321626/"},
		},
		{
			Name: "Wise",
			URL:  "https://wise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1769571,
				Alias: "wiseaccount",
				Name:  "Wise",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "transferwise",
				RepositoryCount: 35,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wise-EI_IE637715.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wise-Reviews-E637715.htm",
			},
			OttaProfileSlug:   "Wise",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://app.otta.com/jobs/UVRcHQhe"},
		},
		{
			Name: "Checkout.com",
			URL:  "https://www.checkout.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3110635,
				Alias: "checkout",
				Name:  "Checkout.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "checkout",
				RepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Checkout-com-EI_IE837487.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Checkout-com-Reviews-E837487.htm",
			},
			OttaProfileSlug:   "Checkout-com",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3920739722/"},
		},
		{
			Name: "Atom bank",
			URL:  "https://www.atombank.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5197064,
				Alias: "atom-bank",
				Name:  "Atom bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:           "atombank",
				RepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Atom-Bank-EI_IE1354887.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Atom-Bank-Reviews-E1354887.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			MainLanguage:      false,
			Vacancies:         []string{"https://www.linkedin.com/jobs/view/3921514553/"},
		},
		//{
		//	Name: "",
		//	URL:  "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:           "",
		//		RepositoryCount: 0,
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	MainLanguage:      false,
		//	Vacancies:         []string{""},
		//},
	}
}
