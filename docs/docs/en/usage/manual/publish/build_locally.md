---
title: Build Package Locally (2/3)
---

## Build in Standalone Mode

Once the model is prepared, you can start to build it locally in standalone mode. By doing so you can verify if your model will work properly in Docker environment.

First switch into the package folder, i.e., where the ```aid.toml``` is located. Then run ```aid build --path=./```. AID will automatically generate required dockerfiles and build the project.

:::caution
Currently build in standalone mode does not support GPU. Please do not use it in production environment. It is intended for testing only.
:::