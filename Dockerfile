FROM golang:1-alpine as builder

COPY go.mod go.sum /usr/src/bouncer/
WORKDIR /usr/src/bouncer
RUN go mod download

COPY . /usr/src/bouncer
RUN go build -o /usr/bin/bouncer .

FROM alpine
COPY --from=builder /usr/bin/bouncer /usr/bin/bouncer
COPY dummy.cert /etc/bouncer/tls/dummy.cert
COPY dummy.key /etc/bouncer/tls/dummy.key

CMD [ "/usr/bin/bouncer", "--cert", "/etc/bouncer/tls/dummy.cert", "--key", "/etc/bouncer/tls/dummy.key" ]
