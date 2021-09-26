FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go test -v
RUN CGO_ENABLED=0 go build -o /searchtermsbackend

EXPOSE 50080

CMD ["/searchtermsbackend"]