package main

import (
	"flag"
	"fmt"
	v4 "jsonschema/drafts/4"
	"os"
	"path"
	"strings"
)

func main() {
	schema := flag.String("scheme", "4", "draft version to use")
	flag.Parse()

	if *schema == "4" {
		ns := v4.New()
		loaded := map[string]bool{}

		for i := 0; i < flag.NArg(); i++ {
			arg := flag.Arg(i)

			if strings.HasPrefix(arg, "./") {
				arg = arg[2:]
			}

			if _, ok := loaded[arg]; ok {
				continue
			}

			if strings.HasSuffix(arg, ".json") {
				_, err := ns.Read(arg)

				if err == nil {
					fmt.Println(
						Color("").
							Cyan(fmt.Sprintf("[%s] => ", arg)).
							Text("loaded...✅"),
					)
				} else {
					fmt.Println(
						Color("").
							Red(fmt.Sprintf("[%s] => ", arg)).
							Text("error...❌"),
					)

					fmt.Println(
						Color("\t").Red("-- " + err.Error()),
					)
				}
			} else if entries, err := os.ReadDir(arg); err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						continue
					}

					if !strings.HasSuffix(entry.Name(), ".json") {
						continue
					}

					path := path.Join(arg, entry.Name())

					if _, ok := loaded[path]; ok {
						continue
					}

					_, err := ns.Read(path)

					if err == nil {
						loaded[path] = true
						fmt.Println(
							Color("").
								Cyan(fmt.Sprintf("[%s] => ", path)).
								Text("loaded...✅"),
						)
					} else {
						fmt.Println(
							Color("").
								Red(fmt.Sprintf("[%s] => ", arg)).
								Text("error...❌"),
						)

						fmt.Println(
							Color("\t").Red("-- " + err.Error()),
						)
					}
				}
			}

			loaded[arg] = true
		}
	}
}
