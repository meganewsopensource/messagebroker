package messagebroker

import (
	"reflect"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Test_receiver_Receive(t *testing.T) {
	const message = "Message test"
	type fields struct {
		messageSent string
		queueName   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				messageSent: message,
				queueName:   "test",
			},
			want:    message,
			wantErr: false,
		},
		{
			name: "Fail, message broker error",
			fields: fields{
				messageSent: message,
				queueName:   "test",
			},
			want:    message,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReceiver(&mockChannel{messages: []string{message},
				errorConsumingMessage: tt.wantErr},
				tt.fields.queueName)
			got, err := r.Receive()
			if (err != nil) != tt.wantErr {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				messageReceived := readMessage(got)
				if !reflect.DeepEqual(tt.fields.messageSent, messageReceived) {
					t.Errorf("Receive() got = %v, want %v", messageReceived, tt.fields.messageSent)
				}
			}
		})
	}
}

func readMessage(messages <-chan amqp.Delivery) string {
	var response string
	for messageReturn := range messages {
		response = string(messageReturn.Body)
		break
	}
	return response
}
