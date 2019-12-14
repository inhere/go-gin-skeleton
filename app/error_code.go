package app

// Predefined const error codes.
// message please see /resources/languages/*.ini
const (
	OK  = 0
	ERR = 2

	baseNum = 10000

	// error codes
	ErrSvrNotAvailable = baseNum + 2100
	ErrServer          = baseNum + 2101
	ErrNoPermission    = baseNum + 2102
	ErrNotFound        = baseNum + 2104
	ErrNoResource      = baseNum + 2105
	ErrParams          = baseNum + 2106
	ErrMissingParams   = baseNum + 2107
	ErrNoRecord        = baseNum + 2108
	ErrOpFail          = baseNum + 2112
	ErrInvalidReq      = baseNum + 2113
	ErrInvalidParam    = baseNum + 2114

	ErrRepeatOp    = baseNum + 2202
	ErrDatabase    = baseNum + 2210
	ErrInvalidData = baseNum + 2211
	ErrCheckFail   = baseNum + 2212

	ErrDupRows    = baseNum + 2406
	ErrInsertFail = baseNum + 2401
	ErrUpdateFail = baseNum + 2403
	ErrDeleteFail = baseNum + 2404
)
