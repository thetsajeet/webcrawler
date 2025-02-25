# go webcrawler

> To rank well in Google Search, websites need to internally link pages one to another.
> For example, a blog post about the benefits of haircuts should probably link to my post about the best places to get haircuts.

A Go CLI tool that generates an internal links report for any website on the internet by crawling each page of the site.

## Setup

```sh
git clone <url>
cd webcrawler
go install
```

## Run

```sh
# build
go build main.go -o ./webcrawler && ./crawler <website> <concurrency> <max_pages>
# development
go run main.go <website> <concurrency> <max_pages>
```

## Run with Docker

```sh
docker build . -t name:version
docker run -e F=site -e S=concurrency -e T=max_pages name:version
```

## Testing

```sh
go test ./...
```
