package metrics

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var tableHeader = table.Row{"类型", "job", "instance", "主机名", "告警信息", "当前值", "预警阈值"}
var tableStyle = table.Style{
	Name: "myNewStyle",
	Box: table.BoxStyle{
		BottomLeft:       "\\",
		BottomRight:      "/",
		BottomSeparator:  "v",
		Left:             "[",
		LeftSeparator:    "{",
		MiddleHorizontal: "-",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      "<",
		PaddingRight:     ">",
		Right:            "]",
		RightSeparator:   "}",
		TopLeft:          "(",
		TopRight:         ")",
		TopSeparator:     "^",
		UnfinishedRow:    " ~~~",
	},
	Color: table.ColorOptions{
		// AutoIndexColumn: nil,
		// FirstColumn:     nil,
		Footer:       text.Colors{text.BgCyan, text.FgBlack},
		Header:       text.Colors{text.BgHiCyan, text.FgBlack},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgWhite, text.FgBlack},
	},
	Format: table.FormatOptions{
		Footer: text.FormatUpper,
		Header: text.FormatUpper,
		Row:    text.FormatDefault,
	},
	Options: table.Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	},
}

var styleCss = `
    table {
	    border-right: 1px solid #000000;
	    border-bottom: 1px solid #000000;
	    text-align: center;
    }
    
    table th {
    	border-left: 1px solid #000000;
    	border-top: 1px solid #000000;
    }
    
    table td {
    	border-left: 1px solid #000000;
    	border-top: 1px solid #000000;
    }`
