package util

import "testing"

func TestEncryptPassword(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []struct {
		name                string
		args                args
		wantEncryptPassword string
		wantErr             bool
	}{
		// TODO: Add test cases.
		{"equal", args{"123123", "123"}, "b822bb93905a9bd8b3a0c08168c427696436cf8bf37ed4ab8ebf41a07642ed1c", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncryptPassword, err := EncryptPassword(tt.args.password, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEncryptPassword != tt.wantEncryptPassword {
				t.Errorf("EncryptPassword() = %v, want %v", gotEncryptPassword, tt.wantEncryptPassword)
			}
		})
	}
}
