package models

import "github.com/siongleng89/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
