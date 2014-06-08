package openweather

import (
	"testing"
)

func TestGetByCity(t *testing.T) {
	_, err := GetByCityName("Belo Horizonte")
	if err != nil {
		t.Error(err.Error())
	}

}
func TestGetById(t *testing.T) {
	_, err := GetById("524901")
	if err != nil {
		t.Error(err.Error())
	}
}
