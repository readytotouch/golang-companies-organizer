// https://jobs.dou.ua/top50/
// https://dou.ua/lenta/articles/top-50-summer-2024/
// https://s.dou.ua/files/lenta/top-50-summer-2024_v2/data/top50-cities.csv
{
    const $companies = document.querySelectorAll("#chart-top50_table a.company-name");

    const companies = Array.from($companies).map($company => ({
        name: $company.textContent.trim(),
        url: $company.href,
    }));

    console.log(companies);
}
{
    const $companies = document.querySelectorAll("#chart-top50_table a.company-name");

    const content = Array.from($companies).map($company =>
        `// ${$company.textContent.trim()} ${$company.href}`
    ).join('\n');

    console.log(content);
}
{
    const $companies = document.querySelectorAll("#chart-top50_table tr");

    const content = Array.from($companies).map($company => {
        const $name = $company.querySelector('a.company-name');

        if ($name === null) {
            return '';
        }

        const employeeCount = $company.querySelector('span.staffTotal-value').textContent.trim().replaceAll(' ', '');

        return `// ${$name.textContent.trim()} ${$name.href}
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   nil,
				Alias: "",
				Name:  "",
			},
			EmployeeCount:   ${employeeCount},
		},`;
    }).join('\n');

    console.log(content);
}
