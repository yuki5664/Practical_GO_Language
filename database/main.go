package database

import (
	"context"
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Hollow World")
}

// 一般的なトランザクション
type Service struct {
	db *sql.DB
}

func (s *Service) UpdateProduct(ctx context.Context, productID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// deferを活用したトランザクション
func (s *Service) UpdateProduct(ctx context.Context, productID string) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx, "...", productID); err != nil {
		return err
	}

	return tx.Commit()
}
