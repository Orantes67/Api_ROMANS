package domain

import (

	"ApiProduct/src/perfumes_personalizados/domain/entities"
)

type IPerfumePersonalizado interface {
	Save(perfumes_personalizados *entities.PerfumePersonalizado)error
	GetAll()([]entities.PerfumePersonalizado,error)
	Delete(id int)error
	Update(id int,updatedPerfumePersonalizado *entities.PerfumePersonalizado)error
}