package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var fileURL = "/Users/cs/note"
var cmdtypora = "/Applications/Typora 2.app/Contents/MacOS/Typora"

func printFile(file os.FileInfo, deep int) {
	for i := 0; i < deep; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("%s\n", file.Name())
}

func printDir(dir []os.FileInfo, root string, deep int) {

	for _, v := range dir {
		name := v.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		isDir := v.IsDir()
		if !isDir {
			printFile(v, deep)
		} else {
			_path := path.Join(root, name)
			printFile(v, deep)
			_dirpath, err := os.Open(_path)
			if err != nil {
				log.Fatal(err)
			}
			defer _dirpath.Close()
			_dir, err := _dirpath.Readdir(0)
			if err != nil {
				log.Fatal(err)
			}
			// 递归目录
			printDir(_dir, _path, deep+1)
		}
	}
}

func main() {

	l := len(os.Args[1:])
	rootDir, err := os.Open(fileURL)
	if err != nil {
		log.Fatalf("open %s error: %v", fileURL, err)
	}
	fs, err := rootDir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	defer rootDir.Close()
	if l == 1 {
		if os.Args[1] == "ll" {
			printDir(fs, fileURL, 0)
		} else if os.Args[1] == "ls" {
			for _, f := range fs {
				printFile(f, 0)
			}
		} else {
			for _, f := range fs {
				if f.Name() == os.Args[1] {
					printDir([]os.FileInfo{f}, fileURL, 0)
				}
			}
		}
	} else if l == 2 {
		if os.Args[1] == "add" {
			for _, f := range fs {
				if f.Name() == os.Args[2] {
					fmt.Println(os.Args[2], " already exits")
					return
				}
			}
			_p := path.Join(fileURL, os.Args[2])
			err := os.Mkdir(_p, 0777)
			if err != nil {
				log.Fatalf("create %s error: %v", _p, err)
			}
		} else if os.Args[1] == "rm" {
			for _, f := range fs {
				if f.Name() == os.Args[2] {
					_p := path.Join(fileURL, os.Args[2])
					err := os.Remove(_p)
					if err != nil {
						log.Fatalf("remove %s error: %v", _p, err)
					}
					return
				}
			}
		} else {
			for _, f := range fs {
				if f.Name() == os.Args[1] {
					if os.Args[2] == "ll" || os.Args[2] == "ls" {
						printDir([]os.FileInfo{f}, fileURL, 0)
					} else {
						_path := path.Join(fileURL, f.Name())
						_dirpath, err := os.Open(_path)
						if err != nil {
							log.Fatal(err)
						}
						defer _dirpath.Close()
						files, err := _dirpath.Readdir(0)
						for _, file := range files {
							if strings.Contains(file.Name(), os.Args[2]) {
								cmd := exec.Command("/Applications/Typora 2.app/Contents/MacOS/Typora", path.Join(_path, file.Name()))
								err := cmd.Start()
								if err != nil {
									log.Fatal(err)
								}
								//cmd.Wait()
							}
						}
					}
				}
			}
		}
	} else if l == 3 {
		if os.Args[1] == "add" {
			for _, f := range fs {
				if f.Name() == os.Args[2] {
					_p := path.Join(fileURL, f.Name())
					_dirpath, err := os.Open(_p)
					if err != nil {
						log.Fatal(err)
					}
					defer _dirpath.Close()
					files, err := _dirpath.Readdir(0)
					for _, file := range files {
						if file.Name() == os.Args[3] {
							fmt.Println(os.Args[2], " already exits")
							return
						}
					}
				}
			}
			_p := path.Join(fileURL, os.Args[2])
			file := path.Join(_p, os.Args[3])
			_, err := os.Create(file)
			if err != nil {
				log.Fatalf("create file %s error : %v", file, err)
			}
			cmd := exec.Command("/Applications/Typora 2.app/Contents/MacOS/Typora", file)
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		} else if os.Args[1] == "rm" {
			for _, f := range fs {
				if f.Name() == os.Args[2] {
					_p := path.Join(fileURL, f.Name())
					_dirpath, err := os.Open(_p)
					if err != nil {
						log.Fatal(err)
					}
					defer _dirpath.Close()
					files, err := _dirpath.Readdir(0)
					for _, file := range files {
						if file.Name() == os.Args[3] {
							_p := path.Join(fileURL, os.Args[2])
							file := path.Join(_p, os.Args[3])
							err := os.Remove(file)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}
			}
		}
	}
}
