# Build
FROM golang:alpine3.18 AS build

WORKDIR /app

COPY . .

RUN go build -o api-svc cmd/server/main.go

# Deploy
FROM alpine:3.18.2
WORKDIR /app

COPY --from=build /app/api-svc .

EXPOSE 8080
ENTRYPOINT [ "/app/api-svc" ]
