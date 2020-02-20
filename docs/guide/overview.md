# Overview

In the past few years, deep learning has played an important role in solving computer vision tasks, such as image caption, image generation, image classification and etc. Developing deep learning applications usually begins from collecting and annotating datasets, finding and optimizing neural networks to evaluating performance on test datasets, and deployments. There appeared many frameworks, such as TensorFlow, PyTorch, Chainer, MXNet, etc.

The growing development of frameworks and libraries has greatly reduced the effort of designing, implementing and evaluating neural networks. However, the bottleneck still exists in large-scale datasets annotation, model sharing and deployments. Creating and annotating large datasets is labour-intensive and involves significant time and costs, which makes it unaffordable for small companies. Besides, the main approach to sharing and redistribution of deep learning models and algorithms is redistributing the source code and weight file, which causes incompatibilities and conflicts and brings extra and prerequisites knowledge when deploying computer vision services.

To reduce the time and efforts for annotating datasets, versioning, deploying and sharing models, we proposed and developed CVFlow for Open Computer Vision Workflow. It is a suite of toolkits and libraries that currently provides 

- Intelligent annotation that can help annotators improve efficiency by providing suggestions.

- Package manager that manages models, datasets, services and serving environment.

- Model hub for downloading, sharing and discovering existing computer vision models.

CVFlow contains two independent but collaborative modules, which are CVPM for package management and model sharing, as well as IAE for intelligent and efficient annotation. In the following, we will describe these modules respectively.

CVPM is the abbreviation of computer vision package manager. It allows users to download, share, and install new computer vision packages. Traditional computer vision libraries often provide some high-level APIs for development such as torch vision, mmcv and etc. On the contrary, CVPM provides a package management system which allows users to download computer vision packages across numerous frameworks and apply them into a Restful HTTP API. By doing so, users can install and deploy computer vision models and services with few clicks or commands, without writing any code. Meanwhile, applying computer vision services into their products would be much easier since CVPM provides a similar restful API across many different libraries. 

Lots of big companies have realised that one of the biggest bottlenecks in modern machine learning applications is the growing need for large, labelled datasets. Creating and annotating the datasets requires significant cost and time. To facilitate such a need, we have included Integrated Annotation Environment (IAE) for intelligent visual objects annotation.

CVFlow was released on [GitHub](https://github.com/autoai-org/aid) from 2018, Since then the proposed workflow has been adopted by several online computer vision platforms.
