package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainPHPContent(platform string) string {
	return Content("src/main.php", platform+"-php", "", "")
}

func BotwayPHPContent() string {
	return Content("packages/bw-php/main.php", "botway", "", "")
}

func ComposerFileContent(botName, platform string) string {
	return Content("composer.json", platform+"-php", botName, "")
}

func PHPTemplate(botName, platform, hostService string) {
	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.php"), []byte(MainPHPContent(platform)), 0644)
	botwayFile := os.WriteFile(filepath.Join(botName, "src", "botway.php"), []byte(BotwayPHPContent()), 0644)
	composerFile := os.WriteFile(filepath.Join(botName, "composer.json"), []byte(ComposerFileContent(botName, platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService, "php.dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "php.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if botwayFile != nil {
		log.Fatal(botwayFile)
	}

	if composerFile != nil {
		log.Fatal(composerFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, "discord")
}
