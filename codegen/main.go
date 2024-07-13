package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const FLAT_DIR = "./flat"
const REEXPORT_FILE = "./flat.go"
const OBJECT_SUFFIX = "T"

func main() {
	files, err := os.ReadDir("./flat")
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		panic("No generated flatbuffer files found, cannot build reexports")
	}

	outputStr := "// @generated\n"

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".go") {
			break
		}

		content, err := os.ReadFile(path.Join(FLAT_DIR, f.Name()))
		if err != nil {
			panic("Couldn't read flatbuffer go file")
		}

		typeName := strings.TrimSuffix(f.Name(), ".go")

		var match string
		isStruct := false
		for _, line := range strings.Split(string(content), "\n") {
			if strings.HasPrefix(line, fmt.Sprintf("type %s byte", typeName)) {
				match = strings.Split(line, " ")[1]
				break
			}
			if strings.HasPrefix(line, fmt.Sprintf("type %s%s struct", typeName, OBJECT_SUFFIX)) {
				match = strings.TrimSuffix(strings.Split(line, " ")[1], OBJECT_SUFFIX)
				isStruct = true
				break
			}
		}

		// a string is "" by default
		if match == "" {
			panic("Couldn't find type definition based on file name in generated flatbuffer files")
		}

		if !isStruct {
			outputStr += fmt.Sprintf("type %s flat.%s\n", match, match)
		} else {
			outputStr += fmt.Sprintf("type %s flat.%s%s\n", match, match, OBJECT_SUFFIX)
		}
	}

	oldContent, err := os.ReadFile(REEXPORT_FILE)
	if err != nil {
		panic("Couldn't read reexport file")
	}

	newContent := strings.Split(string(oldContent), "// @generated")[0] + outputStr
	os.WriteFile(REEXPORT_FILE, []byte(newContent), os.FileMode(0777))
}
