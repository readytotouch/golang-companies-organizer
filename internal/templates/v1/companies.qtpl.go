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

func StreamCompanies(qw422016 *qt422016.Writer) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Golang companies</title>
</head>
<body>
    <h1>Companies that use Golang</h1>
    <table>
        <tr>
            <th>Name</th>
            <th>LinkedIn</th>
            <th>Repositories</th>
            <th>Why</th>
        </tr>
        <tr>
            <td><a href="https://www.google.com/">google.com</a></td>
            <td><a href="https://www.linkedin.com/company/google/">/company/google/</a></td>
            <td><a href="https://github.com/orgs/google/repositories?q=lang:go">GitHub</a>: 157</td>
            <td><a href="https://go.dev/talks/2012/splash.article">Go at Google</a></td>
        </tr>
        <tr>
            <td><a href="https://victoriametrics.com/">victoriametrics.com</a></td>
            <td><a href="https://www.linkedin.com/company/victoriametrics/">/company/victoriametrics/</a></td>
            <td><a href="https://github.com/orgs/VictoriaMetrics/repositories?q=lang:go">GitHub</a>: 10</td>
            <td><a href="https://github.com/VictoriaMetrics/VictoriaMetrics">github.com/cockroachdb/cockroach</a></td>
        <tr>
            <td><a href="https://www.uber.com/">uber.com</a></td>
            <td><a href="https://www.linkedin.com/company/uber-com/">/company/uber-com/</a></td>
            <td><a href="https://github.com/orgs/uber/repositories?q=lang:go">GitHub</a>: 30</td>
            <td><a href="https://www.uber.com/en-TR/blog/data-race-patterns-in-go/">Go at Uber</a></td>
        </tr>
        <tr>
            <td><a href="https://www.cloudflare.com/">cloudflare.com</a></td>
            <td><a href="https://www.linkedin.com/company/cloudflare/">/company/cloudflare/</a></td>
            <td><a href="https://github.com/orgs/cloudflare/repositories?q=lang:go">GitHub</a>: 87</td>
            <td><a href="https://blog.cloudflare.com/tag/go">Go at Cloudflare</a></td>
        </tr>
        <tr>
            <td><a href="https://bitly.com/">bitly.com</a></td>
            <td><a href="https://www.linkedin.com/company/bitly/">/company/bitly/</a></td>
            <td><a href="https://github.com/orgs/bitly/repositories?q=lang:go">GitHub</a>: 11</td>
            <td><a href="https://bitly.com/blog/why-we-write-everything-in-go/">Go at Bitly</a></td>
        </tr>
        <tr>
            <td><a href="https://www.cockroachlabs.com/">cockroachlabs.com</a></td>
            <td><a href="https://www.linkedin.com/company/cockroach-labs/">/company/cockroach-labs/</a></td>
            <td><a href="https://github.com/orgs/cockroachdb/repositories?q=lang:go">GitHub</a>: 92</td>
            <td><a href="https://github.com/cockroachdb/cockroach">github.com/cockroachdb/cockroach</a>
        </td>
        <tr>
            <td><a href="https://www.reddit.com/">reddit.com</a></td>
            <td><a href="https://www.linkedin.com/company/reddit-com/">/company/reddit-com/</a></td>
            <td><a href="https://github.com/orgs/reddit/repositories?q=lang:go">GitHub</a>: 20</td>
            <td><a href="https://app.otta.com/dashboard/jobs/T2ZxNGhO">Senior Software Engineer</a></td>
        </tr>
    </table>
    <h3>LinkedIn</h3>
    <ul>
        <li><a href="https://www.linkedin.com/search/results/PEOPLE/?currentCompany=%5B%22150573%22%2C%229309408%22%2C%22552285%22%2C%22407222%22%2C%221815218%22%2C%2230169914%22%2C%221441%22%5D&network=%5B%22F%22%2C%22S%22%5D">LinkedIn Connections</a></li>
        <li><a href="https://www.linkedin.com/jobs/search/?keywords=%22Golang%20Engineer%22%20OR%20%22Golang%20Software%20Engineer%22%20OR%20%22Golang%20Developer%22%20OR%20%22Go%20Engineer%22%20OR%20%22Go%20Software%20Engineer%22%20OR%20%22Golang%20Developer%22&location=Worldwide">LinkedIn Golang Worldwide Jobs</a></li>
    </ul>
    <footer>
        <p>© 2024 <a href="https://readytotouch.com/">ReadyToTouch</a></p>
    </footer>
</body>
</html>
`)
}

func WriteCompanies(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamCompanies(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func Companies() string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteCompanies(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
