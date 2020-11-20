# hypercloud-operator-go

[![Github All Releases](https://img.shields.io/github/v/release/tmax-cloud/hypercloud-operator-go?include_prereleases)](https://github.com/tmax-cloud/hypercloud-operator-go/releases/latest)
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/tmax-cloud/hypercloud-operator-go)](https://goreportcard.com/report/github.com/tmax-cloud/hypercloud-operator-go)

This is hypercloud operator

## Features
namespaceclaim
roleBindingclaim
resourceQuotaClaim
namespace

## To start developing hypercloud-operator

### Prerequisites

**0. You have a working [Go environment](https://golang.org/doc/install)**
```bash
$ git clone https://github.com/tmax-cloud/hypercloud-operator-go
$ cd hypercloud-operator-go
$ export GO111MODULE=on
$ make
```

**1. Install kubebuilder (version: [2.3.1](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v2.3.1))**
```bash
$ release=2.3.1
$ os=$(go env GOOS)
$ arch=$(go env GOARCH)
$ curl -sL https://go.kubebuilder.io/dl/2.3.1/${os}/${arch} | tar -xz -C /tmp/
$ mv /tmp/kubebuilder_2.3.1_${os}_${arch} /usr/local/kubebuilder
$ export PATH=$PATH:/usr/local/kubebuilder/bin
```

**2. Install operator-sdk (version: [v1.0.0](https://github.com/operator-framework/operator-sdk/releases/tag/v1.0.0))**
```bash
$ release=v1.0.0
$ os=$(go env GOOS)
$ arch=$(arch)
$ curl -OJL https://github.com/operator-framework/operator-sdk/releases/download/${release}/operator-sdk-${release}-${arch}-${os}-gnu
$ chmod +x operator-sdk-${release}-${arch}-${os}-gnu
$ mv operator-sdk-${release}-${arch}-${os}-gnu /usr/local/bin/operator-sdk
```

### TEST
## TEST

