# Synse Emulator Plugin
The emulator plugin for Synse Server. This is the source repo for the emulator
that is built into the Synse Server Docker image and run via the `enable-emulator`
flag.

This plugin is a simple plugin which has no dependencies on any backend hardware
or protocol. It is meant to be a standalone plugin that can run with Synse Server
out of the box allowing you to quickly experiment and develop with Synse Server
without any other dependencies. It may also serve as an example on how to write
plugins of your own.

This plugin supports three devices:
- Temperature (type: `temperature`, model: `emul8-temp`)
- LED (type: `led`, model: `emul8-led`)
- Fan (type: `fan`, model: `emul8-fan`)

All three return dummy data for their reading values. All three support reading,
but only the LED and Fan devices support writing. State is maintained for writable
devices, so any state you set should be readable back out.

## Getting Started

### Getting the Emulator
There are various means by which you can get the emulator plugin.

**Clone the Repo**

If you wish to do development work or extend the emulator for your own purposes,
you can clone the repo or make your own fork.
```
git clone https://github.com/vapor-ware/synse-emulator-plugin.git
```

**Go Get**

Additionally, you can get the repo via `go get`
```
go get github.com/vapor-ware/synse-emulator-plugin
```

**DockerHub**

The Synse Emulator Plugin is available via DockerHub and is pushed there automatically
by our CI, so it should feature the latest builds.
```
docker pull vaporio/emulator-plugin
```

**Synse Server**

This plugin comes built into the Synse Server image as well, where it can be run alongside
Synse Server with the `enable-emulator` command. See the [Synse Server]() repo for more information.

### Building
If you are using the emulator via Docker image, this is already done for you. Otherwise,
you will need to build the emulator binary. This can be done simply via the Makefile
```
make build
```
See the Makefile for other available targets.

### Running
If you are using the emulator as part of Synse Server, you only need to pass `enable-emulator`
to the docker run command for Synse Server.

If you just want to run the binary, `make build` will output the `emulator` binary into the
`build/` subdirectory. The plugin will need some configuration before it can run successfully.
To find out more on how plugins are configued, see the [Synse SDK](). The configurations you
need are provided here in the `config/` directory. While there are multiple ways that the
configuration can be passed to the plugin, the easiest way is simply via environment variable
```
PLUGIN_DEVICE_CONFIG=config ./build/emulator
```
This will run the plugin, at which point you should see info and debug level output similar
to
```
INFO[0000] Plugin Info:
INFO[0000]  Name:        emulator
INFO[0000]  Version:     1.0
INFO[0000]  SDK Version: 0.4.0
INFO[0000]  Git Commit:  3176d68
INFO[0000]  Git Tag:     -
INFO[0000]  Go Version:  go1.9.1
INFO[0000]  Build Date:  2018-01-31T14:34:05
INFO[0000]  OS:          darwin
INFO[0000]  Arch:        amd64
DEBU[0000] Plugin Config:
DEBU[0000]  &config.PluginConfig{Name:"emulator", Version:"1", Debug:true, Settings:config.Settings{LoopDelay:1000, Read:config.ReadSettings{BufferSize:100}, Write:config.WriteSettings{BufferSize:100, PerLoop:5}, Transaction:config.TransactionSettings{TTL:300}}, Network:config.NetworkSettings{Type:"unix", Address:"emulator.sock"}, AutoEnumerate:[]map[string]interface {}{}, Context:map[string]interface {}{}}
INFO[0000] Registered Devices:
INFO[0000]  rack-1-vec-eb100067acb0c054cf877759db376b03 (emul8-temp)
INFO[0000]  rack-1-vec-83cc1efe7e596e4ab6769e0c6e3edf88 (emul8-temp)
INFO[0000]  rack-1-vec-db1e5deb43d9d0af6d80885e74362913 (emul8-temp)
INFO[0000]  rack-1-vec-329a91c6781ce92370a3c38ba9bf35b2 (emul8-temp)
INFO[0000]  rack-1-vec-f97f284037b04badb6bb7aacd9654a4e (emul8-temp)
INFO[0000]  rack-1-vec-eb9a56f95b5bd6d9b51996ccd0f2329c (emul8-fan)
INFO[0000]  rack-1-vec-f52d29fecf05a195af13f14c7306cfed (emul8-led)
INFO[0000]  rack-1-vec-d29e0bd113a484dc48fd55bd3abad6bb (emul8-led)
INFO[0000] --------------------------------
DEBU[0000] starting read-write poller
DEBU[0000] starting data updater
INFO[0000] listening on network unix with address /tmp/synse/procs/emulator.sock
INFO[0000] serving
```

To use the emulator external to Synse Server, you will need to hit its [gRPC API](). The [Synse CLI]()
provides capabilities for easily interacting with plugins in this way.

# Developing
The Makefile provides targets to simplify the development workflow. Use `make help` to list
all available targets, or see the [Makefile](Makefile).

## Setup
When you first get the source (see the Getting Started section), you may need some additional tooling
and the project will need to be vendored. Running
```
make setup
```
Will install additional development tools (`dep`, `gometalinter`) and will initialize vendoring
for the repo if it is not already done.

## Building
The emulator binary can be built for your OS/Arch with
```
make build
```

To build for Linux/AMD64
```
make build-linux
```

## Linting and Formatting
Prior to committing code, the source should be linted and formatted. These operations are
part of the CI workflow, so they will fail a build if not done.
```
# lint via gometalinter
make lint

# format via goimports
make fmt
```

## Troubleshooting
### Debugging
The plugin can be run in debug mode for additional logging. This is done by setting
```yaml
debug: true
```
in the plugin configuration YAML ([config.yml](config.yml))

### Bugs / Issues
If you experience a bug, would like to ask a question, or request a feature open a
[new issue](https://github.com/vapor-ware/synse-emulator-plugin/issues) and provide as much
context as possible.

# License
The Synse Emulator Plugin is licensed under GPLv2. See [LICENSE](LICENSE) for more info.
