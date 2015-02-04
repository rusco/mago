# mago
**ma**rkup in **go**

### Tiny fluent markup builder library for logicful templates in GO


Mago means [Magician](http://pt.wikipedia.org/wiki/Mago) in portuguese.

You write your servercode in [Go](http://www.golang.org), you write your clientcode in [Go](http://www.gopherjs.org/), why not writing your templating code in [Go](https://github.com/rusco/mago) too ?

* strongly typed templates
* natural embedding of markup in your binary
* use go and go fmt everywhere


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
m := Mago().Tag("root").Tag("numbers")
for i := 1; i < 4; i++ {
	m = m.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
}
m = m.End().End()

println(m.String()
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
m := mago.Mago().Tag("a").Att("id","myid").Text("x").Tag("br").End().Text("y").End().Text("c").String()
```

-api still undergoing changes, use on your own risk !
