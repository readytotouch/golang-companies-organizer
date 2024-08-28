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

    const $companies = document.querySelectorAll("#chart-top50_table a.company-name");

    const content = Array.from($companies).map($company =>
        `// ${$company.textContent.trim()} ${$company.href}
        {
			ID:    0,
			IDs:   nil,
			Alias: "",
			Name:  "",
		},`
    ).join('\n');

    console.log(content);
}
