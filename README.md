# Filter Builder For Mongo Go Driver

## Installation

``` go
go get github.com/LegeTech/mongo-filter-builder
```

## Usage

Here shows a simple usage for query a struct with specific filed and value. More usages can be found in test files.

```go
import (
	builder "github.com/tarupo/mongo-filter-builder"
)

// Suppose we have a struct looks like this.
// And we would like to find its data from the mongodb.
type CustomStruct struct {
  Name string `bson:"name"`
  CapName int `bson:"capName"`
}

func main() {
  // To build a simple filter
  filter := builder.New().WantStr("name").Eq("volinda").Build()
  
  ...
  // use the filter in mongo-go-driver's method
  coll.FindOne(ctx, filter)
  
  
  // condtions are in AND mode as default, you can use Or() to compose more condtions.
  andFilter := builder.New().
  WantStr("name").Eq("volinda").
  WantStr("capName").Eq("VOLINDA").Build()
  
  
  // the same as in bson:
  // $or: [
  //   {name: "volinda", capName: "VOLINDA"},
  //   {name: "phile"}
  // ]
  orFilter := builder.New().
  WantStr("name").Eq("volinda").
  WantStr("capName").Eq("VOLINDA").
  Or().
  WantStr("name").Eq("phile").Build()
}

```

