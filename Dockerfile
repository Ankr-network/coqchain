# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""
ARG DEBIAN_FRONTEND=noninteractive
# Build Geth in a stock Go builder container
FROM golang:1.17 as builder
RUN apt-get update && apt-get install -y  apt-utils gcc musl-dev  git libzstd-dev

ADD . /coqchain
RUN cd /coqchain && go run build/ci.go install ./cmd/coq

# Pull Geth into a second stage deploy alpine container
FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y  apt-utils ca-certificates
COPY --from=builder /coqchain/build/bin/coq /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["coq"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"