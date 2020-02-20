package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/mitchellh/go-homedir"
)

var pbTypes = map[string]string{
	"BIGINT":   "int64",
	"INTEGER":  "int32",
	"TEXT":     "string",
	"DOUBLE":   "double",
	"DATETIME": "google.protobuf.Timestamp",
}

// TODO: Read from a file.
var pbExceptions = map[string]string{
	"User.uid":     "uint64",
	"User.gid":     "uint64",
	"Group.gid":    "uint64",
	"Process.uid":  "uint64",
	"Process.gid":  "uint64",
	"Process.euid": "uint64",
	"Process.egid": "uint64",
	"Process.suid": "uint64",
	"Process.sgid": "uint64",
}

type Service struct {
	Tables []*Table
}

type Table struct {
	Name         string
	NameSg       string
	Descr        string
	PbName       string
	PbNameSg     string
	HasTimestamp bool
	Columns      []*Column
}

type Column struct {
	Name    string
	Type    string
	Descr   string
	PbType  string
	PbOrder int
}

var (
	reTable  = regexp.MustCompile(`table_name\((.*)\)`)
	reDescr  = regexp.MustCompile(`description\((.*)\)`)
	reColumn = regexp.MustCompile(`Column\((.*)\)`)
)

func (s *Service) parseFile(file string, dst string, templ *template.Template) error {
	log.Printf("parse file: %s", file)

	// Read schema.
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Parse schema.
	t := Table{Columns: []*Column{}}
	order := 1
	for _, line := range strings.Split(string(b), "\n") {
		switch {
		case reTable.MatchString(line):
			a := strings.Split(reTable.FindStringSubmatch(line)[1], ",")
			t.Name = strings.Trim(strings.TrimSpace(a[0]), "\"")
			t.NameSg = inflection.Singular(t.Name)
			t.PbName = strcase.ToCamel(t.Name)
			t.PbNameSg = inflection.Singular(t.PbName)
		case reDescr.MatchString(line):
			t.Descr = strings.Trim(reDescr.FindStringSubmatch(line)[1], "\"")
		case reColumn.MatchString(line):
			a := strings.Split(reColumn.FindStringSubmatch(line)[1], ",")
			c := Column{
				Name:    strings.Trim(strings.TrimSpace(a[0]), "\""),
				Type:    strings.TrimSpace(a[1]),
				Descr:   strings.Trim(strings.TrimSpace(a[2]), "\""),
				PbOrder: order,
			}
			order++

			if v, ok := pbExceptions[t.PbNameSg+"."+c.Name]; ok {
				c.PbType = v
			} else if v, ok := pbTypes[c.Type]; ok {
				c.PbType = v
			} else {
				return fmt.Errorf("unknown column type: %s", c.Type)
			}

			if c.PbType == "google.protobuf.Timestamp" {
				t.HasTimestamp = true
			}

			t.Columns = append(t.Columns, &c)
		}
	}

	pbFile := filepath.Join(dst, "pb", strings.ToLower(t.NameSg+".proto"))
	log.Printf("template schema file: %s", pbFile)

	// Create file.
	f, err := os.Create(pbFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Template schema.
	if err := templ.ExecuteTemplate(f, "schema", t); err != nil {
		return fmt.Errorf("failed to template schema: %v", err)
	}

	s.Tables = append(s.Tables, &t)

	return nil
}

func (s *Service) Parse(src string, dst string, templ *template.Template, depth int) error {
	depth++
	if depth > 10 {
		return fmt.Errorf("exceeded max depth (10) for: %s", src)
	}

	log.Printf("parse directory: %s", src)

	// Verify path is a directory.
	if fi, err := os.Stat(src); err != nil || !fi.IsDir() {
		return fmt.Errorf("not a directory: %s", src)
	}

	// Glob path.
	paths, err := filepath.Glob(filepath.Join(src, "*"))
	if err != nil {
		return fmt.Errorf("failed to glob path: %v", err)
	}

	for _, p := range paths {
		fi, err := os.Stat(p)
		if err != nil {
			return fmt.Errorf("failed to stat: %v", p)
		}

		switch {
		case fi.IsDir():
			if err := s.Parse(p, dst, templ, depth); err != nil {
				return err
			}
		case fi.Mode().IsRegular():
			if filepath.Ext(p) != ".table" {
				continue
			}
			if err := s.parseFile(p, dst, templ); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// Get GOPATH var.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = "~/go"
	}

	// Define flags.
	var src, dst string
	flag.StringVar(&src, "src", filepath.Join(goPath, "src/github.com/osquery/osquery/specs"), "source path")
	flag.StringVar(&dst, "dst", filepath.Join(goPath, "src/github.com/mickep76/osquery-protobuf"), "destination path")
	flag.Parse()

	// Expand home dir.
	src, _ = homedir.Expand(src)
	dst, _ = homedir.Expand(dst)

	// Compile template's.
	t := template.Must(template.ParseGlob(filepath.Join(dst, "pbgen/*.tmpl")))

	// Parse database schemas and template protobuf.
	s := Service{Tables: []*Table{}}
	if err := s.Parse(src, dst, t, 0); err != nil {
		log.Fatal(err)
	}
}
