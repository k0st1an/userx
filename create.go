package mngus

import (
	"fmt"
	"os/exec"
	"os/user"
)

// CreateUser ..
type CreateUser struct {
	Name          string
	CreateHomeDir bool
}

// Create ..
func (u CreateUser) Create() error {
	_, err := user.Lookup(u.Name)
	if err == nil {
		return fmt.Errorf("user %s already exists", u.Name)
	}

	cmd := exec.Command("useradd")

	if u.CreateHomeDir {
		cmd.Args = append(cmd.Args, "-m")
	}

	cmd.Args = append(cmd.Args, u.Name)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
