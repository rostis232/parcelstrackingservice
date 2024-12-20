#base go image

FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o trackingservice ./cmd/api

RUN chmod +x ./trackingservice

#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app /app

WORKDIR /app

CMD ["/app/trackingservice"]