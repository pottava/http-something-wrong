FROM golang:1.15.6-alpine3.12 AS builder
RUN apk --no-cache add gcc musl-dev git
WORKDIR /go/src/github.com/pottava
ENV APP_PACKAGE="github.com/pottava/http-something-wrong" \
    APP_VERSION=v1.1.0
RUN git clone --depth=1 -b "${APP_VERSION}" "https://${APP_PACKAGE}.git"
WORKDIR /go/src/github.com/pottava/http-something-wrong
RUN CGO_ENABLED=0 GOOS="${GOOS:-linux}" GOARCH="${GOARCH:-amd64}"; \
    go build -ldflags "-s -w" -o /app

FROM alpine:latest
RUN apk --no-cache add tini ca-certificates
COPY --from=builder /app /app
ENV PORT=8080
ENTRYPOINT ["tini", "--", "/app"]
