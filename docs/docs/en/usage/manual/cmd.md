---
title: Command Line Usage
---

## List Something

You can list items installed on your system.

```sh
aid list repositories
aid list images
aid list solvers
aid list containers
```

The first column is always the id of that entity, and should be used in other commands, except build.

## Help

You can get the documents to AID Command Line by

```sh
aid --help
```

You can also print the help message of a single repositories
```sh
aid help {repository_id}
```

## Build

```sh
aid build {vendor_name}/{repository_name}/{solver_name}
```

This command will print the id of the image that it creates.

## Create (Bind local network)

```sh
aid create {image_id} {port}
```

where port is an integer. This command will print the id of the container that it creates.

## Start Serving

```sh
aid start {container_id}
```