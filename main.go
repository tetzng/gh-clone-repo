package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gh clone-repo <repo | owner/repo | repo url> [-u, --upstream-remote-name <string>]")
		os.Exit(1)
	}

	repoArg := os.Args[1]
	otherArgs := os.Args[2:]

	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.HasPrefix(repoArg, "http://") || strings.HasPrefix(repoArg, "https://") {
		parsedUrl, err := url.Parse(repoArg)
		if err != nil {
			fmt.Println("Failed to parse the URL.")
			os.Exit(1)
		}
		repoArg = strings.TrimLeft(parsedUrl.Path, "/")
	}

	split := strings.Split(repoArg, "/")
	if len(split) == 1 {
		split = append([]string{response.Login}, split[0])
	} else if len(split) != 2 {
		fmt.Println("Argument does not match 'owner/repo', 'repo', or 'repo url'.")
		os.Exit(1)
	}

	owner, repo := split[0], split[1]

	if _, err := os.Stat(owner); os.IsNotExist(err) {
		err = os.Mkdir(owner, 0755)
		if err != nil {
			fmt.Printf("Failed to create directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created directory: %s\n", owner)
	}

	err = os.Chdir(owner)
	if err != nil {
		fmt.Printf("Failed to change directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Moved to directory: %s\n", owner)

	cmd := exec.Command("gh", append([]string{"repo", "clone", filepath.Join(owner, repo)}, otherArgs...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to clone repository: %v\n", err)
		os.Exit(1)
	}
}
