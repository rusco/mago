package mago

import "github.com/rusco/mago/tree"

const (
	CONF_INDENT = "_config_indent__true________"
	CONF_EMPTY  = "_config_boolean_attribute___"
)

const ( //common html5 tags
	A          = "a"
	ABBR       = "abbr"
	ADDRESS    = "address"
	AREA       = "area"
	ARTICLE    = "article"
	ASIDE      = "aside"
	AUDIO      = "audio"
	BASE       = "base"
	BDO        = "bdo"
	BGSOUND    = "bgsound"
	BLINK      = "blink"
	BLOCKQUOTE = "blockquote"
	BODY       = "body"
	BR         = "br"
	BUTTON     = "button"
	CANVAS     = "canvas"
	CAPTION    = "caption"
	COL        = "col"
	COLGROUP   = "colgroup"
	COMMAND    = "command"
	COMMENT    = "comment"
	DOCTYPE    = "<!DOCTYPE html>"
	DATALIST   = "datalist"
	DD         = "dd"
	DEL        = "del"
	DETAILS    = "details"
	DIV        = "div"
	DL         = "dl"
	DT         = "dt"
	EMBED      = "embed"
	FIELDSET   = "fieldset"
	FIGURE     = "figure"
	B          = "b"
	I          = "i"
	SMALL      = "small"
	FOOTER     = "footer"
	FORM       = "form"
	HEAD       = "head"
	HEADER     = "header"
	HGROUP     = "hgroup"
	H1         = "h1 "
	H2         = "h2"
	H3         = "h3"
	H4         = "h4"
	H5         = "h5"
	H6         = "h6"
	HR         = "hr"
	HTML       = "html"
	IFRAME     = "iframe"
	ILAYER     = "ilayer"
	IMG        = "img"
	INPUT      = "input"
	INS        = "ins"
	KEYGEN     = "keygen"
	LABEL      = "label"
	LAYER      = "layer"
	LEGEND     = "legend"
	LI         = "li"
	LINK       = "link"
	MAP        = "map"
	MARK       = "mark"
	MARQUEE    = "marquee"
	META       = "meta"
	METER      = "meter"
	MULTICOL   = "multicol"
	NAV        = "nav"
	NOBR       = "nobr"
	NOEMBED    = "noembed"
	NOSCRIPT   = "noscript"
	OBJECT     = "object"
	OL         = "ol"
	OPTGROUP   = "optgroup"
	OPTION     = "option"
	OUTPUT     = "output"
	P          = "p"
	PARAM      = "param"
	CITE       = "cite"
	CODE       = "code"
	DFN        = "dfn"
	EM         = "em"
	KBD        = "kbd"
	SAMP       = "samp"
	STRONG     = "strong"
	VAR        = "var"
	PRE        = "pre"
	PROGRESS   = "progress"
	RUBY       = "ruby"
	Q          = "q"
	SCRIPT     = "script"
	SECTION    = "section"
	SELECT     = "select"
	SPACER     = "spacer"
	SPAN       = "span"
	STYLE      = "style"
	SUB        = "sub"
	SUP        = "sup"
	TABLE      = "table"
	TBODY      = "tbody"
	TD         = "td"
	TEXTAREA   = "textarea"
	TFOOT      = "tfoot"
	TH         = "th"
	THEAD      = "thead"
	TIME       = "time"
	TITLE      = "title"
	TR         = "tr"
	UL         = "ul"
	VIDEO      = "video"
	WBR        = "wbr"
)

const ( //common attributes
	ACCESSKEY       = "accesskey"
	ALIGN           = "align"
	BACKGROUND      = "background"
	BGCOLOR         = "bgcolor"
	CLASS           = "class"
	CONTENTEDITABLE = "contenteditable"
	CONTEXTMENU     = "contextmenu"
	DRAGGABLE       = "draggable"
	HEIGHT          = "height"
	HIDDEN          = "hidden"
	ID              = "id"
	ITEM            = "item"
	ITEMPROP        = "itemprop"
	SPELLCHECK      = "spellcheck"
	SUBJECT         = "subject"
	TABINDEX        = "tabindex"
	VALIGN          = "valign"
	WIDTH           = "width"
	ASYNC           = "async"
	SRC             = "src"
)

type magoCmd struct {
	command string
	args    []interface{}
}

type Mago struct {
	indent bool
	list   []magoCmd
}

func Ma(txt ...string) *Mago {

	list := make([]magoCmd, 0)

	indent := false
	for i, val := range txt {
		if i == 2 {
			break
		}
		if val == CONF_INDENT {
			indent = true
		}
		if val != CONF_INDENT {
			mc := magoCmd{"Text", []interface{}{val}}
			list = append(list, mc)
		}
	}
	return &Mago{indent, list}
}

func (m *Mago) Tag(content string) *Mago {
	mc := magoCmd{"Tag", []interface{}{content}}
	m.list = append(m.list, mc)

	return m
}

func (m *Mago) End() *Mago {
	mc := magoCmd{"End", []interface{}{nil}}
	m.list = append(m.list, mc)

	return m
}

func (m *Mago) Text(content string) *Mago {
	mc := magoCmd{"Text", []interface{}{content}}
	m.list = append(m.list, mc)

	return m
}

func (m *Mago) Go(fn func(mx *Mago)) *Mago {
	fn(m)
	return m
}

func (m *Mago) Att(name, value string) *Mago {
	mc := magoCmd{"Att", []interface{}{name, value}}
	m.list = append(m.list, mc)
	return m
}

func (m *Mago) Fmt(spacer string) *Mago {
	return m
}

func (m *Mago) String() string {

	mt := tree.NewMagoTree()
	for _, cmd := range m.list {

		switch cmd.command {
		case "Tag":
			mt = mt.MtTag(cmd.args[0].(string))
		case "Att":
			mt = mt.MtAtt(cmd.args[0].(string), cmd.args[1].(string))
		case "End":
			mt = mt.MtEnd()
		case "Text":
			mt = mt.MtText(cmd.args[0].(string))
		}
	}
	if m.indent {
		return mt.Indent()
	} else {
		return mt.String()
	}
}

func (m *Mago) Code(markup string) string {
	mt := tree.NewMagoTree()
	return mt.Code(markup)
}
