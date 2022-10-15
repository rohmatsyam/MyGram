package usecase

import (
	"final_zoom/domain"
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckParamIsNumber(t *testing.T) {
	result := CheckParamIsNumber("1")

	if result != nil {
		t.Fail()
	}
	fmt.Println("TestCheckParamIsNumber Eksekusi terhenti")
}

func Test_userUseCase_GetUsersUc(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name    string
		c       userUseCase
		args    args
		want    []*domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetUsersUc(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.GetUsersUc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.GetUsersUc() = %v, want %v", got, tt.want)
			}
		})
	}
}
