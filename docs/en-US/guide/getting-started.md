# Getting Started

## Prerequisite

CVPM requires Python 3 and Pip. Before your installation, please make sure you have both in your system.

You can download Python 3 from [here](https://www.python.org/).

## Automatically Installation

## Manually Installation

### Download Executable Binary File

Go to our [release page](https://github.com/unarxiv/CVPM/releases) to download latest binary file.

Place the binary into your ```PATH```.

If your Python and Pip command are not in your ```PATH```, RUN ```cvpm config``` to set your python and pip location.

### Install CVPM Python Dependency

RUN ```cvpm install cvpm:stable``` or ```cvpm install cvpm:test```. It helps you to install cvpm python dependency into your system.

### Install Packages

RUN ```cvpm install vendor/repo``` to install a package from [model hub](https://hub.autoai.org).

Or, RUN ```cvpm install https://github.com/vendor/repo``` to install a package from GitHub. Make sure the repo you are going to install is CVPM Available.

### Run it!

Before you run any of your installed repo, a daemon process must be enabled. Note that you only need to enable it once by running 
```cvpm daemon install```. It may requires you administrator permission to install on most systems.

Some repos may have more than one solver in it, therefore, you have to know the exact solver name to run. By running ```cvpm run vendor/repo/solver port```, you can start a service.

### Example

For example, we have an official model named Face_Utility with three solvers. If you want to run a face detection solver, after your installation, simply do as following:

```
cvpm install https://github.com/cvmodel/Face_Utility
cvpm repo run cvmodel/Face_Utility/Face_Detection 8080
```