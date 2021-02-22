---
title: Getting Started
---

## Install AID

See [Releases](/docs/releases)

## Install a package

Run ```aid install {Package Name}``` to fetch the required file. It will be downloaded to ```~/.autoai/.aid/models/{Vendor}/{Package}```.

## Build Docker image and create containers

Run ```aid build {Vendor}/{PackageName}/{solverName}```, and run ```aid create aid-{No.}-{Vendor}-{Package}-{SolverName} {PORT}```. Here the ```No.``` will be ```1``` by default (It will also appear in your terminal after you run ```aid build```). You can specify any port here, as long as it is not occupied.

## Start the web server

Run ```docker start {Container_No.}```, here the ```Container_No.``` is the number that was given after ```aid create```.

## Test

Run ```curl -X POST -F input="{Input}" -F file=@test.png http://127.0.0.1:{PORT}/infer``` in your terminal and check the results. Here the PORT is what you specified in above steps.

## Batch Inference

Run ```curl -X POST  -F file=@test.csv http://127.0.0.1:{PORT}/batch``` in your terminal to call the batch inference service.
