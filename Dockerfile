FROM golang:1.22 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o hahaton

FROM scratch
COPY --from=builder /app/hahaton /hahaton

ENTRYPOINT ["/hahaton"]

