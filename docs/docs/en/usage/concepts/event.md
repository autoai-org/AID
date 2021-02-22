---
title: Event
---

***Event***. Event is a protocol used to call external services. Every time a certain operation, such as onImageBuilt, onServiceUp, etc, is finished, an event is triggered. Then AID will iterate over all registered external services (named extensions, we will introduce this later) to check if they support this kind of service. If so, AID will trigger the service to run with the given data.

![7cf0bbba8d6d371f9d1c883698f28bdc317ab530.png](https://i.loli.net/2020/05/06/NxgWSklTVpntLdc.png)

The overall workflow is illustrated above. Overall there are four most important processes and corresponding events: *Install Package*, *Start Service*, *Receive Inference Requests* and *Process events*.

