package db

import "github.com/readytotouch/gocompanies/internal/domain"

func Universities() []domain.University {
	// https://dou.ua/lenta/articles/ukrainian-universities-2023/
	// УКУ +
	// ХНЕУ +
	// КНУ ім. Шевченка +
	// НаУКМА +
	// ОНУ ім. Мечникова +
	// ДУТ https://www.linkedin.com/company/duikt/
	// СумДУ +
	// НТУУ «КПІ ім. І. Сікорського» +
	// ЛНУ ім. Франка +
	// ХНУРЕ +
	// НУ «Львівська політехніка» +
	// ДНУ ім. Гончара +
	// НТУ «Дніпровська політехніка» +
	// ХНУ [м. Хмельницький] +
	// КНЕУ +
	// НУ «Одеська політехніка» (ОНПУ) +
	// ХАІ +
	// ЧНУ ім. Федьковича +
	// ХНУ ім. Каразіна +
	// НТУ «ХПІ» +

	return []domain.University{
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    364340,
					Alias: "taras-shevchenko-national-university-of-kyiv",
					Name:  "Taras Shevchenko National University of Kyiv",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    496320,
					Alias: "national-university-of-kyiv-mohyla-academy",
					Name:  "National University of Kyiv-Mohyla Academy",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    782774,
					Alias: "lviv-polytechnic-national-university",
					Name:  "Lviv Polytechnic National University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    818029,
					Alias: "ukrainian-catholic-university",
					Name:  "Ukrainian Catholic University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    850102,
					Alias: "kharkiv-national-university-of-economics/",
					Name:  "Simon Kuznets Kharkiv National University of Economics",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    1257361,
					Alias: "sumy-state-university",
					Name:  "Sumy State University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15099424,
					Alias: "kharkiv-national-university-of-radioelectronics",
					Name:  "Kharkiv National University of Radioelectronics",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15099711,
					Alias: "national-technical-university-kharkiv-polytechnic-institute-",
					Name:  `National Technical University "Kharkiv Polytechnic Institute"`,
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15101057,
					Alias: "national-aviation-university",
					Name:  "National Aviation University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15101979,
					Alias: "dnipropetrovs'kij-nacional'nij-universitet",
					Name:  "Dnipropetrovs'kij Nacional'nij Universitet",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15102004,
					Alias: "vinnytsia-national-technical-university",
					Name:  "Vinnytsia National Technical University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15250306,
					Alias: "national-technical-university-of-ukraine-'kyiv-pol",
					Name:  "National Technical University of Ukraine 'Kyiv Polytechnic Institute'",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15251128,
					Alias: "national-university-'ivan-franko'%E2%80%8B-lviv",
					Name:  "Ivan Franko National University of Lviv",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    6261241,
					Alias: "odessa-national-polytechnic-university",
					Name:  "Odessa National Polytechnic University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    658198,
					Alias: "national-aerospace-university",
					Name:  "National Aerospace University -'Kharkiv Aviation Institute'",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15099425,
					Alias: "v.-n.-karazin-kharkiv-national-university",
					Name:  "V. N. Karazin Kharkiv National University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    11443062,
					Alias: "chnu",
					Name:  "Yuriy Fedkovych Chernivtsi National University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    80424966,
					Alias: "khnu",
					Name:  "Khmelnitsky National University",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    15101061,
					Alias: "dnipro-university-of-technology",
					Name:  "Dnipro University of Technology",
				},
			},
		},
		{
			LinkedInProfiles: []domain.LinkedInProfile{
				{
					ID:    1198954,
					Alias: "odessa-national-university",
					Name:  "Odessa I.I.Mechnikov National University",
				},
			},
		},
		//{
		//	LinkedInProfiles: []domain.LinkedInProfile{
		//		{
		//			ID:    0,
		//			Alias: "",
		//			Name:  "",
		//		},
		//	},
		//},
	}
}
