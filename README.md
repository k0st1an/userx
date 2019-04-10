# UserX

## Usage

> Tested in Ubuntu/Debian

```golang
package main

import (
	"log"

	"github.com/k0st1an/userx"
)

func main() {
	userc := userx.CreateUser{Name: "Test", CreateHomeDir: true}

	if err := userc.Create(); err != nil {
		log.Fatalln(err)
	}

	userd := userx.DeleteUser{Name: "Test", DeleteEmailHomeDir: true}

	if err := userd.Delete(); err != nil {
		log.Fatalln(err)
	}
}
```
