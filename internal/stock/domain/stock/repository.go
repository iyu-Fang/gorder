package stock

import (
	"context"
	"fmt"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"strings"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("These item not found in stock, ids: %s", strings.Join(e.Missing, ","))
}
