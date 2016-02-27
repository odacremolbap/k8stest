package model

import (
	"html/template"
)

// IndexPage used by a template to render all container data
type IndexPage struct {
	ContainerInfo ContainerInfo
	AppInfo       AppInfo
}

var indexTemplate *template.Template

// GetIndexPage returns index page rendering
func GetIndexPage() (*template.Template, error) {

	if indexTemplate == nil {
		indexTemplate = template.New("index")
		var err error
		indexTemplate, err = indexTemplate.Parse(index)
		if err != nil {
			return nil, err
		}
	}

	return indexTemplate, nil
}

const index string = `
<!doctype html>
<html lang="en">
	<head>
		<title>k8s test</title>
		<meta charset="utf-8">
		<style>
		body {
		font-family: "Lato","sans-serif";
		}

		table {
		font-family: "Lato","sans-serif";	}		/* added custom font-family  */

		table.one {
		margin-bottom: 3em;
		border-collapse:collapse;	}

		td {							/* removed the border from the table data rows  */
		text-align: center;
		width: 10em;
		padding: 1em; 		}

		th {							  /* removed the border from the table heading row  */
		text-align: center;
		padding: 1em;
		background-color: #e8503a;	     /* added a red background color to the heading cells  */
		color: white;		}			      /* added a white font color to the heading text */

		tr {
		height: 1em;	}

		table tr:nth-child(even) {		      /* added all even rows a #eee color  */
					 background-color: #eee;		}

		table tr:nth-child(odd) {		     /* added all odd rows a #fff color  */
		background-color:#fff;		}
		</style>
	</head>
	<body>
<style>
</style>


Host
<table>
  <tr>
    <th>name</th>
    <th>value</th>
  </tr>
  <tr>
    <td>Hostname</td><td>{{.ContainerInfo.Hostname}}</td>
  </tr>
</table>

{{range .ContainerInfo.Interfaces}}
Net Interface {{.Name}}
<table>
  <tr>
    <th>name</th>
    <th>value</th>
  </tr>
{{range .Addresses}}
  <tr>
		<td>IP</dt><td>{{.IP}}</dd>
	</tr>
	<tr>
		<td>Mask</dt><td>{{.Mask}}</dd>
	</tr>
	<tr>
		<td>CIDR</dt><td>{{.CIDR}}</dd>
	</tr>
	<tr>
		<td>Network</dt><td>{{.Network}}</dd>
	</tr>
{{end}}
</table>
{{end}}


Environment variables

<table>
  <tr>
    <th>name</th>
    <th>value</th>
  </tr>
{{range .ContainerInfo.EnvVars}}
<tr>
	<td>{{.Name}}</dt><td>{{.Value}}</dd>
</tr>
{{end}}
</table>


</body></html>
`
