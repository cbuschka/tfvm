# tfvm - Terraform version manager [![Build](https://github.com/cbuschka/tfvm/workflows/build/badge.svg)](https://github.com/cbuschka/tfvm) [![Latest Release](https://img.shields.io/github/release/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/tfvm)](https://goreportcard.com/report/github.com/cbuschka/tfvm) [![codecov](https://codecov.io/gh/cbuschka/tfvm/branch/master/graph/badge.svg)](https://codecov.io/gh/cbuschka/tfvm) [![License](https://img.shields.io/github/license/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/blob/master/license.txt)

### Always the right terraform version for your project

## Features
* download and install terraform version
* select terraform version via .terraform-version
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
echo "0.12.8" > .terraform-version

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

## Contributing
Contributions are welcome!

**By contributing you accept to contribute under the [license](./license.txt) of this project.**

* [Fork](https://github.com/cbuschka/tfvm/fork) and clone the repository
* Create a feature branch (git checkout -b foobar)
* Commit your changes (git commit -am 'Add feature foobar')
* Make sure the project builds and all tests are passing
* Push to the branch (git push origin foobar)
* Create a new pull request and describe what you want to achieve or fix

## License
Copyright (c) 2020 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)

