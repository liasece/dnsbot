package util

import (
	"net"
	"testing"
)

func TestGetHostExternalIP(t *testing.T) {
	type args struct {
		server string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestGetHostExternalIP",
			args: args{
				server: "http://myexternalip.com/raw",
			},
			wantErr: false,
		},
		{
			name: "TestGetHostExternalIP",
			args: args{
				server: "http://myexternalip.com/raw",
			},
			wantErr: false,
		},
		{
			name: "TestGetHostExternalIP",
			args: args{
				server: "http://cip.cc",
			},
			wantErr: false,
		},
		{
			name: "TestGetHostExternalIP",
			args: args{
				server: "ifconfig.me",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHostExternalIP(tt.args.server)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostExternalIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				ip := net.ParseIP(got)
				if ip.IsPrivate() {
					t.Errorf("ip.IsPrivate() = %v", got)
				}
			}
		})
	}
}
