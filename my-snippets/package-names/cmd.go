// The project layout is as follows:
//
// - package folders
// - each folder includes a file with package name different from folder name
// - root file has name cmd.go and includes main package
// - nested packages are imported by their respective folder names, not package names
// - vscode autocomplete recognizes names of imported packages
//   (i.e, packagefolder1 import provides packagename1 package)

package main

import (
	"fmt"
	"my.snippets/package-names/packagefolder1"
	"my.snippets/package-names/packagefolder2"
)

func main() {
	fmt.Println(packagename1.GetCount(), packagename2.GetCount())
}
