#
# Builder Image
#
FROM vaporio/golang:1.13 as builder

#
# Final Image
#
FROM vaporio/scratch-ish:1.0.0

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="vaporio/emulator-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-emulator-plugin" \
      org.label-schema.vendor="Vapor IO"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Build the device configurations directly into the image. This is not
# generally advised, but is acceptable here since the plugin is merely
# an emulator and its config is not tied to anything real.
COPY config/device /etc/synse/plugin/config/device
COPY config.yml    /etc/synse/plugin/config/config.yml

# Copy the executable.
COPY synse-emulator-plugin ./plugin

EXPOSE 5001
ENTRYPOINT ["./plugin"]
