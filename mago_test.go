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
	expected1 := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`
	expected2 := `<parent parentproperty2="5" parentproperty1="true"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected1 != got && expected2 != got {
		t.Errorf("expected Test02: \n%v, or \n%v, but got: \n%v.", expected1, expected2, got)
	}
}

func Test03(t *testing.T) {

	got := Mago().Tag("parent").Att("parentproperty1", "true").Att("parentproperty2", "5").Tag("child1").Att("childproperty1", "c").Text("childbody").End().Tag("child2").Att("childproperty2", "c").Text("childbody").End().End().Tag("script").Text("$.scriptbody();")
	got = got.End()
	gotStr := got.String()
	expected1 := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`
	expected2 := `<parent parentproperty2="5" parentproperty1="true"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected1 != gotStr && expected2 != gotStr {
		t.Errorf("expected Test03: \n%v, or \n%v, but got: \n%v.", expected1, expected2, got)
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

	got, expected := Mago().Tag(A).Text("x").Tag(BR).End().Text("y").End().Text("c").String(), `<a>x<br/>y</a>c`
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

	input := `<a id="myid">x<br/>y</a>c`
	expected := `m := mago.Mago().Tag(A).Att(ID,"myid").Text("x").Tag(BR).End().Text("y").End().Text("c").String()`

	got := Mago().Code(input)
	if expected != got {
		t.Errorf("expected Test07 (Code generation): \n%v, got: \n%v.", expected, got)
	}

	roundrobin := Mago().Tag(A).Att(ID, "myid").Text("x").Tag(BR).End().Text("y").End().Text("c").String()
	if roundrobin != input {
		t.Errorf("expected Test07 (Code generation roundrobin): \n%v, input: \n%v.", roundrobin, input)
	}
}

func Test08(t *testing.T) {

	m := Mago().Tag("root").Tag("numbers")
	for i := 1; i < 4; i++ {
		m = m.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
	}
	m = m.End().End()

	got := m.String()
	expected := `<root><numbers><number class="x1">sometext</number><number class="x2">sometext</number><number class="x3">sometext</number></numbers></root>`
	if expected != got {
		t.Errorf("expected Test08: \n%v, got: \n%v.", expected, got)
	}
}

func Test09(t *testing.T) {

	expected := `<table style="width:100%"><tr><td>row0,col0</td><td>row0,col1</td></tr><tr><td>row1,col0</td><td>row1,col1</td></tr></table>`

	m := Mago().Tag(TABLE).Att(STYLE, "width:100%")
	for row := 0; row < 2; row++ {
		m = m.Tag(TR)
		for col := 0; col < 2; col++ {
			m = m.Tag(TD).Text(fmt.Sprintf("row%d,col%d", row, col)).End()
		}
		m = m.End()
	}
	got := m.End().String()

	if expected != got {
		t.Errorf("expected Test09 (Table): \n%v, got: \n%v.", expected, got)
	}
}

func Test10(t *testing.T) {

	start, end := 1, 5

	got := Mago().Tag("root").Tag("numbers").Exec(
		func(mx *mago) *mago {
			for i := start; i < end; i++ {
				mx = mx.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
			}
			return mx
		}).End().End()

	expected := `<root><numbers><number class="x1">sometext</number><number class="x2">sometext</number><number class="x3">sometext</number><number class="x4">sometext</number></numbers></root>`
	if expected != got.String() {
		t.Errorf("expected Test10 (Exec): \n%v, got: \n%v.", expected, got.String())
	}
}

func Test11(t *testing.T) {

	start, end := 1, 5
	got := Mago().Tag("root").Tag("numbers").Exec(func(mx *mago) *mago {
		for i := start; i < end; i++ {
			mx = mx.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
		}
		return mx
	}).End().End().Indent()

	expected := "<root>\n    <numbers>\n        <number class=\"x1\">sometext</number>\n        <number class=\"x2\">sometext</number>\n        <number class=\"x3\">sometext</number>\n        <number class=\"x4\">sometext</number>\n    </numbers>\n</root>"
	if expected != got {
		t.Errorf("expected Test11 (Indent): \n%v, got: \n%v.", expected, got)
	}
}

func Test12(t *testing.T) {

	expected := `<table style="width:100%"><tr><td>row_0,col_0</td><td>row_0,col_1</td></tr><tr><td>row_1,col_0</td><td>row_1,col_1</td></tr></table>`

	got := Mago().Tag("table").Att("style", "width:100%").Exec(func(mx *mago) *mago {
		for row := 0; row < 2; row++ {
			mx = mx.Tag(TR).Exec(func(my *mago) *mago {
				for col := 0; col < 2; col++ {
					my = my.Tag(TD).Text(fmt.Sprintf("row_%d,col_%d", row, col)).End()
				}
				return my
			}).End()
		}
		return mx
	}).End().String()

	if expected != got {
		t.Errorf("expected Test12 (Table generation with 2 levels of Exec - bad idea !): \n%v, got: \n%v.", expected, got)
	}
}

func Test13(t *testing.T) {

	got := MagoText("<!DOCTYPE html>").
		Tag(HTML).
		Tag(SCRIPT).Att(SRC, "jquery.js").Text("").End().
		Tag(P).Att(ID, "2").End().
		End().String()
	expected := `<!DOCTYPE html><html><script src="jquery.js"></string><p id="2"/></html>`

	if expected != got {
		t.Errorf("expected Test13 (DOCTYPE and SCRIPT tag ): \n%v, got: \n%v.", expected, got)
	}

}

func todo() {
	//put in README.md:
	_ = `tests with gopherjs ?
		.Range() function ?
		use stack internally ?
		empty attibutes ?
		html comments ?
		xss ?
		symbol ?
		html page?`
}
