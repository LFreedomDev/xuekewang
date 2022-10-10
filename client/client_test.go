package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/LFreedomDev/xuekewang/util/random"
)

func NewDefaultClient() *SdkClient {
	return NewSdkClient("test", "test", time.Second*30)
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
			if err != nil {
				t.Errorf("SdkClient.GetSubjects() error = %v", err)
			} else {
				if err = gotRes.Error(); err != nil {
					t.Errorf("SdkClient.GetSubjects() error = %v", err)
				} else {
					fmt.Printf("%+v", gotRes)
				}
			}
		})
	}
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
			if err != nil {
				t.Errorf("SdkClient.GetCoursesAll() error = %v", err)
			} else {
				if err = gotRes.Error(); err != nil {
					t.Errorf("SdkClient.GetCoursesAll() error = %v", err)
				} else {
					fmt.Printf("%+v", gotRes)
				}
			}
		})
	}
}

func TestSdkClient_GetTextBooks(t *testing.T) {
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
			gotRes, err := cli.GetTextBooks(GetTextBooksParams{})
			if err != nil {
				t.Errorf("SdkClient.GetTextBooks() error = %v", err)
			} else {
				if err = gotRes.Error(); err != nil {
					t.Errorf("SdkClient.GetTextBooks() error = %v", err)
				} else {
					fmt.Printf("%+v", gotRes)
				}
			}
		})
	}
}

func TestSdkClient_GetCourseKnowledgePoints(t *testing.T) {
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
			gotRes, err := cli.GetCourseKnowledgePoints(GetCourseKnowledgePointsParams{
				CourseId: 1,
			})
			if err != nil {
				t.Errorf("SdkClient.GetCourseKnowledgePoints() error = %v", err)
			} else {
				if err = gotRes.Error(); err != nil {
					t.Errorf("SdkClient.GetCourseKnowledgePoints() error = %v", err)
				} else {
					fmt.Printf("%+v", gotRes)
				}
			}
		})
	}
}

func TestSdkClient_QuestionPick(t *testing.T) {
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
			gotRes, err := cli.QuestionPick(QuestionPickParams{
				CourseId: 1,
			})
			if err != nil {
				t.Errorf("SdkClient.QuestionPick() error = %v", err)
			} else {
				if err = gotRes.Error(); err != nil {
					t.Errorf("SdkClient.QuestionPick() error = %v", err)
				} else {
					fmt.Printf("%+v", gotRes)
				}
			}
		})
	}
}
