# mago
MArkup in GO

### Tiny fluent markup builder library in GO

Mago means [Magician](http://pt.wikipedia.org/wiki/Mago) in portuguese.

You write your servercode in [Go](http://www.golang.org), you write your clientcode in [Go](http://www.gopherjs.org/), why not writing your templating code in [Go](https://github.com/rusco/mago) too ?


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

MAGO Code Generation example:
```go
input := `<a id="myid">x<br/>y</a>c`
got := Mago().Code(input)
println(got)
```

outputs the string:
```go
m := mago.Mago().Tag("a").Att("id","myid").Text("x").Tag("br").End().Text("y").End().Text("c").String()
```






-stay tuned

