[![Build Status](https://build.vio.sh/buildStatus/icon?job=vapor-ware/synse-emulator-plugin/master)](https://build.vio.sh/blue/organizations/jenkins/vapor-ware%2Fsynse-emulator-plugin/activity)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin?ref=badge_shield)

# Synse Emulator Plugin

A plugin with emulated devices for the Synse platform.

This plugin is designed to be simple and have no dependencies on any backend hardware
or protocol. It is meant to be a standalone plugin that can run with Synse Server
out of the box allowing you to quickly experiment and develop with Synse Server. It
may also serve as an example on how to write plugins of your own.

## Getting Started

### Installing

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
[SDK Documentation](https://synse.readthedocs.io/en/latest/sdk/user/configuration.device/) on configuring
plugin devices.

## Emulator Plugin Configuration

All devices supported by the emulator plugin return dummy data. State is maintained for most writable
devices, so any state that you write should be retrievable via a subsequent read.

### Outputs

Outputs are referenced by name. A single device may have more than one instance
of an output type. A value of `-` in the table below indicates that there is no value
set for that field. The *custom* section describes outputs which this plugin defines
while the *built-in* section describes outputs this plugin uses which are [built-in to
the SDK](https://github.com/vapor-ware/synse-sdk/blob/v3/staging/sdk/output/builtins.go).

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


## Troubleshooting

### Debugging

The plugin can be run in debug mode for additional logging. This is done by setting

```yaml
debug: true
```

in the plugin configuration YAML ([config.yml](config.yml))

The plugin may also be run with the `--debug` flag, e.g.

```
docker run vaporio/emulator-plugin --debug
```

### Bugs / Issues

If you experience a bug, would like to ask a question, or request a feature, open a
[new issue](https://github.com/vapor-ware/synse-emulator-plugin/issues) and provide as much
context as possible.

# License

The Synse Emulator Plugin is licensed under GPLv3. See [LICENSE](LICENSE) for more info.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-emulator-plugin?ref=badge_large)
