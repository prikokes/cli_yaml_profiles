package cmd

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

func createProfile() error {
	cmd := flag.NewFlagSet("create", flag.ExitOnError)
	name := cmd.String("name", "", "Profile name")
	user := cmd.String("user", "", "Username")
	project := cmd.String("project", "", "Project name")

	err := cmd.Parse(os.Args[3:])
	if err != nil {
		return err
	}

	profile := Profile{*user, *project}

	data, err := yaml.Marshal(profile)
	if err != nil {
		return err
	}

	return os.WriteFile(*name+".yaml", data, 0644)
}

func getProfile() error {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	name := cmd.String("name", "", "Profile name")
	err := cmd.Parse(os.Args[3:])

	if err != nil {
		return err
	}

	_, err = os.Stat(*name)
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(*name)

	var profile Profile

	if err := yaml.Unmarshal(data, &profile); err != nil {
		panic(err)
	}

	fmt.Printf("User: %+v\n", profile.User)
	fmt.Println("Project:", profile.Project)

	return nil
}

func deleteProfile() error {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	name := cmd.String("name", "", "Profile name")
	err := cmd.Parse(os.Args[3:])

	if err != nil {
		return err
	}

	err = os.Remove(*name)
	if err != nil {
		return err
	}
	return nil
}

func listProfiles() error {
	files, err := filepath.Glob("*.yaml")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		_, err := os.Stat(file)
		if err != nil {
			panic(err)
		}
		data, err := os.ReadFile(file)

		var profile Profile

		if err := yaml.Unmarshal(data, &profile); err != nil {
			panic(err)
		}
		fmt.Printf("User: %+v\n", profile.User)
		fmt.Println("Project:", profile.Project)
	}
	return nil
}

func profileHandler() error {
	switch os.Args[2] {
	case "create":
		err := createProfile()
		if err != nil {
			return err
		}
	case "get":
		err := getProfile()
		if err != nil {
			return err
		}
	case "delete":
		err := deleteProfile()
		if err != nil {
			return err
		}
	case "list":
		err := listProfiles()
		if err != nil {
			return err
		}
	default:
		panic("Error! No such command provided still.")
	}

	return nil
}

func main() {
	fmt.Println(os.Args[0], os.Args[1], os.Args[2])
	fmt.Println(len(os.Args))

	if os.Args[1] != "profile" && os.Args[1] != "help" {
		panic("Error! No such command provided still.")
	}

	if os.Args[1] == "profile" {
		err := profileHandler()
		if err != nil {
			panic(err)
		}
	}

	if os.Args[1] == "help" {
		fmt.Println("Some information about the commands.")
	}
}
