// Package smtplib provides SMTP status codes and the enhanced status codes
// as constants like the HTTP status codes available in http package.
package smtplib

// SMTP status codes outlined in RFC821.
// See: https://tools.ietf.org/html/rfc821
const (
	StatusSystemStatus            = 211 // 211 System status, or system help reply
	StatusSystemHelp              = 214 // 214 Help message
	StatusServiceReady            = 220 // 220 <domain> Service ready
	StatusServiceClosing          = 221 // 221 <domain> Service closing transmission channel
	StatusAuthenticationSuccess   = 235 // 235 Authentication succeeded
	StatusActionCompleted         = 250 // 250 Requested mail action okay, completed
	StatusUserNotLocalWillForward = 251 // 251 User not local; will forward to <forward-path>
	StatusCannotVerifyUser        = 252 // 252 Cannot verify the user, but it will try to deliver the message anyway

	StatusStartMailInput = 354 // 354 Start mail input; end with <CRLF>.<CRLF>

	StatusServiceNotAvailable              = 421 // 421 <domain> Service not available, closing transmission channel
	StatusPasswordTransitionNeeded         = 432 // 432 A password transition is needed
	StatusActionNotTakenMailboxUnavailable = 450 // 450 Requested mail action not taken: mailbox unavailable
	StatusActionAbortedLocalError          = 451 // 451 Requested action aborted: local error in processing
	StatusActionNotTakenInsufficentStorage = 452 // 452 Requested action not taken: insufficient system storage
	StatusAuthenticationTemporaryFailured  = 454 //	454 Temporary authentication failure
	StatusUnableAccommodateParameters      = 455 //	455 Server unable to accommodate parameters

	StatusSyntaxErrorCommand                   = 500 // 500 Syntax error, command unrecognized
	StatusSyntaxErrorParameters                = 501 // 501 Syntax error in parameters or arguments
	StatusCommandNotImplemented                = 502 // 502 Command not implemented
	StatusCommandBadSequence                   = 503 // 503 Bad sequence of commands
	StatusCommandParameterNotImplemented       = 504 // 504 Command parameter not implemented
	StatusServerDoesNotAcceptMail              = 521 // 521 Server does not accept mail
	StatusEncryptionNeeded                     = 523 // 523 Encryption Needed
	StatusAuthenticationRequired               = 530 // 530 Authentication required
	StatusAuthenticationInvalid                = 535 // 535 Authentication credentials invalid
	StatusActionNotTakenMailboxInaccessible    = 550 // 550 Requested action not taken: mailbox unavailable
	StatusUserNotLocal                         = 551 // 551 User not local; please try <forward-path>
	StatusActionAbortedExceedStorageAllocation = 552 // 552 Requested mail action aborted: exceeded storage allocation
	StatusActionNotTakenMailboxNameNotAllowed  = 553 // 553 Requested action not taken: mailbox name not allowed
	StatusTransactionFailed                    = 554 // 554 Transaction failed
	StatusMessageTooBig                        = 554 // 554 Message too big for system
	StatusDomainDoesNotAcceptMail              = 556 // 556 Domain does not accept mail
)
