## Development guide

```sh
# go get
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/null/v8
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
# Create dotenv
vim configs/.env
# Create db config file
vim configs/sqlboiler.toml
# Generate models
sqlboiler mysql --config configs/sqlboiler.toml
# Bot start
go run cmd/bot/bot.go
# REST Server start
go run cmd/server/server.go
```

## Directory structure
```
/cmd            Run go files
/commands       Discord bot commands
/configs        Security keys
/controllers    REST Server endpoints
/docs           Markdown documents
/events         Discord bot events
/models         ORM Models
/services       Business Logics
```
