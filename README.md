# Hugo Books Tool 

A simple tool to fetch book data for my hugo blog

## How it works

Hugo books will scan the content of a hugo site, and look for any content that contains 'isbn' property, `isbn: "12345"`, in the front matter. It then checks for a [data source](https://gohugo.io/content-management/data-sources/) for book data using the isbn number as a unique key. For any books without a corresponding book entry, HugoBooks will query the Open Library api for missing data.


## Level of AI Usage

Used LLM prompt as a replacement for Google after reading std library docs and std library source code. The goal of this project was to solve a problem and learn as much as possiblie. For me, embracing the hard part of the journey is better path for deep learning.
