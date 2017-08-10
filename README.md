[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d5d832c468cc4307905ca7f3c9d84d67)](https://www.codacy.com/app/sdcplatform/sdc-service-versions-dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ONSdigital/sdc-service-versions-dashboard&amp;utm_campaign=Badge_Grade)

# SDC Microservice Versions Dashboard
This repository contains a dashboard application implemented using [Go](https://golang.org/) that shows which versions of the Survey Data Collection (SDC) platform microservices are deployed to which environment, using information held by the [SDC Microservice Versions Git repository](https://github.com/ONSdigital/sdc-service-versions).

## Prerequisites
* Install the [Godep](https://github.com/tools/godep) package manager using `go get github.com/tools/godep`
* Run `godep get` to download and install the other dependencies managed by Godep

## Copyright
Copyright (C) 2017 Crown Copyright (Office for National Statistics)