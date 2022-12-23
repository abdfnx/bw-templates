package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainRsContent(platform string) string {
	return Content("src/main.rs", platform+"-rust", "", "")
}

func CargoFileContent(botName, platform string) string {
	return Content("Cargo.toml", platform+"-rust", botName, "")
}

func RustTemplate(botName, platform, pm, hostService string) {
	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rs"), []byte(MainRsContent(platform)), 0644)
	cargoFile := os.WriteFile(filepath.Join(botName, "Cargo.toml"), []byte(CargoFileContent(botName, platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService, pm+".dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "rust.md")), 0644)

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if cargoFile != nil {
		log.Fatal(cargoFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	CheckProject(botName, platform)
}
