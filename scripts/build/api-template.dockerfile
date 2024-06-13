# ----------------------------------------------------
# Application build stage
FROM golang:1.18-alpine as builder

WORKDIR /go/src/github.com/celtcoste/go-graphql-api-template
COPY ./ ./

RUN apk update && apk add make git
RUN make build BIN=. GOLANG_ARCH=amd64


# ----------------------------------------------------
# Base production image
FROM alpine AS base

WORKDIR /app/
RUN apk update && apk add ca-certificates
RUN adduser -S api-template

# ----------------------------------------------------
# go-graphql-api-template base image
FROM base AS api-template-base

COPY --from=builder /go/src/github.com/celtcoste/go-graphql-api-template .

USER api-template
ENTRYPOINT [ "/app/go-graphql-api-template" ]

# ----------------------------------------------------
# go-graphql-api-template production image (with locales)
FROM go-graphql-api-template-base AS go-graphql-api-template

USER go-graphql-api-template
ENTRYPOINT [ "/app/go-graphql-api-template" ] 