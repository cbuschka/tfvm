package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/util"
)

func printUsage() {
	buildInfo := build.GetBuildInfo()

	util.Print(`
Usage:	tfvm <command>
	or terraform <terraform command and options>
	or tfenv <tfenv command and options>

THE terraform version manager.

Commands:
  terraform	Run terraform selected via .terraform-version file.
		(Run 'terraform help' for more information.)
  list		List terraform versions.
  install	Install terraform version.
  which		Print selected terraform version.
  use		Change selected version in .terraform-version file.
  uptodate	Check for updates.
  help		Print this usage information.
  version	Print tfvm version.
  info		Print tfvm/runtime info.

For binaries, issues and source code visit %s.
`, buildInfo.Repo)
}
