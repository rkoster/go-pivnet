package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/olekukonko/tablewriter"
	"github.com/pivotal-cf-experimental/go-pivnet"
)

type FileGroupsCommand struct {
	ProductSlug    string `long:"product-slug" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string `long:"release-version" description:"Release version e.g. 0.1.2-rc1"`
}

type FileGroupCommand struct {
	ProductSlug    string `long:"product-slug" description:"Product slug e.g. p-mysql" required:"true"`
	ReleaseVersion string `long:"release-version" description:"Release version e.g. 0.1.2-rc1" required:"true"`
	FileGroupID    int    `long:"file-group-id" description:"Filegroup ID e.g. 1234" required:"true"`
}

type DeleteFileGroupCommand struct {
	ProductSlug string `long:"product-slug" description:"Product slug e.g. p-mysql" required:"true"`
	FileGroupID int    `long:"file-group-id" description:"File group ID e.g. 1234" required:"true"`
}

func (command *FileGroupsCommand) Execute([]string) error {
	client := NewClient()

	if command.ReleaseVersion == "" {
		fileGroups, err := client.FileGroups.List(
			command.ProductSlug,
		)
		if err != nil {
			return err
		}
		return printFileGroups(fileGroups)
	}

	releases, err := client.Releases.List(command.ProductSlug)
	if err != nil {
		return err
	}

	var release pivnet.Release
	for _, r := range releases {
		if r.Version == command.ReleaseVersion {
			release = r
			break
		}
	}

	if release.Version != command.ReleaseVersion {
		return fmt.Errorf("release not found")
	}

	fileGroups, err := client.FileGroups.ListForRelease(
		command.ProductSlug,
		release.ID,
	)
	if err != nil {
		return err
	}

	return printFileGroups(fileGroups)
}

func printFileGroups(fileGroups []pivnet.FileGroup) error {
	switch Pivnet.Format {

	case printAsTable:
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Name",
			"Product File Names",
		})

		for _, fileGroup := range fileGroups {
			var productFileNames []string

			for _, productFile := range fileGroup.ProductFiles {
				productFileNames = append(productFileNames, productFile.Name)
			}

			fileGroupAsString := []string{
				strconv.Itoa(fileGroup.ID),
				fileGroup.Name,
				strings.Join(productFileNames, " "),
			}
			table.Append(fileGroupAsString)
		}
		table.Render()
		return nil
	case printAsJSON:
		b, err := json.Marshal(fileGroups)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", string(b))
		return nil
	case printAsYAML:
		b, err := yaml.Marshal(fileGroups)
		if err != nil {
			return err
		}

		fmt.Printf("---\n%s\n", string(b))
		return nil
	}

	return nil
}

func (command *FileGroupCommand) Execute([]string) error {
	client := NewClient()

	releases, err := client.Releases.List(command.ProductSlug)
	if err != nil {
		return err
	}

	var release pivnet.Release
	for _, r := range releases {
		if r.Version == command.ReleaseVersion {
			release = r
			break
		}
	}

	if release.Version != command.ReleaseVersion {
		return fmt.Errorf("release not found")
	}

	fileGroup, err := client.FileGroups.Get(
		command.ProductSlug,
		release.ID,
		command.FileGroupID,
	)
	if err != nil {
		return err
	}

	return printFileGroups([]pivnet.FileGroup{fileGroup})
}

func (command *DeleteFileGroupCommand) Execute([]string) error {
	client := NewClient()

	fileGroup, err := client.FileGroups.Delete(
		command.ProductSlug,
		command.FileGroupID,
	)
	if err != nil {
		return err
	}

	return printFileGroups([]pivnet.FileGroup{fileGroup})
}
