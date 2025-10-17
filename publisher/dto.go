package publisher

type EventType string

const (
	EventTypeDocument            EventType = "Document"
	EventTypeProduct             EventType = "Product"
	EventTypeProductGroup        EventType = "ProductGroup"
	EventTypeSettings            EventType = "Settings"
	EventTypeMarketplacePurchase EventType = "MarketplacePurchase"
)

type EventSource string

const (
	EventSourceApplication EventSource = "APPLICATION"
	EventSourceDevice      EventSource = "DEVICE"
)

type EventAction string

const (
	EventActionCreated                 EventAction = "created"
	EventActionUpdated                 EventAction = "updated"
	EventActionRemoved                 EventAction = "removed"
	EventActionActivated               EventAction = "Activated"
	EventActionTermsChanged            EventAction = "TermsChanged"
	EventActionDeviceAssignmentChanged EventAction = "DeviceAssignmentChanged"
	EventActionRenewed                 EventAction = "Renewed"
	EventActionCancellationRequested   EventAction = "CancellationRequested"
	EventActionCancelled               EventAction = "Cancelled"
)

type Event[T any] struct {
	ID        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Type      EventType   `json:"type"`
	Source    EventSource `json:"source"`
	Action    EventAction `json:"action"`
	Payload   T           `json:"payload"`
}
