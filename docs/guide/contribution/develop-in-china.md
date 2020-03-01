# Develop in China

Some important build tools are difficult to access in some contries. To help them participate in the development, we provide some documents.

## Golang

The package management of golang, is hard to access, i.e. ```go get``` and related commands may run into problems. To handle this issue, please try (GOPROXY)[https://goproxy.io/] and follow their instructions.

## Docker

```Image pull``` may be slow due to the outage of external bandwidth. We recommend you try the official docker-china registry. To set the registry, simply edit ```/etc/dockeer/daemon.json``` as 

```
{
  "registry-mirrors": ["https://registry.docker-cn.com"]
}
```

There are some other Chinese companies providing private/public registries.

## NPM

NPM might be slow, but it should work. If you want a faster experience, try to put the ```.npmrc``` in your project folder:
```
registry=https://registry.npm.taobao.org
chromedriver_cdnurl=http://npm.taobao.org/mirrors/chromedriver
selenium_cdnurl=http://npm.taobao.org/mirrorss/selenium
phantomjs_cdnurl=http://npm.taobao.org/mirrors/phantomjs
electron_mirror=https://npm.taobao.org/mirrors/electron
sass_binary_site=https://npm.taobao.org/mirrors/node-sass
node_inspector_cdnurl=https://npm.taobao.org/mirrors/node-inspector
```