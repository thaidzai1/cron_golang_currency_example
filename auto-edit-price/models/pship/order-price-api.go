package pship

import (
	cm "gido.vn/gic/cron/auto-edit-price/models"
)

// EditOrderPriceRequest ...
type EditOrderPriceRequest struct {
	ID int64 `json:"id"`
}

// OrderPriceResponse ...
type OrderPriceResponse struct {
	Error      cm.Error      `json:"e"`
	PageInfo   cm.PageInfo   `json:"p"`
	OrderPrice []*OrderPrice `json:"d"`
}

// OrderPrice ...
type OrderPrice struct {
	ID                           int64   `json:"id,string,omitempty"`
	OrderID                      int64   `json:"order_id,string,omitempty"`
	PricingVersion               int64   `json:"pricing_version,string,omitempty"`
	TotalAmountBeforeDiscountVND int     `json:"total_amount_before_discount_vnd,omitempty"`
	TotalAmountAfterDiscountVND  int     `json:"total_amount_after_discount_vnd,omitempty"`
	ChargeableWeight             int     `json:"chargeable_weight,omitempty"`
	ChargeableDistance           int     `json:"chargeable_distance,omitempty"`
	PricingConfigID              int64   `json:"pricing_config_id,string,omitempty"`
	TransportationFeeValue       float64 `json:"transportation_fee_value,omitempty"`
	CreatedAt                    string  `json:"created_at,omitempty"`
	UpdatedAt                    string  `json:"updated_at,omitempty"`
	ActionAdminID                int64   `json:"action_admin_id,string,omitempty"`
	Type                         string  `json:"type,omitempty"`
	PrevTotalPrice               int64   `json:"prev_total_price,string,omitempty"`
	CurrTotalPrice               int64   `json:"curr_total_price,string,omitempty"`
	AdditionalPrice              int64   `json:"additional_price,string,omitempty"`
	PrevTotalPriceVAT            int64   `json:"prev_total_price_vat,string,omitempty"`
	CurrTotalPriceVAT            int64   `json:"curr_total_price_vat,string,omitempty"`
	AdditionalPriceVAT           int64   `json:"additional_price_vat,string,omitempty"`
	WalletTransactionID          int64   `json:"wallet_transaction_id,string,omitempty"`
	WalletTransactionState       string  `json:"wallet_transaction_state,omitempty"`
}
