package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserID(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    UserID
		wantErr bool
	}{
		{
			name: "return parsed ID",
			args: args{
				val: "test_user_id",
			},
			want:    UserID("test_user_id"),
			wantErr: false,
		},
		{
			name: "return error when val is invalid",
			args: args{
				val: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserID(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "ParseUserID() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserID_Validate(t *testing.T) {
	tests := []struct {
		name    string
		id      UserID
		wantErr bool
	}{
		{
			name:    "return no error when value is valid",
			id:      UserID("test_user_id"),
			wantErr: false,
		},
		{
			name:    "return error when value is empty",
			id:      UserID(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.id.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UserID.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	type fields struct {
		id          UserID
		name        string
		displayName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "return no error when value is valid",
			fields: fields{
				id:          UserID("test_user_id"),
				name:        "test_user_name",
				displayName: "test_user_display_name",
			},
			wantErr: false,
		},
		{
			name: "return error when user is invalid",
			fields: fields{
				id:          UserID(""),
				name:        "test_user_name",
				displayName: "test_user_display_name",
			},
			wantErr: true,
		},
		{
			name: "return error when user name is empty",
			fields: fields{
				id:          UserID("test_user_id"),
				name:        "",
				displayName: "test_user_display_name",
			},
			wantErr: true,
		},
		{
			name: "return error when user displayName is empty",
			fields: fields{
				id:          UserID("test_user_id"),
				name:        "test_user_name",
				displayName: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:          tt.fields.id,
				name:        tt.fields.name,
				displayName: tt.fields.displayName,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
