package testsuit

import (
	"mock/gomock"
	"testing"
)

func TestUser_Invoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	personMocker := NewMockPerson(ctrl)
	Invoke(personMocker)
}
