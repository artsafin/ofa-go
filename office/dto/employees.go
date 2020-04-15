package dto

import (
	"html/template"
	"oea-go/common"
)

type paymentCategory struct {
	Name     string
	Total    common.MoneyRub
	Payments []*EmployeePayment
}

func (cat *paymentCategory) AddPayment(payment *EmployeePayment) {
	cat.Payments = append(cat.Payments, payment)
	cat.Total += payment.Amount
}

type EmployeePayment struct {
	Name    string
	Comment template.HTML
	Amount  common.MoneyRub
}

func NewEmployeesPaymentCategories(month string) *EmployeesPaymentCategories {
	return &EmployeesPaymentCategories{
		Month:      month,
		Categories: make(map[string]*paymentCategory),
	}
}

type EmployeesPaymentCategories struct {
	Month      string
	Categories map[string]*paymentCategory
	Total      common.MoneyRub
}

func (cats *EmployeesPaymentCategories) AddItem(cat string, payment *EmployeePayment) {
	if _, hasCategory := cats.Categories[cat]; !hasCategory {
		cats.Categories[cat] = &paymentCategory{
			Name:     cat,
			Total:    common.MoneyRub(0),
			Payments: make([]*EmployeePayment, 0),
		}
	}

	cats.Categories[cat].AddPayment(payment)
	cats.Total += payment.Amount
}