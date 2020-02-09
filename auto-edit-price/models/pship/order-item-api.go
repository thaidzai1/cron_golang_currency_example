package pship

import cm "gido.vn/gic/cron/auto-edit-price/models"

// ListOrderItemsRequest ...
type ListOrderItemsRequest struct {
	Paging cm.Paging
	Filter cm.Filter
}

// ListOrderItemsResponse ...
type ListOrderItemsResponse struct {
	Error      cm.Error     `json:"e"`
	PageInfo   cm.PageInfo  `json:"p"`
	OrderItems []*OrderItem `json:"d"`
}

// OrderItem ...
type OrderItem struct {
	ID                     int64  `json:"id,string,omitempty"`
	Gcode                  string `json:"gcode,omitempty"`
	Gscode                 string `json:"gscode,omitempty"`
	MerchantOrderID        int64  `json:"merchant_order_id,string,omitempty"`
	PackageInfoID          int64  `json:"package_info_id,string,omitempty"`
	SKU                    string `json:"sku,omitempty"`
	SupplierTrackingNumber string `json:"supplier_tracking_number,omitempty"`
	ProductID              int64  `json:"product_id,string,omitempty"`
	ProductName            string `json:"product_name,omitempty"`
	Warehouse              string `json:"warehouse,omitempty"`
	WHInventoryItemID      int64  `json:"wh_inventory_item_id,string,omitempty"`
	ParcelID               int64  `json:"parcel_id,string,omitempty"`
	RealWeight             int    `json:"real_weight,omitempty"`
	RID                    int64  `json:"rid,omitempty"`
	ActionAdminID          int64  `json:"action_admin_id,string,omitempty"`
	State                  string `json:"state,omitempty"`
	StateFinal             int    `json:"state_final,omitempty"`
	CreatedAt              string `json:"created_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
	CancelledAt            string `json:"cancelled_at,omitempty"`
}
