FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /booster-hub-bin

# COPY binary to a container which use lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /booster-hub-bin /booster-hub-bin

# Env for database connection
ENV DATABASE_URL=postgres://test_user:password@localhost:5432/app_db
EXPOSE 3000

USER nonroot:nonroot 

ENTRYPOINT ["/booster-hub-bin"]

