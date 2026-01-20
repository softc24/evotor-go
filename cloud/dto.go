package cloud

import (
	"encoding/json"
	"time"
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

type Employee struct {
	ID             string    `json:"id"`              // Уникальный идентификатор сотрудника в Облаке Эвотор.
	Name           string    `json:"name"`            // Имя сотрудника, которое пользователь указал в Личном кабинете.
	LastName       string    `json:"last_name"`       // Фамилия сотрудника, которую пользователь указал в Личном кабинете.
	PatronymicName string    `json:"patronymic_name"` // Отчество сотрудника, которое пользователь указал в Личном кабинете.
	Phone          string    `json:"phone"`           // Номер телефона сотрудника, который пользователь указал в Личном кабинете.
	Stores         []string  `json:"stores"`          // Массив идентификаторов магазинов, которым соответствует идентификатор сотрудника.
	Role           string    `json:"role"`            // Роль, которая задаёт права доступа сотрудника к функционалу смарт-терминала.
	RoleID         string    `json:"role_id"`         // Уникальный идентификатор роли сотрудника.
	UserID         string    `json:"user_id"`         // Идентификатор пользователя Эвотор.
	CreatedAt      time.Time `json:"created_at"`      // Дата и время создания объекта.
	UpdatedAt      time.Time `json:"updated_at"`      // Дата и время обновления данных объекта.
}

type Store struct {
	ID        string    `json:"id"`         // Уникальный идентификатор магазина в Облаке Эвотор.
	Name      string    `json:"name"`       // Имя магазина, которое пользователь указал в Личном кабинете.
	Address   string    `json:"address"`    // Адрес магазина, который пользователь указал в Личном кабинете.
	UserID    string    `json:"user_id"`    // Идентификатор пользователя Эвотор.
	CreatedAt time.Time `json:"created_at"` // Дата и время создания объекта.
	UpdatedAt time.Time `json:"updated_at"` // Дата и время обновления данных объекта
}

type Role struct {
	ID   string `json:"id"`   // Уникальный идентификатор роли
	Name string `json:"name"` // Название роли
}

type DocumentType string

const (
	DocumentTypeOpenSession    DocumentType = "OPEN_SESSION"
	DocumentTypePosOpenSession DocumentType = "POS_OPEN_SESSION"
	DocumentTypeCloseSession   DocumentType = "CLOSE_SESSION"
	DocumentTypeCashIncome     DocumentType = "CASH_INCOME"
	DocumentTypeCashOutcome    DocumentType = "CASH_OUTCOME"
	DocumentTypeInventory      DocumentType = "INVENTORY"
	DocumentTypeAccept         DocumentType = "ACCEPT"
	DocumentTypeRevaluation    DocumentType = "REVALUATION"
	DocumentTypeWriteOff       DocumentType = "WRITE_OFF"
	DocumentTypeReturn         DocumentType = "RETURN"
	DocumentTypeOpenTare       DocumentType = "OPEN_TARE"
	DocumentTypeSell           DocumentType = "SELL"
	DocumentTypePayback        DocumentType = "PAYBACK"
	DocumentTypeBuy            DocumentType = "BUY"
	DocumentTypeBuyback        DocumentType = "BUYBACK"
	DocumentTypeXReport        DocumentType = "X_REPORT"
	DocumentTypeZReport        DocumentType = "Z_REPORT"
	DocumentTypeCorrection     DocumentType = "CORRECTION"
)

type Document struct {
	Type           DocumentType    `json:"type"`             // Тип документа
	ID             string          `json:"id"`               // Уникальный в рамках магазина идентификатор документа в Облаке Эвотор
	Extras         json.RawMessage `json:"extras"`           // Объект с дополнительной информацией от сторонних приложений
	Number         int             `json:"number"`           // Порядковый номер документа на смарт-терминале
	CloseDate      time.Time       `json:"close_date"`       // Дата сохранения документа на смарт-терминале
	TimezoneOffset int64           `json:"time_zone_offset"` // Смещение часового пояса в миллисекундах относительно UTC
	SessionID      string          `json:"session_id"`       // Уникальный идентификатор смены
	SessionNumber  int             `json:"session_number"`   // Порядковый номер смены на смарт-терминале
	CloseUserID    string          `json:"close_user_id"`    // Уникальный идентификатор сотрудника, создавшего документ
	DeviceID       string          `json:"device_id"`        // Идентификатор смарт-терминала
	StoreID        string          `json:"store_id"`         // Уникальный идентификатор магазина
	UserID         string          `json:"user_id"`          // Идентификатор пользователя Эвотор
	Version        string          `json:"version"`          // Версия схемы документа
	Counterparties json.RawMessage `json:"counterparties"`   // Массив контрагентов (присутствует если установлено приложение "Агентская схема")
	Body           json.RawMessage `json:"body"`             // Объект с основной информацией о документе (структура зависит от типа)
}

type Counterparty struct {
	Addresses []string `json:"addresses"`  // Список адресов контрагента.
	FullName  string   `json:"full_name"`  // Полное наименование контрагента.
	Index     int      `json:"index"`      // Порядковый номер в списке (индекс).
	INN       string   `json:"inn"`        // ИНН контрагента.
	ID        string   `json:"id"`         // Идентификатор контрагента, который задаёт приложение.
	KPP       string   `json:"kpp"`        // КПП контрагента.
	Phones    []string `json:"phones"`     // Список номеров телефонов контрагента.
	Role      string   `json:"role"`       // Роль контрагента.
	ShortName string   `json:"short_name"` // Краткое наименование контрагента.
	Type      string   `json:"type"`       // Тип контрагента.
}

type Choice struct {
	ID   string `json:"id"`   // Идентификатор значения характеристики
	Name string `json:"name"` // Название характеристики
}

type Attribute struct {
	ID      string   `json:"id"`      // Идентификатор характеристик
	Name    string   `json:"name"`    // Название группы характеристик
	Choices []Choice `json:"choices"` // Массив значений характеристик
}

type ProductGroup struct {
	ID         string      `json:"id"`                   // Идентификатор группы, уникальный в рамках магазина
	ParentID   string      `json:"parent_id,omitempty"`  // Уникальный идентификатор группы или группы модификаций
	Name       string      `json:"name"`                 // Название группы, длиной не менее 1 символа и не более 128 символов
	Barcodes   []string    `json:"barcodes,omitempty"`   // Массив штрихкодов для группы модификаций
	Attributes []Attribute `json:"attributes,omitempty"` // Объект, содержащий перечисление возможных характеристик и их значений
	StoreID    string      `json:"store_id"`             // Идентификатор магазина, в базе которого хранится группа
	UserID     string      `json:"user_id"`              // Идентификатор пользователя Эвотор
	CreatedAt  time.Time   `json:"created_at"`           // Дата и время создания объекта
	UpdatedAt  time.Time   `json:"updated_at"`           // Дата и время обновления данных объекта
}

type Product struct {
	Type              string             `json:"type"`                         // Тип товара
	Name              string             `json:"name"`                         // Название товара (услуги) или модификации товара
	Code              string             `json:"code,omitempty"`               // Код товара или модификации товара
	Price             Decimal[Price]     `json:"price"`                        // Отпускная цена товара (услуги) или модификации товара
	CostPrice         *Decimal[Price]    `json:"cost_price,omitempty"`         // Себестоимость товара
	Quantity          *Decimal[Quantity] `json:"quantity,omitempty"`           // Количество товара на складе
	MeasureName       string             `json:"measure_name"`                 // Единица измерения для товара или модификации товара
	IsExcisable       bool               `json:"is_excisable,omitempty"`       // Указывает, является ли товар подакцизным
	IsAgeLimited      bool               `json:"is_age_limited,omitempty"`     // Определяет, установлены ли возрастные ограничения для товара
	Tax               string             `json:"tax"`                          // Код ставки НДС для товара или модификации товара
	AllowToSell       bool               `json:"allow_to_sell"`                // Признак разрешения продажи
	Description       string             `json:"description,omitempty"`        // Описание товара или модификации товара
	ArticleNumber     string             `json:"article_number,omitempty"`     // Артикул товара или модификации товара
	ParentID          string             `json:"parent_id,omitempty"`          // Уникальный идентификатор группы или группы модификаций
	Barcodes          []string           `json:"barcodes,omitempty"`           // Массив штрихкодов
	AttributesChoices json.RawMessage    `json:"attributes_choices,omitempty"` // Выбранные характеристики модификации
	ID                string             `json:"id"`                           // Идентификатор товара, услуги или модификации
	StoreID           string             `json:"store_id"`                     // Идентификатор магазина
	UserID            string             `json:"user_id"`                      // Идентификатор пользователя Эвотор
	CreatedAt         time.Time          `json:"created_at"`                   // Дата и время создания объекта
	UpdatedAt         time.Time          `json:"updated_at"`                   // Дата и время обновления данных объекта
}

type BulkTaskType string

const (
	BulkTaskTypeProduct      BulkTaskType = "product"
	BulkTaskTypeProductGroup BulkTaskType = "product-group"
)

type BulkTaskStatus string

const (
	BulkTaskStatusAccepted  BulkTaskStatus = "ACCEPTED"
	BulkTaskStatusRunning   BulkTaskStatus = "RUNNING"
	BulkTaskStatusCompleted BulkTaskStatus = "COMPLETED"
	BulkTaskStatusDeclined  BulkTaskStatus = "DECLINED"
	BulkTaskStatusFailed    BulkTaskStatus = "FAILED"
)

type BulkTask struct {
	ID         string            `json:"id"`                // Идентификатор задачи
	Type       BulkTaskType      `json:"type"`              // Объект, который обрабатывается в задаче: товар или группа товаров
	Status     BulkTaskStatus    `json:"status"`            // Состояние задачи
	ModifiedAt time.Time         `json:"modified_at"`       // Время обновления состояния задачи
	Details    []json.RawMessage `json:"details,omitempty"` // Детали выполнения задачи (присутствует для COMPLETED, DECLINED, FAILED)
}
