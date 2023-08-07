package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
)

func updateFileOnDrive(srv *drive.Service, fileId, newFilename, mtblDir string) {
	file := &drive.File{
		Name: newFilename,
		// Add other metadata here as needed.
	}

	var dir string
	if strings.Contains(file.Name, "rstrs") {
		dir = "extract/"
	} else {
		dir = "transform/"
	}

	filePath := mtblDir + dir + file.Name

	var jsonFile, err = os.Open(filePath)
	if err != nil {
		log.Fatalf("Cannot open JSON file: %v", err)
	}
	defer jsonFile.Close()

	updatedFile, err := srv.Files.Update(fileId, file).Media(jsonFile).Do()
	if err != nil {
		log.Fatalf("Could not update file: %v", err)
	}

	fmt.Printf("Updated file: %s (%s)\n", updatedFile.Name, updatedFile.Id)
}

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Cannot create the Google Drive service: %v", err)
	}

	armsFileId := "1qyMPIOdTXYIqzg7FHkQUPHmMAQVBhKf7" // Replace with your file ID.
	batsFileId := "1dVOy2Acs-yAQ-TWikGwXNNMOl5p1Y3YG"
	lgrstrsFileId := "1qUJzZPVaiLnLxQ6dPH3YeFh0jKwPF-UZ"

	mtblDir := "/Users/Shared/BaseballHQ/resources/"

	updateFileOnDrive(srv, armsFileId, "arms_trp.json", mtblDir)
	updateFileOnDrive(srv, batsFileId, "bats_trp.json", mtblDir)
	updateFileOnDrive(srv, lgrstrsFileId, "lgrstrs.json", mtblDir)
}
