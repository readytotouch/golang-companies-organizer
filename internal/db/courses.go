package db

import "github.com/readytotouch/golang-companies-organizer/internal/domain"

func Courses() []domain.Course {
	return []domain.Course{
		// CS50
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10604885,
				IDs:   nil,
				Alias: "cs50",
				Name:  "CS50",
			},
			AlumniCount: 4191,
		},
		// Mate academy
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11133657,
				IDs:   nil,
				Alias: "mateacademy",
				Name:  "Mate academy",
			},
			AlumniCount: 4719,
		},
		// IT STEP
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    76509886,
				IDs:   nil,
				Alias: "it-step-academy",
				Name:  "IT STEP Academy",
			},
			AlumniCount: 1418,
		},
		// Lemon School
		// https://www.linkedin.com/company/lemonschool/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// Hillel
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10415579,
				IDs:   nil,
				Alias: "ithillel",
				Name:  "Hillel IT School",
			},
			AlumniCount: 10745,
		},
		// GoIT
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3844136,
				IDs:   nil,
				Alias: "goit-start-your-career-in-it",
				Name:  "GoIT",
			},
			AlumniCount: 15433,
		},
		// CyberBionic Systematics
		// https://www.linkedin.com/company/cbsystematics-development/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// Beetroot Academy
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15159023,
				IDs:   nil,
				Alias: "beetrootacademy",
				Name:  "Beetroot Academy",
			},
			AlumniCount: 3960,
		},
		// Web Academy
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15098547,
				IDs:   nil,
				Alias: "webacademyukraine",
				Name:  "Web Academy",
			},
			AlumniCount: 772,
		},
		// Logos IT Academy
		// https://www.linkedin.com/company/logos-training-center/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// FoxmindEd
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18401271,
				IDs:   nil,
				Alias: "foxminded",
				Name:  "FoxmindEd",
			},
			AlumniCount: 648,
		},
		// Okten School
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    87231241,
				IDs:   nil,
				Alias: "okten-school-it",
				Name:  "OKTEN SCHOOL",
			},
			AlumniCount: 227,
		},
		// Projector
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10987421,
				IDs:   nil,
				Alias: "projector-institute",
				Name:  "Projector Institute",
			},
			AlumniCount: 4711,
		},
		// SkillUP
		// https://www.linkedin.com/company/skillup-ukraine/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// CURSOR.EDUCATION
		// https://www.linkedin.com/company/cursor-education/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// Prog Academy
		// https://www.linkedin.com/company/progacademy/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// A-Level Ukraine
		// https://www.linkedin.com/company/a-level-ukraine/
		// Should be school like https://www.linkedin.com/school/cs50/
		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		IDs:   nil,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	AlumniCount: 0,
		//},
		// Robot_dreams
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    67875206,
				IDs:   nil,
				Alias: "robot-dreams",
				Name:  "Robot_dreams",
			},
			AlumniCount: 192,
		},
	}
}
