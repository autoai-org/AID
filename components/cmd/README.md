# AID Command Line Utility

to get started making adding features:

1. make a local build

```
cd components/cmd
make ci-build
```

2. run commands using the new local build
   examples:

```
./aid --help
```

or

```
./aid up
```

3. make changes and rebuild the relevant component, for example for the cli

```
make build-cli
```
