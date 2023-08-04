#
# Development Dockerfile
#
# This Dockerfile produces an image intended to only be used for
# development and debugging. It should NOT be used in production.
# Development images contain additional tooling that makes it easier
# to exec into a contain and dig into whatever may be going on inside.
#

FROM golang:1.20

WORKDIR /synse

# Build the device configurations directly into the image. This is not
# generally advised, but is acceptable here since the plugin is merely
# an emulator and its config is not tied to anything real. These config
# defaults may be overridden at run time.
COPY config/device  /etc/synse/plugin/config/device
COPY config.yml     /etc/synse/plugin/config/config.yml

# Copy the executable and README information. The executable should be
# built prior to the image build (see Makefile).
COPY synse-emulator-plugin ./plugin
COPY README.md .

EXPOSE 5001
ENTRYPOINT ["./plugin"]
