# Caching Go modules for future builds
# FROM golang:alpine3.21 as modules
# COPY go.mod /modules/
# WORKDIR /modules
# RUN go mod download

# Building app
FROM golang:alpine3.21 as builder
# COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app .

# Final (bare image, no shell)
FROM scratch
COPY --from=builder /bin/app /app 
CMD ["/app"]
EXPOSE 3001