package mago

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

const ( //tags
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

const ( //attributes
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
)

const TAGSATTR = `,a,abbr,address,area,article,aside,audio,base,bdo,bgsound,blink,blockquote,body,br,button,canvas,caption,col,colgroup,command,comment,datalist` +
	`,dd,del,details,div,dl,dt,embed,fieldset,figure,b,i,small,footer,form,head,header,hgroup,h1 ,h2,h3,h4,h5,h6,hr,html,iframe,ilayer,img,input,ins,keygen` +
	`,label,layer,legend,li,link,map,mark,marquee,meta,meter,multicol,nav,nobr,noembed,noscript,object,ol,optgroup,option,output,p,param,` +
	`cite,code,dfn,em,kbd,samp,strong,var,pre,progress,ruby,q,script,section,select,spacer,span,style,sub,sup` +
	`,table,tbody,td,textarea,tfoot,th,thead,time,title,tr,ul,video,wbr` +
	`,accesskey,align,background,bgcolor,class,contenteditable,contextmenu,draggable,height,hidden,id,item,itemprop,spellcheck,subject,tabindex,valign,width,`

type mago struct {
	tagname string
	parent  *mago
	body    string
	attribs map[string]string
	depth   int
}

func Mago() *mago {
	return &mago{"", nil, "", make(map[string]string), 0}
}

func magoChild(tagname string, parent *mago) *mago {
	return &mago{tagname, parent, "", make(map[string]string), parent.depth + 1}
}

func (m *mago) Tag(content string) *mago {
	return magoChild(content, m)
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
	code.WriteString(`m := mago.Mago()`)

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

			if strings.Contains(TAGSATTR, ","+token.Data+",") {
				code.WriteString(`.Tag(` + strings.ToUpper(token.Data) + `)`)
			} else {
				code.WriteString(`.Tag("` + token.Data + `")`)
			}

			for _, v := range token.Attr {

				if strings.Contains(TAGSATTR, ","+v.Key+",") {
					code.WriteString(`.Att(` + strings.ToUpper(v.Key) + `,"` + v.Val + `")`)
				} else {
					code.WriteString(`.Att("` + v.Key + `","` + v.Val + `")`)
				}
			}

		case html.TextToken: // text between start and end tag

			code.WriteString(`.Text("` + token.Data + `")`)

		case html.EndTagToken: // </tag>

			code.WriteString(`.End()`)

		case html.SelfClosingTagToken: // <tag/>

			if strings.Contains(TAGSATTR, ","+token.Data+",") {
				code.WriteString(`.Tag(` + strings.ToUpper(token.Data) + `)`)
			} else {
				code.WriteString(`.Tag("` + token.Data + `")`)
			}
			for _, v := range token.Attr {

				if strings.Contains(TAGSATTR, ","+v.Key+",") {
					code.WriteString(`.Att(` + strings.ToUpper(v.Key) + `,"` + v.Val + `")`)
				} else {
					code.WriteString(`.Att("` + v.Key + `","` + v.Val + `")`)
				}
			}
			code.WriteString(`.End()`)
		}
	}
	code.WriteString(`.String()`)
	return code.String()
}
