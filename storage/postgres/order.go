package postgres

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrder(db *pgxpool.Pool) orderRepo {
	return orderRepo{
		db: db,
	}
}
func (c *orderRepo) Create(ctx context.Context, order models.CreateOrder) (string, error) {
	id := uuid.New()

	query := `INSERT INTO orrders (
        id,
        customer_id, 
        cars_id, 
        from_date, 
        to_date,
        status,
        payment_status,
        amount)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := c.db.Exec(ctx, query, // Use the passed-in ctx here
		id.String(),
		order.CustomerId,
		order.CarId,
		order.FromDate,
		order.ToDate,
		order.Status,
		order.Paid,
		order.Amount)

	if err != nil {
		fmt.Println("you have error while creating :", err)
		return "", err
	}

	return id.String(), nil
}

func (o *orderRepo) GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	var (
		resp   = models.GetAllOrdersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(`and  status ILIKE '%%%v%%'`, req.Search)
	}
	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query := `Select 
		o.id,
		o.from_date,
		o.to_date,
		o.status,
		o.amount,
		o.created_at,
		o.updated_at,
		c.id as car_id,
		c.name as car_name,
		c.brand as car_brand,
		c.engine_cap as car_engine_cap,
		cu.id as customer_id,
		cu.first_name as customer_first_name,
		cu.last_name as customer_last_name,
		cu.gmail as customer_gmail,
		cu.phone as customer_phone
		From orders o JOIN cars c ON o.car_id = c.id
		JOIN customers cu ON o.customer_id = cu.id 	`
	rows, err := o.db.Query(context.Background(), query+filter+``)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			order = models.GetOrder{
				Car:       models.Car{},
				Customers: models.Customers{},
			}
			updateAt sql.NullString
		)

		err := rows.Scan(
			&order.Id,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Amount,
			&order.Created_at,
			&updateAt,
			&order.Car.Id,
			&order.Car.Name,
			&order.Car.Brand,
			&order.Car.EngineCap,
			&order.Customers.Id,
			&order.Customers.First_name,
			&order.Customers.Last_name,
			&order.Customers.Gmail, &order.Customers.Phone)
		if err != nil {
			return resp, err
		}
		order.Updated_at = pkg.NullStringToString(updateAt)
		resp.Orders = append(resp.Orders, order)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}
	countQuery := `SELECT COUNT(*) FROM orders`

	err = o.db.QueryRow(context.Background(), countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (o *orderRepo) GetByID(ctx context.Context, id string) (models.GetOrder, error) {
	response := models.GetOrder{}

	query := `SELECT 
        o.id,
        c.name AS car_name,
        c.brand AS car_brand,
        cu.id AS customer_id,
        cu.first_name AS customer_first_name,
        cu.gmail AS customer_gmail,
        o.from_date,
        o.to_date,
        o.status,
        o.payment_status,
        o.created_at,
        o.updated_at
        FROM orrders o
        JOIN cars c ON o.cars_id=c.id
        JOIN customerss cu ON o.customer_id = cu.id
        WHERE o.id = $1`

	row := o.db.QueryRow(ctx, query, id)

	var order models.GetOrder
	var car models.Car
	var customer models.Customers
	var create, update sql.NullString
	var FromDate, ToDate time.Time

	err := row.Scan(
		&order.Id,
		&car.Name,
		&car.Brand,
		&customer.Id,
		&customer.First_name,
		&customer.Gmail,
		&FromDate,
		&ToDate,
		&order.Status,
		&order.Paid,
		&create,
		&update,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, fmt.Errorf("order with ID %s not found", id)
		}
		return response, err
	}

	order.Created_at = pkg.NullStringToString(create)
	order.Updated_at = pkg.NullStringToString(update)

	order.FromDate = FromDate.Format("2006-01-02")
	order.ToDate = ToDate.Format("2006-01-02")

	order.Car = car
	order.Customers = customer

	response = order

	return response, nil
}

func (c *customerRepo) Delete(ctx context.Context, id string) (string, error) {

	query := `UPDATE orrders set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(context.Background(), query, id)

	if err != nil {
		return id, err
	}

	return id, nil

}

func (c *orderRepo) UpdateOrders(ctx context.Context, order models.CreateOrder) (string, error) {
	queryy := `UPDATE orrders set
	        customer_id=$1,
            cars_id=$2,
            from_date=$3,
            to_date=$4,
			status=$5,
			payment_status=$6
			amount=$7
            updated_at=CURRENT_TIMESTAMP,
			id=8$
        WHERE id=8$ AND deleted_at=0
    `

	_, err := c.db.Exec(context.Background(), queryy,
		order.CustomerId, order.CarId,
		order.FromDate, order.ToDate,
		order.Status, order.Paid, order.Amount, order.Id)
	if err != nil {
		fmt.Println("Error while updating customer:", err)
		return "", err
	}

	return order.Id, nil
}
