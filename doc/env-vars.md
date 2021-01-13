# Environment Variables

| Variable         | Secondary Variable | Meaning        | Default | 
|------------------|--------------------|----------------|---------|
| TFVM_DIR | n/a | tfvm inventory dir | ${HOME}/.tfvm on OS X and windows, ${HOME}/.cache/tfvm/ on Linux |
| TFVM_TERRAFORM_VERSION | TERRAFORM_VERSION | Selected terraform version | |
| TFVM_TERRAFORM_ARCH | TERRAFORM_ARCH | Arch of selected terraform version: amd64, 386 | Platform arch |
| TFVM_TERRAFORM_OS | TERRAFORM_OS | OS of selected terraform version: windows, darwin, linux | Platform os |
| TFVM_TERRAFORM_RELEASES_BASE_URL | n/a | Base URL for downloading terraform binaries | https://releases.hashicorp.com/terraform |
| XDG_CACHE_HOME | n/a | Base dir for inventory dir | $HOME/.cache/tfvm |

