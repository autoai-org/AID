# Install and Use Packages

## Prerequisite

CVPM requires Python 3 and Pip (Package Installer for Python). Before your installation, please make sure you have both in your system. Please refer to [Installing Pip](https://pip.pypa.io/en/stable/installing/) for how to install Pip on your system.

After installing pip, run the following command to config cvpm to use the correct pip and python on your system.

```shell
$ cvpm config
Original Python Location: python
✔ Python(3): █ (type your python 3 path here)
Original Pip Location: pip
✔ Pip(3): █ (type your Pip3 path here)
```

You will see a clear indicator telling you if the executable binary file exists. The files typically exists under ```/usr/bin``` or ```/usr/local/sbin``` folder.

Then install the cvpm base package in Python.

```shell
cvpm install cvpm:py
```

or if you prefer to use original pip to install,

```shell
pip3 install cvpm
```

## Using CLI



## Using Dashboard (Web UI)
