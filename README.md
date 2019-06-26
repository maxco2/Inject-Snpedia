# Introduction
This project is a greasemonkey script and a go program which supply a simple injection on snpedia.com so that you can quick check your
genetic data when you visit a wiki article.

![injection](https://github.com/maxco2/Inject-Snpedia/blob/master/img/injection.png)

# Usage
1. Create your database which should name genetic.db. Only support sqlite3.

```sql
CREATE TABLE "genetic" (
	`rsid`	TEXT NOT NULL,
	`chromosome`	TEXT NOT NULL,
	`position`	INTEGER NOT NULL,
	`genotype`	TEXT NOT NULL,
	`extended_data`	INTEGER NOT NULL
)
```
Import your genetic data, such as 23andme, wegene, etc.

2. Build the go program and run it.
```
```bash
$ go build genetic.go
$ ./genetic
```
3. Install the greasemonkey script.

# What is Snpedia?
SNPedia is a wiki investigating human genetics. See also https://en.wikipedia.org/wiki/SNPedia .
