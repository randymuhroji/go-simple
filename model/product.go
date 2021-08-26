package model

const TableProduct = "product"

type Product struct {
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"product_name" fieldtag:"insert,update"`
	SKU       string  `json:"sku" db:"product_sku" fieldtag:"insert,update"`
	Price     float32 `json:"price" db:"product_price" fieldtag:"insert,update"`
	Qty       int     `json:"qty" db:"product_qty" fieldtag:"insert,update"`
	Deleted   int     `json:"deleted" db:"deleted" fieldtag:"insert,update,delete"`
	CreatedAt string  `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string  `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string  `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string  `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type ProductView struct {
	Id      int     `json:"id" db:"id"`
	Name    string  `json:"name" db:"product_name" fieldtag:"insert,update"`
	SKU     string  `json:"sku" db:"product_sku" fieldtag:"insert,update"`
	Price   float32 `json:"price" db:"product_price" fieldtag:"insert,update"`
	Qty     int     `json:"qty" db:"product_qty" fieldtag:"insert,update"`
	Deleted int     `json:"deleted" db:"deleted" fieldtag:"insert,update,delete"`
}
