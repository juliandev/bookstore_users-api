FROM golang:alpine AS build
WORKDIR /go/src/github.com/juliandev/bookstore_users-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/github.com/juliandev/bookstore_users-api .

FROM scratch
COPY --from=build /go/bin/github.com/juliandev/bookstore_users-api /go/bin/github.com/juliandev/bookstore_users-api
CMD ["/go/bin/github.com/juliandev/bookstore_users-api"]
