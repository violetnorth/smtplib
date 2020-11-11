# smtplib

### Status Codes
```
StatusSystemStatus            = 211
StatusSystemHelp              = 214
StatusServiceReady            = 220
StatusServiceClosing          = 221
StatusActionCompleted         = 250
StatusUserNotLocalWillForward = 251

StatusStartMailInput = 354

StatusServiceNotAvailable              = 421
StatusActionNotTakenMailboxUnavailable = 450
StatusActionAbortedLocalError          = 451
StatusActionNotTakenInsufficentStorage = 452

StatusSyntaxErrorCommand                  = 500
StatusSyntaxErrorParameters               = 501
StatusCommandNotImplemented               = 502
StatusCommandBadSequence                  = 503
StatusCommandParameterNotImplemented      = 504
StatusActionNotTakenMailboxInaccessible   = 550
StatusUserNotLocal                        = 551
StatusActionNotTakenMailboxNameNotAllowed = 553
StatusTransactionFailed                   = 554
```