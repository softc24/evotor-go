package evotor

import "time"

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

type DeviceModel string

const (
	DeviceModel10      DeviceModel = "10"
	DeviceModel72      DeviceModel = "7_2"
	DeviceModel72ATOL  DeviceModel = "7_2_ATOL"
	DeviceModel73      DeviceModel = "7_3"
	DeviceModel5       DeviceModel = "5"
	DeviceModel5I      DeviceModel = "5_I"
	DeviceModelPower   DeviceModel = "POWER"
	DeviceModelUnknown DeviceModel = "UNKNOWN"
)

type Device struct {
	ID      string `json:"id"`       // Уникальный идентификатор смарт-терминала в Облаке Эвотор.
	Name    string `json:"name"`     // Имя смарт-терминала, которое пользователь дал устройству в Личном кабинете.
	StoreID string `json:"store_id"` // Уникальный идентификатор торговой точки в Облаке Эвотор.
	UserID  string `json:"user_id"`  // Идентификатор пользователя Эвотор.

	INN             string      `json:"inn"`              // ИНН владельца терминала.
	TimezoneOffset  int64       `json:"timezone_offset"`  // Часовой пояс, в котором работает смарт-терминал. Представлен в виде смещения в миллисекундах относительно часового пояса UTC.
	IMEI            string      `json:"imei"`             // IMEI смарт-терминала.
	FirmwareVersion string      `json:"firmware_version"` // Версия ПО смарт-терминала.
	Location        Location    `json:"location"`         // GPS-координаты смарт-терминала.
	SerialNumber    string      `json:"serial_number"`    // Серийный номер устройства. Состоит из 14-и символов. Написан на наклейке на обратной стороне Эвотора.
	DeviceModel     DeviceModel `json:"device_model"`     // Модель устройства
	CreatedAt       time.Time   `json:"created_at"`       // Дата и время создания объекта
	UpdatedAt       time.Time   `json:"updated_at"`       // Дата и время обновления данных объекта
}

type Location struct {
	Lng float64 `json:"lng"` // Долгота
	Lat float64 `json:"lat"` // Широта
}

type Event[T any] struct {
	ID        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Type      EventType   `json:"type"`
	Source    EventSource `json:"source"`
	Action    EventAction `json:"action"`
	Payload   T           `json:"payload"`
}
