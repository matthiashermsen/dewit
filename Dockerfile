# Build stage
FROM golang:1.22.3-alpine3.18 AS build
WORKDIR /app

ADD . .

# We need additional packages
# - Makefile for scripts
# - Git for build flags
RUN apk add --no-cache make git

RUN make build

# Run stage
FROM alpine:3.21.0
WORKDIR /app

COPY --from=build /app/build/dewit .

CMD [ "./dewit" ]