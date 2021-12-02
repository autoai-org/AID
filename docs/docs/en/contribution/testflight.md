---
title: Test Flight
---

We do highly appreciate your effort in helping us test our releases! Here we provide a step-by-step protocol on how to proceed with the test process.

In the testing process, please report the bugs, suggestions or any other experience issues to [GitHub issues](https://github.com/autoai-org/AID/issues). If you do not have one, you can also use the alternative [Google Spreadsheet](https://docs.google.com/spreadsheets/d/1kCr-AbiU-SUZyh3QAc3WO0SZQq2NVgQr1jRfGh4NML0/edit?usp=sharing).

# Pre-requisites

You must have Docker installed. Please check [Get Docker](https://docs.docker.com/get-docker/) for how to install Docker on your machine.

You can either use the edge build of AID from [Installation](/docs/getting-started/installation) or get a internal build from the source code directly.

# Testing Steps

## Find Models

We have an experimental new [Model Hub](https://hub.autoai.dev), please try it out!

## Build Models

Probably you will choose `ocr` model, after installation as described on the Model Hub, you need to start to build it by the following command:

```shell
aid build aidmodels/ocr/ocrSolver
```

Then the system will let you know the `id` of the result, and you will need to bind it to a local port:

```shell
aid create {YOUR_ID_HERE} {YOUR_PORT_HERE}
```

For example:

```shell
aid create ef6782b0 8081
```

Then AID will show the new `id` of the service. Please note that it will be **different** from the id before. You can start the service by:

```shell
aid start {YOUR_ID_HERE}
```

Some of the solver names are listed below:

- aidmodels/ocr/ocrSolver
- aidmodels/detectron/detection
- eth-library-lab/speech-analysis/speechSolver

## Testing Predicion with Curl

```sh
curl -X POST -F file=@test.png http://127.0.0.1:{YOUR_PORT_HERE}/infer
```

# Remove Installed Repositories

```shell
rm -rf ~/.autoai/aid
```

The above command will uninstall the repositories on your machine, but will not remove the docker images. If you want to remove the docker images, you can use the following command:

```
docker ps -a # Found all containers
docker stop {container_id} # Stop the container
docker rm {container_id} # Remove the container
docker images # Found all images
docker rmi {image_id} # Remove the image
```

By removing the contents, you can reset your system to the initial state, and make the environment clean enough for next test/other tasks.
