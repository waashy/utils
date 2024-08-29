package errors

import "errors"

var (
	ErrDbConn = errors.New("failed to connect to database")
)
