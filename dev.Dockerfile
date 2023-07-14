#
# Development Dockerfile
#
# This Dockerfile produces an image intended to only be used for
# development and debugging. It should NOT be used in production.
# Development images contain additional tooling that makes it easier
# to exec into a contain and dig into whatever may be going on inside.
#
#
# Builder Image
#
FROM docker.io/library/golang:1.16 as builder

RUN apt-get update && apt-get install -y ca-certificates

RUN useradd -M vaporio

WORKDIR /app

# Download dependencies
COPY go.* ./
RUN go mod download

# Copy the project into the builder
COPY . ./

# Disable dynamic linking. The binary won't work on scratch otherwise
ENV CGO_ENABLED=0

# Build the binary.
RUN make build-linux


RUN mkdir -p /etc/synse/plugin/config \
 && mkdir -p /etc/synse/plugin/config \
 && chown -R vaporio /etc/synse \
 && chown -R vaporio /app

#
# Final Image
#
FROM docker.io/library/debian:stable-slim

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="vaporio/emulator-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-emulator-plugin" \
      org.label-schema.vendor="Vapor IO"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd

# Build the device configurations directly into the image. This is not
# generally advised, but is acceptable here since the plugin is merely
# an emulator and its config is not tied to anything real.
COPY --chown=vaporio config/device /etc/synse/plugin/config/device
COPY --chown=vaporio config.yml    /etc/synse/plugin/config/config.yml

# Copy the executable.
COPY --chown=vaporio --from=builder /app/synse-emulator-plugin /app/plugin


EXPOSE 5001

USER vaporio
ENTRYPOINT ["/app/plugin"]
