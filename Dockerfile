FROM iron/go:1.9-dev as builder
WORKDIR /go/src/github.com/vapor-ware/synse-emulator-plugin
COPY . .

RUN apk --update add ca-certificates
RUN make build CGO_ENABLED=0


FROM scratch

LABEL maintainer="vapor@vapor.io"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/vapor-ware/synse-emulator-plugin/build/emulator ./plugin
COPY config.yml    /etc/synse/plugin/config/config.yml

# Build the device configurations directly into the image. This is not
# generally advised, but is acceptable here since the plugin is merely
# an emulator and its config is not tied to anything real.
COPY config/device /etc/synse/plugin/config/device

EXPOSE 5001

ENTRYPOINT ["./plugin"]