package mysql

import (
	"database/sql"

	"github.com/sandronister/crud-truck/internal/entity"
)

type linkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *linkRepository {
	return &linkRepository{db: db}
}

func (r *linkRepository) Delete(driverId, truckID string) error {
	_, err := r.db.Exec("UPDATE link set deleted=1 WHERE driver_id = ? AND truck_id = ?", driverId, truckID)
	if err != nil {
		return err
	}

	return nil
}

func (r *linkRepository) ExistsLink(driverID, truckID string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM links WHERE driver_id = ? AND truck_id = ? and deleted=0)`

	err := r.db.QueryRow(query, driverID, truckID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *linkRepository) GetTrucksByDriver(id string) (*entity.DriverTrucks, error) {
	rows, err := r.db.Query("SELECT "+
		"D.id AS DriverID, "+
		"D.name AS DriverName, "+
		"D.license_id AS LicenseID, "+
		"T.id AS truck_id, "+
		"T.Brand AS brand, "+
		"T.Model as model, "+
		"T.Year as year, "+
		"T.license_plate as licensePlate, "+
		"T.created_at as created_at, "+
		"T.updated_at as updated_at "+
		"FROM "+
		"links AS L INNER JOIN "+
		"drivers AS D ON L.driver_id=D.id AND D.deleted=0 INNER JOIN "+
		"trucks AS T ON L.truck_id=T.id AND T.deleted=0 "+
		"WHERE L.driver_id=?", id)

	if err != nil {
		return nil, err
	}

	var (
		Trucks                          []entity.Truck
		DriverId, DriverName, LicenseID string
	)

	for rows.Next() {
		var truck entity.Truck
		err := rows.Scan(&DriverId, &DriverName, &LicenseID, &truck.ID, &truck.Brand, &truck.Model, &truck.Year, &truck.LicensePlate, &truck.CreatedAt, &truck.UpdatedAt)

		if err != nil {
			return nil, err
		}

		Trucks = append(Trucks, truck)
	}

	return entity.New(DriverId, DriverName, LicenseID, Trucks), nil

}

func (r *linkRepository) Save(link *entity.Link) error {
	query := `INSERT INTO links (id, driver_id, truck_id) VALUES (?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(link.ID, link.DriverID, link.TruckID)
	if err != nil {
		return err
	}

	return nil
}
