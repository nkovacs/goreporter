package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DirList(projectPath string, suffix, expect string) (dirs map[string]string, err error) {
	dirs = make(map[string]string, 0)
	_, err = os.Stat(projectPath)
	if err != nil {
		log.Fatal("dir path is invalid")
	}
	err = filepath.Walk(projectPath, func(subPath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(subPath, suffix) {
			sepIdx := strings.LastIndex(subPath, system)
			var dir string
			if sepIdx == -1 {
				dir = "."
			} else {
				dir = subPath[0:sepIdx]
			}
			if ExpectPkg(expect, dir) {
				return nil
			}
			dirs[PackageAbsPath(dir)] = dir
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return dirs, nil
}

func ExpectPkg(expect, pkg string) bool {
	if expect == "" {
		return false
	}
	expects := strings.Split(expect, ",")
	for _, va := range expects {
		if strings.Contains(pkg, va) {
			return true
		}
	}
	return false
}

func PackageAbsPath(path string) (packagePath string) {
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal("package path is invalid")
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}
	packagePathIndex := strings.Index(absPath, "src")
	if -1 != packagePathIndex {
		packagePath = absPath[(packagePathIndex + 4):]
	}

	return packagePath
}

func PackageAbsPathExceptSuffix(path string) (packagePath string) {
	if strings.LastIndex(path, system) <= 0 {
		path, _ = os.Getwd()
	}
	path = path[0:strings.LastIndex(path, system)]
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}
	packagePathIndex := strings.Index(absPath, "src")
	if -1 != packagePathIndex {
		packagePath = absPath[(packagePathIndex + 4):]
	}

	return packagePath
}

func projectName(projectPath string) (project string) {
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		log.Println(err)
	}
	projectPathIndex := strings.LastIndex(absPath, system)
	if -1 != projectPathIndex {
		project = absPath[(projectPathIndex + 1):len(absPath)]
	}

	return project
}

func absPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
		return path
	}
	return absPath
}
