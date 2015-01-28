package mago

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

type mago struct {
	tagname string
	parent  *mago
	body    string
	attribs map[string]string
	depth   int
}

func Maco() *mago {
	return &mago{"", nil, "", make(map[string]string), 0}
}

func MacoInner(tagname string, parent *mago) *mago {
	return &mago{tagname, parent, "", make(map[string]string), parent.depth + 1}
}

func (m *mago) Tag(content string) *mago {
	return MacoInner(content, m)
}

func (m *mago) End() *mago {

	m.parent.Text(m.String())
	return m.parent
}

func (m *mago) Text(content string) *mago {
	m.body += content
	return m
}

func (m *mago) Att(name, value string) *mago {
	m.attribs[name] = value
	return m
}

func (m *mago) String() string {

	var b bytes.Buffer
	if len(m.tagname) != 0 {
		b.WriteString(`<` + m.tagname)
	}

	if len(m.attribs) > 0 {
		for k, v := range m.attribs {
			b.WriteString(` ` + k + `="` + v + `"`)
		}
	}

	if len(m.body) > 0 {
		if len(m.tagname) != 0 || len(m.attribs) > 0 {
			b.WriteString(`>`)
		}
		b.WriteString(m.body)
		if len(m.tagname) != 0 {
			b.WriteString(`</` + m.tagname + `>`)
		}
	} else if len(m.tagname) != 0 {
		b.WriteString(`/>`)
	}
	return b.String()
}

func (m *mago) Code(markup string) string {

	var code bytes.Buffer
	code.WriteString(`m := mago.Maco()`)

	r := (strings.NewReader(markup))
	d := html.NewTokenizer(r)

	for {
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := d.Token()

		switch tokenType {
		case html.StartTagToken: // <tag>

			code.WriteString(`.Tag("` + token.Data + `")"`)
			for _, v := range token.Attr {
				code.WriteString(`.Att("` + v.Key + `,` + v.Val + `)`)
			}

		case html.TextToken: // text between start and end tag

			code.WriteString(`.Text(` + token.Data + `")"`)

		case html.EndTagToken: // </tag>

			code.WriteString(`.End()`)

		case html.SelfClosingTagToken: // <tag/>

			code.WriteString(`.Tag("` + token.Data + `")"`)
			for _, v := range token.Attr {
				code.WriteString(`.Att("` + v.Key + `,` + v.Val + `")"`)
			}
			code.WriteString(`.End()`)

		}
	}
	code.WriteString(`.String()`)
	return code.String()
}
