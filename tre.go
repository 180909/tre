package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var args []string
	args = os.Args[1:]

	switch len(args) {
	case 0:
		e := os.Stderr
		e.WriteString("\nplease input folder name\n\n")
		e.Close()
		os.Exit(1)
	case 1:
		printAll(args[0])
	default:
		e := os.Stderr
		e.WriteString("\nplease don't input folder name more than one\n\n")
		e.Close()
		os.Exit(1)
	}

	os.Exit(0)
}

func printAll(arg string) {
	o := os.Stdout
	defer func() {
		o.Close()
		os.Exit(0)
	}()

	argType := checkArg(arg, o)
	if argType == 2 {
		o.WriteString("\ntre: " + arg + " is not a directory \n")
		o.Close()
		return
	}

	allName(arg, o)

}

func allName(prefix string, o io.StringWriter) {
	fs, err := ioutil.ReadDir(prefix)
	if err != nil {
		e := os.Stderr
		e.WriteString(err.Error() + "\n")
		e.Close()
		os.Exit(1)
	}

	for _, f := range fs {
		o.WriteString(prefix + "/" + f.Name() + "\n")
		if f.IsDir() {
			allName(prefix+"/"+f.Name(), o)
		}
	}
	return
}

// check arg exist or is file or is dir
// if arg not exist, exit program
// if arg is dir, return 1
// if arg is file, return 2
func checkArg(arg string, o io.StringWriter) int {
	s, err := os.Stat(arg)
	if err != nil {
		if os.IsNotExist(err) {
			o.WriteString("\ntre: no such file or directory: " + arg + " \n")

			os.Exit(0)
		}
		e := os.Stderr
		e.WriteString(err.Error())
		e.Close()
		os.Exit(1)
	}
	if s.IsDir() {
		return 1
	} else {
		return 0
	}
}
