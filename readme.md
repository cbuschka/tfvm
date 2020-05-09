# tfvm - Terraform version manager [![Build](https://api.travis-ci.com/cbuschka/tfvm.svg?branch=master)](https://travis-ci.com/github/cbuschka/tfvm) [![Latest Release](https://img.shields.io/github/release/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/tfvm)](https://goreportcard.com/report/github.com/cbuschka/tfvm) [![License](https://img.shields.io/github/license/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/blob/master/license.txt)

### Always the right terraform version for your project

## Features
* download and install terraform version
* select terraform version via .tfvmrc (or .terraform-version for compatibility to tfenv)
* invoke selected terraform through tfvm

## System Requirements
* 64-Bit Linux, MacOS X or Windows

## Installation
```
curl -sL https://raw.githubusercontent.com/cbuschka/tfvm/master/install.sh -o - | bash
```

Installs into ${HOME}/bin or /usr/local/bin if executed as root.

### or
* Download tfvm from [latest release](https://github.com/cbuschka/tfvm/releases/latest) to your bin dir
* Make executable ```chmod o+x tfvm```
* Optionally link to tfvm as terraform ```ln -s tfvm terraform```

## Usage

### List terraform versions
```
tfvm list
```

### Configure terraform version
```
echo "0.12.8" > .tfvmrc

tfvm which
```

"latest" and [semver ranges](https://github.com/hashicorp/go-version#version-constraints), e.g. ```>= 0.12.1, <0.12.10```, are also supported.

### Invoke selected terraform
```
tfvm terraform version
```

### Tip: Call selected terraform directly
```
ln -s tfvm terraform

terraform version
```

### More commands via tfvm usage info
```
tfvm help
```

### Environment Variables
[Supported Environment Variables](./doc/env-vars.md)

## Similar tools
* [tfenv](https://github.com/tfutils/tfenv) [(comparison)](./doc/tfvm-vs-tfenv.md)
* [renatoassis01/tfvm](https://github.com/renatoassis01/tfvm)
* [jkehres/tfvm](https://github.com/jkehres/tfvm)
* [terraformw](https://objectpartners.com/2017/12/21/use-a-terraform-wrapper-script-to-easily-manage-terraform-installations/)
* [tfswitch](https://github.com/warrensbox/terraform-switcher)
* [terraform-installer](https://github.com/robertpeteuil/terraform-installer)
* [tfw](https://github.com/stormbeta/tfw)

## License
Copyright (c) 2020 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)

