package application

import (
	"ApiProduct/src/pagos/domain"
	"ApiProduct/src/pagos/domain/entities"
)

type UpdatePago struct{
	db domain.IPagos
}

func NewUpdatePago(db domain.IPagos) *UpdatePago{
	return &UpdatePago{db:db}
}
func (upay *UpdatePago) Execute(id int,updatedPagos *entities.Pagos)error{
	err := upay.db.Update(id,updatedPagos)
	if err != nil {
		return err
	}
	return nil
}