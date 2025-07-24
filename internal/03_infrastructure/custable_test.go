// internal\03_infrastructure\custable_test.go

package infrastructure

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetRecentActiveCustomers_Success(t *testing.T) {
	// 1. Mock DB, sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := &Queries{
		GetRecentActiveCustomers: "SELECT * FROM customers WHERE ...", // หรืออะไรก็ได้
	}

	// 2. Mock rows return
	rows := sqlmock.NewRows([]string{"BillerID", "AccountNum", "CompanyBank", "Name", "CustomerStatus", "CreatedDateTime"}).
		AddRow("001", "123456", "SCB", "Alice", 1, time.Now()).
		AddRow("002", "987654", "KBank", "Bob", 0, time.Now())

	mock.ExpectQuery("SELECT \\* FROM customers"). // regex match
		WillReturnRows(rows)

	repo := NewCustableRepository(db, queries)

	// 3. Run
	result, err := repo.GetRecentActiveCustomers(context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Alice", result[0].Name)
	assert.Equal(t, "Bob", result[1].Name)

	// 4. Check all expectations met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetRecentActiveCustomers_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := &Queries{
		GetRecentActiveCustomers: "SELECT * FROM customers WHERE ...",
	}

	mock.ExpectQuery("SELECT \\* FROM customers").
		WillReturnError(sql.ErrConnDone)

	repo := NewCustableRepository(db, queries)
	result, err := repo.GetRecentActiveCustomers(context.Background())
	assert.Nil(t, result)
	assert.EqualError(t, err, sql.ErrConnDone.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetRecentActiveCustomers_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := &Queries{
		GetRecentActiveCustomers: "SELECT * FROM customers WHERE ...",
	}

	// Mock row มี datatype ผิด จะทำให้ Scan error
	rows := sqlmock.NewRows([]string{"BillerID", "AccountNum", "CompanyBank", "Name", "CustomerStatus", "CreatedDateTime"}).
		AddRow("001", "123456", "SCB", "Alice", "NOT_INT", time.Now()) // CustomerStatus เป็น int, ใส่ string เพื่อ error

	mock.ExpectQuery("SELECT \\* FROM customers").
		WillReturnRows(rows)

	repo := NewCustableRepository(db, queries)
	result, err := repo.GetRecentActiveCustomers(context.Background())
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
