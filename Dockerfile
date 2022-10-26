# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.19-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /coqchain
RUN cd /coqchain && go run build/ci.go install ./cmd/coq

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /coqchain/build/bin/coq /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["coq"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
