package main

import (
	"bufio"
	"fmt"
	gops "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)



func cd(dir string) {
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err.Error())
	}
}

func pwd() string {
	if wd, err := os.Getwd(); err != nil {
		fmt.Println(err.Error())
		return ""
	} else {
		fmt.Println(wd)
		return wd
	}
}

func echo(args []string) {
	for _, word := range args {
		fmt.Print(word + " ")
	}
	fmt.Println()
}

func ps() {
	if pcs, err := gops.Processes(); err != nil {
		fmt.Println(err.Error())
	} else {
		for _, pc := range pcs {
			fmt.Println(pc.Pid(), pc.Executable())
		}
	}
}

func kill(pid int) {
	if p, err := os.FindProcess(pid); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		if err := p.Kill(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func execute(query string) {
	commands := strings.Split(query, " | ")
	for _, command := range commands {
		commandSlice := strings.Split(command, " ")
		switch commandSlice[0] {
		case "pwd":
			pwd()
		case "cd":
			cd(commandSlice[1])
		case "echo":
			echo(commandSlice[1:])
		case "ps":
			ps()
		case "kill":
			if pid, err := strconv.Atoi(commandSlice[1]); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				kill(pid)
			}
		}
	}

}

func shell() {
	wd, _ := os.Getwd()
	fmt.Print(wd + "> ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); fmt.Print(wd + "> ") {
		if query := scanner.Text(); query != "\\quit" {
			execute(query)
		} else {
			break
		}
		wd, _ = os.Getwd()
	}
}

func main() {
	shell()
}