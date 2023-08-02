# Build
FROM golang:alpine3.18 AS build

WORKDIR /app

COPY . .

RUN go build -o vid-processor-svc cmd/video-processor/main.go

# Deploy
FROM alpine:3.18.2
WORKDIR /app

COPY --from=build /app/vid-processor-svc .

RUN apk add ffmpeg

ENTRYPOINT [ "/app/vid-processor-svc" ]
