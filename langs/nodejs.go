package templates

import (
	"log"
	"os"
	"path/filepath"
)

func isTypescript(platform string, isTs bool) (string, string, string) {
	if isTs {
		return "nodejs-ts", "ts", MainTSContent(platform)
	} else {
		return "nodejs", "js", MainJSContent(platform)
	}
}

func MainJSContent(platform string) string {
	return Content("main.js", platform+"-nodejs", "", "")
}

func MainTSContent(platform string) string {
	return Content("main.ts", platform+"-nodejs-ts", "", "")
}

func NodejsTemplate(botName, pm, platform string, isTs bool) {
	tmpName, ext, content := isTypescript(platform, isTs)

	mainFile := os.WriteFile(filepath.Join(botName, "src", "main."+ext), []byte(content), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, pm+".dockerfile", platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "nodejs.md")), 0644)
	packageFile := os.WriteFile(filepath.Join(botName, "package.json"), []byte(Content("package.json", platform+"-"+tmpName, "", "")), 0644)

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if packageFile != nil {
		log.Fatal(packageFile)
	}

	if isTs {
		tsConfigFile := os.WriteFile(filepath.Join(botName, "tsconfig.json"), []byte(Content("tsconfig.json", platform+"-nodejs-ts", "", "")), 0644)

		if tsConfigFile != nil {
			log.Fatal(tsConfigFile)
		}
	}

	CheckProject(botName, platform)
}
