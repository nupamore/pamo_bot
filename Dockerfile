# Build Stage
FROM golang:1.14-alpine AS builder

WORKDIR /go/src/pamo-bot
COPY . .

RUN go get github.com/volatiletech/sqlboiler/v4
RUN go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql

RUN sqlboiler mysql --config configs/sqlboiler.toml
RUN go build cmd/bot/bot.go

# Runtime Stage
FROM golang:1.14-alpine

WORKDIR /go/src/pamo-bot
COPY --from=builder /go/src/pamo-bot/configs/.env configs/.env
COPY --from=builder /go/src/pamo-bot/bot bot

CMD ["./bot"]
