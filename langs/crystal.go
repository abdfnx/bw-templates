package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainCrContent() string {
	return Content("src/main.cr", "discord-crystal", "", "")
}

func ShardFileContent(botName string) string {
	return Content("shard.yml", "discord-crystal", botName, "")
}

func CrystalTemplate(botName, hostService string) {
	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.cr"), []byte(MainCrContent()), 0644)
	shardFile := os.WriteFile(filepath.Join(botName, "shard.yml"), []byte(ShardFileContent(botName)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "crystal.dockerfile", "discord")), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources("discord", "crystal.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if shardFile != nil {
		log.Fatal(shardFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, "discord")
}
