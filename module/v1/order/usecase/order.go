package usecase

import (
	"fmt"
	"go-simple/config"
	"go-simple/model"
	"go-simple/module/v1/order/repo"
	"go-simple/utl/middleware/request"
)

func OrderList(conf config.Configuration) (users []model.OrderView, err error) {
	db := conf.MysqlDB
	return repo.GetOrderList(db)
}

// create order
func CreateOrder(conf config.Configuration, p *model.OrderPayload) (user model.OrderPayload, err error) {
	var (
		total = 0
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	// total
	for i, item := range p.Items {
		subTotal := item.Qty * int(item.ProductPrice)
		total = total + subTotal
		p.Items[i].SubTotal = float32(subTotal)
	}

	cd := fmt.Sprintf("ORDER-%s", request.Order())
	// validation and generate code

	p.Order.Code = cd
	p.Order.TotalAmount = float32(total)
	p.Order.TotalPaid = float32(total) - (float32(total) * p.Order.Discount)
	p.Order.Status = "waiting_payment"

	// save order
	orderSqlResult, _ := repo.CreateOrder(tx, &p.Order)
	OrderId, err := orderSqlResult.LastInsertId()

	if err != nil {
		tx.Rollback()
		return
	}
	p.Order.Id = int(OrderId)

	// insert into detail
	for _, item := range p.Items {
		item.OrderId = p.Order.Id
		_, err = repo.CreateOrderDetail(tx, &item)
		if err != nil {
			tx.Rollback()
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *p, nil
}
