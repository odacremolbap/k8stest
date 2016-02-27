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
		indexTemplate, err = indexTemplate.Parse(
			`<html><head><title>k8s test</title></head><body>
      <style>
      dt {
        float: left;
        clear: left;
        width: 6rem;
        margin: 0;
        padding: 0 .5em 0 0;
      }
      dt:after {
        content: ":";
      }
      dd {
        margin: 0;
        padding: 0;
      }
      </style>
      <dt>Hostname</dt><dd>{{.ContainerInfo.Hostname}}</dd>

      {{range .ContainerInfo.Interfaces}}
      <dt>Interface</dt><dl>{{.Name}}</dl>
      {{range .Addresses}}
      <dt>IP</dt><dd>{{.IP}}</dd>
      <dt>Mask</dt><dd>{{.Mask}}</dd>
      <dt>CIDR</dt><dd>{{.CIDR}}</dd>
      <dt>Network</dt><dd>{{.Network}}</dd>
      {{end}}
      <br></br>
      {{end}}

      {{range .ContainerInfo.EnvVars}}
      <dt>{{.Name}}</dt><dd>{{.Value}}</dd>
      {{end}}

      </body></html>
`)
		if err != nil {
			return nil, err
		}
	}

	return indexTemplate, nil
}
