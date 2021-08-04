---
title: Perform Inference (3/3)
---

## Build Image and Container

After installing the packages, you need to build it into an image which consists of all dependencies, pretrained model and source code. Use ```aid build [vendor name]/[package name]/[solver name]``` to build image. Once finished, it will return a reference number for the built image. You can check images by running ```aid list images```.

Now, you can start to create a corresponding container that bind the local IP to this image. Use ```aid create [image ID] [port number]``` to create a container. A container ID will be given if succeeded. You can also check the ID by ```aid list containers``` command.

## Deploy to HTTP Service

To start the container, ```aid start [container ID]``` will help to deploy HTTP service to the specified port number. Now you can check if the contianer run on the port successfully by using ```aid list containers```. If running properly, the status of Running will be TRUE.

## Test the Model

Now you can test the model by ```aid infer [container ID] [argument]="content/path"```. The solver will give the results. You can use ```aid help [package ID]``` to check the meaning of results.