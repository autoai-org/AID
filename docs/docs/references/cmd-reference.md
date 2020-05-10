---
title: Command Line Reference
---

The command line utility, ```components/cmd```, is the most important tool to interact with AID. It supports all features, and gives users the maximal usability and flexibility to control MLOps pipeline.

After successfully installed the command line interface on your computer, you can use ```aid help``` to fetch a global manual. Its something like:

```
> aid help
NAME:
   AID - One-Stop AIOps System.

USAGE:
   aid [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
   help, h  Shows a list of commands or help for one command
   cluster:
     cluster, cl  Cluster Management -  Add Node/Upload Image/etc
   daemon:
     up  Server Up
   image:
     image, img  Manage images - Export/Import/List prebuilt images
   packages:
     install, i  Install Package
     new         Initialize a new package
   runtime:
     build, b       Build Image
     create, cc     Create Container
     generate, gen  Generate Runners and Dockerfile
   self:
     config, conf     Reset Configurations to default value
     init             Initialize all configurations, database and web ui
     interactive, in  Interactive Mode
   storage:
     db, db  database
   utility:
     list, ls  aid list [image|container|config]

GLOBAL OPTIONS:
   --license      print the license (default: false)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

You can find all commands, their description and examples below:

| Command                             | Description                                                              | Example                                                          |
| ----------------------------------- | ------------------------------------------------------------------------ | ---------------------------------------------------------------- |
| `list {entity}`                     | List entities in AID, now supports images/containers/config/nodes.       | `aid list images`                                                |
| `install {link}`                    | Install third pre-built models from GitHub or AIDHub.                    | `aid install https://github.com/aidmodels/sentiment-analysis`    |
| `build {vendor}/{package}/{solver}` | Build docker images with certain vendor/package/solver triplet.          | `aid build aidmodels/sentiment-analysis/sentimentSolver`         |
| `create {image} {port}`             | Create docker containers with the image, and make it listen on the port. | `aid create aid-1-cyy-sentiment_prediction-sentimentsolver 8081` |
