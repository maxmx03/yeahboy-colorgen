package vim

import (
	"log"
	"path/filepath"
	"yeahboy/scheme"
	"yeahboy/system"
	"yeahboy/vim/template"
)

type projectStructure struct {
	colors struct {
		dir         string
		colorscheme string
	}
	docs struct {
		dir            string
		colorschemeTxt string
	}
}

func Create(scheme scheme.Scheme) {
	var project projectStructure

	var root = "build/vim9"
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.colorscheme = filepath.Join(project.colors.dir, scheme.Name+".vim")
	project.docs.dir = filepath.Join(root, "docs")
	project.docs.colorschemeTxt = filepath.Join(project.docs.dir, scheme.Name+".txt")

	var dirs []string
	dirs = append(dirs, project.colors.dir)
	dirs = append(dirs, project.docs.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}

	createFile(project.colors.colorscheme, scheme, template.Colors())
	createFile(project.docs.colorschemeTxt, scheme, template.Docs())
	log.Printf("%v.vim created successfully", scheme.Name)
}

func Update() {}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}