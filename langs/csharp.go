package templates

import (
	"log"
	"os"
	"path/filepath"
)

func MainCsContent(platform string) string {
	return Content("src/Main.cs", platform+"-csharp", "", "")
}

func BotCSharpProj(platform string) string {
	return Content(platform+"-csharp.csproj", platform+"-csharp", "", "")
}

func CsharpTemplate(botName, platform, hostService string) {
	mainFile := os.WriteFile(filepath.Join(botName, "src", "Main.cs"), []byte(MainCsContent(platform)), 0644)
	csprojFile := os.WriteFile(filepath.Join(botName, botName+".csproj"), []byte(BotCSharpProj(platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "csharp.dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "csharp.md")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if csprojFile != nil {
		log.Fatal(csprojFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, platform)

}
