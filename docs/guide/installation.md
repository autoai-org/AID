# Installation

## Command Line Utility

### Prerequisite

The command line utility is written in Golang. In order to build the utility, you will need to have ```Golang > 1.12```.

We use ```go mod``` to manage third-party dependencies, so you will need to have it enabled.

### Build from Source

To install cutting-edge version of AID, we recommend you to build it from source.

1. Pull/Clone/Download the whole repository. (`git pull https://github.com/autoai-org/AID.git`)
2. Switch into the command line utility source code. (`cd aid/components/cmd`)
3. To build for testing, switch to the ```entry``` folder. (`cd entry`), and build the source code with ```go build```.
4. To build for production and multiplatform, you can manually set ```go build envs```, but we recommend you stay in ```aid/components/cmd``` and run ```build.sh```.
5. You will have to binary ready inside entry for testing or in ```aid/components/cmd/dist``` for production.

Note that building from source is not promised to be stable and may incur unknown issues, bugs. Use it with CAUTION!