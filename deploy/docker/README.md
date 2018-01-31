# Emulator Plugin Deployment
This directory contains examples for docker-compose based deployments of the
containerized Emulator Plugin and containerized Synse Server (v2.0).

## Deployments
As a general note - the Emulator Plugin does not need to be containerized. In
fact, it is built in to the Synse Server 2.0 image itself so that it can run
alongside it, providing an easy way to demo/play with Synse Server.

In this case, we use a containerized version of the same emulator in order to
give us a plugin that is not dependent on any platform or hardware. This makes 
the examples here a good place to get started.

There are two docker-compose based deployments here. The only difference between
the two is how Synse Server and the Plugin are configured to communicate. Currently,
we support communication via
- TCP port
- UNIX socket

The two deployments here give an example of how to run a plugin with Synse Server
in those cases.

The difference is entirely a configuration difference. See the compose files and
the corresponding configuration files in the `config/` subdirectory to see how
the two deployments differ.


## Setup
Note that the images required for these examples are not available pre-built.
They must be built locally. See the sections below for how to build the two
required images.

### vaporio/synse-server:2.0
For the Synse Server 2.0 image:

- checkout the `v2.0-dev` branch of `vapor-ware/synse-server-internal`
- do a `make build`
- create a `vaporio/synse-server:2.0` tag from the results of that build

### vaporio/plugin-emulator
for the Emulator Plugin image:

- checkout `master` of `vapor-ware/synse-plugins-internal` and make sure it
  is up to date
- `cd` into the `emulator/` subdirectory
- do a `make docker` to build the image 
- *(note: as of writing this, the changes for containerizing the emulator are
  still in a PR but they should make it into master shortly.)*


## Usage
Running either of the example is pretty straightforward. There are Makefile
targets to make it easy. To run the deployment that uses TCP-based communication
between Synse Server and the plugin, run
```
make tcp
```

To run the deployment that uses UNIX socket-based communication between 
Synse Server and the plugin, run
```
make unix
```

Functionally, both should behave the same, but the configuration for each 
case is different. See the compose files and their related config files 
(in the `config/` subdirectory) to see the differences between them.

Once one of the deployments is running, you can test out that Synse Server
is reachable.
```
curl localhost:5000/synse/test
```

If successful, you are ready to go. Next, perform a scan to see everything
that is available via the plugin:
```
curl localhost:5000/synse/2.0/scan
```

This should give back a set of devices - in particular, a fake fan device,
two fake LED devices, and a bunch of fake temperature devices. If you look at the
log output of the Emulator Plugin , you should see that these results match up
with what that plugin had registered on startup. 