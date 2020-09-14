FROM golang:1.14-alpine as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app

WORKDIR /go/bin

ARG app_db_username
ENV APP_DB_USERNAME $app_db_username
ARG app_db_password
ENV APP_DB_PASSWORD $app_db_password
ARG app_db_name
ENV APP_DB_NAME $app_db_name

ENTRYPOINT ["/go/bin/app"]