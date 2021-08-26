package usecase

import (
	"errors"
	"go-simple/config"
	"go-simple/model"
	orderRepo "go-simple/module/v1/order/repo"
	"go-simple/module/v1/payment/repo"
)

func PaymentList(conf config.Configuration) (users []model.Payment, err error) {
	db := conf.MysqlDB
	return repo.GetPaymentList(db)
}

// create  payment
func CreatePayment(conf config.Configuration, p *model.Payment) (user model.Payment, err error) {
	var (
		payload = model.Payment{}
	)

	if p.OrderId == 0 {
		err = errors.New("invalid order id")
		return
	}

	// validation order
	order, err := orderRepo.GetDetailOrder(conf.MysqlDB, p.OrderId)
	if err != nil {
		return
	}

	if order.Id == 0 {
		err = errors.New("invalid order id")
		return
	}

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.CreatePayment(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	// update order status
	order.Status = "paid"
	_, err = orderRepo.UpdateOrder(tx, &order)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *p, nil
}
