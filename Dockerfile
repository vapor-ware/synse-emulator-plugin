#
# Builder Image
#
# FROM docker.io/library/golang:1.16 as builder
FROM docker.io/library/debian:stable-slim as builder

RUN apt-get update && apt-get install -y ca-certificates

RUN useradd -M vaporio

RUN mkdir -p /etc/synse/plugin/config \
 && mkdir -p /etc/synse/plugin/config \
 && chown -R vaporio /etc/synse

#
# Final Image
#
FROM scratch

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
COPY --chown=vaporio synse-emulator-plugin /plugin

EXPOSE 5001

USER vaporio
ENTRYPOINT ["/plugin"]