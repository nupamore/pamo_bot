# pamo_bot (beta)

### [Bot Invite Link](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

### [Discord Community](https://discordapp.com/channels/681300669973790824/681482057787899905/681491948258721803)

[한국어설명서](README_ko.md) |
[日本語説明書](README_ja.md)

---

## Development guide

### Generate models
```sh
# Create dotenv
vim .env
# Create db config file
vim sqlboiler.toml
# Generate models
sqlboiler mysql
# Bot start
go run pamo_bot.go
```
