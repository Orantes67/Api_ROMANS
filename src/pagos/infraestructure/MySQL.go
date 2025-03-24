package infraestructure


import (
	"ApiProduct/src/core"
	"ApiProduct/src/pagos/domain/entities"
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
func (mysql *MySQL) Save(pagos *entities.Pagos) error {
	query := "INSERT INTO notificaciones_de_pagos (descripcion) VALUES ( ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query,pagos.GetDescripcion())
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
func (mysql *MySQL) GetAll() ([]entities.Pagos, error) {
	query := "SELECT idNotification, descripcion FROM notificaciones_de_pagos"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var pagos []entities.Pagos

	for rows.Next() {
		var p entities.Pagos
		if err := rows.Scan(&p.ID,&p.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		pagos = append(pagos, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return pagos, nil
}

// Implementación del método Update
func (mysql *MySQL) Update(id int, updatedPagos *entities.Pagos) error {
	query := "UPDATE notificaciones_de_pagos SET  descripcion = ? WHERE idNotification = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, updatedPagos.GetDescripcion(), id)
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
	query := "DELETE FROM notificaciones_de_pagos WHERE idNotification = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Pagos con ID %d eliminado correctamente", id)
	} else {
		log.Printf("[MySQL] - No se encontró el perfume con ID %d", id)
	}
	return nil
}

