package v1

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/readytotouch/gocompanies/internal/domain"
)

type Company = domain.Company

func linkedinConnections(companies []Company) string {
	ids := make([]string, len(companies))
	for i, c := range companies {
		for _, l := range c.LinkedInProfiles {
			ids[i] = strconv.FormatInt(int64(l.ID), 10)
		}
	}

	companyQueryParam, _ := json.Marshal(ids)

	values := url.Values{
		"currentCompany": {string(companyQueryParam)},
		"network":        {`["F","S"]`},
		"keywords":       {`"Go" OR "Golang"`},
	}

	return "https://www.linkedin.com/search/results/PEOPLE/?" + values.Encode()
}
