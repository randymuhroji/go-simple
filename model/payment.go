package model

const TablePayment = "payment"

type Payment struct {
	Id        int     `json:"id" db:"id"`
	Method    string  `json:"method" db:"payment_method" fieldtag:"insert"`
	OrderId   int     `json:"order_id" db:"order_id" fieldtag:"insert"`
	TotalPaid float32 `json:"total_paid" db:"payment_total_paid" fieldtag:"insert"`
	CreatedAt string  `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string  `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string  `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string  `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}
