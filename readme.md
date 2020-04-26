# tfvm - Terraform version manager [![Build](https://api.travis-ci.com/cbuschka/tfvm.svg?branch=master)](https://travis-ci.com/github/cbuschka/tfvm) [![Latest Release](https://img.shields.io/github/release/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/tfvm)](https://goreportcard.com/report/github.com/cbuschka/tfvm) [![License](https://img.shields.io/github/license/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/blob/master/license.txt)

### Always the right terraform version for your project

## Features
* configure terraform via .tfvmrc
* invoke configured terraform through tfvm 
* download and install terraform version

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

### Invoke configured terraform
```
tfvm terraform version
```

### Tip: Call configured terraform directly
```
ln -s tfvm terraform

terraform version
```

### More commands via tfvm usage info
```
tfvm help
```

## Similar tools
* [tfenv](https://github.com/tfutils/tfenv) - similar to tfvm and more mature than this (2020-04-26)
* [renatoassis01/tfvm](https://github.com/renatoassis01/tfvm) - like this, configuration similar to nvm
* [jkehres/tfvm](https://github.com/jkehres/tfvm) - like this, no releases yet
* [terraformw](https://objectpartners.com/2017/12/21/use-a-terraform-wrapper-script-to-easily-manage-terraform-installations/) - simple solution, no releases yet

## License
Copyright (c) 2020 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)

