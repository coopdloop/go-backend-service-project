from golang:1.22.0 AS build-stage
  WORKDIR /app

  COPY go.mod ./ 
  RUN go mod download
  
  RUN apt-get update -yq \
    && apt-get -yq install curl gnupg ca-certificates openssl \
    && apt-get update -yq
  RUN update-ca-certificates

  COPY Makefile ./
  COPY cmd/ ./cmd
  COPY internal/ ./internal

  RUN CGO_ENABLED=0 GOOS=linux go build -o /run cmd/server/main.go

FROM build-stage AS run-test
  RUN go test -v ./...

FROM scratch AS run-release-stage
  WORKDIR /app
  COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

  COPY --from=build-stage /run /run
  EXPOSE $PORT

  CMD ["/run/main"]
