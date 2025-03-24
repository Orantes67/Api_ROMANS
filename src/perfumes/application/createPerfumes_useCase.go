package application

import (
	"ApiProduct/src/perfumes/domain"
	"ApiProduct/src/perfumes/domain/entities"
)

type CreatePerfume struct{
	db domain.IPerfume
}

func NewPerfume(db domain.IPerfume) *CreatePerfume{
	return &CreatePerfume{db:db}
}

func (cp *CreatePerfume) Execute(NewPerfume *entities.Perfume) error{
	err := cp.db.Save(NewPerfume)
	if err != nil {
		return err
	}
	return nil
}