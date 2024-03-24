package v1

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/readytotouch/gocompanies/internal/domain"
)

type Company = domain.Company

const (
	keywordsSkills = `"Go" OR "Golang"`
	keywordsTitles = `"Golang Engineer" OR "Golang Software Engineer" OR "Golang Developer" OR "Go Engineer" OR "Go Software Engineer" OR "Golang Developer"`
)

func linkedinConnectionsURL(companies []Company, ukrainianUniversity bool) string {
	companyQueryParam, _ := json.Marshal(toIDs(companies))

	values := url.Values{
		"currentCompany": {string(companyQueryParam)},
		"network":        {`["F","S"]`},
		"keywords":       {keywordsSkills},
	}

	if ukrainianUniversity {
		values["schoolFilter"] = []string{`["364340","496320","782774","818029","850102","1257361","15099424","15099711","15101057","15101979","15102004","15250306"]`}
	}

	return "https://www.linkedin.com/search/results/PEOPLE/?" + values.Encode()
}

func linkedinJobsURL(companies []Company, keywords string) string {
	values := url.Values{
		"keywords": {keywords},
		"location": {"Worldwide"},
	}

	if len(companies) > 0 {
		values["f_C"] = []string{strings.Join(toIDs(companies), ",")}
	}

	return "https://www.linkedin.com/jobs/search/?" + values.Encode()
}

func toIDs(companies []Company) []string {
	ids := make([]string, 0, len(companies))
	for _, company := range companies {
		for _, profile := range company.LinkedInProfiles {
			ids = append(ids, strconv.FormatInt(int64(profile.ID), 10))
		}
	}
	return ids
}
