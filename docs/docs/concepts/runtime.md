---
title: Runtime
---

***Runtime***. Runtime is where the solver program actually runs. In the past, we let the solver run on bare-metal, without any isolation across different packages. That old approach led to several problems, especially the incompatibility of dependencies across two packages. Thus, in the latest version, we allow users to use container/docker as their runtime, which greatly reduced the effort in managing dependencies.