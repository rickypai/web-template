package repo

import (
	"context"
	"fmt"
	"time"

	dbModel "github.com/rickypai/web-template/api/dbmodels/phone"
)

type Writer struct {
	db dbModel.Querier
}

func NewWriter(db dbModel.DBTX) *Writer {
	return &Writer{
		db: dbModel.New(db),
	}
}

func (r *Writer) CreateOne(ctx context.Context, rpcInstance *createReqT) (*modelT, error) {
	now := time.Now()

	dbResult, err := r.db.CreateOne(ctx, dbModel.CreateOneParams{
		Name:              rpcInstance.GetName(),
		ManufacturerID:    int32(rpcInstance.GetManufacturerId()),
		OperatingSystemID: int32(rpcInstance.GetOperatingSystemId()),
		CreatedAt:         now,
		ModifiedAt:        now,
	})
	if err != nil {
		return nil, fmt.Errorf("error inserting into database: %w", err)
	}

	// TODO: either return record or hydrate

	return toRPCModel(dbResult), nil
}
