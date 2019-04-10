package userx

import (
	"fmt"
	"os/exec"
	"os/user"
)

// DeleteUser ..
type DeleteUser struct {
	Name               string
	Force              bool
	DeleteEmailHomeDir bool
}

// Delete ..
func (u DeleteUser) Delete() error {
	_, err := user.Lookup(u.Name)
	if err != nil {
		return fmt.Errorf("user %s is not exists", u.Name)
	}

	cmd := exec.Command("userdel")

	if u.Force {
		cmd.Args = append(cmd.Args, "-f")
	}

	if u.DeleteEmailHomeDir {
		cmd.Args = append(cmd.Args, "-r")
	}

	cmd.Args = append(cmd.Args, u.Name)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
