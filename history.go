package main

import (
	"encoding/json"
	"io"
	"sort"
	"text/template"
)

type History struct {
	responses map[string]*Response
}

func NewHistory() *History {
	return &History{responses: map[string]*Response{}}
}

func (h *History) Add(resp *Response) {
	h.responses[resp.Referer] = resp
}

func (h *History) Get(referer string) *Response {
	response, _ := h.responses[referer]
	return response
}

func (h *History) Has(referer string) bool {
	_, ok := h.responses[referer]
	return ok
}

func (h *History) Len() int {
	return len(h.responses)
}

func (h *History) Keys() []string {
	keys := []string{}
	for key := range h.responses {
		keys = append(keys, key)
	}
	return keys
}

func (h *History) Values() []*Response {
	responses := []*Response{}
	for _, response := range h.responses {
		responses = append(responses, response)
	}
	return responses
}

const txtTemplate = `{{range .}}{{.Referer}}
{{range .Links}}  [link ] {{.}}
{{end}}{{range .Assets}}  [asset] {{.}}
{{end}}{{end}}`

func (h *History) WriteAsText(w io.Writer) error {
	responses := h.Values()
	sort.Sort(ResponseByReferer(responses))

	t := template.Must(template.New("template").Parse(txtTemplate))
	return t.Execute(w, responses)
}

const htmlTemplate = `<html>
<head><title>Sitemap</title></head>
<body>
<table>
  <tr>
    <td>Page</td>
    <td>Links</td>
    <td>Assets</td>
  </tr>
  {{range .}}<tr>
    <td><a href="{{.Referer}}">{{.Referer}}</a></td>
    <td>
      {{range .Links}}<a href="{{.}}">{{.}}</a><br/>
      {{end}}
    </td>
    <td>
      {{range .Assets}}<a href="{{.}}">{{.}}</a><br/>
      {{end}}
    </td>
  </tr>{{end}}
</table>
</body>
</html>
`

func (h *History) WriteAsHTML(w io.Writer) error {
	responses := h.Values()
	sort.Sort(ResponseByReferer(responses))

	t := template.Must(template.New("template").Parse(htmlTemplate))
	return t.Execute(w, responses)
}

func (h *History) WriteAsJSON(w io.Writer) error {
	responses := h.Values()
	sort.Sort(ResponseByReferer(responses))
	b, err := json.Marshal(responses)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
