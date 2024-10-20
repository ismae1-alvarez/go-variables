package users

import (
	"fmt"
	"time"
	"udemy-1/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(10, "Pablo", time.Now(), true)

	fmt.Println(u)
}
