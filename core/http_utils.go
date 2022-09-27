package core

type HttpStatusCode int

const (
	HttpStatusContinue                      HttpStatusCode = 100 // RFC 9110, 15.2.1
	HttpStatusSwitchingProtocols            HttpStatusCode = 101 // RFC 9110, 15.2.2
	HttpStatusProcessing                    HttpStatusCode = 102 // RFC 2518, 10.1
	HttpStatusEarlyHints                    HttpStatusCode = 103 // RFC 8297
	HttpStatusOK                            HttpStatusCode = 200 // RFC 9110, 15.3.1
	HttpStatusCreated                       HttpStatusCode = 201 // RFC 9110, 15.3.2
	HttpStatusAccepted                      HttpStatusCode = 202 // RFC 9110, 15.3.3
	HttpStatusNonAuthoritativeInfo          HttpStatusCode = 203 // RFC 9110, 15.3.4
	HttpStatusNoContent                     HttpStatusCode = 204 // RFC 9110, 15.3.5
	HttpStatusResetContent                  HttpStatusCode = 205 // RFC 9110, 15.3.6
	HttpStatusPartialContent                HttpStatusCode = 206 // RFC 9110, 15.3.7
	HttpStatusMultiStatus                   HttpStatusCode = 207 // RFC 4918, 11.1
	HttpStatusAlreadyReported               HttpStatusCode = 208 // RFC 5842, 7.1
	HttpStatusIMUsed                        HttpStatusCode = 226 // RFC 3229, 10.4.1
	HttpStatusMultipleChoices               HttpStatusCode = 300 // RFC 9110, 15.4.1
	HttpStatusMovedPermanently              HttpStatusCode = 301 // RFC 9110, 15.4.2
	HttpStatusFound                         HttpStatusCode = 302 // RFC 9110, 15.4.3
	HttpStatusSeeOther                      HttpStatusCode = 303 // RFC 9110, 15.4.4
	HttpStatusNotModified                   HttpStatusCode = 304 // RFC 9110, 15.4.5
	HttpStatusUseProxy                      HttpStatusCode = 305 // RFC 9110, 15.4.6
	HttpStatusTemporaryRedirect             HttpStatusCode = 307 // RFC 9110, 15.4.8
	HttpStatusPermanentRedirect             HttpStatusCode = 308 // RFC 9110, 15.4.9
	HttpStatusBadRequest                    HttpStatusCode = 400 // RFC 9110, 15.5.1
	HttpStatusUnauthorized                  HttpStatusCode = 401 // RFC 9110, 15.5.2
	HttpStatusPaymentRequired               HttpStatusCode = 402 // RFC 9110, 15.5.3
	HttpStatusForbidden                     HttpStatusCode = 403 // RFC 9110, 15.5.4
	HttpStatusNotFound                      HttpStatusCode = 404 // RFC 9110, 15.5.5
	HttpStatusMethodNotAllowed              HttpStatusCode = 405 // RFC 9110, 15.5.6
	HttpStatusNotAcceptable                 HttpStatusCode = 406 // RFC 9110, 15.5.7
	HttpStatusProxyAuthRequired             HttpStatusCode = 407 // RFC 9110, 15.5.8
	HttpStatusRequestTimeout                HttpStatusCode = 408 // RFC 9110, 15.5.9
	HttpStatusConflict                      HttpStatusCode = 409 // RFC 9110, 15.5.10
	HttpStatusGone                          HttpStatusCode = 410 // RFC 9110, 15.5.11
	HttpStatusLengthRequired                HttpStatusCode = 411 // RFC 9110, 15.5.12
	HttpStatusPreconditionFailed            HttpStatusCode = 412 // RFC 9110, 15.5.13
	HttpStatusRequestEntityTooLarge         HttpStatusCode = 413 // RFC 9110, 15.5.14
	HttpStatusRequestURITooLong             HttpStatusCode = 414 // RFC 9110, 15.5.15
	HttpStatusUnsupportedMediaType          HttpStatusCode = 415 // RFC 9110, 15.5.16
	HttpStatusRequestedRangeNotSatisfiable  HttpStatusCode = 416 // RFC 9110, 15.5.17
	HttpStatusExpectationFailed             HttpStatusCode = 417 // RFC 9110, 15.5.18
	HttpStatusTeapot                        HttpStatusCode = 418 // RFC 9110, 15.5.19 (Unused)
	HttpStatusMisdirectedRequest            HttpStatusCode = 421 // RFC 9110, 15.5.20
	HttpStatusUnprocessableEntity           HttpStatusCode = 422 // RFC 9110, 15.5.21
	HttpStatusLocked                        HttpStatusCode = 423 // RFC 4918, 11.3
	HttpStatusFailedDependency              HttpStatusCode = 424 // RFC 4918, 11.4
	HttpStatusTooEarly                      HttpStatusCode = 425 // RFC 8470, 5.2.
	HttpStatusUpgradeRequired               HttpStatusCode = 426 // RFC 9110, 15.5.22
	HttpStatusPreconditionRequired          HttpStatusCode = 428 // RFC 6585, 3
	HttpStatusTooManyRequests               HttpStatusCode = 429 // RFC 6585, 4
	HttpStatusRequestHeaderFieldsTooLarge   HttpStatusCode = 431 // RFC 6585, 5
	HttpStatusUnavailableForLegalReasons    HttpStatusCode = 451 // RFC 7725, 3
	HttpStatusInternalServerError           HttpStatusCode = 500 // RFC 9110, 15.6.1
	HttpStatusNotImplemented                HttpStatusCode = 501 // RFC 9110, 15.6.2
	HttpStatusBadGateway                    HttpStatusCode = 502 // RFC 9110, 15.6.3
	HttpStatusServiceUnavailable            HttpStatusCode = 503 // RFC 9110, 15.6.4
	HttpStatusGatewayTimeout                HttpStatusCode = 504 // RFC 9110, 15.6.5
	HttpStatusHTTPVersionNotSupported       HttpStatusCode = 505 // RFC 9110, 15.6.6
	HttpStatusVariantAlsoNegotiates         HttpStatusCode = 506 // RFC 2295, 8.1
	HttpStatusInsufficientStorage           HttpStatusCode = 507 // RFC 4918, 11.5
	HttpStatusLoopDetected                  HttpStatusCode = 508 // RFC 5842, 7.2
	HttpStatusNotExtended                   HttpStatusCode = 510 // RFC 2774, 7
	HttpStatusNetworkAuthenticationRequired HttpStatusCode = 511 // RFC 6585, 6
)
