# tfvm - Terraform version manager [![build](https://api.travis-ci.com/cbuschka/tfvm.svg?branch=master)](https://travis-ci.com/github/cbuschka/tfvm) [![latest release](https://img.shields.io/github/tag/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/releases) [![go report](https://goreportcard.com/badge/github.com/cbuschka/tfvm)](https://goreportcard.com/report/github.com/cbuschka/tfvm) [![license](https://img.shields.io/github/license/cbuschka/tfvm.svg)](https://github.com/cbuschka/tfvm/blob/master/license.txt)

### Configure required terraform version via adding .tfvmrc file to your project

## Features
* configure terraform via .tfvmrc
* invoke configured terraform through tfvm 
* download and install terraform version

## TODOs
* create release (linux, windows, osx)
* add shell completion support

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

## License
Copyright (c) 2020 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)

