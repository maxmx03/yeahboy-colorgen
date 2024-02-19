package nvim

import (
	"log"
	"path/filepath"
	"schemecraft/nvim/template"
	"schemecraft/scheme"
	"schemecraft/system"
	"strings"
)

type projectStructure struct {
	colors struct {
		dir  string
		file string
	}
	lua struct {
		dir     string
		file    string
		config  string
		palette struct {
			dir  string
			file string
		}
		highlights struct {
			dir  string
			file string
		}
		utils string
		color string
	}
	tests struct {
		dir       string
		loadSpec  string
		setupSpec string
	}
	dockerfile string
	compose    string
	shell      string
}

var project projectStructure

func setProject(scheme scheme.Scheme, isMainTheme bool, root string) {
	var schemeName = strings.Split(scheme.GetName(), "-")[0]
	project.colors.dir = filepath.Join(root, "colors")
	project.colors.file = filepath.Join(project.colors.dir, scheme.GetName()+".lua")
	project.lua.dir = filepath.Join(root, "lua", schemeName)
	project.lua.palette.dir = filepath.Join(project.lua.dir, "palette")
	project.lua.highlights.dir = filepath.Join(project.lua.dir, "highlights")

	if isMainTheme {
		project.lua.palette.file = filepath.Join(project.lua.palette.dir, "init.lua")
		project.lua.highlights.file = filepath.Join(project.lua.highlights.dir, "init.lua")
	} else {
		project.lua.palette.file = filepath.Join(project.lua.palette.dir, scheme.GetName()+".lua")
		project.lua.highlights.file = filepath.Join(project.lua.highlights.dir, scheme.GetName()+".lua")
	}

	project.lua.file = filepath.Join(project.lua.dir, "init.lua")
	project.lua.utils = filepath.Join(project.lua.dir, "utils.lua")
	project.lua.config = filepath.Join(project.lua.dir, "config.lua")
	project.lua.color = filepath.Join(project.lua.dir, "color.lua")
	project.tests.dir = filepath.Join(root, "tests")
	project.tests.loadSpec = filepath.Join(project.tests.dir, "load_spec.lua")
	project.tests.setupSpec = filepath.Join(project.tests.dir, "setup_spec.lua")
	project.dockerfile = filepath.Join(root, "Dockerfile")
	project.compose = filepath.Join(root, "docker-compose.yml")
	project.shell = filepath.Join(root, "shell.nix")
}

func createProjectDirs() {
	var dirs []string
	dirs = append(dirs, project.colors.dir)
	dirs = append(dirs, project.lua.dir)
	dirs = append(dirs, project.lua.palette.dir)
	dirs = append(dirs, project.lua.highlights.dir)
	dirs = append(dirs, project.tests.dir)

	for _, dir := range dirs {
		var err error = system.CreateDir(dir)

		if err != nil {
			panic(err)
		}
	}
}

func createProjectFiles(scheme scheme.Scheme) {
	createFile(project.compose, scheme, template.DockerCompose())
	createFile(project.dockerfile, scheme, template.DockerFile())
	createFile(project.shell, scheme, template.NixShell())
	createFile(project.tests.setupSpec, scheme, template.SetupSpec())
	createFile(project.tests.loadSpec, scheme, template.LoadSpec())
	createFile(project.colors.file, scheme, template.Colors())
	createFile(project.lua.utils, scheme, template.Utils())
	createFile(project.lua.config, scheme, template.Config())
	createFile(project.lua.palette.file, scheme, template.Palette())
	createFile(project.lua.highlights.file, scheme, template.Highlights())
	createFile(project.lua.color, scheme, template.Color())
	createFile(project.lua.file, scheme, template.Root())
}

func Create(scheme scheme.Scheme, isMainTheme bool) {
	var schemeName = strings.Split(scheme.GetName(), "-")[0]
	var root string = "build"
	root = filepath.Join(root, schemeName+".nvim")
	setProject(scheme, isMainTheme, root)
	createProjectDirs()
	createProjectFiles(scheme)
	log.Printf("%v.lua created successfully", scheme.GetName())
}

func Update(scheme scheme.Scheme, isMainTheme bool) {
	var root string
	setProject(scheme, isMainTheme, root)
	createProjectFiles(scheme)
	log.Printf("%v.lua updated successfully", scheme.GetName())
}

func createFile(file string, scheme scheme.Scheme, schemeTemplate string) {
	var err error = system.WriteTemplateFile(file, scheme, schemeTemplate)

	if err != nil {
		panic(err)
	}
}
