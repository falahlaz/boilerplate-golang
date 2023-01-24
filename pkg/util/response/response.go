package response

import "github.com/falahlaz/boilerplate-golang/internal/abstraction"

type Meta struct {
	Success bool                        `json:"success" default:"true"`
	Message string                      `json:"message" default:"OK"`
	Info    *abstraction.PaginationInfo `json:"info"`
}
