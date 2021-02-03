package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uui "github.com/satori/go.uuid"
)

type Bank struct {
	Base     `valid:required`
	Code     string     `json:"code"`
	Name     string     `json:"name"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {

	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uui.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
