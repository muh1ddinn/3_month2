package postgres

import (
	models "cars_with_sql/api/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
)

type carRepo struct {
	DATA *pgxpool.Pool
}

func Newwcar(DATA *pgxpool.Pool) carRepo {
	return carRepo{
		DATA: DATA,
	}
}

/*
======================================
create (body) id ,err
update(body) id,err
delete(id) err
get(id) body,err
getALL(serach) []body,count,err

======================================
*/
func (c *carRepo) Create(ctx context.Context, car models.Car) (string, error) {

	id := uuid.New()

	query := `INSERT INTO cars(
        id,
        name,
        brand,
        model,
		year,
        hourse_power,
        colour,
        engine_cap)
    VALUES($1,$2,$3,$4,$5,$6,$7,$8) `

	res, err := c.DATA.Exec(context.Background(), query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.Year,
		car.HorsePower, car.Colour, car.EngineCap)

	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)

	}
	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

// //////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////

func (c carRepo) GetByidcar(ctx context.Context, id string) (models.Car, error) {
	car := models.Car{}
	err := c.DATA.QueryRow(context.Background(), `SELECT 
    id,
    name,
    brand,
    model,
    year,
    hourse_power,
    colour,
    engine_cap FROM cars
    where id=$1`, id).Scan(&car.Id, &car.Name, &car.Brand, &car.Model, &car.Year, &car.HorsePower, &car.Colour, &car.EngineCap)

	if err != nil {
		fmt.Println("error while scaning rows: ", err)
		return car, err
	}

	return car, nil
}

// func (c carRepo) GetByIdCar(id string) (models.Car, error) {
// 	car := models.Car{}

// 	row := c.DATA.QueryRow(context.Background(), `
//         SELECT
//             id,
//             name,
//             brand,
//             model,
//             year,
//             hourse_power,
//             colour,
//             engine_cap
//         FROM cars
//         WHERE id=$1
//     `, id)

// 	if err := row.Scan(&car.Id, &car.Name, &car.Brand, &car.Model, &car.Year, &car.HorsePower, &car.Colour, &car.EngineCap); err != nil {
// 		fmt.Println("error while getting car by ID: ", err)
// 		if err == sql.ErrNoRows {
// 			return car, fmt.Errorf("car with ID %s not found", id)
// 		}
// 		return car, err
// 	}

// 	return car, nil
// }

////////////////////////////////////////////////////////////

func (c carRepo) GetAllCarS(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.DATA.Query(
		context.Background(),
		`select
				count(id) OVER(),
				id,
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 `+filter)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			id           sql.NullString
			name         sql.NullString
			brand        sql.NullString
			model        sql.NullString
			year         sql.NullInt64
			hourse_power sql.NullInt64
			colour       sql.NullString
			engine_cap   sql.NullFloat64
			created_at   sql.NullString
			updateAt     sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&brand,
			&model,
			&year,
			&hourse_power,
			&colour,
			&engine_cap,
			&created_at,
			&updateAt); err != nil {
			return resp, err
		}
		resp.Cars = append(resp.Cars, models.Car{
			Id:         id.String,
			Name:       name.String,
			Brand:      brand.String,
			Model:      model.String,
			Year:       int(year.Int64),
			HorsePower: int(hourse_power.Int64),
			Colour:     colour.String,
			EngineCap:  float32(engine_cap.Float64),
			CreatedAt:  created_at.String,
			UpdatedAt:  updateAt.String,
		})
	}

	return resp, nil
}

func (c *carRepo) Deletecar(ctx context.Context, id string) (string, error) {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.DATA.Exec(context.Background(), query, id)

	if err != nil {
		return id, err
	}

	return id, nil

}

func (c *carRepo) UpdateCar(ctx context.Context, car models.UpdateCarRequest) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at=0
	`

	_, err := c.DATA.Exec(context.Background(), query,
		car.Name, car.Brand,
		car.Model, car.HorsePower,
		car.Colour, car.EngineCap, car.ID)

	if err != nil {
		return "", err
	}

	return car.ID, nil
}
