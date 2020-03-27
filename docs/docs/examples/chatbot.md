---
title: Chatbot
---
```Aidmodels/chatbot``` is provided as an example for chatbot.

## Installation

* **[Install package]** Run ```aid install https://github.com/aidmodels/Chatbot``` to fetch the required file. It will be downloaded to ```~/.autoai/.aid/models/aidmodels/Chatbot```.

* **[Generate required files]** Switch into the directory ```~/.autoai/.aid/models/aidmodels/Chatbot```, run ```aid gen``` to generate the required dockerfile and runners.

* **[Build Docker image and create containers]** Run ```aid build chatbotSolver``` inside the package directory, and run ```aid create aid-{No.}-aicamp-chatbot-chatbotsolver {PORT}```. Here the ```No.``` will be ```1``` by default (It will also appear in your terminal after you run ```aid build```). You can specify any port here, as long as it is not occupied.

## Start and Provision

* **[Start the container]** Run ```docker start {Container_No.}```, here the ```Container_No.``` is the number that was given after ```aid create```.

## Test

* **[Test the Inference Service]** Run ```curl -X POST -F input="three plus five is" http://127.0.0.1:{PORT}/infer``` in your terminal and check the results. Here the PORT is what you specified in Step 3.

![](/img/docs/chatbot-example.png)