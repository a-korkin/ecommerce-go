package services

import (
	"github.com/a-korkin/ecommerce/internal/core/models"
	"github.com/jmoiron/sqlx"
)

type BillService struct {
	DB *sqlx.DB
}

func NewBillService(db *sqlx.DB) *BillService {
	return &BillService{DB: db}
}

func (s *BillService) GetBillByUser(userID string) (*models.Bill, error) {
	sql := `
select uuid_generate_v4() as id, sum(o.amount * p.price) as total_sum
from public.orders as o
inner join public.products as p on p.id = o.product_id
where o.user_id = $1::uuid`
	row, err := s.DB.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		bill := models.Bill{}
		if err = row.Scan(&bill.ID, &bill.TotalSum); err != nil {
			return nil, err
		}
		return &bill, nil
	}
	return nil, nil
}
