package notifications

import (
	"encoding/json"
	"fmt"
	"time"
)

type Notification struct {
	ID        int       `json:"id"`
	Data      Data      `json:"notification"`
	Timestamp time.Time `json:"timestamp"`
	Read      bool      `json:"read"`
}

type Thumbnail struct {
	Tiny  string `json:"tiny"`
	Small string `json:"small"`
}

type Data interface {
	// TODO maybe should be made 'real interface', which will allow
	// to use typed channels, type checking and semantic dispatching
	// instead of typecase:

	// Serialize() []byte
	// Describe() (string, string)
}

type notificationWrapper struct {
	Notification Data `json:"notification"`
}

type messageWrapper struct {
	Message Data `json:"message"`
}

type walletWrapper struct {
	Message Data `json:"wallet"`
}

type messageReadWrapper struct {
	MessageRead Data `json:"messageRead"`
}

type messageTypingWrapper struct {
	MessageRead Data `json:"messageTyping"`
}

type OrderNotification struct {
	Type              string    `json:"type"`
	Title             string    `json:"title"`
	BuyerId           string    `json:"buyerId"`
	BuyerBlockchainId string    `json:"buyerBlockchainId"`
	Thumbnail         Thumbnail `json:"thumbnail"`
	Timestamp         int       `json:"timestamp"`
	OrderId           string    `json:"orderId"`
}

type PaymentNotification struct {
	Type         string `json:"type"`
	OrderId      string `json:"orderId"`
	FundingTotal uint64 `json:"fundingTotal"`
}

type OrderConfirmationNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type OrderDeclinedNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type OrderCancelNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type RefundNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type FulfillmentNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type CompletionNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type DisputeOpenNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type DisputeUpdateNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type DisputeCloseNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type DisputeAcceptedNotification struct {
	Type      string    `json:"type"`
	OrderId   string    `json:"orderId"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type FollowNotification struct {
	Type   string `json:"type"`
	PeerId string `json:"peerId"`
}

type UnfollowNotification struct {
	Type   string `json:"type"`
	PeerId string `json:"peerId"`
}

type ModeratorAddNotification struct {
	Type   string `json:"type"`
	PeerId string `json:"peerId"`
}

type ModeratorRemoveNotification struct {
	Type   string `json:"type"`
	PeerId string `json:"peerId"`
}

type StatusNotification struct {
	Status string `json:"status"`
}

type ChatMessage struct {
	MessageId string    `json:"messageId"`
	PeerId    string    `json:"peerId"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatRead struct {
	MessageId string `json:"messageId"`
	PeerId    string `json:"peerId"`
	Subject   string `json:"subject"`
}

type ChatTyping struct {
	PeerId  string `json:"peerId"`
	Subject string `json:"subject"`
}

type IncomingTransaction struct {
	Txid          string    `json:"txid"`
	Value         int64     `json:"value"`
	Address       string    `json:"address"`
	Status        string    `json:"status"`
	Memo          string    `json:"memo"`
	Timestamp     time.Time `json:"timestamp"`
	Confirmations int32     `json:"confirmations"`
	OrderId       string    `json:"orderId"`
	Thumbnail     string    `json:"thumbnail"`
	Height        int32     `json:"height"`
	CanBumpFee    bool      `json:"canBumpFee"`
}

func wrap(i interface{}) interface{} {
	switch i.(type) {
	case OrderNotification:
		n := i.(OrderNotification)
		n.Type = "order"
		return notificationWrapper{n}
	case PaymentNotification:
		n := i.(PaymentNotification)
		n.Type = "payment"
		return notificationWrapper{n}
	case OrderConfirmationNotification:
		n := i.(OrderConfirmationNotification)
		n.Type = "orderConfirmation"
		return notificationWrapper{n}
	case OrderCancelNotification:
		n := i.(OrderCancelNotification)
		n.Type = "cancel"
		return notificationWrapper{n}
	case RefundNotification:
		n := i.(RefundNotification)
		n.Type = "refund"
		return notificationWrapper{n}
	case FulfillmentNotification:
		n := i.(FulfillmentNotification)
		n.Type = "fulfillment"
		return notificationWrapper{n}
	case CompletionNotification:
		n := i.(CompletionNotification)
		n.Type = "orderComplete"
		return notificationWrapper{n}
	case DisputeOpenNotification:
		n := i.(DisputeOpenNotification)
		n.Type = "disputeOpen"
		return notificationWrapper{n}
	case DisputeUpdateNotification:
		n := i.(DisputeUpdateNotification)
		n.Type = "disputeUpdate"
		return notificationWrapper{n}
	case DisputeCloseNotification:
		n := i.(DisputeCloseNotification)
		n.Type = "disputeClose"
		return notificationWrapper{n}
	case DisputeAcceptedNotification:
		n := i.(DisputeAcceptedNotification)
		n.Type = "disputeAccepted"
		return notificationWrapper{n}
	case FollowNotification:
		n := i.(FollowNotification)
		n.Type = "follow"
		return notificationWrapper{n}
	case UnfollowNotification:
		n := i.(UnfollowNotification)
		n.Type = "unfollow"
		return notificationWrapper{n}
	case ModeratorAddNotification:
		n := i.(ModeratorAddNotification)
		n.Type = "moderatorAdd"
		return notificationWrapper{n}
	case ModeratorRemoveNotification:
		n := i.(ModeratorRemoveNotification)
		n.Type = "moderatorRemove"
		return notificationWrapper{n}
	case ChatMessage:
		return messageWrapper{i.(ChatMessage)}
	case ChatRead:
		return messageReadWrapper{i.(ChatRead)}
	case ChatTyping:
		return messageTypingWrapper{i.(ChatTyping)}
	case IncomingTransaction:
		return walletWrapper{i.(IncomingTransaction)}
	default:
		return i
	}
}

func Serialize(i interface{}) []byte {
	w := wrap(i)
	if _, ok := w.([]byte); ok {
		return w.([]byte)
	}
	b, _ := json.MarshalIndent(w, "", "    ")
	return b
}

func Describe(i interface{}) (string, string) {
	var head, body string
	switch i.(type) {
	case OrderNotification:
		head = "Order received"

		n := i.(OrderNotification)
		var buyer string
		if n.BuyerBlockchainId != "" {
			buyer = n.BuyerBlockchainId
		} else {
			buyer = n.BuyerId
		}
		form := "You received an order \"%s\".\n\nOrder ID: %s\nBuyer: %s\nThumbnail: %s\nTimestamp: %d"
		body = fmt.Sprintf(form, n.Title, n.OrderId, buyer, n.Thumbnail.Small, n.Timestamp)

	case PaymentNotification:
		head = "Payment received"

		n := i.(PaymentNotification)
		form := "Payment for order \"%s\" received (total %d)."
		body = fmt.Sprintf(form, n.OrderId, n.FundingTotal)

	case OrderConfirmationNotification:
		head = "Order confirmed"

		n := i.(OrderConfirmationNotification)
		form := "Order \"%s\" has been confirmed."
		body = fmt.Sprintf(form, n.OrderId)

	case OrderCancelNotification:
		head = "Order cancelled"

		n := i.(OrderCancelNotification)
		form := "Order \"%s\" has been cancelled."
		body = fmt.Sprintf(form, n.OrderId)

	case RefundNotification:
		head = "Payment refunded"

		n := i.(RefundNotification)
		form := "Payment refund for order \"%s\" received."
		body = fmt.Sprintf(form, n.OrderId)

	case FulfillmentNotification:
		head = "Order fulfilled"

		n := i.(FulfillmentNotification)
		form := "Order \"%s\" was marked as fulfilled."
		body = fmt.Sprintf(form, n.OrderId)

	case CompletionNotification:
		head = "Order completed"

		n := i.(CompletionNotification)
		form := "Order \"%s\" was marked as completed."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeOpenNotification:
		head = "Dispute opened"

		n := i.(DisputeOpenNotification)
		form := "Dispute around order \"%s\" was opened."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeUpdateNotification:
		head = "Dispute updated"

		n := i.(DisputeUpdateNotification)
		form := "Dispute around order \"%s\" was updated."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeCloseNotification:
		head = "Dispute closed"

		n := i.(DisputeCloseNotification)
		form := "Dispute around order \"%s\" was closed."
		body = fmt.Sprintf(form, n.OrderId)
	}
	return head, body
}
