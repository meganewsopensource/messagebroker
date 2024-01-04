package messagebroker

import (
	"context"
	"testing"
	"time"
)

func Test_sender_Send(t *testing.T) {
	type fields struct {
		queueName string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				queueName: "test",
			},
			args: args{
				message: "Message test",
			},
			wantErr: false,
		},
		{
			name: "Fail, error while seending message",
			fields: fields{
				queueName: "test",
			},
			args: args{
				message: "Message test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		t.Run(tt.name, func(t *testing.T) {
			s := NewSender(&mockChannel{[]string{}, false, tt.wantErr},
				tt.fields.queueName,
				ctx)
			if err := s.Send(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
