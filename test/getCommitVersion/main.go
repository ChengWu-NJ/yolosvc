package main

import (
	_ "embed"
	"fmt"
	"runtime/debug"
)

var Commit1 = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("%#v", info.Main)
		for _, setting := range info.Settings {
			fmt.Printf("%#v\n", setting)
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

//go:generate sh -c "printf %s $(go list -m -versions -json $(go list -m)@$(git branch --show-current) | jq -r .Version) > commit.txt"
//go:embed commit.txt
var CommitVersion string

func main() {
	fmt.Println(Commit1)
	fmt.Println(CommitVersion)
}
