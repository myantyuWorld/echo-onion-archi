package auth

import (
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

func IsAdmin(p model.Profile) bool {
	return p.Role == model.Admin
}
