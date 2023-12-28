package custom_errs

import "github.com/pkg/errors"

var DBErrCreateWithID = errors.New("create with object id not allowed")
var DBErrIDConversion = errors.New("id conversion error")
