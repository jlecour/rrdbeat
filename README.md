# Rrdbeat

Welcome to Rrdbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/jlecour`

## Getting Started with Rrdbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.6
* [Glide](https://github.com/Masterminds/glide) >= 0.10.0

### Init Project
To get running with Rrdbeat, run the following command:

```
make init
```

To commit the first version before you modify it, run:

```
make commit
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Rrdbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/jlecour/rrdbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Rrdbeat run the command below. This will generate a binary
in the same directory with the name rrdbeat.

```
make
```


### Run

To run Rrdbeat with debugging output enabled, run:

```
./rrdbeat -c rrdbeat.yml -e -d "*"
```


### Test

To test Rrdbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Package

To be able to package Rrdbeat the requirements are as follows:

 * [Docker Environment](https://docs.docker.com/engine/installation/) >= 1.10
 * $GOPATH/bin must be part of $PATH: `export PATH=${PATH}:${GOPATH}/bin`

To cross-compile and package Rrdbeat for all supported platforms, run the following commands:

```
cd dev-tools/packer
make deps
make images
make
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/rrdbeat.template.json and etc/rrdbeat.asciidoc

```
make update
```


### Cleanup

To clean  Rrdbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Rrdbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/jlecour
cd ${GOPATH}/github.com/jlecour
git clone https://github.com/jlecour/rrdbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
