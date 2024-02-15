package cya_game

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose You Own GoVenture</title>

</head>
<body>
<section class="page">
    <h1>
        {{.Title}}
    </h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</section>
	<style>
		  body {
			font-family: helvetica, arial;
		  }
		  h1 {
			text-align:center;
			position:relative;
		  }
		  .page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FCF6FC;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #797;
		  }
		  ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		  }
		  li {
			padding-top: 10px;
		  }
		  a,
		  a:visited {
			text-decoration: underline;
			color: #555;
		  }
		  a:active,
		  a:hover {
			color: #222;
		  }
		  p {
			text-indent: 1em;
		  }
	</style>
</body>
</html>`

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	//Checking against no path
	if path == "" || path == "/" {
		path = "/intro"
	}

	//Filtering out the slash that comes first
	path = path[1:]

	// Passing in ["intro"]. This does check the value in the map being there.
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)

		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	//Decoding the file
	d := json.NewDecoder(r)
	//Story map from story.go
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
