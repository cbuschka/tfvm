# tfvm vs. tfenv vs. terraform-switcher

| feature          | tfvm | tfenv | terraform-switcher | 
|------------------|------|-------|--------------------|
| cool logo        | no | no | [cool logo](https://github.com/warrensbox/terraform-switcher#terraform-switcher) |
| platform support | [win32/64, linux32/64, linux/arm64, mac32/64, mac/arm64](https://github.com/cbuschka/tfvm/releases/latest) | [win64, linux64/arm, mac64](https://github.com/tfutils/tfenv#support) | [linux32/64, mac32/64, win64](https://github.com/warrensbox/terraform-switcher#installation) |
| installation     | [download binary](https://github.com/cbuschka/tfvm/releases/latest) or [installer script](https://github.com/cbuschka/tfvm/blob/main/readme.md#installation) | bunch of files share, libexec | [magic bash script](https://github.com/warrensbox/terraform-switcher#general-linux), [brew](https://github.com/warrensbox/terraform-switcher#homebrew), [snap](https://github.com/warrensbox/terraform-switcher#snapcraft-for-centos-ubuntu-linux-mint-rhel-debian-fedora) |
| update           | [download binary](https://github.com/cbuschka/tfvm/releases/latest) or [installer script](https://github.com/cbuschka/tfvm/blob/main/readme.md#installation) | [update via brew](https://github.com/tfutils/tfenv#automatic) or [update git repo](https://github.com/tfutils/tfenv#upgrading) | see installation |
| configuration    | put anywhere in your path | [add ~/.tfenv/bin/ to your path](https://github.com/tfutils/tfenv#manual) | put anywhere in your path |
| home brew package | no | [yes](https://github.com/tfutils/tfenv#automatic) | [yes](https://github.com/warrensbox/terraform-switcher#homebrew) |
| terraform auto install  | yes | [yes](https://github.com/tfutils/tfenv#tfenv_auto_install) | yes |
| .terraform-version file | [.terraform-version in current dir or up the tree](https://github.com/cbuschka/tfvm#configure-terraform-version) | [.terraform-version in current dir or up the tree](https://github.com/tfutils/tfenv#terraform-version) | [.terraform-version in current dir](https://github.com/warrensbox/terraform-switcher#instead-of-a-tfswitchrc-file-a-terraform-version-file-may-be-used-for-compatibility-with-tfenv-and-other-tools-which-use-it), [.tfswitch.toml](https://github.com/warrensbox/terraform-switcher#use-tfswitchtoml-file--for-non-admin---users-with-limited-privilege-on-their-computers), [.tfswitchrc in current dir](https://github.com/warrensbox/terraform-switcher#use-tfswitchrc-file), [version.tf in current dir](https://github.com/warrensbox/terraform-switcher#use-versiontf-file) |
| terraform binary checksum verification | no | [yes](https://github.com/tfutils/tfenv#tfenv-install-version) | no |
| version ranges   | [yes](https://github.com/cbuschka/tfvm#configure-terraform-version) | [no](https://github.com/tfutils/tfenv#min-required) | no |
| override hashicorp repo | [yes](https://github.com/cbuschka/tfvm/blob/main/doc/env-vars.md#environment-variables) | [yes](https://github.com/tfutils/tfenv#tfenv_remote) | no |
| override terraform arch | [yes](https://github.com/cbuschka/tfvm/blob/main/doc/env-vars.md#environment-variables) | [yes](https://github.com/tfutils/tfenv#tfenv_arch) | no |
| override terraform os | [yes](https://github.com/cbuschka/tfvm/blob/main/doc/env-vars.md#environment-variables) | [yes](https://github.com/tfutils/tfenv#tfenv_arch) | no |
| compliant with xdg dir structure | yes | no (stores all below ~/.tfenv/bin/) | [no](https://github.com/warrensbox/terraform-switcher/issues/80) |
| github actions support | [yes](https://github.com/cbuschka/setup-tfvm) | no, maual setup required | no, manual setup required |
| darwin/arm64 with fallback to darwin/amd64 | [yes](https://github.com/cbuschka/tfvm/issues/25) | no | no |

