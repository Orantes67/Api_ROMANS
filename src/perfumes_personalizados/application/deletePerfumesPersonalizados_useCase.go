package application

import (
	"ApiProduct/src/perfumes_personalizados/domain"
)

type DeletePerfumePersonalizado struct{
	db domain.IPerfumePersonalizado
}

func NewDeletePerfumePersonalizado(db domain.IPerfumePersonalizado) *DeletePerfumePersonalizado{
	return &DeletePerfumePersonalizado{db: db}
}


func (dpp *DeletePerfumePersonalizado) Execute(id int)error{
	err := dpp.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}