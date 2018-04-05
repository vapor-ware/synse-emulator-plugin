FROM iron/go:dev as builder
WORKDIR /go/src/github.com/vapor-ware/synse-emulator-plugin
COPY . .
RUN make build GIT_TAG="" GIT_COMMIT=""


FROM iron/go
MAINTAINER Vapor IO <eng@vapor.io>

WORKDIR /plugin

COPY --from=builder /go/src/github.com/vapor-ware/synse-emulator-plugin/build/emulator ./plugin
COPY config.yml .
COPY config/proto /etc/synse/plugin/config/proto

CMD ["./plugin"]
