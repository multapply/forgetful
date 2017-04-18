package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func scan(ignore []string, target string, caseSensitive bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		// skip anything we want to ignore
		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignore {
				if d == dir {
					return filepath.SkipDir
				}
			}
		} else { // otherwise we can parse
			err := parse(path, target, caseSensitive)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// parse - opens file 'filename' and attempts to parse it
func parse(filename string, target string, caseSensitive bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		// scan line for targets
		line := strings.TrimSpace(scanner.Text())
		if !caseSensitive {
			line = strings.ToUpper(line)
			target = strings.ToUpper(target)
		}
		if len(line) > 1 && line[:2] == "//" {
			if strings.Contains(line, target) {
				fmt.Printf("'%s' found at %s:%d\n", target, filename, i)
			}
		}
		i++
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	// check for no args
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("[ERROR] You must supply at least a path to start scanning from...")
		return
	}

	if len(args) > 4 {
		fmt.Println("[ERROR] You supplied too many arguments!")
		fmt.Println("        Argument format is: <startingpath> <target> <OPTIONAL -C flag>")
		return
	}

	root := os.Args[1]
	caseSensitive := false
	var target string
	for idx, arg := range args {
		if arg == "-C" {
			caseSensitive = true
			switch idx {
			case 3:
				target = args[2]
			case 2:
				if len(args) == 4 {
					target = args[3]
				} else {
					target = "TODO"
				}
			}
		}
	}

	if len(args) == 2 {
		target = "TODO"
	}

	fmt.Printf("Scanning for '%s'...\n\n", target)

	ignore := []string{".git"}
	err := filepath.Walk(root, scan(ignore, target, caseSensitive))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDone scanning!\n")
}
