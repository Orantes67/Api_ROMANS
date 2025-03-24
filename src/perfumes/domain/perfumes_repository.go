package domain


import (
    "ApiProduct/src/perfumes/domain/entities"
)

type IPerfume interface {
    Save(perfume *entities.Perfume) error
    GetAll() ([]entities.Perfume, error)
    Delete(id int) error
    Update(id int, updatedPerfume *entities.Perfume) error
}

