package db

import "github.com/readytotouch/gocompanies/internal/domain"

func FaangCompanyGroup() domain.FaangCompanyGroup {
	return domain.FaangCompanyGroup{
		FaangCompanies: []domain.LinkedInProfile{
			{
				ID: 0,
				IDs: []int{
					10667,    // Meta
					289891,   // WhatsApp
					2289109,  // Instagram
					76987811, // Facebook
				},
				Alias: "meta",
				Name:  "Meta",
			},
			{
				ID: 0,
				IDs: []int{
					162479, // Apple
				},
				Alias: "apple",
				Name:  "Apple",
			},
			{
				ID: 0,
				IDs: []int{
					1586,     // Amazon
					2320329,  // Twitch
					2382910,  // Amazon Web Services (AWS)
					4787585,  // Ring
					11091426, // Prime Video
					78392228, // Amazon Games
					80073065, // Amazon Music
				},
				Alias: "amazon",
				Name:  "Amazon",
			},
			{
				ID: 0,
				IDs: []int{
					165158, // Netflix
				},
				Alias: "netflix",
				Name:  "Netflix",
			},
			{
				ID: 0,
				IDs: []int{
					1441,  // Google
					16140, // YouTube
				},
				Alias: "google",
				Name:  "Google",
			},
			{
				ID: 0,
				IDs: []int{
					1035,    // Microsoft
					1337,    // LinkedIn
					1418841, // GitHub
				},
				Alias: "microsoft",
				Name:  "Microsoft",
			},
			{
				ID: 0,
				IDs: []int{
					3608, // NVIDIA
				},
				Alias: "nvidia",
				Name:  "NVIDIA",
			},
		},
		OtherCompanies: []domain.LinkedInProfile{
			// https://www.linkedin.com/company/openai/

			// https://www.linkedin.com/company/intel-corporation/
			// https://www.linkedin.com/company/amd/

			// https://www.linkedin.com/company/oracle/

			// https://www.linkedin.com/company/ibm/
			// https://www.linkedin.com/company/hashicorp/

			// https://www.linkedin.com/company/uber-com/

			// https://www.linkedin.com/company/gitlab-com/

			// https://www.linkedin.com/company/red-hat/
		},
	}
}
