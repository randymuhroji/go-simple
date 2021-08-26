package model

const TableOrder = "orders"
const TableOrderDetail = "order_detail"

type Order struct {
	Id          int     `json:"id" db:"id"`
	Code        string  `json:"code" db:"order_code" fieldtag:"insert"`
	TotalAmount float32 `json:"total_amount" db:"order_total_amount" fieldtag:"insert"`
	TotalPaid   float32 `json:"total_paid" db:"order_total_paid" fieldtag:"insert"`
	Discount    float32 `json:"discount" db:"order_discount" fieldtag:"insert"`
	Status      string  `json:"status" db:"order_status" fieldtag:"insert,update"`
	CreatedAt   string  `json:"created_at,omitempty" db:"created_at"`
	CreatedBy   string  `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt   string  `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy   string  `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type OrderPayload struct {
	Order
	Items []OrderDetail `json:"items"`
}

type OrderView struct {
	Id          int     `json:"id" db:"id"`
	Code        string  `json:"code" db:"order_code" fieldtag:"insert"`
	TotalAmount float32 `json:"total_amount" db:"order_total_amount" fieldtag:"insert"`
	TotalPaid   float32 `json:"total_paid" db:"order_total_paid" fieldtag:"insert"`
	Discount    float32 `json:"discount" db:"order_discount" fieldtag:"insert"`
	Status      string  `json:"status" db:"order_status" fieldtag:"insert,update"`
}

type OrderDetail struct {
	Id           int     `json:"id" db:"id"`
	ProductSKU   string  `json:"sku" db:"product_sku" fieldtag:"insert"`
	OrderId      int     `json:"order_id" db:"order_id" fieldtag:"insert"`
	ProductPrice float32 `json:"price" db:"product_price" fieldtag:"insert"`
	SubTotal     float32 `json:"sub_total" db:"sub_total" fieldtag:"insert"`
	Qty          int     `json:"qty" db:"qty" fieldtag:"insert"`
	CreatedAt    string  `json:"created_at,omitempty" db:"created_at"`
	CreatedBy    string  `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt    string  `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy    string  `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type OrderDetailView struct {
	Id           int     `json:"id" db:"id"`
	ProductSKU   string  `json:"sku" db:"product_sku" fieldtag:"insert"`
	OrderId      int     `json:"int" db:"order_id" fieldtag:"insert"`
	ProductPrice float32 `json:"price" db:"product_price" fieldtag:"insert"`
	SubTotal     float32 `json:"sub_total" db:"sub_total" fieldtag:"insert"`
	Qty          int     `json:"qty" db:"product_qty" fieldtag:"insert"`
}
