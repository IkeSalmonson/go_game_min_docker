FROM golang:1.21.3 AS build

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY . .

CMD ["go", "run", "./cmd/main.go"]