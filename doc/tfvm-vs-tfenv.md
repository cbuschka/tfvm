# tfvm vs. tfenv

| feature          | tfvm | tfenv | 
|------------------|------|-------|
| platform support | [win32/64, linux32/64, mac32/64](https://github.com/cbuschka/tfvm/releases/tag/v0.13.0) | [win64, linux64/arm, mac64](https://github.com/tfutils/tfenv#support) |
| installation     | [single binary](https://github.com/cbuschka/tfvm/releases/latest), [optionally magic bash script](https://github.com/cbuschka/tfvm/tree/conni_tfvm-vs-tfenv#installation) | bunch of files share, libexec |
| update           | download latest binary | [update via brew](https://github.com/tfutils/tfenv#automatic) or [update git repo](https://github.com/tfutils/tfenv#upgrading) |
| configuration    | put anywhere in your path | [add ~/.tfenv/bin/ to your path](https://github.com/tfutils/tfenv#manual) | 
| home brew package | no | [yes](https://github.com/tfutils/tfenv#automatic) |
| terraform auto install  | yes | [yes](https://github.com/tfutils/tfenv#tfenv_auto_install) |
| .terraform-version file | yes | [yes](https://github.com/tfutils/tfenv#terraform-version) |
| .tfvmrc file | yes | no |
| terraform binary checksum verification | no | [yes](https://github.com/tfutils/tfenv#tfenv-install-version) |
| version ranges   | yes | [no](https://github.com/tfutils/tfenv#min-required) |
| override hashicorp repo | yes | [yes](https://github.com/tfutils/tfenv#tfenv_remote) |
| override terraform arch | no | [yes](https://github.com/tfutils/tfenv#tfenv_arch) |

