package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Monkey-Mouse/mo2log/logmodel"
	"github.com/Monkey-Mouse/mo2log/service/logservice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_server_Count(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *logservice.UserID
	}
	now := time.Now()
	id := primitive.NewObjectID()
	col.InsertMany(context.TODO(), []interface{}{
		logmodel.LogModel{
			OperatorID: id,
			CreateTime: now,
		},
		logmodel.LogModel{
			OperatorID: id,
			CreateTime: now.AddDate(0, 0, -2),
		},
		logmodel.LogModel{
			OperatorID: id,
			CreateTime: now.Add(-time.Hour),
		},
	})
	tests := []struct {
		name    string
		s       *server
		args    args
		wantNum *logservice.Num
		wantErr bool
	}{
		{name: "Test count", s: &server{}, args: args{context.TODO(), &logservice.UserID{UserId: id[:]}}, wantNum: &logservice.Num{Num: 2}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, err := tt.s.Count(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNum, tt.wantNum) {
				t.Errorf("server.Count() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
	col.DeleteMany(context.TODO(), bson.M{"operator_id": id})
}
