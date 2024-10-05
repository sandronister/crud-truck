package mysql

import (
	"database/sql"

	"github.com/sandronister/crud-truck/internal/entity"
)

type driverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *driverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) FindAll() (*[]entity.Driver, error) {
	drivers := []entity.Driver{}

	rows, err := r.db.Query("SELECT id, name, license_id, created_at,updated_at FROM drivers WHERE deleted = 0")
	if err != nil {
		return &drivers, err
	}

	for rows.Next() {
		var driver entity.Driver
		err = rows.Scan(&driver.ID, &driver.Name, &driver.LicenseID, &driver.CreatedAt, &driver.UpdatedAt)
		if err != nil {
			return &drivers, err
		}

		drivers = append(drivers, driver)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &drivers, nil
}

func (r *driverRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE drivers set deleted=1 WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *driverRepository) ExistsLicenseID(license string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM drivers WHERE license_id = ? AND deleted = 0)`
	err := r.db.QueryRow(query, license).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *driverRepository) FindById(id string) (*entity.Driver, error) {
	var driver entity.Driver

	row := r.db.QueryRow("SELECT id, name, license_id, created_at,updated_at FROM drivers WHERE id = ? AND deleted=0", id)
	err := row.Scan(&driver.ID, &driver.Name, &driver.LicenseID, &driver.CreatedAt, &driver.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (r *driverRepository) Save(driver *entity.Driver) error {
	stmt, err := r.db.Prepare("INSERT INTO drivers(id, name, license_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(driver.ID, driver.Name, driver.LicenseID)
	if err != nil {
		return err
	}
	return nil
}

func (r *driverRepository) Update(driver *entity.Driver) error {
	_, err := r.db.Exec("UPDATE drivers SET name = ?, license_id = ? WHERE id = ?", driver.Name, driver.LicenseID, driver.ID)
	if err != nil {
		return err
	}

	return nil
}
