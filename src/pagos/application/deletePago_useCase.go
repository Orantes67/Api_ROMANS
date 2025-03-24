package application

import (
	"ApiProduct/src/pagos/domain"

)

type DeletePagos struct{
	db domain.IPagos
}
func NewDeletePagos(db domain.IPagos) *DeletePagos{
	return &DeletePagos{db:db}
}


func (dp *DeletePagos) Execute(id int) error{
	err := dp.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}