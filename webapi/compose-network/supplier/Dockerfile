FROM golang:alpine3.21 as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app .

FROM scratch
COPY --from=builder /bin/app /app
CMD ["/app"]
EXPOSE 3001