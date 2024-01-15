# check-debian-reboot-required

## Description

Mackerel check plugin to detect being required reboot on Debian

## Synopsis

```sh
check-debian-reboot-required
```

## Installation

```sh
sudo mkr plugin install Arthur1/check-debian-reboot-required@x.y.z
```

## Setting for mackerel-agent

```
[plugin.checks.check-debian-reboot-required]
command = ["check-debian-reboot-required"]
```

## Usage

### Options

```
  -critical
        create critical check report when reboot required
  -dir string
        directory of reboot-required file [for debug] (default "/var/run/")
  -warning
        create warning check report when reboot required (default true)
```
