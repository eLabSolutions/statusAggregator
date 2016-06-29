# Introduction

This app was developed to solve a simple problem. At AWS the Load Balancer Health
Check allows you to perform simple health checks directed at a single port.
However, with Docker (or even without) it is common to run many services
on a single node.

We wanted to take a node out of commission if any of a number of essential
services failed on that node. This status aggregator allows us to do this
with a very simple, very lightweight app that can monitor local services
and present an aggregated health status to the outside world (e.g the
AWS Health Check).

# Notes
Set the SACONFIGFILE environment variable to point to the config file (see example.json).

Will return 500 when the sites have not yet been checked, 503 if any of them timeout, otherwise
it will return the highest status code.

# Docker

See [DockerHub page](https://hub.docker.com/r/elab/statusaggregator/)

# Use
```
go get
go build
SACONFIGFILE=./example.json ./statusAggregator
```

And in another terminal you can test it:
```
curl -i localhost:18888

HTTP/1.1 500 Internal Server Error
Date: Wed, 29 Jun 2016 14:27:19 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

... wait a few seconds for sites to be checked ...

curl -i localhost:18888

HTTP/1.1 200 OK
Date: Wed, 29 Jun 2016 14:27:23 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

# Config Options
JSON configuration file with the following options:

| Name | Description |
| ------ |-----|
| Port | Port to listen on |
| CheckFrequency | Frequency (in seconds) to check each site |
| Timeout | Maximum time to wait for a site check |
| Sites | Array of URL strings |
| JsonLogs | Whether or not the logs should output as JSON |

## Example

```json
{
  "Port": 18888,
  "CheckFrequency": 10,
  "Timeout": 3,
  "Sites": [
    "https://www.google.com",
    "https://www.yahoo.com"
  ],
  "JsonLogs": false
}
```

# License
GPL-3; see LICENSE file.

Copyright 2016, MARC Technologies, LLC., a subsidiary of eLab® Solutions Corporation.