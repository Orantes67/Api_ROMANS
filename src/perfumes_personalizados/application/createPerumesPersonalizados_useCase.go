package application

import (
	"ApiProduct/src/perfumes_personalizados/domain"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
)

type CreatePerfumePersonalizado struct{
	db domain.IPerfumePersonalizado
}

func NewPerfumePersonalizado(db domain.IPerfumePersonalizado)*CreatePerfumePersonalizado{
	return &CreatePerfumePersonalizado{db:db}

}


func(cpp *CreatePerfumePersonalizado)Execute(NewPerfumePersonalizado *entities.PerfumePersonalizado)error{
	err := cpp.db.Save(NewPerfumePersonalizado)
	if err != nil{
		return err
	}
	return nil
}