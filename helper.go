package migorm

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

func getMigrationNames() []string {
	names := make([]string, 0, len(pool.migrations))
	for k := range pool.migrations {
		names = append(names, k)
	}

	sort.Strings(names)

	return names
}

// сheck the existence of a file in the directory with migrations
func checkFileExists(dir string, name string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		split := strings.Split(f.Name(), "_")

		if name == strings.Join(split[1:], "_") {
			return fmt.Errorf("file %s already exists in dir: %s", name, dir)
		}
	}

	return nil
}

//
func getTemplate() (*template.Template, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("template caller")
	}

	tmpl, err := template.ParseFiles(path.Dir(filename) + "/template")
	if err != nil {
		return nil, fmt.Errorf("parse template : %v", err)
	}

	return tmpl, nil
}
