package repo

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/rickypai/web-template/api/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader_ListByPrefix(t *testing.T) {
	instanceA := &modelT{Name: "A"}
	instanceAAC := &modelT{Name: "AAC"}
	instanceABC := &modelT{Name: "ABC"}
	instanceAaa := &modelT{Name: "Aaa"}
	instanceB := &modelT{Name: "B"}
	instanceC := &modelT{Name: "C"}

	type fields struct {
		existing []*modelT
	}
	type args struct {
		prefix string
		limit  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*modelT
		wantErr bool
	}{
		{
			"matches items that has the prefix (case-insensitive)",
			fields{
				existing: []*modelT{
					instanceA,
					instanceAaa,
					instanceAAC,
					instanceABC,
					instanceB,
					instanceC,
				},
			},
			args{
				prefix: "a",
				limit:  10,
			},
			nil, // TODO: implement case-insensitive match
			false,
		},
		{
			"matches items that has the prefix",
			fields{
				existing: []*modelT{
					instanceA,
					instanceAaa,
					instanceAAC,
					instanceABC,
					instanceB,
					instanceC,
				},
			},
			args{
				prefix: "A",
				limit:  10,
			},
			[]*modelT{
				instanceA,
				instanceAAC,
				instanceABC,
				instanceAaa,
			},
			false,
		},
		{
			"matches items that has the prefix up to a limit",
			fields{
				existing: []*modelT{
					instanceA,
					instanceAaa,
					instanceAAC,
					instanceABC,
					instanceB,
					instanceC,
				},
			},
			args{
				prefix: "A",
				limit:  2,
			},
			[]*modelT{
				instanceA,
				instanceAAC,
			},
			false,
		},
		{
			"returns nil if there are no matches",
			fields{
				existing: []*modelT{},
			},
			args{
				prefix: "Z",
				limit:  10,
			},
			nil,
			false,
		},
		{
			"returns nil if limit is 0",
			fields{
				existing: []*modelT{
					instanceA,
					instanceAaa,
					instanceAAC,
					instanceABC,
					instanceB,
					instanceC,
				},
			},
			args{
				prefix: "A",
				limit:  0,
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			db, err := config.TestDB(ctx)
			require.NoError(t, err)

			tx, err := db.Begin()
			require.NoError(t, err)
			defer tx.Rollback() //nolint:errcheck

			w := NewWriter(tx)
			for _, existing := range tt.fields.existing {
				// TODO: use translation method
				_, err := w.CreateOne(ctx, &createReqT{
					Name: existing.GetName(),
				})
				require.NoError(t, err)
			}

			r := NewReader(tx)

			got, err := r.ListByPrefix(ctx, tt.args.prefix, tt.args.limit)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// clear out transient fields for equality assertion
			for _, result := range got {
				result.CreatedAt = nil
				result.ModifiedAt = nil
				result.Id = 0
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReader_GetManyByIDs(t *testing.T) {
	numInstances := 100
	existing := make([]*modelT, 0, numInstances)

	ctx := context.Background()

	db, err := config.TestDB(ctx)
	require.NoError(t, err)

	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback() //nolint:errcheck

	w := NewWriter(tx)
	for i := 0; i < numInstances; i++ {
		newInstance, err := w.CreateOne(ctx, &createReqT{
			Name: strconv.Itoa(rand.Int()),
		})
		require.NoError(t, err)

		existing = append(existing, newInstance)
	}

	type args struct {
		ids []int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*modelT
		wantErr bool
	}{
		{
			"returns items with matching ids",
			args{
				ids: []int64{
					existing[0].GetId(),
					existing[1].GetId(),
					existing[2].GetId(),
				},
			},
			[]*modelT{
				existing[0],
				existing[1],
				existing[2],
			},
			false,
		},
		{
			"returns items with matching ids sorted by id order", // TODO: return by requested order?
			args{
				ids: []int64{
					existing[99].GetId(),
					existing[54].GetId(),
					existing[2].GetId(),
				},
			},
			[]*modelT{
				existing[2],
				existing[54],
				existing[99],
			},
			false,
		},
		{
			"returns items with matching ids deduped", // TODO: maybe not dedup?
			args{
				ids: []int64{
					existing[2].GetId(),
					existing[2].GetId(),
					existing[2].GetId(),
					existing[23].GetId(),
					existing[23].GetId(),
					existing[23].GetId(),
				},
			},
			[]*modelT{
				existing[2],
				existing[23],
			},
			false,
		},
		{
			"returns items with matching ids deduped in order", // TODO: maybe not dedup?
			args{
				ids: []int64{
					existing[23].GetId(),
					existing[2].GetId(),
					existing[23].GetId(),
					existing[2].GetId(),
					existing[23].GetId(),
					existing[2].GetId(),
				},
			},
			[]*modelT{
				existing[2],
				existing[23],
			},
			false,
		},
		{
			"returns only found items", // TODO: maybe error if not all found?
			args{
				ids: []int64{
					rand.Int63(),
					existing[0].GetId(),
					existing[1].GetId(),
					existing[2].GetId(),
					existing[23].GetId(),
					rand.Int63(),
				},
			},
			[]*modelT{
				existing[0],
				existing[1],
				existing[2],
				existing[23],
			},
			false,
		},
		{
			"returns nothing if none are found", // TODO: maybe error if not all found?
			args{
				ids: []int64{
					rand.Int63(),
					rand.Int63(),
					rand.Int63(),
					rand.Int63(),
					rand.Int63(),
					rand.Int63(),
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(tx)

			got, err := r.GetManyByIDs(ctx, tt.args.ids)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
