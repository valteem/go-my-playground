FROM golang:alpine3.21 as modules
COPY go.mod /modules/
WORKDIR /modules
RUN go mod download

FROM golang:alpine3.21 as tester
COPY --from=modules /go/pkg /go/pkg
RUN mkdir -p /work
COPY go.mod /work
COPY go.sum /work
COPY /app /work/app
WORKDIR /work
CMD ["/bin/sh"]