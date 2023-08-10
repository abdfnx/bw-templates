package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainCContent() string {
	return Content("src/main.c", "discord-c", "", "")
}

func BWCContent(botName string) string {
	return Content("packages/bwc/main.h", "botway", botName, "")
}

func CRunPsFileContent() string {
	return Content("run.ps1", "discord-c", "", "")
}

func CTemplate(botName string) {
	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.c"), []byte(MainCContent()), 0644)
	botwayHeaderFile := os.WriteFile(filepath.Join(botName, "src", "botway.h"), []byte(BWCContent(botName)), 0644)
	runPsFile := os.WriteFile(filepath.Join(botName, "run.ps1"), []byte(CRunPsFileContent()), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "c-discord.dockerfile", "discord")), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources("discord", "c.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if botwayHeaderFile != nil {
		log.Fatal(botwayHeaderFile)
	}

	if runPsFile != nil {
		log.Fatal(runPsFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, "discord")
}
