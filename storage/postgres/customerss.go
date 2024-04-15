package postgres

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/pkg"
	"cars_with_sql/pkg/check"
	"cars_with_sql/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type customerRepo struct {
	db     *pgxpool.Pool
	logger logger.ILogger
}

func Newcustomer(db *pgxpool.Pool, log logger.ILogger) customerRepo {
	return customerRepo{
		db:     db,
		logger: log,
	}
}

// func (c *customerRepo) CreateCus(ctx context.Context, customers models.Customers) (string, error) {
// 	// Validate the password
// 	if err := check.ValidatePassword(customers.Password); err != nil {
// 		return "", err
// 	}

// 	// Generate hashed password
// 	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(customers.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.logger.Error("failed to generate customer new password", logger.Error(err))
// 		return "", err
// 	}

// 	// Generate new UUID for customer ID
// 	id := uuid.New()

// 	// SQL query to insert new customer
// 	query := `INSERT INTO customerss (
//         id,
//         first_name,
//         last_name,
//         gmail,
//         phone,
//         is_blocked,
//         password,
//         login
//     ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// 	// Execute the query
// 	_, err := c.db.Exec(context.Background(), query,
// 		id.String(),
// 		customers.First_name,
// 		customers.Last_name,
// 		customers.Gmail,
// 		customers.Phone,
// 		customers.Is_blocked,
// 		newHashedPassword,
// 		customers.Login)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.logger.Error("you have error while creating :", logger.Error(err))
// 		return "", err
// 	}

// 	return id.String(), nil
// }

func (c *customerRepo) CreateCus(ctx context.Context, customers models.Customers) (string, error) {
	// Validate the password
	if err := check.ValidatePassword(customers.Password); err != nil {
		fmt.Println(customers.Gmail)

		return "", err
	}
	fmt.Println(customers.Gmail)
	if err := check.Validategmail(customers.Gmail); err != nil {
		return "", err
	}

	// if err := check.Validatenumber(customers.Phone); err != nil {
	// 	return "", err
	// }

	// Generate hashed password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(customers.Password), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Error("failed to generate customer new password", logger.Error(err))
		return "", err
	}

	// Generate new UUID for customer ID
	id := uuid.New()

	// SQL query to insert new customer
	query := `INSERT INTO customerss (
        id,
        first_name,
        last_name,
        gmail,
        phone,
        is_blocked,
        password,
        login
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Execute the query
	_, err = c.db.Exec(context.Background(), query,
		id.String(),
		customers.First_name,
		customers.Last_name,
		customers.Gmail,
		customers.Phone,
		customers.Is_blocked,
		newHashedPassword,
		customers.Login)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		c.logger.Error("you have error while creating :", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

func (c *customerRepo) UpdateCustomer(ctc context.Context, customer models.Customers) (string, error) {

	fmt.Println(customer.Gmail)
	if err := check.Validategmail(customer.Gmail); err != nil {
		return "", err
	}
	queryy := `UPDATE customerss set
            first_name=$1,
            last_name=$2,
            gmail=$3,
            phone=$4,
            is_blocked=$5,
            updated_at=CURRENT_TIMESTAMP,
			id=$6
        WHERE id=$6 AND deleted_at=0
    `

	_, err := c.db.Exec(context.Background(), queryy,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone,
		customer.Is_blocked, customer.Id)
	if err != nil {
		c.logger.Error("Error while updating customer:", logger.Error(err))
		return "", err
	}

	return customer.Id, nil
}
func (c *customerRepo) GetAllCustomers(ctx context.Context, req models.GetAllCustomerRequest) (models.GetAllCustomersResponse, error) {
	resp := models.GetAllCustomersResponse{}
	filter := ""

	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND first_name ILIKE '%%%v%%' `, req.Search)
	}

	fmt.Println("filter: ", filter)

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)
	query := `
        SELECT 
		count(id) OVER(),
            id,
            first_name,
            last_name,
            gmail,
            phone,
            is_blocked
        FROM 
            customerss 
        WHERE 
            deleted_at = 0` + filter

	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customers
		var order models.GetOrder
		var orders models.GetAllOrdersResponse
		var updatedAt sql.NullString
		var created_at sql.NullString

		err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.First_name,
			&customer.Last_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_blocked,
		)
		if err != nil {
			c.logger.Error("error while scanning customer info: ", logger.Error(err))
			return resp, err
		}

		query := `
            SELECT 
			count(id) OVER(),
			id,
                from_date,
                to_date,
                status,
                payment_status,
              created_at,
                updated_at
            FROM 
                orrders 
            WHERE 
                customer_id = $1`

		orderRows, err := c.db.Query(context.Background(), query, customer.Id)
		if err != nil {
			c.logger.Error("error while querying orders: ", logger.Error(err))
			return resp, err
		}
		defer orderRows.Close()

		for orderRows.Next() {
			var fromDate, toDate time.Time
			err := orderRows.Scan(
				&orders.Count,
				&order.Id,
				&fromDate,
				&toDate,
				&order.Status,
				&order.Paid,
				&created_at,
				&updatedAt,
			)
			if err != nil {
				c.logger.Error("error while scanning order info: ", logger.Error(err))
				return resp, err
			}
			order.FromDate = fromDate.Format("2006-01-02 15:04:05")
			order.ToDate = toDate.Format("2006-01-02 15:04:05")
			order.Updated_at = pkg.NullStringToString(updatedAt)
			order.Created_at = pkg.NullStringToString(created_at)

			customer.Order = append(customer.Order, order)
		}

		if err := orderRows.Err(); err != nil {
			c.logger.Error("error while iterating over order rows: ", logger.Error(err))
			return resp, err
		}

		resp.Customers = append(resp.Customers, customer)
	}
	if err := rows.Err(); err != nil {
		c.logger.Error("error while iterating over rows: ", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (c *customerRepo) GetByIDCustomer(ctcx context.Context, id string) (models.Customers, error) {

	query := `SELECT 
        id,
        first_name,
        last_name,
        gmail,
        phone,
        is_blocked
        FROM customerss WHERE id=$1 AND deleted_at=0`

	row := c.db.QueryRow(context.Background(), query, id)

	customer := models.Customers{}

	err := row.Scan(

		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone,
		&customer.Is_blocked)
	if err != nil {
		fmt.Println(err, "_______________________")
		c.logger.Error("error while getting id customer err: ", logger.Error(err))
		return customer, err
	}

	query = `SELECT 
        id,
        from_date,
        to_date,
        status,
        payment_status,
        created_at,
        updated_at
        FROM orrders WHERE customer_id=$1`

	rows, err := c.db.Query(context.Background(), query, id)
	if err != nil {
		return models.Customers{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.GetOrder
		var updatedAt sql.NullString
		err := rows.Scan(

			&order.Id,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Created_at,
			&updatedAt,
		)
		if err != nil {
			c.logger.Error("error while scanning getbyid: ", logger.Error(err))
			return customer, err
		}
		order.Updated_at = pkg.NullStringToString(updatedAt)
		customer.Order = append(customer.Order, order)
	}
	if err := rows.Err(); err != nil {
		c.logger.Error("error while iterating over rows: ", logger.Error(err))
		return customer, err
	}

	return customer, nil
}

func (c *customerRepo) DeleteCustomer(ctx context.Context, id string) (string, error) {

	query := `UPDATE customerss set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(context.Background(), query, id)

	if err != nil {
		c.logger.Error("failed to  delete customer from database", logger.Error(err))

		return id, err
	}

	return id, nil

}
func (c *customerRepo) GetAllCustomerCars(ctx context.Context, req models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	var (
		resp   = models.GetAllCustomerCarsResponse{}
		filter = ""
	)

	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(` AND ca.name ILIKE '%%%v%%'`, req.Search)
	}

	if req.Id != "" {
		filter += fmt.Sprintf(` AND cu.id = '%v'`, req.Id)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	fmt.Println("filter:", filter)

	query := `
	SELECT 
		cu.id as customer_id,
		ca.name as car_name,
		o.amount as price
	FROM 
		customerss cu 
	JOIN 
		orrders o ON cu.id = o.customer_id
	JOIN 
		cars ca ON ca.id = o.cars_id
	`

	rows, err := c.db.Query(context.Background(), query, filter+``)
	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var customer models.GetAllCustomerCars
		if err := rows.Scan(
			&customer.CustomerID,
			&customer.CarName,
			&customer.Price,
		); err != nil {
			return resp, err
		}
		resp.Customer = append(resp.Customer, customer)
	}

	if err = rows.Err(); err != nil {
		c.logger.Error("failed to scan customer cars from database", logger.Error(err))

		return resp, err
	}

	countQuery := `SELECT COUNT(*) FROM customerss cu JOIN orrders o ON  cu.id = o.customer_id Join cars ca  ON ca.id=o.cars_id;
	` + filter
	err = c.db.QueryRow(context.Background(), countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *customerRepo) Login(ctx context.Context, login models.Changepasswor) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE login = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query,
		login.Phone,
	).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect login")
		}
		c.logger.Error("failed to get customer password from database", logger.Error(err))
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(login.OldPassword))
	if err != nil {
		return "", errors.New("password mismatch")
	}

	return "Logged in successfully", nil
}

func (c *customerRepo) GetPassword(ctx context.Context, phone string) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE phone = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query, phone).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect phone")
		} else {
			c.logger.Error("failed to get customer password from database", logger.Error(err))
			return "", err
		}
	}

	return hashedPass, nil
}

// func (c *customerRepo) ChangePassword(ctx context.Context, password models.Changepasswor) (string, error) {
// 	var hashedPass string

// 	query := `SELECT password
// 	FROM customerss
// 	WHERE login = $1 AND deleted_at = 0`

// 	err := c.db.QueryRow(ctx, query,
// 		password.Login,
// 	).Scan(&hashedPass)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return "", errors.New("incorrect login")
// 		}
// 		fmt.Println(err)
// 		c.logger.Error("failed to get customer password from database", logger.Error(err))
// 		return "", err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password.Password))
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", errors.New("password mismatch")
// 	}

// 	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.logger.Error("failed to generate customer new password", logger.Error(err))
// 		return "", err
// 	}

// 	query = `UPDATE customers SET
// 		password = $1,
// 		updated_at = CURRENT_TIMESTAMP
// 	WHERE login = $2 AND deleted_at = 0`

// 	_, err = c.db.Exec(ctx, query, newHashedPassword, password.Login)
// 	if err != nil {
// 		c.logger.Error("failed to change customer password in database", logger.Error(err))
// 		return "", err
// 	}

// return "Password changed successfully", nil

func (c *customerRepo) ChangePassword(ctx context.Context, password models.Changepasswor) (string, error) {
	// Validate the new password
	if err := check.ValidatePassword(password.NewPassword); err != nil {
		fmt.Println(err, "ttttttttt")
		return "", err

	}

	query := `SELECT 
        password
        FROM customerss WHERE login=$1 AND deleted_at=0`

	row := c.db.QueryRow(context.Background(), query, password.Login)

	customer := models.Changepasswor{}

	err := row.Scan(
		&customer.OldPassword,
	)
	if err != nil {
		fmt.Println(err, "_===")
		c.logger.Error("error while getting customer's old password: ", logger.Error(err))
		return customer.OldPassword, err
	}

	// Compare the hash of the old password stored in the database with the provided old password

	// Generate a new hashed password for the new password provided
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Error("failed to generate new hashed password", logger.Error(err))
		return "", err
	}

	fmt.Println(newHashedPassword) // Print hashed password only when there's no error

	err = bcrypt.CompareHashAndPassword([]byte(newHashedPassword), []byte(password.OldPassword))
	if err == nil {
		fmt.Println(err, "")
		fmt.Println(newHashedPassword)
		return "", errors.New("password mismatch")
	}

	if len(password.NewPassword) < 8 {
		return "", errors.New("new password must be at least 8 characters long")
	}
	// Update the password in the database
	query = `UPDATE customerss SET 
        password = $1, 
        updated_at = CURRENT_TIMESTAMP 
    WHERE login = $2 AND deleted_at = 0`

	_, err = c.db.Exec(ctx, query, newHashedPassword, password.Login)
	if err != nil {
		fmt.Println(err, "===================")
		c.logger.Error("failed to change customer password in database", logger.Error(err))
		return "", err
	}

	return "Password changed successfully", nil
}

func (c *customerRepo) GetByLogin(ctx context.Context, login string) (models.Customers, error) {
	var (
		firstname sql.NullString
		lastname  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
	)

	query := `SELECT 
		id, 
		first_name, 
		last_name, 
		phone,
		email,
		created_at, 
		updated_at,
		password
		FROM customerss WHERE phone = $1 AND deleted_at = 0`

	row := c.db.QueryRow(ctx, query, login)

	customer := models.Customers{
		Order: []models.GetOrder{},
	}

	err := row.Scan(
		&customer.Id,
		&firstname,
		&lastname,
		&phone,
		&email,
		&createdat,
		&updatedat,
		&customer.Password,
	)

	if err != nil {
		c.logger.Error("failed to scan customer by LOGIN from database", logger.Error(err))
		return models.Customers{}, err
	}

	customer.First_name = firstname.String
	customer.Last_name = lastname.String
	customer.Phone = phone.String
	customer.Gmail = email.String
	customer.Created_at = createdat.String
	customer.Updated_at = updatedat.String

	return customer, nil
}

func (c *customerRepo) Checklogin(ctx context.Context, gmail string) (models.CustomerRegisterRequest, error) {

	fmt.Println("gmail.mail:", gmail)
	query := `SELECT 
        gmail
       FROM customerss WHERE gmail=$1 AND deleted_at=0`

	row := c.db.QueryRow(ctx, query, gmail)

	customer := models.CustomerRegisterRequest{}

	err := row.Scan(
		&customer.Mail,
	)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
