package application

import (
	"ApiProduct/src/pagos/domain"
	"ApiProduct/src/pagos/domain/entities"
)

type CreatePago struct{
	db domain.IPagos
}

func NewPago(db domain.IPagos) *CreatePago{
	return &CreatePago{db: db}
}

func (cpay *CreatePago)Execute(NewPago *entities.Pagos) error{
	err := cpay.db.Save(NewPago)
	if err != nil {
		return err
	}
	return nil
}