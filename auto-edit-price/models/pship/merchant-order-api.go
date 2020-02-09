package pship

import cm "gido.vn/gic/cron/auto-edit-price/models"

// ListMerchantOrdersRequest ...
type ListMerchantOrdersRequest struct {
	Paging cm.Paging
	Filter cm.Filter
}

// MerchantOrdersResponse ...
type MerchantOrdersResponse struct {
	Error         cm.Error         `json:"e"`
	PageInfo      cm.PageInfo      `json:"p"`
	MerchantOrder []*MerchantOrder `json:"d"`
}

// MerchantOrder ...
type MerchantOrder struct {
	ID                     int64  `json:"id,string,omitempty"`
	Code                   string `json:"code,omitempty"`
	SupplierOrderNumber    string `json:"supplier_order_number,omitempty"`
	SupplierTrackingNumber string `json:"supplier_tracking_number,omitempty"`
	SupplierCourier        string `json:"supplier_courier,omitempty"`
	SupplierCouponCode     string `json:"supplier_coupon_code,omitempty"`
	InvoiceLink            string `json:"invoice_link,omitempty"`
	CouponCode             string `json:"coupon_code,omitempty"`
	Country                string `json:"country,omitempty"`
	Platform               string `json:"platform,omitempty"`
	UseInsurance           bool   `json:"use_insurance,omitempty"`
	FirstMileMethod        string `json:"first_mile_method,omitempty"`
	LastMileMethod         string `json:"last_mile_method,omitempty"`
	NoteOrder              string `json:"note_order,omitempty"`
	NoteShipping           string `json:"note_shipping,omitempty"`
	NoteAdminMerchant      string `json:"note_admin_merchant,omitempty"`
	NoteAdmin              string `json:"note_admin,omitempty"`
	NoteCancel             string `json:"note_cancel,omitempty"`
	InvoiceValueVND        int    `json:"invoice_value_vnd,omitempty"`
	IsCBE                  bool   `json:"is_cbe,omitempty"`
	IsTest                 int    `json:"is_test,omitempty"`
	Flow                   string `json:"flow,omitempty"`
	RID                    int64  `json:"rid,omitempty"`
	ActionAdminID          int64  `json:"action_admin_id,omitempty"`
	State                  string `json:"state,omitempty"`
	StateFinal             int    `json:"state_final,omitempty"`
	CBEOrderDate           string `json:"cbe_order_date,omitempty"`
	CreatedAt              string `json:"created_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
	CancelledAt            string `json:"cancelled_at,omitempty"`
	CODAmountVND           int    `json:"cod_amount_vnd,omitempty"`
	EstimatedDeliveryDate  string `json:"estimated_delivery_date,omitempty"`
	MerchantType           string `json:"merchant_type,omitempty"`
	SortingCode            string `json:"sorting_code,omitempty"`
	CustomsInvoiceNo       string `json:"customs_invoice_no,omitempty"`
	CustomsInvoiceDate     string `json:"customs_invoice_date,omitempty"`
	CustomsInvoiceLink     string `json:"customs_invoice_link,omitempty"`
	CustomsPONo            string `json:"customs_po_no,omitempty"`
	CustomsPODate          string `json:"customs_po_date,omitempty"`
	CustomsPOLink          string `json:"customs_po_link,omitempty"`
	CustomsCOLink          string `json:"customs_co_link,omitempty"`
	CustomsCQLink          string `json:"customs_cq_link,omitempty"`
	CustomsImportDuty      string `json:"customs_import_duty,omitempty"`
	CustomsVAT             int64  `json:"customs_vat,string,omitempty"`
	CustomsNote            string `json:"customs_note,omitempty"`
	CustomsHBLNo           string `json:"customs_hbl_no,omitempty"`
	CustomsMBLNo           string `json:"customs_mbl_no,omitempty"`
	DepartureDate          string `json:"departure_date,omitempty"`
	TruckNo                string `json:"truck_no,omitempty"`
	LetterOfAuthorization  string `json:"letter_of_authorization,omitempty"`
	DeliveryOrderNo        string `json:"delivery_order_no,omitempty"`
	CustomsDeclarationNo   string `json:"customs_declaration_no,omitempty"`
}
