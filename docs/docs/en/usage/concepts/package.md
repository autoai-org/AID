---
title: Package
---

- ***Solver***. The solver is the minimal unit to solve a certain machine learning task. It is ideally a single Python class that extends from `aid.Solver`, and implement the `__init__` function (where to load the pretrained file, initialize the inference class, etc) and `infer` function (where actually do the prediction). In addition, it can optionally implement `train` function (This is still in discussion), which will allow users to upload a training dataset and get their fine-tuned model. With the class defined, AID will be able to automatically turn a solver into a web service. AID will also create a `batch` function, which reads a `.csv` file and call the `infer` function row by row.

Each solver has several different status.

* Running. The solver is running and listening on a specific port.
* Ready. The solver has been built into docker images, and can start running at any time.
* (Fully) Installed. The solver is fully installed, i.e. the code and artifacts have been installed on the machine.
* (Partial) Installed. The docker image of the solver is installed, but AID cannot find its corresponding code.

- ***Package***. The package is a set of solvers. This is because several solvers can share something in common when performing machine learning tasks. For example, when performing face recognition, in some cases we will have a face embedding model. With the model, we can calculate the distances between two faces, but also the distance between the new face to the faces space. The latter can be used for face detection tasks. Conceptually, the package is such a set of solvers that have similar tasks, like face recognition and face detection, and in our implementation, they will share same environment variables, enabled extensions, dependencies, etc. The required file for a package can be seen as below:

![0bf58c41e81106ac7387f6bde1af0662767378ab.PNG](https://i.loli.net/2020/05/06/7T3OuXoeCxS5Hsp.png)

