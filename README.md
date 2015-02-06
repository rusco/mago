# mago
**ma**rkup in **go**

### Tiny fluent markup builder library for logicful templates in GO


Mago means [Magician](http://pt.wikipedia.org/wiki/Mago) in portuguese.

You write your servercode in [Go](http://www.golang.org), you write your clientcode in [Go](http://www.gopherjs.org/), why not writing your templating code in [Go](https://github.com/rusco/mago) too ?

* strongly typed templates
* natural embedding of markup in your binary
* use go, go fmt and go test in your view layer


***

To create this:
```xml
<root>
  <numbers>
    <number class="x1">sometext</number>
    <number class="x2">sometext</number>
    <number class="x3">sometext</number>
  </numbers>
</root>
```

You have to write this:
```go
m := Ma().Tag("root").Tag("numbers")
for i := 1; i < 4; i++ {
	m.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
}
m = m.End().End()

println(m.String()
```

Complete Example:

To create this:
```xml
<table style="width:100%">
    <tr>
        <td>row_0,col_0</td>
        <td>row_0,col_1</td>
        <td>row_0,col_2</td>
        <td>row_0,col_3</td>
    </tr>
    <tr>
        <td>row_1,col_0</td>
        <td>row_1,col_1</td>
        <td>row_1,col_2</td>
        <td>row_1,col_3</td>
    </tr>
</table>
```

Your code would look like this:
```go
package main

import (
	"fmt"
	m "github.com/rusco/mago"
)

func main() {
	table := m.Ma(m.CONF_INDENT).Tag(m.TABLE).Att(m.STYLE, "width:100%").Go(func(mx *m.Mago) {
		for row := 0; row < 2; row++ {
			mx.Tag(m.TR).Go(func(my *m.Mago) {
				for col := 0; col < 4; col++ {
					my.Tag(m.TD).Text(fmt.Sprintf("row_%d,col_%d", row, col)).End()
				}
			}).End()
		}
	}).End()
	fmt.Printf("%s", table.String())
}
```



***

MAGO Code Generation support for easy development of templates:
```go
input := `<a id="myid">x<br/>y</a>c`
got := Mago().Code(input)
println(got)
```

outputs the string:
```go
m := mago.Ma().Tag("a").Att("id","myid").Text("x").Tag("br").End().Text("y").End().Text("c").String()
```

-api still undergoing changes, use on your own risk !
