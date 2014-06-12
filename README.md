openweather
===========

http://openweathermap.org/ wrapper in Go.

Example usage:

```
package main

import (
	"fmt"
	o "github.com/humbhenri/openweather"
	"os"
)

func main() {
	res, _ := o.GetByCityName(os.Args[1])
	fmt.Printf("%+v\n", res)
}
```

And run like:

```
OPENWEATHER_API_KEY=$(cat api_key.txt) go run main.go "Rio de Janeiro"
```

