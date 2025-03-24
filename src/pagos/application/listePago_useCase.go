package application

import (
	"ApiProduct/src/pagos/domain"
	"ApiProduct/src/pagos/domain/entities"
)

type ListPagos struct{
	db domain.IPagos
}

func NewListPagos(db domain.IPagos) *ListPagos{
	return &ListPagos{db: db}
}

func (vpay *ListPagos) Execute()([]entities.Pagos,error){
	res, err := vpay.db.GetAll()
	if err != nil {
		return nil, err
	}
	return res,nil
}
