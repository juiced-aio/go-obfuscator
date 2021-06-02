package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: gobfuscate [flags] pkg_name out_path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	if !obfuscate(inPath, outPath) {
		os.Exit(1)
	}
}

func obfuscate(inPath, outPath string) bool {
	log.Println("Creating output directory...")
	gopath, err := filepath.Abs(outPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create output directory:", err)
		return false
	}
	if err := os.Mkdir(gopath, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create output directory:", err)
		return false
	}

	log.Println("Copying files into output directory...")
	if err := CopyDirectory(inPath, gopath); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to copy files into output directory:", err)
		return false
	}

	// var n NameHasher
	// buf := make([]byte, 32)
	// rand.Read(buf)
	// n = buf

	// log.Println("Obfuscating package names...")
	// if err := ObfuscatePackageNames(gopath, n); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Failed to obfuscate package names:", err)
	// 	return false
	// }
	log.Println("Obfuscating strings...")
	if err := ObfuscateStrings(gopath); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to obfuscate strings:", err)
		return false
	}
	// log.Println("Obfuscating symbols...")
	// if err := ObfuscateSymbols(gopath, n); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Failed to obfuscate symbols:", err)
	// 	return false
	// }

	log.Println("Obfuscation complete.")
	return true
}
