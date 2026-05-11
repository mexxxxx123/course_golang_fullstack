package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func PrintErrors(errors validate.Errors) string {
	var res string
	for _, v := range errors.Errors {
		res = res + strings.Join(v, ", ") + "\n"
	}
	return res
}
