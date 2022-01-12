# Filter Builder For MongoDB Go Driver

## Installation

``` go
go get github.com/JsyTech/mongo-filter-builder
```

## Usage

Here shows a simple usage for query a struct with specific filed and value. More usages can be found in test files.

```go
import (
	builder "github.com/JsyTech/mongo-filter-builder"
)

// Suppose we have a struct looks like this.
// And we would like to find its data from the mongodb.
type CustomStruct struct {
  Name string `bson:"name"`
  CapName int `bson:"capName"`
}

func main() {
  // To build a simple filter
  filter := builder.New().Str("name").Eq("volinda").Build()
  
  ...
  // use the filter in mongo-go-driver's method
  coll.FindOne(ctx, filter)
  
  
  // condtions are in AND mode as default, you can use Or() to compose more condtions.
  andFilter := builder.New().
  Str("name").Eq("volinda").
  Str("capName").Eq("VOLINDA").Build()
  
  
  // the same as in bson:
  // $or: [
  //   {name: "volinda", capName: "VOLINDA"},
  //   {name: "phile"}
  // ]
  orFilter := builder.New().
  Str("name").Eq("volinda").
  Str("capName").Eq("VOLINDA").
  Or().
  Str("name").Eq("phile").Build()
}

```

