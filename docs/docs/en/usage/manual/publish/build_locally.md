---
title: Build Package Locally (2/3)
---

## Build in Standalone Mode

Once the model is prepared, you can start to build it locally in standalone mode. By doing so you can verify if your model will work properly in Docker environment.

First switch into the package folder, i.e., where the `aid.toml` is located. Then run `aid build --path=./`. AID will automatically generate required dockerfiles and build the project.

:::caution
Currently build in standalone mode does not support GPU. Please do not use it in production environment. It is intended for testing only.
:::

After the build is done, you can create the docker container by running `docker create -p {host_port}:8080 {image_hash}`. Then you will see the hash to the created container in the terminal. You can use `docker start {container_hash}` to start the container, and test it with your own input. You probably need some http client, e.g., curl for testing.

:::caution
The build process does not considers the installation process, e.g., if the pretrained models can be downloaded properly. Please make sure the installation process is done properly, by checking the content in your pretrained file.
:::
