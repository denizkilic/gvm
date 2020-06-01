package commands

import (
	"github.com/olimpias/gvm/internal/filesystem"
)

type PackageUser interface {
	UseGoPackage(version string) error
}

type UseCommand struct {
	packageUser PackageUser
	version     string
}

func NewUseCommand(fileManager PackageUser, version string) *UseCommand {
	return &UseCommand{packageUser: fileManager, version: version}
}

func (u *UseCommand) Validate() error {
	return filesystem.ValidateOperation()
}

func (u *UseCommand) Apply() error {
	return u.packageUser.UseGoPackage(u.version)
}
