package infraestructure

import (
	"ApiProduct/src/core"
	"ApiProduct/src/perfumes/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}
// Constructor para inicializar la conexión a MySQL
func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}
// Implementación del método Save
func (mysql *MySQL) Save(perfume *entities.Perfume) error {
	query := "INSERT INTO perfumes (name, descripcion) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, perfume.GetName(), perfume.GetDescripcion())
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume agregado con éxito")
	}
	return nil
}
// Implementación del método GetAll
func (mysql *MySQL) GetAll() ([]entities.Perfume, error) {
	query := "SELECT idPerfume, name, descripcion FROM perfumes"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var perfumes []entities.Perfume

	for rows.Next() {
		var p entities.Perfume
		if err := rows.Scan(&p.ID, &p.Name, &p.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		perfumes = append(perfumes, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return perfumes, nil
}

// Implementación del método Update
func (mysql *MySQL) Update(id int, updatedPerfume *entities.Perfume) error {
	query := "UPDATE perfumes SET name = ?, descripcion = ? WHERE idPerfume = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, updatedPerfume.GetName(), updatedPerfume.GetDescripcion(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume con ID %d actualizado correctamente", id)
	} else {
		log.Printf("[MySQL] - No se encontró el perfume con ID %d", id)
	}
	return nil
}

// Implementación del método Delete
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM perfumes WHERE idPerfume = ?"
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

