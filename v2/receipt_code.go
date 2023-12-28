package v2

type ReceiptCode int

const (
	MessageCorrectlyTransmitted ReceiptCode = 100
	// General errors
	InvalidEnvelopeSyntax ReceiptCode = 200
	DuplicateMessageID    ReceiptCode = 201
	NoPayloadFound        ReceiptCode = 202
	MessageTooOldToSend   ReceiptCode = 203
	MessageExpired        ReceiptCode = 204

	// Sender and Recipient errors
	UnknownSenderID                    ReceiptCode = 300
	UnknownRecipientID                 ReceiptCode = 301
	UnknownPhysicalSenderID            ReceiptCode = 302
	InvalidMessageType                 ReceiptCode = 303
	InvalidMessageClass                ReceiptCode = 304
	NotAllowedToSend                   ReceiptCode = 310
	NotAllowedToReceive                ReceiptCode = 311
	UserCertificateNotValid            ReceiptCode = 312
	OtherRecipientsNotAllowedToReceive ReceiptCode = 313

	// Size and Limit errors
	MessageSizeExceedsLimit ReceiptCode = 330

	// Network and Service errors
	NetworkError                     ReceiptCode = 400
	OSCIHubNotReachable              ReceiptCode = 401
	FolderNotReachable               ReceiptCode = 402
	LoggingServiceNotReachable       ReceiptCode = 403
	AuthorizationServiceNotReachable ReceiptCode = 404

	// Internal and Receiving errors
	InternalError        ReceiptCode = 500
	ErrorDuringReceiving ReceiptCode = 501

	// Success and Expiry messages
	MessageSuccessfullySent ReceiptCode = 601
	MessageExpiresSoon      ReceiptCode = 701
)
