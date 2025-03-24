package application

import (
	"ApiProduct/src/perfumes/domain"
	"ApiProduct/src/perfumes/domain/entities"
)

type UpdatePerfume struct{
	db domain.IPerfume
}

func NewUpdatePerfume(db domain.IPerfume) *UpdatePerfume{
	return &UpdatePerfume{db: db}
}


func (up *UpdatePerfume) Execute(id int, updatedPerfumes *entities.Perfume) error {
    err := up.db.Update(id, updatedPerfumes)
    if err != nil {
        return err
    }
    return nil
}