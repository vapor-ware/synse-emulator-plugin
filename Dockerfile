FROM iron/go:dev as builder
WORKDIR /go/src/github.com/vapor-ware/synse-emulator-plugin
COPY . .

RUN make build CGO_ENABLED=0


FROM scratch

LABEL maintainer="vapor@vapor.io"

COPY --from=builder /go/src/github.com/vapor-ware/synse-emulator-plugin/build/emulator ./plugin
COPY config.yml   /etc/synse/plugin/config.yml
COPY config/proto /etc/synse/plugin/config/proto

EXPOSE 5001

ENTRYPOINT ["./plugin"]