# metalsecure

[![Build Status](https://github.com/devops-metalflow/metalsecure/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/devops-metalflow/metalsecure/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/devops-metalflow/metalsecure/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/devops-metalflow/metalsecure)
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-metalflow/metalsecure)](https://goreportcard.com/report/github.com/devops-metalflow/metalsecure)
[![License](https://img.shields.io/github/license/devops-metalflow/metalsecure.svg)](https://github.com/devops-metalflow/metalsecure/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/devops-metalflow/metalsecure.svg)](https://github.com/devops-metalflow/metalsecure/tags)



## Introduction

*metalsecure* is the worker of [metalflow](https://github.com/devops-metalflow/metalflow) written in Go.



## Prerequisites

- Go >= 1.18.0



## Run

```bash
version=latest make build
./bin/metalsecure --listen-url=:19094
```



## Docker

```bash
version=latest make docker
docker run ghcr.io/devops-metalflow/metalsecure:latest --listen-url=:19094
```



## Usage

```
usage: metalsecure --listen-url=LISTEN-URL [<flags>]

metalsecure

Flags:
  --help                   Show context-sensitive help (also try --help-long and --help-man).
  --version                Show application version.
  --listen-url=LISTEN-URL  Listen URL (host:port)
```



## Protobuf

```json
{
  "apiVersion": "v1",
  "kind": "metalsecure",
  "metadata": {
    "name": "metalsecure"
  },
  "spec": {
    "name": "foo"
  }
}
```



## License

Project License can be found [here](LICENSE).



## Reference
