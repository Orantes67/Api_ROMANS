package domain

import (
	"ApiProduct/src/pagos/domain/entities"
)

type IPagos interface{
	Save(pagos *entities.Pagos)error
	GetAll()([]entities.Pagos,error)
	Delete(id int)error
	Update(id int,updatedPago *entities.Pagos) error
}