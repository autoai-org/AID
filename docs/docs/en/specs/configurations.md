---
title: Configuration Files
---

## aid.toml

The ```aid.toml``` is the main entrance of a specific machine learning model. It has the following format:

``` toml
[package]
name="sentiment_prediction"
vendor="aidmodels"
tagline="Sentiment Analysis"
key=value // optional parameters

[[solvers]] 
name="sentimentSolver"
class="sentiment_prediction/solver/SentimentSolver"
```

The optional parameters includes:

```toml
gpu=true // indicate that a GPU is suggested to run.
```

:::note
Currently all the optional parameters are package-wise, i.e. you cannot set a specific parameter for certain solvers.
:::

## pretrained.toml

The ```pretrained.toml``` indicates where to find and download the pretrained weights and other assets. It has the following format:

```toml
[[models]]
name = "sentiment"
url = "https://github.com/aidmodels/sentiment-analysis/releases/download/v0.1/model.h5"

[[models]]
name = "some_file..."
url = "https://github.com/aidmodels/sentiment-analysis/releases/download/v0.1/somefile.h5"
```

AID will automatically fetch the pretrained file and put them ```/pretrailed``` folder under the project root.

:::note
Though you can download files when initialising the solvers, just like [this repository](https://github.com/aidmodels/image_encoding), it is not suggested to do so as it will not allow AID to read the pretrained weights for analysis and possibly conversion in the future. In the meanwhile, downloading when initialising the solvers may require to download multiple times if you have re-created the container, which is less efficient and causes extra network traffic.
:::