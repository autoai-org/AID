# Install and Use Packages

## Prerequisite


## Using CLI



## Using Dashboard (Web UI)

## Manually

To manually install a package, you need to:

1. **Prepare Package**. First, download the package by using ```git``` or other utilities, such as ```curl, wget```.
2. **Generate Dockerfile and Runner**. Switch the package root repository, run ```aid gen```.
3. **[Docker Runtime Only] Build Docker Image**. ```aid build {solverName}```.