# log4j-collector
![Build Status](https://github.com/bluestoneag/log4j-collector/workflows/CI/badge.svg) 
[![Go Report Card](https://goreportcard.com/badge/github.com/bluestoneag/log4j-collector)](https://goreportcard.com/report/github.com/bluestoneag/log4j-collector) 
![GitHub top language](https://img.shields.io/github/languages/top/bluestoneag/log4j-collector)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bluestoneag/log4j-collector) 
![open issues](https://img.shields.io/github/issues-raw/bluestoneag/log4j-collector)
![license](https://img.shields.io/github/license/bluestoneag/log4j-collector)

This is a simple log4j collector that will collect logs from a HTTP JSON API REST call and writes the data inside a simple sqlite3 db file.  
The frontend can be found at:
- https://github.com/bluestoneag/log4j-collector-frontend

## Env Variables

* `DB_NAME` - The name of the database file. Default: `log4j-collector.db`
* `DB_PATH` - The path of the database file. Default is `./db`.

**Note:** The database file will be created at the given `DB_PATH` or if not provided at the default path with the default name. Make sure you add .db to the end of the `DB_NAME` name.

## Docker

The docker image is available at:
```
docker pull ghcr.io/bluestoneag/log4j-collector:latest
```

Port mapping:
```
docker run -d -p 8080:8080 ghcr.io/bluestoneag/log4j-collector:latest
```

## API Documentation
The Api documentation is available at:
- [API Documentation](docs/)

## Authors
This project is maintained by [bluestoneag](https://github.com/bluestoneag) and [netrics](https://netrics.ch).  
Feel free to contribute!