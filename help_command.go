package tfvm

func printUsage() {
	Print(`
Usage:	tfvm <command>
	or terraform <terraform command and options>

A terraform version manager.

Commands:
  terraform	Run terraform configured via .tfvmrc/.terraform-version file.
		(Run 'terraform help' for more information.)
  list		List terraform versions.
  install	Install terraform version.
  which		Print selected terraform version.
  uptodate	Check for updates.
  help		Print this usage information.
  version	Print tfvm version.
  info		Print tfvm/runtime info.

For binaries, issues and source code visit https://github.com/cbuschka/tfvm.
`)
}
