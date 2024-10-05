package mysql

import (
	"database/sql"

	"github.com/sandronister/crud-truck/internal/entity"
)

type truckRepository struct {
	db *sql.DB
}

func NewTruckRepository(db *sql.DB) *truckRepository {
	return &truckRepository{
		db: db,
	}
}

func (r *truckRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE truck set deleted=1 WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *truckRepository) ExistsLicensePlate(licensePlate string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM trucks WHERE license_plate = ? and deleted = 0)`
	err := r.db.QueryRow(query, licensePlate).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *truckRepository) FindAll() (*[]entity.Truck, error) {
	rows, err := r.db.Query("SELECT id,brand,model,year,license_plate FROM trucks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trucks []entity.Truck
	for rows.Next() {
		var t entity.Truck
		err := rows.Scan(&t.ID, &t.Brand, &t.Model, &t.Year, &t.LicensePlate)
		if err != nil {
			return nil, err
		}

		trucks = append(trucks, t)
	}

	return &trucks, nil
}

func (r *truckRepository) FindById(id string) (*entity.Truck, error) {
	var truck entity.Truck

	row := r.db.QueryRow("SELECT id, brand, model, year, license_plate FROM trucks WHERE id = ?", id)
	err := row.Scan(&truck.ID, &truck.Brand, &truck.Model, &truck.Year, &truck.LicensePlate)
	if err != nil {
		return nil, err
	}

	return &truck, nil
}

func (r *truckRepository) Save(truck *entity.Truck) error {
	stmt, err := r.db.Prepare("INSERT INTO trucks (id, brand, model, year,license_plate) VALUES (?, ?, ?, ?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(truck.ID, truck.Brand, truck.Model, truck.Year, truck.LicensePlate)
	if err != nil {
		return err
	}

	return nil
}

func (r *truckRepository) Update(truck *entity.Truck) error {
	stmt, err := r.db.Prepare("UPDATE trucks SET brand = ?, model = ?, year = ?, license_plate = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(truck.Brand, truck.Model, truck.Year, truck.LicensePlate, truck.ID)
	if err != nil {
		return err
	}

	return nil
}
