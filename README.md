# Go in memory cache

Helps you keep in mind what you need


## Example

```go
package main

import (
	"fmt"
	"github.com/urcop/GoMemoryCache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
```

### Output

```
42
<nil>
```