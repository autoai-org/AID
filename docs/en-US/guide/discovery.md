# Discovery

## Overview

Discovery is one of the most important part of cvpm, which allows thirdparties to provide model registries on their behalf, users to query models with SDKs and etc.

It is an open source service, locates in the ```discovery``` folder, and deployed as a docker-based distributed service for use. 

## Infrastructure

Thanks to our generous [sponsors](https://cvpm.autoai.org/en-US/guide/credits.html), we could be able to deploy and maintain our discovery services on various kind of cloud services. 

Our main infrastructure depends on the following cloud services:

* **AWS DynamoDB** for providing database.
* **Open Stack** for providing virtual machines.
* **Docker Cloud** for providing container build and deploy services. 
* **Cloudflare** for DNS.

### Region

Our servers locates in difference regions and by using load balancer, it should be able to help users in different areas with enjoying the benefit. We now support the following supported area:

* **Hong Kong** (Provided by Cyberport Hong Kong)
* **Mainland China** (Provided by Tencent Cloud)
* **Nova, US** (Provided by Open Source Lab, Oregon State University)

But for now, our storage service in only available in the following area and our intranet.

* **Nova, US**(Provided by Open Source Lab, Oregon State University)
* **Mainland China** (Provided by DiDi Cloud)

Other *paid* models are currently served by Amazon S3 and therefore have a better bandwidth support.

## Services

Our Discovery Service provides the following services:

* **Model Discovery**
* **Registry Management**
* **Pretrained Discovery**
* **User Authentication**
* **Other Utilities**

## API Reference