package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var registry = make(map[string]IStructure)

func createStructure(t byte) IStructure {
	switch t {
	case 'S': return NewStack()
	case 'Q': return NewQueue()
	case 'F': return NewSList()
	case 'L': return NewDList()
	case 'M': return NewArray()
	}
	return nil
}

func saveToFile(filename string) {
	out, err := os.Create(filename)
	if err != nil { return }
	defer out.Close()
	for name, s := range registry {
		fmt.Fprintf(out, "%c %s %s\n", s.GetType(), name, s.Save())
	}
	fmt.Println("Saved to", filename)
}

func loadFromFile(filename string) {
	in, err := os.Open(filename)
	if err != nil { return }
	defer in.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { continue }
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 3 { continue }
		t := parts[0][0]
		name := parts[1]
		data := parts[2]
		if _, exists := registry[name]; !exists {
			if s := createStructure(t); s != nil {
				registry[name] = s
			}
		}
		if s, exists := registry[name]; exists {
			s.Load(data)
		}
	}
	fmt.Println("Loaded from", filename)
}

func executeCommand(query string) {
	parts := strings.Fields(query)
	if len(parts) == 0 { return }
	cmd := parts[0]

	if cmd == "PRINT" {
		if len(parts) < 2 { return }
		if s, ok := registry[parts[1]]; ok { s.Print() } else { fmt.Println("Structure not found") }
		return
	}

	if len(parts) < 2 { return }
	name := parts[1]
	t := cmd[0]
	if _, exists := registry[name]; !exists {
		if s := createStructure(t); s != nil { registry[name] = s } else { return }
	}
	registry[name].Execute(cmd, parts[2:])
}

func main() {
	var filename, query string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--file" && i+1 < len(os.Args) { filename = os.Args[i+1]; i++ }
		if os.Args[i] == "--query" && i+1 < len(os.Args) { query = os.Args[i+1]; i++ }
	}

	if filename != "" { loadFromFile(filename) }

	if query != "" {
		executeCommand(query)
	} else {
		fmt.Println("DBMS ready. Type EXIT to quit.")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() { break }
			line := scanner.Text()
			if line == "EXIT" { break }
			executeCommand(line)
		}
	}

	if filename != "" { saveToFile(filename) }
}