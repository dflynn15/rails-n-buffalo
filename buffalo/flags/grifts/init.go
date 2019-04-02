package grifts

import (
	"github.com/dflynn15/rails-n-buffalo/buffalo/flags/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
