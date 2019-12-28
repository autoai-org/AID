# Installation

## Install from Package Manager

## Download Binary

We periodically release CVFlow for different platforms, including Linux, Mac and experimental Windows. You can download the release file from our GitHub Repositories.

[CVPM Releases](https://github.com/unarxiv/cvpm/releases)

[IAE Releases](https://github.com/unarxiv/iae/releases)

## Build from Source

If you are looking for cutting edge features, we would recommend you build cvpm and IAE from source.

### Prepare Build Environment

To build IAE from source, you will need at least node v10.0+. We only test the build process with latest node. For installation instructions, please refer to [NodeJS Downloads](https://nodejs.org/en/download/)

To build CVPM from source, you will need at least Golang compiler at least v1.12. The build status can be viewed at [Travis](https://travis-ci.org/unarxiv/CVPM). For installation of Golang, please refer to [Go Downloads](https://golang.org/dl/)

### Prepare Source Code

``` bash
# download IAE code to ./IAE
git clone https://github.com/unarxiv/IAE.git  
# download cvpm code to ./cvpm
git clone https://github.com/unarxiv/cvpm.git 
```

### Install Dependencies

To install dependencies of IAE, please follow the instructions below:

```shell
cd iae/app
# If you are using yarn as package manager for NodeJS.
yarn
# If you are using NPM as package manager for NodeJS.
npm install
```

As for CVPM, when building it for the first time, the process will automatically install all dependencies.

### Build

To serve IAE in development mode,

```shell
cd iae/app
# If you are using yarn as package manager for NodeJS.
yarn run electron:serve
# If you are using NPM as package manager for NodeJS.
npm run electron:serve
```

To build IAE for production,

```shell
cd iae/app
# If you are using yarn as package manager for NodeJS.
yarn run electron:build
# If you are using NPM as package manager for NodeJS.
npm run electron:build
```

To build CVPM,

```shell
cd cvpm/core/cmd
go build
```

To build CVPM for different platforms, we suggest,

```shell
cd cvpm/core/cmd
./build.sh cvpm
```

it will build cvpm for Windows, Mac OS X, Linux and FreeBSD and stores the binary under cvpm/core/cvpm/dist folder.

You can also set ```GOARCH``` and ```GOOS``` manually to build CVPM for different platforms, for more details, please refer to [Go Cross Compile](https://golangcookbook.com/chapters/running/cross-compiling/)
