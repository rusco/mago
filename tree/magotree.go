package tree

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

type magoTree struct {
	tagname string
	parent  *magoTree
	body    string
	attribs map[string]string
	depth   int
}

func NewMagoTree() *magoTree {

	return &magoTree{"", nil, "", make(map[string]string), 0}
}

func magoTreeChild(tagname string, parent *magoTree) *magoTree {

	return &magoTree{tagname, parent, "", make(map[string]string), parent.depth + 1}
}

func (m *magoTree) MtTag(content string) *magoTree {

	return magoTreeChild(content, m)
}

func (m *magoTree) MtEnd() *magoTree {

	m.parent.MtText(m.String())
	return m.parent
}

func (m *magoTree) MtGo(fn func(mx *magoTree) *magoTree) *magoTree {

	m = fn(m)
	return m
}

func (m *magoTree) MtText(content string) *magoTree {

	m.body += content
	return m
}

func (m *magoTree) MtAtt(name, value string) *magoTree {

	m.attribs[name] = value
	return m
}

func (m *magoTree) String() string {

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
		if m.tagname == "script" {
			b.WriteString(`></script>`)
		} else {
			b.WriteString(`/>`)
		}
	}
	return b.String()
}

func (m *magoTree) Indent() string {

	markup := m.String()
	r := strings.NewReader(markup)
	d := html.NewTokenizer(r)
	prevToken := html.CommentToken

	retStr, depth := "", 0

	for {
		tt := d.Next()
		tokenString := string(d.Raw())

		if tt == html.TextToken {
			strippedNewlines := strings.Trim(tokenString, "\n")
			if len(strippedNewlines) == 0 {
				continue
			}
		}
		if tt == html.EndTagToken {
			depth -= 1
		}
		if tt != html.TextToken {
			if prevToken != html.TextToken {
				retStr += "\n"
				for i := 0; i < depth; i++ {
					retStr += "    "
				}
			}
		}
		retStr += tokenString

		if tt == html.ErrorToken {
			break //last token
		} else if tt == html.StartTagToken {
			depth += 1
		}
		prevToken = tt
	}
	return strings.Trim(retStr, "\n")
}

func (m *magoTree) Code(markup string) string {

	var code bytes.Buffer
	code.WriteString(`mago.Ma()`)

	r := strings.NewReader(markup)
	d := html.NewTokenizer(r)

	for {
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := d.Token()

		switch tokenType {
		case html.StartTagToken: // <tag>

			code.WriteString(`.Tag("` + token.Data + `")`)
			for _, v := range token.Attr {
				code.WriteString(`.Att("` + v.Key + `","` + v.Val + `")`)
			}

		case html.TextToken: // text between start and end tag

			code.WriteString(`.Text("` + token.Data + `")`)

		case html.EndTagToken: // </tag>

			code.WriteString(`.End()`)

		case html.SelfClosingTagToken: // <tag/>

			code.WriteString(`.Tag("` + token.Data + `")`)
			for _, v := range token.Attr {
				code.WriteString(`.Att("` + v.Key + `","` + v.Val + `")`)

			}
			code.WriteString(`.End()`)
		}
	}
	code.WriteString(`.String()`)
	return code.String()
}
