// Package decls lists contract, interface, and library declarations in Solidity
// source. It is a lexer, not a parser: comments and strings are skipped so a
// mention of "contract" inside either is not a declaration.
package decls

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Decl is one top-level type declaration.
type Decl struct {
	// Kind is "contract", "abstract contract", "interface", or "library".
	Kind string
	Name string
}

// Extract parses src and returns declarations in source order.
func Extract(src []byte) ([]Decl, error) {
	var decls []Decl
	i := 0
	for i < len(src) {
		word, next, ok := nextIdent(src, i)
		if !ok {
			break
		}
		switch word {
		case "abstract":
			second, after, ok2 := nextIdent(src, next)
			if ok2 && second == "contract" {
				name, afterName, okName := nextIdent(src, after)
				if !okName {
					return decls, nil
				}
				decls = append(decls, Decl{Kind: "abstract contract", Name: name})
				i = afterName
				continue
			}
			i = next
		case "contract", "interface", "library":
			name, afterName, okName := nextIdent(src, next)
			if !okName {
				return decls, nil
			}
			decls = append(decls, Decl{Kind: word, Name: name})
			i = afterName
		default:
			i = next
		}
	}
	return decls, nil
}

// ExtractPaths reads Solidity files and directories. Names are de-duplicated,
// first occurrence wins, files are visited in path order.
func ExtractPaths(paths []string) ([]Decl, error) {
	files, err := solFiles(paths)
	if err != nil {
		return nil, err
	}
	var decls []Decl
	seen := map[string]bool{}
	for _, file := range files {
		buf, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		found, err := Extract(buf)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		for _, d := range found {
			key := d.Kind + "\t" + d.Name
			if seen[key] {
				continue
			}
			seen[key] = true
			decls = append(decls, d)
		}
	}
	return decls, nil
}

// Format writes one declaration per line: "<kind>\t<name>".
func Format(decls []Decl) string {
	var b strings.Builder
	for _, d := range decls {
		b.WriteString(d.Kind)
		b.WriteByte('\t')
		b.WriteString(d.Name)
		b.WriteByte('\n')
	}
	return b.String()
}

func solFiles(paths []string) ([]string, error) {
	var files []string
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			files = append(files, path)
			continue
		}
		err = filepath.Walk(path, func(p string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			if strings.HasSuffix(p, ".sol") {
				files = append(files, p)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(files)
	return files, nil
}

// nextIdent returns the next identifier at or after i, skipping whitespace,
// comments, and strings.
func nextIdent(src []byte, i int) (string, int, bool) {
	for i < len(src) {
		r, size := utf8.DecodeRune(src[i:])
		if r == utf8.RuneError && size == 1 {
			i++
			continue
		}
		if unicode.IsSpace(r) {
			i += size
			continue
		}
		if r == '/' && i+1 < len(src) {
			switch src[i+1] {
			case '/':
				i += 2
				for i < len(src) && src[i] != '\n' {
					i++
				}
				continue
			case '*':
				i += 2
				for i+1 < len(src) && !(src[i] == '*' && src[i+1] == '/') {
					i++
				}
				if i+1 < len(src) {
					i += 2
				} else {
					i = len(src)
				}
				continue
			}
		}
		if r == '"' {
			i = skipString(src, i+size)
			continue
		}
		if isIdentStart(r) {
			start := i
			i += size
			for i < len(src) {
				r2, size2 := utf8.DecodeRune(src[i:])
				if !isIdentCont(r2) {
					break
				}
				i += size2
			}
			return string(src[start:i]), i, true
		}
		i += size
	}
	return "", i, false
}

func skipString(src []byte, i int) int {
	for i < len(src) {
		if src[i] == '\\' && i+1 < len(src) {
			i += 2
			continue
		}
		if src[i] == '"' || src[i] == '\n' {
			return i + 1
		}
		i++
	}
	return i
}

func isIdentStart(r rune) bool {
	return r == '_' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isIdentCont(r rune) bool {
	return isIdentStart(r) || (r >= '0' && r <= '9')
}
