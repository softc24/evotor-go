package cloud

import "encoding/json"

type BodyAccept struct {
	Positions   []AcceptPosition `json:"positions"`   // Массив товаров к приёмке.
	SupplierID  string           `json:"supplier_id"` // Идентификатор поставщика. Кассир выбирает поставщика из выпадающего списка перед сохранением документа.
	Sum         Decimal[Price]   `json:"sum"`         // Общая стоимость принятых товарных позиций. До двух знаков в дробной части.
	Description string           `json:"description"` // Основание для приёмки товара.
}

type BodyInventory struct {
	Positions         []InventoryPosition `json:"positions"`          // Массив товарных позиций.
	CompleteInventory bool                `json:"complete_inventory"` // Определяет является инветаризация полной или нет.
	Description       string              `json:"description"`        // Основание для приёмки товара.
}

type BodyRevaluation struct {
	BaseDocumentID     string                `json:"base_document_id"`     // Идентификатор документа приёмки, на основании которого осуществляется переоценка.
	BaseDocumentNumber int                   `json:"base_document_number"` // Номер документа, на основании которого осуществляется переоценка.
	Positions          []RevaluationPosition `json:"positions"`            // Массив товарных позиций для переоценки.
	Description        string                `json:"description"`          // Основание для приёмки товара.
}

type BodyReturn struct {
	Positions        []ReturnPosition `json:"positions"`         // Массив товарных позиций для возврата поставщику.
	Sum              Decimal[Price]   `json:"sum"`               // Общая стоимость товарных позиций, которые будут возвращены поставщику.
	CounterpartyName string           `json:"counterparty_name"` // ФИО поставщика.
	Description      string           `json:"description"`       // Основание для возврата товаров поставщику.
}

type BodyWriteOff struct {
	Positions   []WriteOffPosition `json:"positions"`   // Массив списанных товарных позиций.
	Sum         Decimal[Price]     `json:"sum"`         // Общая стоимость списанных товарных позиций.
	Description string             `json:"description"` // Основание для списания товарных позиций.
}

type ExtraKey struct {
	Identity    string `json:"identity"`
	AppID       string `json:"app_id"`
	Description string `json:"description"`
}

type Position struct {
	ProductID string     `json:"product_id"` // Идентификатор товара, уникальный в рамках магазина.
	Code      string     `json:"code"`       // Код товара или модификации товара.
	ExtraKeys []ExtraKey `json:"extra_keys"` // Массив объектов, соответствующих дополнительным полям (extras) товара, на момент добавления товара в документ.
}

type PositionID struct {
	ID   int    `json:"id"`   // Порядковый номер транзакции регистрации позиции в документе.
	UUID string `json:"uuid"` // Уникальный идентификатор товарной позиции в чеке.
}

type PositionQuantity struct {
	Quantity        Decimal[Quantity] `json:"quantity"`         // Количество товара, над которыми выполняется операция. Всегда положительное число. До трёх знаков в дробной части.
	InitialQuantity Decimal[Quantity] `json:"initial_quantity"` // Остаток товара до выполнения операции. До трёх знаков в дробной части.
}

type PositionProductDetails struct {
	ProductName  string `json:"product_name"`   // Наименование товара.
	ProductType  string `json:"product_type"`   // Тип товара.
	MeasureName  string `json:"measure_name"`   // Единица измерения из карточки товара.
	IsAgeLimited bool   `json:"is_age_limited"` // Определяет, установлены ли возрастные ограничения для товара.
	IsExcisable  bool   `json:"is_excisable"`   // Указывает, является ли товар подакцизным.

	BarCode string          `json:"bar_code"` // Штрихкод, по которому товар добавили в документ приёмки.
	Mark    json.RawMessage `json:"mark"`     // Марка алкогольного товара.

	Price     Decimal[Price] `json:"price"`      // Отпускная стоимость единицы товара.
	CostPrice Decimal[Price] `json:"cost_price"` // Закупочная стоимость единицы товара.

	TareVolume             Decimal[Quantity] `json:"tare_volume"`               // Ёмкость тары алкогольной продукции в литрах. До трёх знаков в дробной части.
	AlcoholProductKindCode int64             `json:"alcohol_product_kind_code"` // Код вида алкогольной продукции ФСРАР.
	AlcoholByVolume        Decimal[Quantity] `json:"alcohol_by_volume"`         // Крепость алкогольной продукции. До трёх знаков в дробной части.
}

type AcceptPosition struct {
	Position
	PositionID
	PositionQuantity

	Sum Decimal[Price] `json:"sum"` // Закупочная стоимость товарной позиции (price*quantity). До двух знаков в дробной части.
}

type InventoryPosition struct {
	Position
	PositionID
	PositionQuantity
}

type RevaluationPrice struct {
	Before Decimal[Price] `json:"before"` // Отпускная цена до переоценки.
	After  Decimal[Price] `json:"after"`  // Отпускная цена после переоценки.
	Accept Decimal[Price] `json:"accept"` // Цена приёмки.
}

type RevaluationPosition struct {
	Position
	PositionID

	Price RevaluationPrice `json:"price"` // Данные о ценах товара.
}

type ReturnPosition struct {
	AcceptPosition

	ResultSum Decimal[Price] `json:"result_sum"` // Отпускная стоимость товарной позиции с учётом скидок.
}

type WriteOffPosition struct {
	AcceptPosition
}

// BodySell describes the payload for a sell document.
type BodySell struct {
	DocDiscounts    []DocDiscount    `json:"doc_discounts"`     // Массив скидок на документ.
	ResultSum       Decimal[Price]   `json:"result_sum"`        // Итоговая сумма чека с учётом налогов и скидок.
	Payments        []Payment        `json:"payments"`          // Массив платежей.
	CustomerEmail   string           `json:"customer_email"`    // Адрес электронной почты покупателя.
	CustomerPhone   string           `json:"customer_phone"`    // Номер телефона покупателя.
	Positions       []SellPosition   `json:"positions"`         // Массив товарных позиций.
	Sum             Decimal[Price]   `json:"sum"`               // Итоговая сумма с учётом налогов.
	PrintGroups     []PrintGroup     `json:"print_groups"`      // Печатные группы.
	PosPrintResults []PosPrintResult `json:"pos_print_results"` // Результаты печати ККТ (массив объектов с фискальными данными).
}

type AgentRequisites struct {
	CounterpartyIndexes  []int  `json:"counterparty_indexes"`  // Индексы контрагентов из списка в заголовке документа.
	OperationDescription string `json:"operation_description"` // Описание операции контрагента.
}

type AttributesChoice struct {
	AttributeID   string `json:"attribute_id"`   // Идентификатор характеристики.
	AttributeName string `json:"attribute_name"` // Название характеристики.
	ChoiceID      string `json:"choice_id"`      // Идентификатор значения характеристики.
	ChoiceValue   string `json:"choice_value"`   // Текст значения характеристики.
}

type SettlementMethod struct {
	Type string `json:"type"`
}

type Tax struct {
	Type      string         `json:"type"`       // Ставка НДС.
	Sum       Decimal[Price] `json:"sum"`        // Сумма НДС на товарную позицию без применения скидок.
	ResultSum Decimal[Price] `json:"result_sum"` // Сумма НДС на товарную позицию с учетом скидок.
}

type DocDistributedDiscount struct {
	DiscountType    string           `json:"discount_type"`    // Тип скидки — сумовая или процентная.
	DiscountSum     Decimal[Price]   `json:"discount_sum"`     // Сумма скидки, которая была распределена на позицию в момент применения скидки на документ.
	DiscountPercent Decimal[Percent] `json:"discount_percent"` // Процент скидки, которая была распределена на позицию в момент применения скидки на документ.
}

type PositionDiscount struct {
	DiscountType    string           `json:"discount_type"`    // Тип скидки — сумовая или процентная.
	DiscountSum     Decimal[Price]   `json:"discount_sum"`     // Сумма скидки на позицию, примененной в момент добавления позиции в документ.
	DiscountPercent Decimal[Percent] `json:"discount_percent"` // Процент скидки на позицию в момент добавления позиции в документ.
	DiscountPrice   Decimal[Price]   `json:"discount_price"`   // Стоимость одной единицы товара в позиции, с учётом примененной скидки на позицию (в момент добавления позиции в документ).
}

type SplittedPosition struct {
	Quantity               Decimal[Quantity] `json:"quantity"`                 // Количество товара, над которыми выполняется операция.
	ResultPrice            Decimal[Price]    `json:"result_price"`             // Конечная стоимость единицы позиции после всех калькуляций в чеке (скидки на чек и скидки на позицию).
	ResultSum              Decimal[Price]    `json:"result_sum"`               // Отпускная стоимость товарной позиции с учётом скидок.
	PositionDiscount       Decimal[Price]    `json:"position_discount"`        // Сумма скидки на позицию, применённой в момент добавления позиции в документ.
	DocDistributedDiscount Decimal[Price]    `json:"doc_distributed_discount"` // Сумма скидки, которая была распределена на позицию в момент применения скидки на документ.
	TaxType                string            `json:"tax_type"`                 // Ставка НДС.
	ResultTaxSum           Decimal[Price]    `json:"result_tax_sum"`           // Сумма НДС на товарную позицию с учетом скидок.
}

type SellPosition struct {
	Position
	PositionID
	PositionQuantity
	PositionProductDetails

	DocDistributedDiscount DocDistributedDiscount `json:"doc_distributed_discount"` // Сумма скидки на документ, распределенная на товарную позицию.
	SplittedPositions      []SplittedPosition     `json:"splitted_positions"`       // Массив разделённых позиций.
	AgentRequisites        AgentRequisites        `json:"agent_requisites"`         // Агентские реквизиты.
	AttributesChoices      []AttributesChoice     `json:"attributes_choices"`       // Массив, содержащий объекты с описанием характеристик и их значений на момент формирования документа.
	PositionDiscount       PositionDiscount       `json:"position_discount"`        // Объект скидки на позицию, изначально применённой к конкретной позиции.
	SubPositions           []SellPosition         `json:"sub_positions"`            // Перечень позиций в рамках позиции в чеке.

	ResultSum   Decimal[Price] `json:"result_sum"`   // Отпускная стоимость товарной позиции с учётом скидок.
	Sum         Decimal[Price] `json:"sum"`          // Отпускная стоимость товарной позиции (price*quantity).
	ResultPrice Decimal[Price] `json:"result_price"` // Конечная стоимость единицы позиции после всех калькуляций в чеке (скидки на чек и скидки на позицию).
	Excise      Decimal[Price] `json:"excise"`       // Сумма акциза, за акцизный товар.
	Tax         Tax            `json:"tax"`          // Информация о расчете НДС на товарную позицию.

	QuantityInPackage Decimal[Quantity] `json:"quantityInPackage"` // Количество товара в упаковке всего.

	SettlementMethod SettlementMethod `json:"settlement_method"`
	PrintGroupID     string           `json:"print_group_id"` // Идентификатор печатной группы при разделении чека.

}

type DocDiscount struct {
	DiscountType    string           `json:"discount_type"`    // Тип скидки - суммовая или процентная.
	DiscountSum     Decimal[Price]   `json:"discount_sum"`     // Сумма скидки.
	DiscountPercent Decimal[Percent] `json:"discount_percent"` // Процент скидки.
	Coupon          string           `json:"coupon,omitempty"` // Идентификатор купона (для случаев применения скидки через штатный функционал скидочных купонов).
}

type PaymentAppInfo struct {
	AppID string `json:"app_id"` // Идентификатор приложения, которое произвело платёж электронными средствами.
	Name  string `json:"name"`   // Наименование платежной системы.
}

type PaymentPart struct {
	PrintGroupID string         `json:"print_group_id"` // Идентификатор печатной группы при разделении чека.
	PartSum      Decimal[Price] `json:"part_sum"`       // Часть платежа, которая пошла на оплату данной печатной группы.
	Change       Decimal[Price] `json:"change"`         // Сдача.
}

type Payment struct {
	ID      string         `json:"id"`       // Уникальный идентификатор платежа.
	Type    string         `json:"type"`     // Способ платежа.
	Sum     Decimal[Price] `json:"sum"`      // Итоговая сумма к оплате данным способом оплаты.
	AppInfo PaymentAppInfo `json:"app_info"` // Объект с информацией для электронных платежей, в случае типа оплаты ELECTRON.
	Parts   []PaymentPart  `json:"parts"`    // Массив платежей для разделённой оплаты.
}

type PrintGroup struct {
	ID             string `json:"id"`              // Идентификатор печатной группы.
	Type           string `json:"type"`            // Способ оплаты.
	OrgName        string `json:"org_name"`        // Наименование организации.
	OrgInn         string `json:"org_inn"`         // ИНН организации.
	OrgAddress     string `json:"org_address"`     // Адрес организации.
	TaxationSystem string `json:"taxation_system"` // Используемая система налогообложения.
}

type PosPrintResult struct {
	PrintGroupID         string         `json:"print_group_id"`         // Идентификатор печатной группы.
	ReceiptNumber        int            `json:"receipt_number"`         // Номер чека в ККТ.
	DocumentNumber       int            `json:"document_number"`        // Номер документа в ККТ.
	SessionNumber        int            `json:"session_number"`         // Номер смены в ККТ.
	CheckSum             Decimal[Price] `json:"check_sum"`              // Итоговая сумма для печати на документе, с учётом скидок.
	ReceiptDate          string         `json:"receipt_date"`           // Дата печати в ККТ на момент печати документа, в формате “DDMMYYYY”.
	FnSerialNumber       string         `json:"fn_serial_number"`       // Заводской номер фискального накопителя.
	KktSerialNumber      string         `json:"kkt_serial_number"`      // Серийный номер ККТ.
	KktRegNumber         string         `json:"kkt_reg_number"`         // Регистрационный номер ККТ.
	FiscalSignDocNumber  string         `json:"fiscal_sign_doc_number"` // Номер фискального признака документа.
	FiscalDocumentNumber int            `json:"fiscal_document_number"` // Фискальный номер документа в ККТ из ФН.
}

type BodyPayback struct {
	BodySell

	BaseDocumentID     string `json:"base_document_id,omitempty"`
	BaseDocumentNumber int    `json:"base_document_number,omitempty"`
}

type BodyBuy struct {
	BodySell
}

type BuyBack struct {
	BodySell

	BaseDocumentID     string `json:"base_document_id,omitempty"`
	BaseDocumentNumber int    `json:"base_document_number,omitempty"`
}
