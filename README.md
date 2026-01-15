# Hugo Books Tool 

A simple tool to fetch book data for my hugo blog

## How it works

Hugo books will scan the content of a hugo site, and look for any content that contains 'isbn' property in the front matter. It will then check for a [data source](https://gohugo.io/content-management/data-sources/) for book data. It will ensure that a data entry exists for each found isbn number, utilizing openLibrary api for any missing books data
