package application

import (
	"ApiProduct/src/perfumes/domain"
)

type DeletePerfume struct{
	db domain.IPerfume
}

func NewDeletePerfume(db domain.IPerfume) *DeletePerfume{
	return &DeletePerfume{db:db}
}

func (dp *DeletePerfume) Execute(id int) error{
	err := dp.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}