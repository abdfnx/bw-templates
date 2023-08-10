package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainDartContent(platform string) string {
	return Content("src/main.dart", platform+"-dart", "", "")
}

func PubspecFileContent(botName, platform string) string {
	return Content("pubspec.yaml", platform+"-dart", botName, "")
}

func DartTemplate(botName, platform string) {

	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.dart"), []byte(MainDartContent(platform)), 0644)
	pubspecFile := os.WriteFile(filepath.Join(botName, "pubspec.yaml"), []byte(PubspecFileContent(botName, platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "dart.dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "dart.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if pubspecFile != nil {
		log.Fatal(pubspecFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, platform)

}
