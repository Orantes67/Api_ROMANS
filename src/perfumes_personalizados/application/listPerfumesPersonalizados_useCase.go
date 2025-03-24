package application

import (
	"ApiProduct/src/perfumes_personalizados/domain"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
)

type ListPerfumePersonalizados struct{
	db domain.IPerfumePersonalizado
}

func NewListPerfumePersonalizado(db domain.IPerfumePersonalizado) *ListPerfumePersonalizados{
	return &ListPerfumePersonalizados{db:db}
}

func (vpp *ListPerfumePersonalizados) Execute()([]entities.PerfumePersonalizado,error){
	res,err := vpp.db.GetAll()
	if  err != nil {
		return nil,err
	}
	return res,nil
}