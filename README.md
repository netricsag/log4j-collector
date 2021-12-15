# log4j-collector
This is a simple log4j collector that will collect logs from a HTTP JSON API REST call and writes the data inside a simple sqlite3 db file.  
The frontend can be found at:
- https://github.com/bluestoneag/log4j-collector-frontend

## Env Variables

* `DB_NAME` - The name of the database file. Default: `log4j-collector.db`
* `DB_PATH` - The path of the database file. Default is `./db`.

**Note:** The database file will be created at the given path or if not provided at the default path with the default name. Make sure you add .db to the end of the file name.

## API Documentation
The Api documentation is available at:
- [API Documentation](docs/)