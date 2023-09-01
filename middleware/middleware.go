package middleware

import (
    "myapp/data"

    "github.com/folkersj/celeritas"
)

type Middleware struct {
    App    *celeritas.Celeritas
    Models data.Models
}
