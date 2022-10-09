package internal

import (
	"reflect"
	"testing"
	"time"

	"xuekewang/internal/util/random"
)

func NewDefaultClient() *SdkClient {
	return NewSdkClient("100931645431743400", "exJWPj6YGujk3v5wNKo2r0HEqTU0xBgI", time.Second*30)
}

func TestSdkClient_GetSubjects(t *testing.T) {
	random.Init(time.Now().UnixNano())

	tests := []struct {
		name    string
		wantRes interface{}
		wantErr bool
	}{
		{"", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := NewDefaultClient()
			gotRes, err := cli.GetSubjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("SdkClient.GetSubjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SdkClient.GetSubjects() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

	time.Sleep(time.Second)
}

func TestSdkClient_GetCoursesAll(t *testing.T) {
	random.Init(time.Now().UnixNano())
	tests := []struct {
		name    string
		wantRes interface{}
		wantErr bool
	}{
		{"", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := NewDefaultClient()
			gotRes, err := cli.GetCoursesAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SdkClient.GetCoursesAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SdkClient.GetCoursesAll() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

	time.Sleep(time.Second)
}
