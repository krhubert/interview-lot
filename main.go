package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"parking_lot/database"
	"parking_lot/exec"
	"parking_lot/lot/parser"
	"parking_lot/shell"
	"parking_lot/version"
)

// Storage types.
const (
	MemoryStorage string = "memory"
	FileStorage   string = "file"
)

// Program flags and args.
var (
	storage      = flag.String("storage", MemoryStorage, "type of storage [memory|file]")
	storageFile  = flag.String("storage-file", "", "file to store database")
	printVersion = flag.Bool("version", false, "print version and exit")
	sourceFile   string
)

// parseFlags parses flags and validate them. It returns on invalid flag.
func parseFlags() error {
	flag.Parse()

	if len(os.Args) > 2 {
		return fmt.Errorf("give only one source file")
	}

	if len(os.Args) == 2 {
		sourceFile = os.Args[1]
		// COMMENT for tests
		// if !strings.HasSuffix(sourceFile, ".lot") {
		// 	return fmt.Errorf("file %s is not lot source file", filepath.Base(sourceFile))
		// }
		f, err := os.Open(sourceFile)
		if err != nil {
			return fmt.Errorf("source file %s: %s", filepath.Base(sourceFile), err)
		}
		f.Close()
	}

	if *storage != MemoryStorage && *storage != FileStorage {
		return fmt.Errorf("invalid %q storage type: use \"memory\" or \"file\"", *storage)
	}

	if *storage == MemoryStorage && *storageFile != "" {
		return fmt.Errorf(`--storage-file flag is not allowed with "memory" storage`)
	}

	if *storage == FileStorage && *storageFile == "" {
		return fmt.Errorf(`with "file" storage type flag --storage-file is required`)
	}

	return nil
}

// initDatabase creates database based on set flags
func initDatabase() (*database.Database, error) {
	var db *database.Database

	if *storage == MemoryStorage {
		db = database.NewDatabase(database.NewMemoryWriter())
	} else if *storage == FileStorage {
		w, err := database.NewFileWriter(*storageFile)
		if err != nil {
			return nil, fmt.Errorf("creating file storage error: %s", err)
		}
		db = database.NewDatabase(w)
	}
	return db, nil
}

// processSourceFile parse and process lot source file.
func processSourceFile(db *database.Database) error {
	e := exec.NewExecutor()

	content, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return fmt.Errorf("reading source code error: %s", err)
	}

	program, err := parser.Parse(string(content))
	if err != nil {
		return fmt.Errorf("parsing source code error: %s", err)
	}

	e.Execute(program, db)
	return nil
}

// startShell starts interactive shell for processing lot source.
func startShell(db *database.Database) error {
	e := exec.NewExecutor()

	shell := shell.NewShell()
	for {
		line, err := shell.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}

		program, err := parser.Parse(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing error: %s\n", err)
			continue
		}

		e.Execute(program, db)
	}
}

func main() {
	if err := parseFlags(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *printVersion {
		fmt.Println("parking_lot", version.Version)
		os.Exit(0)
	}

	db, err := initDatabase()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if sourceFile != "" {
		err = processSourceFile(db)
	} else {
		err = startShell(db)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
