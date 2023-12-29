package custom_errs

import "github.com/pkg/errors"

var DBErrCreateWithID = errors.New("create with object id not allowed")
var DBErrIDConversion = errors.New("id conversion error")
var ServerError = errors.New("server error")
var ParamError = errors.New("param error")
