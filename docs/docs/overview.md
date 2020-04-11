---
title: Overview
---

In the past few years, deep learning has played an important role in solving computer vision tasks, such as image caption, image generation, image classification and etc. Developing deep learning applications usually begins from collecting and annotating datasets, finding and optimizing neural networks to evaluating performance on test datasets, and deployments. There appeared many frameworks, such as TensorFlow, PyTorch, Chainer, MXNet, etc.

The growing development of frameworks and libraries has greatly reduced the effort of designing, implementing and evaluating neural networks. However, the bottleneck still exists in large-scale datasets annotation, model sharing and deployments. Creating and annotating large datasets is labour-intensive and involves significant time and costs, which makes it unaffordable for small companies. Besides, the main approach to sharing and redistribution of deep learning models and algorithms is redistributing the source code and weight file, which causes incompatibilities and conflicts and brings extra and prerequisites knowledge when deploying computer vision services.

To reduce the time and efforts for annotating datasets, versioning, deploying and sharing models, we proposed and developed A.I.D for Open Machine Learning Workflow. It is a suite of toolkits and libraries that currently provides 

- Intelligent annotation that can help annotators improve efficiency by providing suggestions.

- Package manager that manages models, datasets, services and serving environment.

- Model hub for downloading, sharing and discovering existing computer vision models.

A.I.D was released on [GitHub](https://github.com/autoai-org/aid) from 2018.
