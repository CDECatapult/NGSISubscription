# NGSI-import
Script to subscribe to NGSI data sources for the SynchroniCity Historical API from a list.

[![License badge](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

+ [Introduction](#def-introduction)
+ [How to Build](#def-build)
+ [Configuration](#def-conf)
+ [How to Use](#def-use)
+ [License](#def-license)

---

<br>

<a name="def-introduction"></a>
## Introduction

This tool is part of the EU H2020 [SynchroniCity](https://synchronicity-iot.eu) project and it is an add on for the [SynchroniCity IoT Data Marketplace](https://github.com/caposseleDigicat/SynchroniCityDataMarketplace).

- You will find the source code of this tool in GitHub [here](https://github.com/caposseleDigicat/NGSISubscription)

Thanks to this tool you will be able to subscribe to NGSI data sources directly on the Historical API. This tool reads from a file a list of NGSI entity types and creates the respective subscription on Orion CB. 

<a name="def-build"></a>
## How to Build

Inside the `bin` folder you will find 3 different version already compiled for `Mac OSX`, `Windows`, and `Linux`. If you wish to recompile it, just read the following instructions. 

Requirements: [Go Programming Language](https://golang.org/doc/install)

To build the binary for your `OS` and `architecture` just run:

```
GOOS=[OS] GOARCH=[ARCH] go build main.go utility.go
```

where [OS] and [ARCH] are your Operating System and Architecture respectively. Some examples for OSX, Windows and Linux (with 32bit architecture) are:

```
GOOS=darwin GOARCH=386 go build main.go utility.go

GOOS=windows GOARCH=386 go build main.go utility.go

GOOS=linux GOARCH=386 go build main.go utility.go
```

These commads will create a `main` (`main.exe` for windows) executable file. For your reference, you can find a similar set of commands inside the script `build.sh`.

<a name="def-conf"></a>
## Configuration

This tool requires a configuration file `config.json` to be filled as:

```
"orion_url": "",
"cygnus_url": "",
"fiware_service":"",
"fiware_servicepath":""
```

- `orion_url`: the url of Orion Contex Broker (e.g, http://orion-cb:1026).
- `cygnus_url`: the url of Cygnus (e.g, http://cygnus:5050).
- `fiware_service`: the Fiware-Service used on the Orion Context Broker while registering the NGSI entities (e.g, manchester). 
- `fiware_servicepath`: the Fiware-ServicePath used on the Orion Context Broker while registering the NGSI entities (e.g, /). 
As an example:

```
"orion_url": "http://<orion_ip>:1026",
"cygnus_url": "http://<cygnus_ip>:5050",
"fiware_service":"manchester",
"fiware_servicepath":"/"
```

Please note that the configuration file must be saved in the same directory of the executable.

<a name="def-use"></a>
## How to Use

Before using the tool, you need to create a file containing the list of data sources that you wish to subscribe to.
This file should be created as list of [Entity Type] [Entity ID]. For your reference, the following is an example of the content of such file:

```
AirQualityObserved
BikeHireDockingStation
CrowdFlowObserved
NoiseLevelObserved
OffStreetParking
OnStreetParking
TrafficFlowObserved
WeatherObserved
```

After creating this file, you can run the tool and pass the path of the file (e.g., ./datasourcelist.dat) as argument:

```
./bin/subscribe ./datasourcelist.dat
```

You will see an output similar to:

```
 ----------------------------------
|        NGSI Subscription         |
 ----------------------------------

The following data sources will be imported:
type1
type2 
type2 
type3 
WARNING: Are you sure? (yes/no)
```

By typing `no` the program will terminate while by typing `yes` the tool will subscribe to the respective data sources specification on the Orion CB.

## License

The MIT License
 
Copyright (C) 2018 Digital Catapult.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
