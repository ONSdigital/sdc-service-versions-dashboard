[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d5d832c468cc4307905ca7f3c9d84d67)](https://www.codacy.com/app/sdcplatform/sdc-service-versions-dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ONSdigital/sdc-service-versions-dashboard&amp;utm_campaign=Badge_Grade)

# SDC Microservice Versions Dashboard
This repository contains a dashboard application implemented using [Go](https://golang.org/) that shows which versions of the Survey Data Collection (SDC) platform microservices are deployed to which environment, using information held by the [SDC Microservice Versions Git repository](https://github.com/ONSdigital/sdc-service-versions).

## Prerequisites
* Install the [Godep](https://github.com/tools/godep) package manager using `go get github.com/tools/godep`
* Run `godep get` to download and install the other dependencies managed by Godep

## Building
Install Go and ensure your `GOPATH` environment variable is set (usually it's `~/go`).

### Make
A Makefile is provided for compiling the code:

```
make
```

The compiled executable is placed within the `build` directory tree.

## Running
First compile the code using `make` then execute the binary in the background using `./versions-dashboard &` from within the `bin` directory within the `build` directory tree.

The following environment variable may be overridden:

| Environment Variable | Purpose            | Default Value  |
| :------------------- | :----------------- | :------------- |
| PORT                 | HTTP listener port | :8080          |

## Cleaning
To clobber the `build` directory tree that's created when running `make`, run:

```
make clean
```

## Copyright
Copyright (C) 2017 Crown Copyright (Office for National Statistics)