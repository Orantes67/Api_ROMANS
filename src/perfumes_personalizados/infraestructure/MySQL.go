package infraestructure

import (
	"ApiProduct/src/core"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(perfumepersonalizado *entities.PerfumePersonalizado) error {
	query := "INSERT INTO perfumes_personalizados (nombre, descripcion, ingredientes, fecha_creacion) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query,
		perfumepersonalizado.GetName(),
		perfumepersonalizado.GetDescripcion(),
		perfumepersonalizado.GetIngredientes(),
		perfumepersonalizado.GetFecha_creacion(),
	)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume agregado con éxito")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.PerfumePersonalizado, error) {
	query := "SELECT idPersonalizado, nombre, descripcion, ingredientes, fecha_creacion FROM perfumes_personalizados"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var perfumes_personalizados []entities.PerfumePersonalizado

	for rows.Next() {
		var pp entities.PerfumePersonalizado
		if err := rows.Scan(&pp.ID, &pp.Name, &pp.Descripcion, &pp.Ingredientes, &pp.Fecha_creacion); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		perfumes_personalizados = append(perfumes_personalizados, pp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return perfumes_personalizados, nil
}

func (mysql *MySQL) Update(id int, updatedPerfumePersonalizado *entities.PerfumePersonalizado) error {
	query := "UPDATE perfumes_personalizados SET nombre = ?, descripcion = ?, ingredientes = ? WHERE idPersonalizado = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query,
		updatedPerfumePersonalizado.GetName(),
		updatedPerfumePersonalizado.GetDescripcion(),
		updatedPerfumePersonalizado.GetIngredientes(),
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume Personalizado con ID %d actualizado correctamente", id)
	} else {
		log.Printf("[MySQL] - No se encontró el perfume con ID %d", id)
	}
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM perfumes_personalizados WHERE idPersonalizado = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume con ID %d eliminado correctamente", id)
	} else {
		log.Printf("[MySQL] - No se encontró el perfume con ID %d", id)
	}
	return nil
}
