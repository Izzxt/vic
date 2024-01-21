package main

import (
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
)

type data struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	PackageName string `json:"packageName"`
	Path        string `json:"path"`
}

func main() {
	var d []data
	wd, _ := os.Getwd()
	file, err := os.ReadFile(filepath.Join(wd, "packets", "packets.json"))
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &d)

	for _, v := range d {
		if checkFileExists(filepath.Join(wd, "packets", v.Type, v.Path, v.Name+".go")) {
			continue
		}

		if err := os.MkdirAll(filepath.Join(wd, "packets", v.Type, v.Path), os.ModePerm); err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(wd, "packets", v.Type, v.Path, v.Name+".go"))
		if err != nil {
			panic(err)
		}

		t := template.Must(template.New("packet").Parse(tmpt))
		t.Execute(f, v)

		f.Close()
	}

}

func checkFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

var tmpt = `{{if eq .Type "incoming"}}
package {{.PackageName}}
{{else if eq .Type "outgoing"}}
package {{.PackageName}}
{{end}}
import (
	"github.com/Izzxt/vic/core"{{if eq .Type "outgoing"}}
	"github.com/Izzxt/vic/packets/outgoing"
){{end}}

type {{.Name}} struct {}
{{if eq .Type "incoming"}}
func (*{{.Name}}) Execute(client core.HabboClient, in core.IncomingPacket) {}
{{else if eq .Type "outgoing"}}
func (*{{.Name}}) GetId() uint16 {
	return outgoing.{{.Name}}
}

func (c *{{.Name}}) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	return compose
}
{{end}}
`
