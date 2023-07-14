[![Build Status](https://github.com/vapor-ware/synse-emulator-plugin/workflows/build/badge.svg)](https://github.com/vapor-ware/synse-emulator-plugin/actions)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin?ref=badge_shield)

# Synse Emulator Plugin

A plugin with emulated devices for the Synse platform.

This plugin is designed to be simple and have no dependencies on any backend hardware
or protocol. It is meant to be a standalone plugin that can run with Synse Server
out of the box allowing you to quickly experiment and develop with Synse Server. It
may also serve as an example on how to write plugins of your own.

## Getting Started

### Getting

You can install the emulator plugin via a [release](https://github.com/vapor-ware/synse-emulator-plugin/releases)
binary or via Docker image

```
docker pull vaporio/emulator-plugin
```

If you wish to use a development build, fork and clone the repo and build the plugin
from source.

### Running

A [compose file](docker-compose.yml) is included in this repo which provides a basic example of
how to run the emulator plugin in conjunction with Synse Server. You may also run the emulator on
its own

```
docker run -d --name emulator -p 5001:5001 vaporio/emulator-plugin
```

and use the [Synse CLI](https://github.com/vapor-ware/synse-cli) to query the plugin's gRPC API.

The emulator plugin ships with some default built-in [device configurations](config/device). If you
wish to run the plugin with a different set of plugins, be sure to read through the
[SDK Documentation](https://synse.readthedocs.io/en/latest/sdk/configuration/device/) on configuring
plugin devices.

## Emulator Plugin Configuration

All devices supported by the emulator plugin return dummy data. State is maintained for most writable
devices, so any state that you write should be retrievable via a subsequent read.

### Outputs

Outputs are referenced by name. A single device may have more than one instance
of an output type. A value of `-` in the table below indicates that there is no value
set for that field. The *custom* section describes outputs which this plugin defines
while the *built-in* section describes outputs this plugin uses which are [built-in to
the SDK](https://synse.readthedocs.io/en/latest/sdk/concepts/reading_outputs/#built-ins).

**Custom**

| Name    | Description                                      | Unit  | Type    | Precision |
| ------- | ------------------------------------------------ | :---: | ------- | :-------: |
| airflow | A measure of airflow, in millimeters per second. | mm/s  | `speed` | 3         |

**Built-in**

| Name          | Description                                        | Unit  | Type          | Precision |
| ------------- | -------------------------------------------------- | :---: | ------------- | :-------: |
| color         | A color, represented as an RGB string.             | -     | `color`       | -         |
| direction     | A measure of directionality.                       | -     | `direction`   | -         |
| humidity      | A measure of humidity, as a percentage.            | %     | `humidity`    | 2         |
| kilowatt-hour | A measure of energy, in kilowatt-hours.            | kWh   | `energy`      | 3         |
| pascal        | A measure of pressure, in Pascals.                 | Pa    | `pressure`    | 3         |
| rpm           | A measure of frequency, in revolutions per minute. | RPM   | `frequency`   | 2         |
| state         | A generic description of state.                    | -     | `state`       | -         |
| status        | A generic description of status.                   | -     | `status`      | -         |
| temperature   | A measure of temperature, in degrees Celsius.      | C     | `temperature` | 2         |
| voltage       | A measure of voltage, in Volts.                    | V     | `voltage`     | 5         |
| watt          | A measure of power, in Watts.                      | W     | `power`       | 3         |

### Device Handlers

Device Handlers are referenced by name.

| Name        | Description                                 | Outputs                   | Read  | Write | Bulk Read | Listen |
| ----------- | ------------------------------------------- | ------------------------- | :---: | :---: | :-------: | :----: |
| airflow     | A handler for emulated airflow devices.     | `airflow`                 | ✓     | ✓     | ✗         | ✗      |
| energy      | A handler for emulated energy devices.      | `kilowatt-hour`           | ✓     | ✓     | ✗         | ✗      |
| fan         | A handler for emulated fan devices.         | `direction`, `rpm`        | ✓     | ✓     | ✗         | ✗      |
| fan-multi   | A handler for emulated fan devices with multiple RPM readings. | `direction`, `rpm` | ✓ | ✓ | ✗     | ✗      |
| humidity    | A handler for emulated humidity devices.    | `humidity`, `temperature` | ✓     | ✓     | ✗         | ✗      |
| led         | A handler for emulated LED devices.         | `color`, `state`          | ✓     | ✓     | ✗         | ✗      |
| lock        | A handler for emulated lock devices.        | `status`                  | ✓     | ✓     | ✗         | ✗      |
| power       | A handler for emulated power devices.       | `watt`                    | ✓     | ✓     | ✗         | ✗      |
| pressure    | A handler for emulated pressure devices.    | `pascal`                  | ✓     | ✓     | ✗         | ✗      |
| temperature | A handler for emulated temperature devices. | `temperature`             | ✓     | ✓     | ✗         | ✗      |
| voltage     | A handler for emulated voltage devices.     | `voltage`                 | ✓     | ✓     | ✗         | ✗      |

### Write Values

This plugin supports the following values when writing to a device via a handler.

The emulator enables writing on devices where it would generally not otherwise be possible,
allowing the user to adjust the reading values and ranges at runtime without having to re-deploy.
Handlers set up this way will have the `min`, `max`, and `current` write actions available.

| Handler     | Write Action  | Write Data | Description |
| ----------- | :-----------: | :--------: | ----------- |
| airflow     | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| energy      | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| fan         | `speed`       | `int`      | The speed to set the fan to. |
| humidity    | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| led         | `state`       | `string`: *on*, *off*, *blink* | The LED power state. |
|             | `color`       | `string`   | RGB hex color string. |
| lock        | `lock`        | -          | Lock the door. |
|             | `unlock`      | -          | Unlock the door. |
|             | `pulseUnlock` | -          | Unlock the door for a short time, then lock it again. |
| power       | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| pressure    | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| temperature | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |
| voltage     | `min`         | `int`      | The minimum bound for readings to be generated within. |
|             | `max`         | `int`      | The maximum bound for readings to be generated within. |
|             | `current`     | `int`      | The static current reading value. |

### Device Configurations

In addition to allowing values to be set via writing to devices, the behaviors/data ranges
for device readings may also be configured in the device configuration YAML for certain
devices. These options should be set under a device's `data` configuration, e.g.

```yaml
version: 3
devices:
  - type: temperature
    context:
      model: emul8-temp
    instances:
      - info: Synse Temperature Sensor 1
        data:
          id: 1
          min: 20
          max: 40
```

The table below describes the supported configuration values for each type. None of the values
specified below are required to be set. If omitted, the plugin falls back on sane defaults.
Definitions for the settings follow:

* `min`: The lower bound for a random walk. The walk will not subceed this value.
* `max`: The upper bound for a random walk. The walk will not exceed this value.
* `step`: The maximum size that a random walk may step by. For each iteration, the step is randomly chosen
   between this max step size and 0.
* `seed`: The starting value for the device.

| Device Type | Setting | Type  | Default |
| :---------: | :-----: | :---: | :-----: |
| airflow     | `min`   | int   | -100 |
|             | `max`   | int   | 100 |
|             | `step`  | int   | 4 |
| current     | `min`   | int   | 0 |
|             | `max`   | int   | 30 |
|             | `step`  | int   | 4 |
| energy      | `min`   | int   | 0 |
|             | `max`   | int   | 100000 |
| fan         | `seed`  | int   | 0 |
| frequency   | `min`   | int   | 0 |
|             | `max`   | int   | 60 |
|             | `step`  | int   | 4 |
| humidity    | `min`   | int   | 0 |
|             | `max`   | int   | 100 |
|             | `step`  | int   | 4 |
| power       | `min`   | int   | 1000 |
|             | `max`   | int   | 3000 |
|             | `step`  | int   | 4 |
| pressure    | `min`   | int   | -5 |
|             | `max`   | int   | 5 |
|             | `step`  | int   | 4 |
| temperature | `min`   | int   | 0 |
|             | `max`   | int   | 100 |
|             | `step`  | int   | 4 |
| voltage     | `min`   | int   | 100 |
|             | `max`   | int   | 500 |
|             | `step`  | int   | 0 |

## Compatibility

Below is a table describing the compatibility of plugin versions with Synse platform versions.

|             | Synse v2 | Synse v3 |
| ----------- | -------- | -------- |
| plugin v1.x | ✗        | ✗        |
| plugin v2.x | ✓        | ✗        |
| plugin v3.x | ✗        | ✓        |

## Troubleshooting

### Debugging

The plugin can be run in debug mode for additional logging. This is done by:

- Setting the `debug` option  to `true` in the plugin configuration YAML ([config.yml](config.yml))

  ```yaml
  debug: true
  ```

- Passing the `--debug` flag when running the binary/image

  ```
  docker run vaporio/emulator-plugin --debug
  ```

- Running the image with the `PLUGIN_DEBUG` environment variable set to `true`

  ```
  docker run -e PLUGIN_DEBUG=true vaporio/emulator-plugin
  ```

### Building

To build the production image locally, you'll first need to [install the goreleaser binary](https://goreleaser.com/install/).

```
make docker
```

This will run `goreleaser` to build a local image tagged as the current release, i.e., `docker.io/vaporio/emulator-plugin:3.4.1`.

### Developing

A [development/debug Dockerfile](Dockerfile.dev) is provided in the project repository to enable
building image which may be useful when developing or debugging a plugin. Unlike the slim `scratch`-based
production image, the development image uses a Debian base, bringing with it all the standard command line
tools one would expect. To build a development image:

```
make docker-dev
```

The built image will be tagged using the format `dev-{COMMIT}`, where `COMMIT` is the short commit for
the repository at the time. This image is not published as part of the CI pipeline, but those with access
to the Docker Hub repo may publish manually.

## Contributing / Reporting

If you experience a bug, would like to ask a question, or request a feature, open a
[new issue](https://github.com/vapor-ware/synse-emulator-plugin/issues) and provide as much
context as possible. All contributions, questions, and feedback are welcomed and appreciated.

# License

The Synse Emulator Plugin is licensed under GPLv3. See [LICENSE](LICENSE) for more info.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin?ref=badge_large)
