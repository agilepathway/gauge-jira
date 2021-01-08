Gauge-Jira
==========

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

Publishes Gauge specifications to Jira. This is a plugin for [gauge](https://gauge.org/).

Installation
------------

```
gauge install jira
```
To install a specific version of jira plugin use the ``--version`` flag.

```
gauge install jira --version $VERSION
```

### Offline Installation

Download the plugin zip from the [Github Releases](https://github.com/agilepathway/gauge-jira/releases).

use the ``--file`` or ``-f`` flag to install the plugin from  zip file.

```
gauge install jira --file ZIP_FILE_PATH
```

### Build from Source

#### Requirements
* [Golang](http://golang.org/)

#### Compiling

```
go run build/make.go
```

For cross-platform compilation

```
go run build/make.go --all-platforms
```

#### Installing
After compilation

```
go run build/make.go --install
```

### Creating distributable

Note: Run after compiling

```
go run build/make.go --distro
```

For distributable across platforms: Windows and Linux for both x86 and x86_64

```
go run build/make.go --distro --all-platforms
```

Contributing
------------

See the [CONTRIBUTING.md](./CONTRIBUTING.md)

License
-------

`Gauge-Jira` is released under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text.
