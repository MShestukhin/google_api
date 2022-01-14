package google_api

const (
	ServiceName         = "google_api"
	CompleteServiceName = "ai.somin.service." + ServiceName

	GoogleErrorCode = "google_api_error" // Micro error has instagram api error details

	TokenExpiredErrorCode  = 401 // Token expired or revoked
	UnknownGoogleErrorCode = 501
	UnknownErrorCode       = 500
)
