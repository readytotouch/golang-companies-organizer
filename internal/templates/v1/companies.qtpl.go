// Code generated by qtc from "companies.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamCompanies(qw422016 *qt422016.Writer, companies []Company, universities []University) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<head>
    <title>Golang companies</title>
    <meta name="description" content="Golang companies">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="author" type="text/plain" href="https://gocompanies.readytotouch.com/humans.txt"/>

	`)
	streamfavicon(qw422016)
	qw422016.N().S(`

    <!-- Google tag (gtag.js) -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-VCD3QKRT9Z"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-VCD3QKRT9Z');
    </script>
    <style>
        td {
            padding-top:5px;
            padding-bottom:5px;
            padding-right:5px;
        }

        td:first-child {
            padding-left:5px;
            padding-right:0;
        }
    </style>
</head>
<body>
    <h1>Companies that use Golang <img width="20px" src="./icons/go.svg"><iframe src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=gocompanies&type=star&count=true" frameborder="0" scrolling="0" width="150" height="20" title="GitHub"></iframe></h1>
    <table border="1">
        <tr>
            <th>Name</th>
            <th><img width="20px" src="./icons/linkedin.svg"/> LinkedIn</th>
            <th><img width="20px" src="./icons/github.svg"/> GitHub</th>
            <th>Glassdoor</th>
            <th><img width="20px" src="./icons/otta.png"/> Otta</th>
            <th><img width="20px" src="./icons/youtube.svg"/> YouTube</th>
            <th>Vacancies</th>
            <th>Go main language</th>
            <th>SimilarWeb</th>
            <th>Actions</th>
        </tr>
        `)
	for _, company := range companies {
		qw422016.N().S(`
            <tr>
                <td>
                    <a href="`)
		qw422016.E().S(company.URL)
		qw422016.N().S(`">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</a>
                </td>
                <td>
                    <a href="https://www.linkedin.com/company/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" title="`)
		qw422016.E().S(company.LinkedInProfile.Name)
		qw422016.N().S(`">Overview</a>,
                    <a href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, universities))
		qw422016.N().S(`" title="`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`">Connections</a>,
                    <a href="`)
		qw422016.E().S(linkedinJobsURL([]Company{company}, keywordsSkills))
		qw422016.N().S(`" title="`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`">Jobs</a>
                </td>
                <td>
`)
		if company.GitHubProfile.Login != "" {
			qw422016.N().S(`                        <a href="https://github.com/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`">Overview</a>, <a href="https://github.com/orgs/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`/repositories?q=lang:go">Repositories</a>: `)
			qw422016.N().D(company.GitHubProfile.RepositoryCount)
			qw422016.N().S(`
`)
		}
		qw422016.N().S(`                </td>
                <td>
                    <a href="`)
		qw422016.E().S(company.GlassdoorProfile.OverviewURL)
		qw422016.N().S(`">Overview</a>, <a href="`)
		qw422016.E().S(company.GlassdoorProfile.ReviewsURL)
		qw422016.N().S(`">Reviews</a>
                </td>
                <td>
`)
		if company.OttaProfileSlug != "" {
			qw422016.N().S(`                        <a href="https://app.otta.com/companies/`)
			qw422016.E().S(company.OttaProfileSlug)
			qw422016.N().S(`">Overview</a>
`)
		}
		qw422016.N().S(`                </td>
                <td>
`)
		if company.YouTubeChannelURL != "" {
			qw422016.N().S(`                        <a href="`)
			qw422016.E().S(company.YouTubeChannelURL)
			qw422016.N().S(`">Overview</a>
`)
		}
		qw422016.N().S(`                </td>
                <td>
`)
		for i, vacancy := range company.Vacancies {
			qw422016.N().S(`                        <a href="`)
			qw422016.E().S(vacancy)
			qw422016.N().S(`">#`)
			qw422016.N().D(i)
			qw422016.N().S(`</a>
`)
		}
		qw422016.N().S(`                </td>
                <td>
                    `)
		if company.MainLanguage {
			qw422016.N().S(`✔️ TRUE`)
		}
		qw422016.N().S(`
                </td>
                <td>
                    <a href="`)
		qw422016.E().S(similarwebURL(company.URL))
		qw422016.N().S(`">Overview</a>
                </td>
                <td>
                    <ul>
                        <li><a href="`)
		qw422016.E().S(googleCompanyTitlesJobsURL(company))
		qw422016.N().S(`">Search in Google by titles</a></li>
                        <li><a href="`)
		qw422016.E().S(googleCompanySkillsJobsURL(company))
		qw422016.N().S(`">Search in Google by skills</a></li>
                    </ul>
                </td>
            </tr>
        `)
	}
	qw422016.N().S(`
    </table>
    <h3>LinkedIn</h3>
    <ul>
        <li><a href="`)
	qw422016.E().S(linkedinConnectionsURL(companies, nil))
	qw422016.N().S(`">LinkedIn Connections ["Go" OR "Golang"] [Companies]</a></li>
        <li><a href="`)
	qw422016.E().S(linkedinConnectionsURL(companies, universities))
	qw422016.N().S(`">LinkedIn Connections ["Go" OR "Golang"] [Companies] [Ukrainian University]</a></li>
        <li><a href="`)
	qw422016.E().S(linkedinJobsURL(nil, keywordsTitles))
	qw422016.N().S(`">LinkedIn Jobs [Golang] [Worldwide]</a></li>
        <li><a href="`)
	qw422016.E().S(linkedinJobsURL(companies, keywordsSkills))
	qw422016.N().S(`">LinkedIn Jobs [Golang] [Companies] [Worldwide]</a></li>
    </ul>
    <h3>Other</h3>
    <ul>
        <li><a href="https://justjoin.it/all-locations/go">justjoin.it/all-locations/go</a></li>
        <li><a href="https://remoteok.com/remote-golang-jobs">remoteok.com/remote-golang-jobs</a></li>
        <li><a href="https://nofluffjobs.com/golang">nofluffjobs.com/golang</a></li>
        <li><a href="https://www.golangprojects.com/">golangprojects.com</a></li>
        <li><a href="https://golang.cafe/">golang.cafe</a></li>
        <li><a href="https://startup.jobs/golang-jobs">startup.jobs/golang-jobs</a></li>
    </ul>
    <footer>
        <p>© 2024 <a href="https://readytotouch.com/">ReadyToTouch</a></p>
    </footer>
</body>
</html>
`)
}

func WriteCompanies(qq422016 qtio422016.Writer, companies []Company, universities []University) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamCompanies(qw422016, companies, universities)
	qt422016.ReleaseWriter(qw422016)
}

func Companies(companies []Company, universities []University) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteCompanies(qb422016, companies, universities)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
