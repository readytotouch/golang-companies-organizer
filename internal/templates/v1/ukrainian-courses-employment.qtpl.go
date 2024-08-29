// Code generated by qtc from "ukrainian-courses-employment.qtpl". DO NOT EDIT.
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

func StreamUkrainianCoursesEmployment(qw422016 *qt422016.Writer, courses []Course, douCompanies []DouCompany, faangCompanyGroup FaangCompanyGroup) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<head>
    <title>Ukrainian courses employment</title>
    <meta name="description" content="Ukrainian courses employment">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="author" type="text/plain" href="https://golang-companies-organizer.readytotouch.com/humans.txt"/>

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
    <h1>Ukrainian courses employment</h1>
    <table border="1">
        <tr>
            <th>Name</th>
            <th>Alumni</th>
            <th>TOP-50 UA (current company)</th>
            <th>TOP-50 UA (past company)</th>
            <th>FAANG (current company)</th>
            <th>FAANG (past company)</th>
            <th>Rate</th>
        </tr>
        `)
	for i, course := range courses {
		qw422016.N().S(`
            <tr>
                <td>
                    `)
		qw422016.N().D(i)
		qw422016.N().S(` | <a href="https://www.linkedin.com/school/`)
		qw422016.E().S(course.LinkedInProfile.Alias)
		qw422016.N().S(`/">`)
		qw422016.E().S(course.LinkedInProfile.Name)
		qw422016.N().S(`</a>
                </td>
                <td><a href="https://www.linkedin.com/school/`)
		qw422016.E().S(course.LinkedInProfile.Alias)
		qw422016.N().S(`/people/">`)
		qw422016.N().D(course.AlumniCount)
		qw422016.N().S(`</a></td>
                <td><a href="`)
		qw422016.E().S(linkedinCourseDouConnectionsURL(course, douCompanies, nil))
		qw422016.N().S(`">`)
		qw422016.N().D(course.DouCurrentCount)
		qw422016.N().S(`</a></td>
                <td><a href="`)
		qw422016.E().S(linkedinCourseDouConnectionsURL(course, nil, douCompanies))
		qw422016.N().S(`">`)
		qw422016.N().D(course.DouPastCount)
		qw422016.N().S(`</a></td>
                <td><a href="`)
		qw422016.E().S(linkedinCourseConnectionsURL(course, faangCompanyGroup.FaangCompanies, nil))
		qw422016.N().S(`">`)
		qw422016.N().D(course.FaangCurrentCount)
		qw422016.N().S(`</a></td>
                <td><a href="`)
		qw422016.E().S(linkedinCourseConnectionsURL(course, nil, faangCompanyGroup.FaangCompanies))
		qw422016.N().S(`">`)
		qw422016.N().D(course.FaangPastCount)
		qw422016.N().S(`</a></td>
                <td>
                    (`)
		qw422016.N().D(course.DouCurrentCount)
		qw422016.N().S(` + `)
		qw422016.N().D(course.DouCurrentCount)
		qw422016.N().S(` + `)
		qw422016.N().D(course.FaangCurrentCount)
		qw422016.N().S(` + `)
		qw422016.N().D(course.FaangPastCount)
		qw422016.N().S(`) / `)
		qw422016.N().D(course.AlumniCount)
		qw422016.N().S(` = `)
		qw422016.N().FPrec(courseRate(course), 2)
		qw422016.N().S(`%
                </td>
            </tr>
        `)
	}
	qw422016.N().S(`
    </table>

    <h2>FAANG companies</h2>
    <ul>
        `)
	for _, company := range faangCompanyGroup.FaangCompanies {
		qw422016.N().S(`
            <li><a href="https://www.linkedin.com/company/`)
		qw422016.E().S(company.Alias)
		qw422016.N().S(`/">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</a></li>
        `)
	}
	qw422016.N().S(`
    </ul>

    <footer>
        <p>© 2024 <a href="https://readytotouch.com/">ReadyToTouch</a></p>
    </footer>
</body>
</html>
`)
}

func WriteUkrainianCoursesEmployment(qq422016 qtio422016.Writer, courses []Course, douCompanies []DouCompany, faangCompanyGroup FaangCompanyGroup) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamUkrainianCoursesEmployment(qw422016, courses, douCompanies, faangCompanyGroup)
	qt422016.ReleaseWriter(qw422016)
}

func UkrainianCoursesEmployment(courses []Course, douCompanies []DouCompany, faangCompanyGroup FaangCompanyGroup) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteUkrainianCoursesEmployment(qb422016, courses, douCompanies, faangCompanyGroup)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

// www.linkedin.com/co
