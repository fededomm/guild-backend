FROM golang:1.21.0 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

## Deploy
FROM busybox

## Copiare i certificati per connettersi a postgres
COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app/main /opt/main

EXPOSE 8000
CMD ["/opt/main"]