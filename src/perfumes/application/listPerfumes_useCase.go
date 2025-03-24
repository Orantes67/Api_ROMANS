package application

import (
	"ApiProduct/src/perfumes/domain"
	"ApiProduct/src/perfumes/domain/entities"
)

type ListPerfume struct{
	db domain.IPerfume
}

func NewListPerfume(db domain.IPerfume) *ListPerfume{
	return &ListPerfume{db:db}
}

func (vp *ListPerfume) Execute() ([]entities.Perfume, error) {
	res, err := vp.db.GetAll()
	if err != nil {
		return nil, err // Si hay error, devolvemos nil como lista y pasamos el error
	}
	return res, nil // Si todo está bien, devolvemos la lista y nil como error
}