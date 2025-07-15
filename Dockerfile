FROM golang:1.24-alpine AS builder

WORKDIR /go/src/github.com/Scrin/bambulab-exporter/
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /go/bin/bambulab-exporter

FROM alpine

COPY --from=builder /go/bin/bambulab-exporter /usr/local/bin/bambulab-exporter

CMD ["bambulab-exporter"]
