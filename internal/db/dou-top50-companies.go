package db

import "github.com/readytotouch/golang-companies-organizer/internal/domain"

func DouTop50Companies() []domain.DouCompany {
	// top50.js
	return []domain.DouCompany{
		// EPAM Ukraine https://jobs.dou.ua/companies/epam-systems/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4972,
				IDs:   nil,
				Alias: "epam-systems",
				Name:  "EPAM Systems",
			},
			EmployeeCount: 9600,
		},
		// SoftServe https://jobs.dou.ua/companies/softserve/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6549,
				IDs:   nil,
				Alias: "softserve",
				Name:  "SoftServe",
			},
			EmployeeCount: 7335,
		},
		// GlobalLogic Ukraine https://jobs.dou.ua/companies/globallogic/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    164008,
				IDs:   nil,
				Alias: "globallogic",
				Name:  "GlobalLogic",
			},
			EmployeeCount: 6026,
		},
		// Ajax Systems https://jobs.dou.ua/companies/ajax-systems/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10469202,
				IDs:   nil,
				Alias: "ajax-systems",
				Name:  "Ajax Systems",
			},
			EmployeeCount: 3411,
		},
		// Luxoft https://jobs.dou.ua/companies/luxoft/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5300,
				IDs:   nil,
				Alias: "luxoft",
				Name:  "Luxoft",
			},
			EmployeeCount: 3400,
		},
		// Evoplay https://jobs.dou.ua/companies/evoplay/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	EmployeeCount: 2903,
		//},
		// Genesis https://jobs.dou.ua/companies/genesis-technology-partners/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2416279,
				IDs:   nil,
				Alias: "genesis-technology-partners",
				Name:  "Genesis Tech",
			},
			EmployeeCount: 2789,
		},
		// Intellias https://jobs.dou.ua/companies/intellias/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    49445,
				IDs:   nil,
				Alias: "intellias",
				Name:  "Intellias",
			},
			EmployeeCount: 2305,
		},
		// DataArt https://jobs.dou.ua/companies/dataart/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10882,
				IDs:   nil,
				Alias: "dataart",
				Name:  "DataArt",
			},
			EmployeeCount: 2300,
		},
		// ZONE3000 https://jobs.dou.ua/companies/zone3000/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7970644,
				IDs:   nil,
				Alias: "zone-3000",
				Name:  "ZONE3000",
			},
			EmployeeCount: 2290,
		},
		// Ciklum https://jobs.dou.ua/companies/ciklum/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9183,
				IDs:   nil,
				Alias: "ciklum",
				Name:  "Ciklum",
			},
			EmployeeCount: 2170,
		},
		// NIX https://www.nixsolutions.com/ru/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	EmployeeCount: 1960,
		//},
		// Sigma Software https://jobs.dou.ua/companies/sigma-software/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68326,
				IDs:   nil,
				Alias: "sigma-software-group",
				Name:  "Sigma Software Group",
			},
			EmployeeCount: 1870,
		},
		// Infopulse & Tietoevry Ukraine https://jobs.dou.ua/companies/infopulse/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    8642,
				IDs:   nil,
				Alias: "infopulse",
				Name:  "Infopulse",
			},
			EmployeeCount: 1827,
		},
		// N-iX https://jobs.dou.ua/companies/n-ix/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    31921,
				IDs:   nil,
				Alias: "n-ix",
				Name:  "N-iX",
			},
			EmployeeCount: 1640,
		},
		// ELEKS https://jobs.dou.ua/companies/eleks/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13412,
				IDs:   nil,
				Alias: "eleks",
				Name:  "ELEKS",
			},
			EmployeeCount: 1606,
		},
		// Capgemini Engineering https://jobs.dou.ua/companies/lohika-systems/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    72092703,
				IDs:   nil,
				Alias: "capgemini-engineering",
				Name:  "Capgemini Engineering",
			},
			EmployeeCount: 1115,
		},
		// Netpeak Group https://jobs.dou.ua/companies/netpeak-group/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11869488,
				IDs:   nil,
				Alias: "netpeak-group",
				Name:  "Netpeak Group",
			},
			EmployeeCount: 1096,
		},
		// AUTODOC https://jobs.dou.ua/companies/autodoc/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7703911,
				IDs:   nil,
				Alias: "autodoc",
				Name:  "AUTODOC",
			},
			EmployeeCount: 1092,
		},
		// Avenga https://jobs.dou.ua/companies/avenga/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    52144789,
				IDs:   nil,
				Alias: "avenga",
				Name:  "Avenga",
			},
			EmployeeCount: 1019,
		},
		// SKELAR https://jobs.dou.ua/companies/skelar/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    34230094,
				IDs:   nil,
				Alias: "skelartech",
				Name:  "SKELAR",
			},
			EmployeeCount: 1000,
		},
		// ALLSTARSIT https://jobs.dou.ua/companies/allstars-it/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5019449,
				IDs:   nil,
				Alias: "allstarsit",
				Name:  "ALLSTARSIT",
			},
			EmployeeCount: 987,
		},
		// Temabit Fozzy Group https://jobs.dou.ua/companies/fozzy/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    44542,
				IDs:   nil,
				Alias: "fozzy-group",
				Name:  "Fozzy Group",
			},
			EmployeeCount: 954,
		},
		// Playrix https://jobs.dou.ua/companies/playrix/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    213036,
				IDs:   nil,
				Alias: "playrix-entertainment",
				Name:  "Playrix",
			},
			EmployeeCount: 932,
		},
		// Grid Dynamics Group https://jobs.dou.ua/companies/grid-dynamics/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    250774,
				IDs:   nil,
				Alias: "grid-dynamics",
				Name:  "Grid Dynamics",
			},
			EmployeeCount: 900,
		},
		// Intecracy Group https://jobs.dou.ua/companies/intecracy-group-consortium/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18440329,
				IDs:   nil,
				Alias: "intecracygroup",
				Name:  "Intecracy Group",
			},
			EmployeeCount: 887,
		},
		// EVO https://jobs.dou.ua/companies/evo/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1267199,
				IDs:   nil,
				Alias: "evo.company",
				Name:  "EVO",
			},
			EmployeeCount: 872,
		},
		// Metinvest Digital https://jobs.dou.ua/companies/metinvest-digital/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39306,
				IDs:   nil,
				Alias: "metinvest",
				Name:  "Metinvest",
			},
			EmployeeCount: 858,
		},
		// SQUAD https://jobs.dou.ua/companies/squad/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    75409516,
				IDs:   nil,
				Alias: "squad-ua",
				Name:  "SQUAD",
			},
			EmployeeCount: 836,
		},
		// Trinetix https://jobs.dou.ua/companies/trinetix/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5382653,
				IDs:   nil,
				Alias: "trinetix-inc",
				Name:  "Trinetix",
			},
			EmployeeCount: 823,
		},
		// Room 8 Group https://jobs.dou.ua/companies/room-8-group/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71267943,
				IDs:   nil,
				Alias: "room-8-group",
				Name:  "Room 8 Group",
			},
			EmployeeCount: 815,
		},
		// Playtech https://jobs.dou.ua/companies/playtech/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18717,
				IDs:   nil,
				Alias: "playtech",
				Name:  "Playtech",
			},
			EmployeeCount: 809,
		},
		// Netcracker https://jobs.dou.ua/companies/netcracker/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    165155,
				IDs:   nil,
				Alias: "netcrackertech",
				Name:  "Netcracker Technology",
			},
			EmployeeCount: 800,
		},
		// MEGOGO https://jobs.dou.ua/companies/megogonet-/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2582114,
				IDs:   nil,
				Alias: "megogo-net",
				Name:  "MEGOGO",
			},
			EmployeeCount: 790,
		},
		// Onseo https://jobs.dou.ua/companies/onseo/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7073475,
				IDs:   nil,
				Alias: "onseo",
				Name:  "Onseo",
			},
			EmployeeCount: 770,
		},
		// GR8 Tech https://jobs.dou.ua/companies/gr8-tech/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	EmployeeCount: 721,
		//},
		// Uklon https://jobs.dou.ua/companies/uklon/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3177450,
				IDs:   nil,
				Alias: "uklontech",
				Name:  "Uklon",
			},
			EmployeeCount: 720,
		},
		// Ubisoft Ukraine https://jobs.dou.ua/companies/ubisoft/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2528,
				IDs:   nil,
				Alias: "ubisoft",
				Name:  "Ubisoft",
			},
			EmployeeCount: 720,
		},
		// Plarium https://jobs.dou.ua/companies/plarium/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1949408,
				IDs:   nil,
				Alias: "plarium",
				Name:  "Plarium",
			},
			EmployeeCount: 718,
		},
		// Playtika https://jobs.dou.ua/companies/playtika-ua/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1919232,
				IDs:   nil,
				Alias: "playtika",
				Name:  "Playtika",
			},
			EmployeeCount: 695,
		},
		// SPD Technology https://jobs.dou.ua/companies/spd-technology/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9998811,
				IDs:   nil,
				Alias: "spdtechnology",
				Name:  "SPD Technology",
			},
			EmployeeCount: 645,
		},
		// Innovecs https://jobs.dou.ua/companies/innovecs/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2430847,
				IDs:   nil,
				Alias: "innovecs",
				Name:  "Innovecs",
			},
			EmployeeCount: 640,
		},
		// Nova Digital https://jobs.dou.ua/companies/nova-digital/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    53101536,
				IDs:   nil,
				Alias: "nova-poshta-digital-company",
				Name:  "Nova Digital",
			},
			EmployeeCount: 635,
		},
		// Svitla Systems https://jobs.dou.ua/companies/svitla-systems-inc/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    592431,
				IDs:   nil,
				Alias: "svitla-systems-inc-",
				Name:  "Svitla Systems, Inc.",
			},
			EmployeeCount: 627,
		},
		// MODUS X https://jobs.dou.ua/companies/modus-x/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14058705,
				IDs:   nil,
				Alias: "modusx",
				Name:  "modusx",
			},
			EmployeeCount: 625,
		},
		// GeeksForLess https://jobs.dou.ua/companies/geeksforless/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2165356,
				IDs:   nil,
				Alias: "geeksforless",
				Name:  "GeeksForLess Inc.",
			},
			EmployeeCount: 600,
		},
		// Creatio https://jobs.dou.ua/companies/creatio/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20353970,
				IDs:   nil,
				Alias: "creatioglobal",
				Name:  "Creatio",
			},
			EmployeeCount: 600,
		},
		// Wix https://jobs.dou.ua/companies/wix/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    97250,
				IDs:   nil,
				Alias: "wix-com",
				Name:  "Wix",
			},
			EmployeeCount: 600,
		},
		// Binotel https://jobs.dou.ua/companies/binotel/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2903672,
				IDs:   nil,
				Alias: "binotel",
				Name:  "Binotel LLC.",
			},
			EmployeeCount: 587,
		},
		// Viseven Group https://jobs.dou.ua/companies/Viseven/
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2907491,
				IDs:   nil,
				Alias: "viseven",
				Name:  "Viseven",
			},
			EmployeeCount: 586,
		},
	}
}
