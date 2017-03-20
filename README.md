URL SHORTENER
==============

Url shortener written in Go. 

Uses [BoltDB](https://github.com/boltdb/bolt)

## Configuration

You can change urlshortener domain (p.io by default) by editing docker-compose.yml `URLSHORTENER_DOMAIN` environment variable 


## Run

This will build and run urlshortener API at [http://localhost](http://localhost):
```bash
$ docker-compose up -d
```

## API
#### Add Url to shorten
##### Request
```
PUT /add
{"url": "http://www.example.com/test"}
```
##### Response
```
{"url": "p.io/fsDETr"}
```
