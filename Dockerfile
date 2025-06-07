FROM golang:1.23-alpine3.20 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux

ARG BUILD_VERSION="0.0.0"
ARG BUILD_COMMIT=none
ARG BUILD_DATE=unknown

WORKDIR /www

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build --ldflags "-X 'go-api-cms/internal/version.Version=${BUILD_VERSION}' \
                        -X 'go-api-cms/internal/version.Commit=${BUILD_COMMIT}' \
                        -X 'go-api-cms/internal/version.BuildTime=${BUILD_DATE}' \
                        -extldflags -static" -o main .

FROM alpine:3.20

RUN apk add --no-cache tzdata

WORKDIR /www

COPY --from=builder /www/main /www/
COPY --from=builder /www/database/migrations/ /www/database/migrations/
COPY --from=builder /www/development.sh /www/development.sh

RUN chmod +x development.sh

ENTRYPOINT ["/www/main"]
