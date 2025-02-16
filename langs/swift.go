package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainSwiftContent(platform string) string {
	return Content("Sources/bwbot/main.swift", platform+"-swift", "", "")
}

func BotwaySwiftContent(botName string) string {
	return Content("packages/botway-swift/main.swift", "botway", botName, "")
}

func PackageSwiftFileContent(botName, platform string) string {
	return Content("Package.swift", platform+"-swift", botName, "")
}

func SwiftTemplate(botName, platform string) {
	if err := os.Mkdir(filepath.Join(botName, "Sources"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "Sources", botName), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	mainFile := os.WriteFile(filepath.Join(botName, "Sources", botName, "main.swift"), []byte(MainSwiftContent(platform)), 0644)
	botwaySwiftFile := os.WriteFile(filepath.Join(botName, "Sources", botName, "botway.swift"), []byte(BotwaySwiftContent(botName)), 0644)
	packageSwiftFile := os.WriteFile(filepath.Join(botName, "Package.swift"), []byte(PackageSwiftFileContent(botName, platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "swift.dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "swift.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if botwaySwiftFile != nil {
		log.Fatal(botwaySwiftFile)
	}

	if packageSwiftFile != nil {
		log.Fatal(packageSwiftFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, platform)
}
