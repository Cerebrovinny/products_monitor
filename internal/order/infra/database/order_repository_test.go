package database

import (
	"database/sql"
	"github.com/Cerebrovinny/products_monitor/internal/order/entity"
	"github.com/stretchr/testify/suite"
	"testing"
	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL PRIMARY KEY, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL)")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenISave_ThenIShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repository := NewOrderRepository(suite.Db)
	err = repository.Save(order)
	suite.NoError(err)

	//fill orderResult using the select and guarantee that the order is the same as the NewOrder
	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
}
