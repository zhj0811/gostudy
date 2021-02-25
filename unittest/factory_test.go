package unittest

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_unittest "github.com/zhj0811/gostudy/unittest/mocks"
)

func TestPrintln(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_unittest.NewMockFactory(ctrl)
	m.EXPECT().Println(gomock.Any()).Do(func(args ...string) {
		fmt.Println(args)
	})

	m.Println("abc")
}
