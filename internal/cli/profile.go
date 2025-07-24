package cli

import (
	"fmt"
	"mws/internal/storage"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	profileName  string
	userName     string
	projectName  string
	profileStore = storage.NewFileStorage("")
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Profile management operations",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateProfileArgs(); err != nil {
			return err
		}
		return profileStore.Create(profileName, userName, projectName)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get profile details",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := profileStore.Get(profileName)
		if err != nil {
			return err
		}
		fmt.Printf("Profile: %s\nUser: %s\nProject: %s\n",
			filepath.Base(profileName), p.User, p.Project)
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles := profileStore.List()
		if len(profiles) == 0 {
			fmt.Println("No profiles found")
			return
		}
		for name, p := range profiles {
			fmt.Printf("%s:\n  User: %s\n  Project: %s\n\n", name, p.User, p.Project)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profileStore.Delete(profileName)
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Available command: 
  Create profile: profile create --name=NAME --user=USER --project=PROJECT
  Get profile: profile get --name=NAME
  List all profiles: profile list
  Delete profile: profile delete --name=NAME
  Print commands documentation: help`)
	},
}

func init() {
	createCmd.Flags().StringVarP(&profileName, "name", "n", "", "Profile name")
	createCmd.Flags().StringVarP(&userName, "user", "u", "", "Username")
	createCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("user")
	createCmd.MarkFlagRequired("project")

	for _, c := range []*cobra.Command{getCmd, deleteCmd} {
		c.Flags().StringVarP(&profileName, "name", "n", "", "Profile name")
		c.MarkFlagRequired("name")
	}

	profileCmd.AddCommand(createCmd, getCmd, listCmd, deleteCmd)
}

func validateProfileArgs() error {
	if profileName == "" || userName == "" || projectName == "" {
		return fmt.Errorf("all arguments are required")
	}
	return nil
}
