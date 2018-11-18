package revolut

// Revolut errors
const (
	errBadRequest   = "bad request - check syntax"
	errUnauthorized = "not authorized - check the API key"
	errForbidden    = "resource or action can't be accessed with supplied key"
	errNotFound     = "unknown resource - check spelling"
	errDisallowed   = "you tried to access an endpoint with an invalid method"
	errUnacceptable = "you requested a format that isn't JSON"
	errHammer       = "you're sending too many requests too quickly"
	errInternal     = "internal server error - try again later"
	errUnavailable  = "service unavailable - offline for maintenance"
)

// Internal errors
const (
	// ErrKeyFormat means the API key string is mangled.
	ErrKeyFormat = "API key has the wrong format - not starting with sand_ or prod_"
)

func codeToError(code int) string {
	var msg = ""
	switch code {
	case 400:
		msg = errBadRequest
	case 401:
		msg = errUnauthorized
	case 403:
		msg = errForbidden
	case 404:
		msg = errNotFound
	case 405:
		msg = errDisallowed
	case 406:
		msg = errUnacceptable
	case 429:
		msg = errHammer
	case 500:
		msg = errInternal
	case 501:
		msg = errUnavailable
	}
	return msg
}
