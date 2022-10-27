package alert

import "github.com/jedib0t/go-pretty/v6/table"

var tableHeader = table.Row{"类型", "job", "instance", "主机名", "告警信息", "当前值", "预警阈值"}

// var styleCss = `
//     table {
// 	    border-right: 1px solid #000000;
// 	    border-bottom: 1px solid #000000;
// 	    text-align: center;
//     }

//     table th {
//     	border-left: 1px solid #000000;
//     	border-top: 1px solid #000000;
//     }

//     table td {
//     	border-left: 1px solid #000000;
//     	border-top: 1px solid #000000;
//     }`
var styleCss = `
    table {
    	text-align: center;
    	font-family: verdana, arial, sans-serif;
    	font-size: 11px;
    	color: #333333;
    	border-width: 1px;
    	border-color: #666666;
    	border-collapse: collapse;
    }
    
    table th {
    	border-width: 1px;
    	padding: 8px;
    	border-style: solid;
    	border-color: #666666;
    	background-color: #e69900;
    }
    
    table td {
    	border-width: 1px;
    	padding: 8px;
    	border-style: solid;
    	border-color: #666666;
    	background-color: #f2f2f2;
  }`
