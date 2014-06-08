openweather
===========

<<<<<<< HEAD
http://openweathermap.org/ wrapper in Go.

Example usage:

```
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
```
=======
http://openweathermap.org/ api wrapper for golang
>>>>>>> b68016836ba2a1c974f8d38501554d1f999afdfb
