package repo

import (
	"context"
	"fmt"

	manufacturerPb "github.com/rickypai/web-template/api/protobuf/manufacturer"
	operatingSystemPb "github.com/rickypai/web-template/api/protobuf/operating_system"
	cursorPkg "github.com/rickypai/web-template/api/server/cursor"
)

type HydratedReader struct {
	repo     *Reader
	hydrator *Hydrator
}

func NewHydratedReader(repo *Reader, manufacturerClient manufacturerPb.ManufacturerReaderClient, operatingSystemClient operatingSystemPb.OperatingSystemReaderClient) *HydratedReader {
	return &HydratedReader{
		repo:     repo,
		hydrator: NewHydrator(manufacturerClient, operatingSystemClient),
	}
}

func (r *HydratedReader) ListByPage(ctx context.Context, req cursorPkg.PageRequest) ([]*modelT, *cursorPkg.PageResult, error) {
	results, pageResult, err := r.repo.ListByPage(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	err = r.hydrator.HydrateMany(ctx, results)
	if err != nil {
		return nil, nil, fmt.Errorf("hydrating: %w", err)
	}

	return results, pageResult, nil
}

func (r *HydratedReader) ListByCursor(ctx context.Context, req cursorPkg.CursorRequest) ([]*modelT, *cursorPkg.CursorResult, error) {
	results, cursorResult, err := r.repo.ListByCursor(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	err = r.hydrator.HydrateMany(ctx, results)
	if err != nil {
		return nil, nil, fmt.Errorf("hydrating: %w", err)
	}

	return results, cursorResult, nil
}

func (r *HydratedReader) GetOneByID(ctx context.Context, id int64) (*modelT, error) {
	result, err := r.repo.GetOneByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = r.hydrator.HydrateOne(ctx, result)
	if err != nil {
		return nil, fmt.Errorf("hydrating: %w", err)
	}

	return result, nil
}

func (r *HydratedReader) GetManyByIDs(ctx context.Context, ids []int64) ([]*modelT, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	results, err := r.repo.GetManyByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	err = r.hydrator.HydrateMany(ctx, results)
	if err != nil {
		return nil, fmt.Errorf("hydrating: %w", err)
	}

	return results, nil
}

func (r *HydratedReader) ListByPrefix(ctx context.Context, prefix string, limit int64) ([]*modelT, error) {
	results, err := r.repo.ListByPrefix(ctx, prefix, limit)
	if err != nil {
		return nil, err
	}

	err = r.hydrator.HydrateMany(ctx, results)
	if err != nil {
		return nil, fmt.Errorf("hydrating: %w", err)
	}

	return results, nil
}
