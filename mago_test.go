package mago

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {

	t01 := Mago().Tag("a").Text("x").Tag("b").Text("y").End().Text("c").End().String()
	got, expected := t01, `<a>x<b>y</b>c</a>`
	if expected != got {
		t.Errorf("expected Test01: \n%v, got: \n%v.", expected, got)
	}
}

func Test02(t *testing.T) {

	got := Mago().Tag("parent").Att("parentproperty1", "true").Att("parentproperty2", "5").Tag("child1").Att("childproperty1", "c").Text("childbody").End().Tag("child2").Att("childproperty2", "c").Text("childbody").End().End().Tag("script").Text("$.scriptbody();").End().String()
	expected := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected != got {
		t.Errorf("expected Test02: \n%v, got: \n%v.", expected, got)
	}
}

func Test03(t *testing.T) {

	got := Mago().Tag("parent").Att("parentproperty1", "true").Att("parentproperty2", "5").Tag("child1").Att("childproperty1", "c").Text("childbody").End().Tag("child2").Att("childproperty2", "c").Text("childbody").End().End().Tag("script").Text("$.scriptbody();")
	got = got.End()
	expected := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected != got.String() {
		t.Errorf("expected Test03: \n%v, got: \n%v.", expected, got)
	}
}

func Test04(t *testing.T) {

	m := Mago().Tag("root").Tag("numbers")
	for i := 1; i < 4; i++ {
		m = m.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
	}
	m = m.End().End()

	got := m.String()
	expected := `<root><numbers><number class="x1">sometext</number><number class="x2">sometext</number><number class="x3">sometext</number></numbers></root>`

	if expected != got {
		t.Errorf("expected Test04: \n%v, got: \n%v.", expected, got)
	}
}

func Test05(t *testing.T) {

	got, expected := Mago().Tag("a").Text("x").Tag("br").End().Text("y").End().Text("c").String(), `<a>x<br/>y</a>c`
	if expected != got {
		t.Errorf("expected Test05: \n%v, got: \n%v.", expected, got)
	}
}

func Test06(t *testing.T) {

	got, expected := Mago().Tag("a").End().String(), `<a/>`

	if expected != got {
		t.Errorf("expected Test06: \n%v, got: \n%v.", expected, got)
	}
}

func Test07(t *testing.T) {

	input, output := `<a id="myid">x<br/>y</a>c`, Mago().Tag("a").Text("x").Tag("br").End().Text("y").End().Text("c").String()
	//continue here

	got := Mago().Code(input)
	println("got: ", got)
	_, _ = input, output
}

func Test08(t *testing.T) {

	table := `<table style="width:100%"><tr><td>h1</td><td>h2</td><td>h3</td></tr><tr><td>line1, col1</td><td>line1, col2</td><td>line1, col3</td></tr></table>`
	_ = table

	m := Mago().Tag("table").Att("style", "width:100%")
	for row := 0; row < 10; row++ {
		m = m.Tag("tr")
		for col := 0; col < 10; col++ {
			m = m.Tag("td").Text(fmt.Sprintf("R=%d, C=%d", row, col)).End()
		}
		m = m.End()
	}
	mstr := m.End().String()
	_ = mstr

	//continue here ...

}

func todo() {
	//put in README.md:
	_ = 	`tests with gopherjs
		better test setup
		round robin tests
		empty attibutes
		indent
		xss
		tool: html page for code generation`
}
