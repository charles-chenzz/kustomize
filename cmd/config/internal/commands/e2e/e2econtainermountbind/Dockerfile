# Copyright 2019 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.20-bullseye
ENV CGO_ENABLED=0
WORKDIR /go/src/

# download modules
# COPY go.mod go.mod
# COPY go.sum go.sum
# RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/function ./cmd/config/internal/commands/e2e/e2econtainermountbind

FROM alpine:latest
COPY --from=0 /usr/local/bin/function /usr/local/bin/function
CMD ["function"]
