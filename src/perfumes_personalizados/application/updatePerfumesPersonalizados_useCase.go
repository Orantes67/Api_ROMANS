package application


import (
	"ApiProduct/src/perfumes_personalizados/domain"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
)

type UpdatePerfumePersonalizados struct{
	db domain.IPerfumePersonalizado
}

func NewUpdatePerfumePersonalizado(db domain.IPerfumePersonalizado) *UpdatePerfumePersonalizados{
	return &UpdatePerfumePersonalizados{db: db}
}

func (up UpdatePerfumePersonalizados) Execute(id int,updatePerfumesPersonalizado *entities.PerfumePersonalizado)error{
	err := up.db.Update(id,updatePerfumesPersonalizado)
	if err !=nil{
		return err
	}
	return nil
}