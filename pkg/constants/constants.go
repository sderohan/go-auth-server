package constants

// STATUS CODES
const (
	BAD_REQUEST           = 400
	INTERNAL_SERVER_ERROR = 500
	SUCCESS               = 200
)

// MESSAGES
// VIPHER
const CONFIG_KEY_NOT_SET = "Key: %s is not set"

// HTTP HELPER
const REQUEST_ERROR = "error occured while processing the request, please retry!"

// HTTP DATA CONTENT
const (
	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json"
)
