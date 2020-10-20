## Development guide

```sh
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
