package mago

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {

	t01 := Ma().Tag("a").Text("x").Tag("b").Text("y").End().Text("c").End().String()
	got, expected := t01, `<a>x<b>y</b>c</a>`
	if expected != got {
		t.Errorf("expected Test01: \n%v, got: \n%v.", expected, got)
	}
}

func Test02(t *testing.T) {

	got := Ma().Tag("parent").Att("parentproperty1", "true").Att("parentproperty2", "5").Tag("child1").Att("childproperty1", "c").Text("childbody").End().Tag("child2").Att("childproperty2", "c").Text("childbody").End().End().Tag("script").Text("$.scriptbody();").End().String()
	expected1 := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`
	expected2 := `<parent parentproperty2="5" parentproperty1="true"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected1 != got && expected2 != got {
		t.Errorf("expected Test02: \n%v, or \n%v, but got: \n%v.", expected1, expected2, got)
	}
}

func Test03(t *testing.T) {

	got := Ma().Tag("parent").Att("parentproperty1", "true").Att("parentproperty2", "5").Tag("child1").Att("childproperty1", "c").Text("childbody").End().Tag("child2").Att("childproperty2", "c").Text("childbody").End().End().Tag("script").Text("$.scriptbody();")
	got = got.End()
	gotStr := got.String()
	expected1 := `<parent parentproperty1="true" parentproperty2="5"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`
	expected2 := `<parent parentproperty2="5" parentproperty1="true"><child1 childproperty1="c">childbody</child1><child2 childproperty2="c">childbody</child2></parent><script>$.scriptbody();</script>`

	if expected1 != gotStr && expected2 != gotStr {
		t.Errorf("expected Test03: \n%v, or \n%v, but got: \n%v.", expected1, expected2, got)
	}
}

func Test04(t *testing.T) {

	m := Ma().Tag("root").Tag("numbers").Go(func(mx *Mago) {
		for i := 1; i < 4; i++ {
			mx.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
		}
	}).End().End()

	got := m.String()
	expected := `<root><numbers><number class="x1">sometext</number><number class="x2">sometext</number><number class="x3">sometext</number></numbers></root>`

	if expected != got {
		t.Errorf("expected Test04: \n%v, got: \n%v.", expected, got)
	}
}

func Test05(t *testing.T) {

	got, expected := Ma().Tag(A).Text("x").Tag(BR).End().Text("y").End().Text("c").String(), `<a>x<br/>y</a>c`
	if expected != got {
		t.Errorf("expected Test05: \n%v, got: \n%v.", expected, got)
	}
}

func Test06(t *testing.T) {

	got, expected := Ma().Tag("a").End().String(), `<a/>`

	if expected != got {
		t.Errorf("expected Test06: \n%v, got: \n%v.", expected, got)
	}
}

func Test07(t *testing.T) {

	input := `<a id="myid">x<br/>y</a>c`
	expected := `mago.Ma().Tag("a").Att("id","myid").Text("x").Tag("br").End().Text("y").End().Text("c").String()`

	got := Ma().Code(input)
	if expected != got {
		t.Errorf("expected Test07 (Code generation): \n%v, got: \n%v.", expected, got)
	}

	roundrobin := Ma().Tag(A).Att(ID, "myid").Text("x").Tag(BR).End().Text("y").End().Text("c").String()
	if roundrobin != input {
		t.Errorf("expected Test07 (Code generation roundrobin): \n%v, input: \n%v.", roundrobin, input)
	}
}

func Test08(t *testing.T) {
	got := Ma("<!DOCTYPE html>").
		Tag(HTML).
		Fmt("").Tag(SCRIPT).Att(SRC, "jquery.js").Text("").End().
		Fmt("").Tag(P).Att(ID, "2").End().
		End().String()
	expected := `<!DOCTYPE html><html><script src="jquery.js"></script><p id="2"/></html>`

	if expected != got {
		t.Errorf("expected Test08 (DOCTYPE and SCRIPT tag ): \n%v, got: \n%v.", expected, got)
	}
}

func Test09(t *testing.T) {

	expected := `<table style="width:100%"><tr><td>row0,col0</td><td>row0,col1</td></tr><tr><td>row1,col0</td><td>row1,col1</td></tr></table>`

	m := Ma().Tag(TABLE).Att(STYLE, "width:100%")
	for row := 0; row < 2; row++ {
		m.Tag(TR)
		for col := 0; col < 2; col++ {
			m.Tag(TD).Text(fmt.Sprintf("row%d,col%d", row, col)).End()
		}
		m.End()
	}
	got := m.End().String()

	if expected != got {
		t.Errorf("expected Test09 (Table): \n%v, got: \n%v.", expected, got)
	}
}

func Test10(t *testing.T) {

	start, end := 1, 5

	got := Ma().Tag("root").Tag("numbers").Go(
		func(mx *Mago) {
			for i := start; i < end; i++ {
				mx.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
			}
		}).End().End()

	expected := `<root><numbers><number class="x1">sometext</number><number class="x2">sometext</number><number class="x3">sometext</number><number class="x4">sometext</number></numbers></root>`
	if expected != got.String() {
		t.Errorf("expected Test10 (Go): \n%v, got: \n%v.", expected, got.String())
	}
}

func Test11(t *testing.T) {

	start, end := 1, 5
	got := Ma(CONF_INDENT).Tag("root").Tag("numbers").Go(func(mx *Mago) {
		for i := start; i < end; i++ {
			mx.Tag("number").Att("class", "x"+fmt.Sprintf("%d", i)).Text("sometext").End()
		}
	}).End().End().String()

	expected := "<root>\n    <numbers>\n        <number class=\"x1\">sometext</number>\n        <number class=\"x2\">sometext</number>\n        <number class=\"x3\">sometext</number>\n        <number class=\"x4\">sometext</number>\n    </numbers>\n</root>"
	if expected != got {
		t.Errorf("expected Test11 (Indent): \n%v, got: \n%v.", expected, got)
	}
}

func Test12(t *testing.T) {

	expected := `<table style="width:100%"><tr><td>row_0,col_0</td><td>row_0,col_1</td></tr><tr><td>row_1,col_0</td><td>row_1,col_1</td></tr></table>`

	got := Ma().Tag("table").Att("style", "width:100%").Go(func(mx *Mago) {
		for row := 0; row < 2; row++ {
			mx.Tag(TR).Go(func(my *Mago) {
				for col := 0; col < 2; col++ {
					my.Tag(TD).Text(fmt.Sprintf("row_%d,col_%d", row, col)).End()
				}
			}).End()
		}
	}).End().String()

	if expected != got {
		t.Errorf("expected Test12 (Table generation with 2 levels of Go): \n%v, got: \n%v.", expected, got)
	}
}

func Test13(t *testing.T) {

	tx1 := Ma().Tag("x13")
	tx1.Att("a", "1")
	tx1.Att("b", "2")
	tx1.Att("c", "3").Go(func(mx *Mago) {
		mx.Att("go", "rocks")
	})
	tx1.Att("d", "4")
	tx1.End()

	gotLen := len(tx1.String())
	if gotLen != 41 {
		t.Errorf("expected Test13.1 (sequential Att Calls): should return 41 character and not %d.", gotLen)
	}

	tx2 := Ma().Tag("y").Att("a", "sx1").Go(func(mx *Mago) {
		mx.Att("b", "2")
		for i := 0; i < 4; i++ {
			str := "a" + fmt.Sprintf("%d", i)
			mx.Att(str, str)
		}
	}).End().String()

	gotLen2 := len(tx2)
	if gotLen2 != 50 {
		t.Errorf("expected Test13.2 (sequential Att Calls): should return 50 character and not %d.", gotLen)
	}

	tx3 := Ma(CONF_INDENT).Tag("yyy").Att("zzz", "1").Go(func(mx *Mago) {
		mx.Tag("ttt")
		for i := 0; i < 3; i++ {
			mx.Att(fmt.Sprintf("data%d", i), fmt.Sprintf("val%d", i))
		}
		mx.End()
	}).End().String()

	gotLen3 := len(tx3)
	if gotLen3 != 70 {
		t.Errorf("expected Test13.3 (sequential Att/Tag Calls): should return 70 character and not %d.", gotLen)
	}

}

func Test14(t *testing.T) {

	tx1 := Ma(CONF_INDENT).Tag("a").Att("a1", "1")
	tx1.Tag("b").Att("idx", "2").End().Go(func(mx *Mago) {
		mx.Tag("ttt")
		for i := 0; i < 3; i++ {
			mx.Att(fmt.Sprintf("data%d", i), fmt.Sprintf("val%d", i))
		}
		mx.End()
	}).Tag("moretags").End().End()

	gotLen1 := len(tx1.String())
	if gotLen1 != 98 {
		t.Errorf("expected Test14.1 should return 98 character and not %d.", gotLen1)
	}

	tx2 := Ma(CONF_INDENT).Tag(TABLE).Att(STYLE, "width:100%").Go(func(mx *Mago) {
		for row := 0; row < 3; row++ {
			mx.Tag(TR).Go(func(my *Mago) {
				for col := 0; col < 12; col++ {
					my.Tag(TD).Text(fmt.Sprintf("row_%d,col_%d", row, col)).End()
				}
			}).End()
		}
	}).End()

	gotLen2 := len(tx2.String())
	if gotLen2 != 1142 {
		t.Errorf("expected Test14.2 should return 1142 character and not %d.", gotLen2)
	}

}

func todo() {
	//put in README.md:
	_ = `
		empty attibutes
		tagnames with spaces 
		html comments 
		xss 
		symbol/png
		html page`
}
