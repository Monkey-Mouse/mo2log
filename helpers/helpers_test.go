package helpers

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBytesToMongoID(t *testing.T) {
	bs := primitive.NewObjectID()
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want primitive.ObjectID
	}{
		{name: "test", args: args{bs[:]}, want: bs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToMongoID(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToMongoID() = %v, want %v", got, tt.want)
			}
		})
	}
}
