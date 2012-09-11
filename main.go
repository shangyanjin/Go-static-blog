/*
This program creates a static html blog. The contents of POSTDIR
will be processed and be put into BLOGURL/language/year/month/day/basename.html.
The contents of POSTDIR should be basename.markdown files which are markdown
formatted with some differences. Here is an example;
---
title: This is my first post.
datetime: 2012/09/02 14:15
language: english

Hi there! This is *my* fist post. And I use **markdown!**. Here is a code
for you;
{% codeblock lang:go %}
func main() {
    fmt.Println("hello world!")
}
{% endcodeblock %}
---
"---" parts aren't included in file.The part from start to first empty line is header 
section. This section contains post metadata. Order of fields doesn't matter. But 
all of the fields shown above must be present for posts. You can also use keywords 
header. Keywords are comma separated words that will be added to the head section 
of the post. The actual content is following part. The {% xxx %} are the parts to 
be preprocessed before rendering markdown. In this example, the codeblock will be 
used for syntax highlighting. There are also other preprossing directives like 
{% img "link" %} for putting images and {% include "something.html" %} for including 
other files. Note that included files will not be preprosessed. You can use this to 
insert arbitrary text and html to your posts and pages. Those files should be under 
INCLUDEDIR. Post are also used to create index.html on DEPLOYDIR.

Rest of your site comes from SITEDIR directory. markdown files will be
processed as above, but no special directory tree will be created for them.
html, js and xml files will be preprocessed by go template library. 

Your posts and pages can also have optional layout header. layouts are
html templates that will be used as a container to rendered page. Rendered
page will be put in {{.Page.child}} in layout. They reside in LAYOUTDIR. 
Layouts can be nested. Default layout for posts are post.html, which in 
turn uses default.html as it's layout. Default layout for other markdown 
files are default.html. You can use layout: none header to prevent this 
behaviour. Default layout for xml and html files are none.You can add 
layout header for them if you want.
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"io"
	"io/ioutil"
	"os"
    "os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
)

type Page map[string]string

type Posts []Page

// Sort interface
func (p Posts) Len() int           { return len(p) }
func (p Posts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Posts) Less(i, j int) bool { return p[i]["datetime"] > p[j]["datetime"] }

var (
	fileCache   = make(map[string]string)
	codeBlockRE = *regexp.MustCompile("(?smU)(^{% codeblock lang\\:([^ ]+) %}(.*){% endcodeblock %})")
	imgRE       = *regexp.MustCompile("{% img (\\S+?) %}")
	includeRE   = *regexp.MustCompile("(?sU){% include (\\S+?) %}")
	relativeRE   = *regexp.MustCompile("(?sU){% relative (\\S+?) %}")

	AllPosts  = make(Posts, 0, 128) // cap of 128 seems reasonable.
	postindex = 0
)

const (
	POSTDIR    = "_posts/" // posts, there are special types of pages
	SITEDIR    = "_site/"  // other pages and static content
	LAYOUTDIR  = "_layouts/"
	INCLUDEDIR = "_includes/"
	DEPLOYDIR  = ".."    // root of site will be here. Relative to where you execute this program.
	BLOGURL    = "/blog" // _posts will end up in filepath.Join(DEPLOYDIR, BLOGURL, language ...  

	highlightpath = "C:\\Program Files (x86)\\WinHighlight\\highlight.exe"
)

// load file from filesystem, or return cached copy.
func loadFile(name string) string {
	str, ok := fileCache[name]
	if ok {
		return str
	}
	f, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintf("Can't read %s", name))
	}
	fileCache[name] = string(f)
	return fileCache[name]
}

// return a pointer to named Page. name is filepath.
func loadPage(name string) Page {
	f := loadFile(name)
	p := make(Page)
	p["filename"] = name
	p["content"] = f
	p.ProcessHeaders()
	return p
}

func loadLayout(filename string, current Page) Page {
	p := loadPage(filename)
	for k, v := range current {
		if !strings.Contains("content|layout|filename", k) {
			p[k] = v
		}
	}
	return p
}

func highlight(source, lang string) string {
    if lang == "go" { // if it is go sources, run it thrugh gofmt
        cmd := exec.Command("gofmt")
        cmd.Stdin = strings.NewReader(source)
        out := bytes.Buffer{}
        stderr := bytes.Buffer{}
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        if err := cmd.Run(); err != nil {
            panic(fmt.Sprintf("go formatting failed: \n%s\n%s\n%s\n",source,err.Error(), stderr.String()))
        }
        source = out.String()
    }
	cmd := exec.Command(highlightpath, "--syntax", lang, "--fragment", "--encoding=utf-8", "--enclose-pre")
	cmd.Stdin = strings.NewReader(source)
	var out = bytes.Buffer{}
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Unable to colorize: %s", err.Error()))
	}
	return out.String()
}

func (p Page) ProcessHeaders() {
	/*
	   Headers looks like http headers, and separated from actual content of
	   the file with a blank line. Put each header as seperate key in p
	   and remove headers from Page["content"]
	*/
	lines := strings.Split(p["content"], "\n")
	var rest_starts int
	for i, l := range lines {
		ind := strings.Index(l, ":")
		if ind == -1 {
			rest_starts = i
			break
		}
		p[strings.TrimSpace(l[:ind])] = strings.TrimSpace(l[ind+1:])
	}

	// Trim extra "\r" because we will join remaining lines with "\r\n"
	for i, l := range lines[rest_starts:] {
		lines[rest_starts+i] = strings.TrimRight(l, "\r")
	}

	if lines[rest_starts] == "" { // Don't include empty line
		rest_starts++
	}

	p["content"] = strings.Join(lines[rest_starts:], "\r\n")
}

func (p Page) ExpandMacros() {
	// {% img "foo.png "%} -> <img src="foo.png" />
	content := imgRE.ReplaceAllString(p["content"], "<img src=\"$1\" />")

	codeblock := func(s string) string {
		m := codeBlockRE.FindAllStringSubmatch(s, -1)
		if m == nil {
			panic(fmt.Sprintf("Bad codeblock [%s]", s))
		}
		language := m[0][2]
		source := m[0][3]
		return highlight(source, language)
	}

	content = codeBlockRE.ReplaceAllStringFunc(content, codeblock)

	include := func(s string) string {
		m := includeRE.FindAllStringSubmatch(s, -1)
		if m == nil {
			panic(fmt.Sprintf("Invalid include %s", s))
		}
		return loadFile(filepath.Join(INCLUDEDIR, m[0][1]))
	}
	
    abs_filename, err := filepath.Abs(p["targetfile"])
	if err != nil {
		panic(err.Error())
	}
    
    // find a relative link
    relative := func(s string) string {
        m := relativeRE.FindAllStringSubmatch(s, -1)
        if m == nil {
            panic("Cannot find submatches in function relative.")
        }
        target := m[0][1]
        if target[0] == byte('/') {
            localpath := filepath.Join(strings.Split(target[1:], "/")...)
			targetOnDisk := filepath.Join(DEPLOYDIR, localpath)
			// take absolute url of targetOnDisk, because it break rel otherwise
			targetOnDisk, err := filepath.Abs(targetOnDisk)
			if err != nil {
				panic(err.Error())
			}
			rel, err := filepath.Rel(filepath.Dir(abs_filename), targetOnDisk)
			if err != nil {
				panic(err.Error())
			}

			// turn "\" to "/" on windows, do nothing otherwise
			return filepath.ToSlash(rel)
        }
        return target

    }
    content = relativeRE.ReplaceAllStringFunc(content, relative)
	// {% include filename.html %}
	p["content"] = includeRE.ReplaceAllStringFunc(content, include)

}

func (p Page) ProcessMarkup() {
	p.ExpandMacros()
	p["content"] = string(blackfriday.MarkdownCommon([]byte(p["content"])))
}

func (p Page) Render() {
	// include said file
	include := func(filename string) string {
		inc := loadLayout(filepath.Join(INCLUDEDIR, filename), p)
		inc.Render()
		return inc["content"]
	}
	abs_filename, err := filepath.Abs(p["targetfile"])
	if err != nil {
		panic(err.Error())
	}
	// find relative url
	relative := func(url string) string {
		if url[0] == byte('/') { // abs url
			// turn "/" to "\" on windows, do nothing otherwise
			localpath := filepath.Join(strings.Split(url[1:], "/")...)
			targetOnDisk := filepath.Join(DEPLOYDIR, localpath)
			// take absolute url of targetOnDisk, because it break rel otherwise
			targetOnDisk, err := filepath.Abs(targetOnDisk)
			if err != nil {
				panic(err.Error())
			}
			rel, err := filepath.Rel(filepath.Dir(abs_filename), targetOnDisk)
			if err != nil {
				panic(err.Error())
			}

			// turn "\" to "/" on windows, do nothing otherwise
			return filepath.ToSlash(rel)
		}
		// if it doesn't start with "/" treat it as relative url
		return url
	}

	funcs := template.FuncMap{"include": include, "relative": relative}

	type Data struct {
		Page  Page
		Posts Posts
	}

	tpl := template.Must(template.New(p["filename"]).Funcs(funcs).Parse(p["content"]))
	b := new(bytes.Buffer)
	err = tpl.Execute(b, Data{p, AllPosts})
	if err != nil {
		panic(err.Error())
	}
	p["content"] = b.String()

	if layoutname, ok := p["layout"]; ok {
		layout := loadLayout(filepath.Join(LAYOUTDIR, layoutname), p)
		layout["child"] = p["content"]
		layout.Render()
		p["content"] = layout["content"]
	}
}

// read a file from POSTDIR, process it, and write it to DEPLOYDIR
func processPost(path string) {
	p := loadPage(path)

	p["index"] = fmt.Sprintf("%d", postindex)
	postindex++


	t, err := time.Parse("2006/01/02 15:04", p["datetime"])
	if err != nil {
		panic(err.Error())
	}

	nameparts := []string{DEPLOYDIR,
		BLOGURL,
		p["language"],
		t.Format("2006"),
		t.Format("01"),
		t.Format("02"),
		strings.Replace(p["filename"][len(POSTDIR):], ".markdown", ".html", -1)}

	p["targetfile"] = filepath.Join(nameparts...)

	if _, ok := p["language"]; !ok {
		p["language"] = "english"
	}

	p[p["language"]] = "yes"

	if _, ok := p["layout"]; !ok {
		p["layout"] = filepath.Join(p["language"], "post.html") // default layout for posts are language/post.html
	}

	p.ProcessMarkup()

	// Save this so that we can link it in index page
	p["url"] = strings.Join(nameparts[1:], "/")
	p.Render()

	dir := filepath.Dir(p["targetfile"])
	if os.MkdirAll(dir, os.ModePerm) != nil {
		panic(fmt.Sprintf("can't create %s", dir))
	}
	if err := ioutil.WriteFile(p["targetfile"], []byte(p["content"]), os.ModePerm); err != nil {
		panic(err.Error())
	}
	// Add this post's address to these list, because we will create index page from them.
	AllPosts = append(AllPosts, p)
}

func processPosts() {
	processfunc := func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ".markdown" {
			processPost(path)
		}
		return err
	}
	filepath.Walk(POSTDIR, processfunc)
	sort.Sort(AllPosts)
}

func processPage(path string) {
	p := loadPage(path)
    var needsmarkup = false
	if filepath.Ext(path) == ".markdown" {
        needsmarkup = true
		path = path[:len(path)-len(".markdown")] + ".html"
	}
	d := filepath.Join(DEPLOYDIR, filepath.Dir(path[len(SITEDIR):]))
	p["targetfile"] = filepath.Join(d, filepath.Base(path))
    if needsmarkup {
        p.ProcessMarkup()
    }
	p.Render()
	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		panic(err.Error())
	}
	if err := ioutil.WriteFile(p["targetfile"], []byte(p["content"]), os.ModePerm); err != nil {
		panic(err.Error())
	}
}

func CopyFileContents(source, destination string) error {
	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dest.Close()
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	return err

}

func doRest() {
	processfunc := func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			if strings.Contains(".html.markdown.js.xml", filepath.Ext(path)) {
				// first process and write to deploydir
				processPage(path)
			} else {
				// Just copy file as is to deploydir
				targetfile := filepath.Join(DEPLOYDIR, path[len(SITEDIR):])
				if err := os.MkdirAll(filepath.Dir(targetfile), os.ModePerm); err != nil {
					panic(err.Error())
				}
				if err := CopyFileContents(path, targetfile); err != nil {
					panic(err.Error())
				}
			}
		}
		return err
	}
	if err := filepath.Walk(SITEDIR, processfunc); err != nil {
		panic(err.Error())
	}
}

func main() {
	processPosts()
	doRest() // Process all other things :)
}
