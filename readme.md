# tfvm - Terraform version manager [![Build](https://api.travis-ci.com/cbuschka/tfvm.svg?branch=master)](https://travis-ci.com/github/cbuschka/tfvm) [![Latest Release](https://img.shields.io/github/release/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/tfvm)](https://goreportcard.com/report/github.com/cbuschka/tfvm) [![License](https://img.shields.io/github/license/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/blob/master/license.txt)

### Always the right terraform version for your project

## Features
* download and install terraform version
* select terraform version via .tfvmrc (or .terraform-version for compatibility to tfenv)
* invoke selected terraform through tfvm

## Installation
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

## Similar tools
* [tfenv](https://github.com/tfutils/tfenv) - similar to tfvm
* [renatoassis01/tfvm](https://github.com/renatoassis01/tfvm) - like this, configuration similar to nvm
* [jkehres/tfvm](https://github.com/jkehres/tfvm) - like this, no releases yet
* [terraformw](https://objectpartners.com/2017/12/21/use-a-terraform-wrapper-script-to-easily-manage-terraform-installations/) - simple solution, no releases yet
* [tfswitch](https://github.com/warrensbox/terraform-switcher) - looks good, no windows support
* [terraform-installer](https://github.com/robertpeteuil/terraform-installer)

## License
Copyright (c) 2020 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)

