package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/entity"
	"github.com/ahmdyaasiin/workshop-ci-cd/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	ctx := context.Background()
	now := time.Now()

	type args struct {
		ctx     context.Context
		keyword string
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(pr *mock.MockIRProduct)
		want       []dto.ResponseGetProduct
		wantErr    bool
	}{
		{
			name: "success get all products",
			args: args{
				ctx:     ctx,
				keyword: "ayam",
			},
			beforeTest: func(pr *mock.MockIRProduct) {
				pr.EXPECT().
					All(gomock.Any(), "ayam").
					Return([]entity.Product{
						{
							ID:        2,
							Name:      "Soto Ayam",
							Price:     13000,
							CreatedAt: now,
							UpdatedAt: now,
						},
						{
							ID:        4,
							Name:      "Ayam Goreng",
							Price:     15000,
							CreatedAt: now,
							UpdatedAt: now,
						},
					}, nil)
			},
			want: []dto.ResponseGetProduct{
				{
					ID:    2,
					Name:  "Soto Ayam",
					Price: 13000,
				},
				{
					ID:    4,
					Name:  "Ayam Goreng",
					Price: 15000,
				}},
			wantErr: false,
		},
		{
			name: "failed get all products",
			args: args{
				ctx:     ctx,
				keyword: "ayam",
			},
			beforeTest: func(pr *mock.MockIRProduct) {
				pr.EXPECT().
					All(gomock.Any(), "ayam").
					Return([]entity.Product{}, fmt.Errorf("db down"))
			},
			want:    []dto.ResponseGetProduct{},
			wantErr: true,
		},
	}

	for _, _t := range tests {
		t.Run(_t.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)

			pr := mock.NewMockIRProduct(ctrl)
			u := New(pr)

			if _t.beforeTest != nil {
				_t.beforeTest(pr)
			}

			r, err := u.All(_t.args.ctx, _t.args.keyword)

			if _t.wantErr {
				assert.Error(t, err, "expected an error but got nil")
				return
			}

			assert.NoError(t, err, "unexpected error from usecase")
			assert.Equal(t, _t.want, r, "products result mismatch")
		})
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	now := time.Now()

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(pr *mock.MockIRProduct)
		want       dto.ResponseGetProduct
		wantErr    bool
	}{
		{
			name: "success get product details",
			args: args{
				ctx: ctx,
				id:  "4",
			},
			beforeTest: func(pr *mock.MockIRProduct) {
				pr.EXPECT().
					Get(gomock.Any(), "4").
					Return(entity.Product{
						ID:        4,
						Name:      "Ayam Goreng",
						Price:     15000,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
			},
			want: dto.ResponseGetProduct{
				ID:    4,
				Name:  "Ayam Goreng",
				Price: 15000,
			},
			wantErr: false,
		},
		{
			name: "failed get product details",
			args: args{
				ctx: ctx,
				id:  "4",
			},
			beforeTest: func(pr *mock.MockIRProduct) {
				pr.EXPECT().
					Get(gomock.Any(), "4").
					Return(entity.Product{}, fmt.Errorf("db down"))
			},
			want:    dto.ResponseGetProduct{},
			wantErr: true,
		},
	}

	for _, _t := range tests {
		t.Run(_t.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)

			pr := mock.NewMockIRProduct(ctrl)
			u := New(pr)

			if _t.beforeTest != nil {
				_t.beforeTest(pr)
			}

			r, err := u.Get(_t.args.ctx, _t.args.id)

			if _t.wantErr {
				assert.Error(t, err, "expected an error but got nil")
				return
			}

			assert.NoError(t, err, "unexpected error from usecase")
			assert.Equal(t, _t.want, r, "products result mismatch")
		})
	}
}
