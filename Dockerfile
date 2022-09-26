# Build
FROM golang:1.16-alpine AS build

WORKDIR /app
RUN apk add build-base
COPY . . 
RUN go mod vendor
RUN GOOS=linux go build -o app

# Deploy
FROM alpine:latest
WORKDIR /workspace
COPY --from=build /app/app ./app
COPY .env.prod .env.prod
EXPOSE 8080
CMD ["/workspace/app", "-env", "prod"]
