---
title: Installation
---

## Prerequisite

AID relies on several dependencies to contain the machine learning algorithms. In order to install AID, your system must satisfy the following requirements:
* **Operation Systems**: Linux (generic) or MacOS. We recommend to use Linux as your base operation system. If you are using Windows, please consider using Windows Subsystems for Linux (aka WSL).
* **Docker**: The system must have installed docker on the system. For instructions, please take a look at [Get Docker](https://docs.docker.com/get-docker/).
  
## Command Line Utility

You can install the edge releases of the command line utility by running the command below:

```
curl https://releases.autoai.org/aid/install.sh | sudo bash -s -- edge
```

:::caution
Running an unknown scripts with root privilege may damage your computer.

Do not directly copy and run the above command to your terminal, especially if you do not trust your network provider. You can first download the ```.sh``` file and audit its content, and then run the ```.sh``` file with root privilege.

We can only ensure the content is free from damage on our server, but your network might be hijacked.
:::

### Build from Source

Build instructions will be released soon.