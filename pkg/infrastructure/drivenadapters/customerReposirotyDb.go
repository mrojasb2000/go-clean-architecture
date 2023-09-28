package drivenadapters

import (
	"Customers/internal/domain/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]model.Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := c.client.Query(findAllSql)
	if err != nil {
		log.Println(fmt.Sprintf("Error while querying customer table %s\n", err.Error()))
		return nil, err
	}

	customers := make([]model.Customer, 0)
	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println(fmt.Sprintf("Error while scanning customers %s\n", err.Error()))
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (c CustomerRepositoryDb) FindById(id int64) (model.Customer, error) {

	findByIdSql := fmt.Sprintf("select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = %v", id)

	customer := model.Customer{}
	rows, err := c.client.Query(findByIdSql)
	if err != nil {
		log.Println(fmt.Sprintf("Error while querying customer table %s\n", err.Error()))
		return customer, err
	}

	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println(fmt.Sprintf("Error while scanning customer %s\n", err.Error()))
			return customer, err
		}
		customer = c
	}
	return customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:lemontech@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
