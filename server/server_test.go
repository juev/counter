package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/gojuno/minimock/v3"
	"github.com/juev/counter/proto/counter"
)

func TestImplementation_AddDomain(t *testing.T) {
	type fields struct {
		UnimplementedCounterServer counter.UnimplementedCounterServer
	}
	type args struct {
		ctx    context.Context
		domain *counter.Domain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *counter.Response
		wantErr bool
	}{
		{
			name:   "1",
			fields: fields{},
			args: args{
				ctx:    nil,
				domain: &counter.Domain{Domain: "juev.org"},
			},
			want: &counter.Response{
				Status: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Implementation{
				UnimplementedCounterServer: tt.fields.UnimplementedCounterServer,
			}

			// mock
			mc := minimock.NewController(t)
			expectCmdNil := redis.StringCmd{}
			expectCmdNil.SetErr(redis.Nil)

			expectCmd := redis.StringCmd{}
			expectCmd.SetVal("1")

			expectCmdReturn := redis.StatusCmd{}
			expectCmdReturn.SetErr(nil)

			rangeIn = func() int { return 1 }
			rdb = NewRedisMock(mc).
				GetMock.When(tt.args.ctx, "juev.org").Then(&expectCmdNil).
				GetMock.When(tt.args.ctx, "1").Then(&expectCmdNil).
				SetMock.When(tt.args.ctx, "juev.org", 1).Then(&expectCmdReturn).
				SetMock.When(tt.args.ctx, "1", 0).Then(&expectCmdReturn)

			got, err := s.AddDomain(tt.args.ctx, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDomain() got = %v, want %v", got, tt.want)
			}
		})
	}
}
