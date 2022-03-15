FROM golang:1.17-bullseye AS builder
WORKDIR /src

COPY go.mod ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o /bin/dummy .

FROM gcr.io/distroless/base-debian10
WORKDIR /app

COPY --from=builder /bin/dummy ./
EXPOSE 8989

CMD ["./dummy"]