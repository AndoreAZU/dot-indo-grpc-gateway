# build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --update make
RUN make install
RUN make generate
RUN go build -v -o go-docker /app/cmd/main.go
# final stage
FROM alpine:latest
EXPOSE 8080 8081
RUN apk add -U tzdata
WORKDIR /app
COPY --from=builder /app/go-docker .
COPY . .
ENTRYPOINT ./go-docker