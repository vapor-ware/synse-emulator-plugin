FROM vaporio/golang:1.11 as builder
WORKDIR /go/src/github.com/vapor-ware/synse-emulator-plugin
COPY . .

# If the vendored dependencies are not present in the docker build context,
# we'll need to do the vendoring prior to building the binary.
RUN if [ ! -d "vendor" ]; then make dep; fi
RUN make build CGO_ENABLED=0


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/vapor-ware/synse-emulator-plugin/build/emulator ./plugin

# Build the device configurations directly into the image. This is not
# generally advised, but is acceptable here since the plugin is merely
# an emulator and its config is not tied to anything real.
COPY config/device /etc/synse/plugin/config/device
COPY config.yml    /etc/synse/plugin/config/config.yml

# Image Metadata -- http://label-schema.org/rc1/
# This should be set after the dependency install so we can cache that layer
ARG BUILD_DATE
ARG BUILD_VERSION
ARG VCS_REF

LABEL maintainer="vapor@vapor.io" \
      org.label-schema.schema-version="1.0" \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="vaporio/emulator-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-emulator-plugin" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor="Vapor IO" \
      org.label-schema.version=$BUILD_VERSION

EXPOSE 5001
ENTRYPOINT ["./plugin"]