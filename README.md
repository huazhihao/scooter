# 🛵 Scooter: A fast and fully featured reverse proxy

![license Apache-2.0](https://img.shields.io/github/license/huazhihao/scooter)
[![Build Status](https://travis-ci.org/huazhihao/scooter.svg?branch=master)](https://travis-ci.org/huazhihao/scooter)
[![GoDoc](https://godoc.org/github.com/huazhihao/scooter?status.svg)](https://godoc.org/github.com/huazhihao/scooter)
![GoReport](https://goreportcard.com/badge/github.com/huazhihao/scooter)

`Scooter` is a lightweight L4+L7 reverse proxy and load balancer written in Go. It provides nginx-like functionalities with little effort on setup, and better integration with modern monitoring tools.

## Contents

- [Installation](#installation)
- [Quick start](#quick-start)
- [Benchmarks](#benchmarks)
- [Configuration examples](#configuration-examples)
    - [Reverse proxy with custom headers](#reverse-proxy-with-custom-headers)
    - [API gateway with version-path mapping](#api-gateway-with-version-path-mapping)
    - [Secured load balancer with weighted backends](#secured-load-balancer-with-weighted-backends)
    - [TCP proxy as a ssh relay server](#tcp-proxy-as-a-ssh-relay-server)
    - [Scooter API endpoint](#scooter-api-endpoint)
    - [Scooter prometheus endpoint](#scooter-prometheus-endpoint)
- [Migrate from nginx to scooter](#migrate-from-nginx-to-scooter)


## Installation

```sh
$ go get -u github.com/huazhihao/scooter
```

## Quick start

```sh
# assume the following content in scooter.yaml file
$ cat scooter.yaml
```

```yaml
http:
- address: ":8000"
  rules:
    - url: "http://example.com"
```

```
# run scooter and visit 0.0.0.0:8000 on browser
$ scooter --config scooter.yaml --debug
```

## Benchmarks

```
// TODO
```

## Configuration examples

### Reverse proxy with custom headers

```yaml
http:
- name: reverse-proxy
  address: ":8000"
  rules:
    - url: "http://localhost:8000"
      headers:
        - key: Host
          value: $proxy_host
        - key: X-Real-IP
          value: $client_ip
```

### API gateway with version-path mapping

```yaml
http:
- name: api-gateway
  address: ":8090"
  rules:
    - path: /
      url: "http://api-v2/"
    - path: /v1/
      url: "http://api-v1/"
```

### Secured load balancer with weighted backends

```yaml
https:
- name: load-balancer
  address: ":8443"
  tls:
    cert: ./test.pem
    key: ./test-key.pem
  rules:
    - url: "http://127.0.0.1:8001/"
      weight: 1
    - url: "http://127.0.0.1:8002/"
      weight: 10
```

### TCP proxy as a ssh relay server

```yaml
tcp:
- name: tcp-relay
  protocol: tcp
  address: ":2022"
  remote: "127.0.0.1:22"
```

### Scooter API endpoint

```yaml
api:
  address: ":9000"
```

### Scooter prometheus endpoint

```yaml
metrics:
  prometheus:
    address: ":8081"
```

## Migrate from nginx to scooter

Find more examples at [Migrate from nginx to scooter](https://github.com/huazhihao/scooter/blob/master/migrate-from-nginx.md)

