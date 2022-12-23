package templates

import (
	"log"
	"os"
	"path/filepath"
)

func FindDppCmakeContent() string {
	return Content("cmake/FindDPP.cmake", "discord-cpp", "", "")
}

func BWCPPFileContent(botName string) string {
	return Content("packages/bwpp/main.hpp", "botway", botName, "")
}

func MainIncludeFileContent() string {
	return Content("include/bwbot/bwbot.h", "discord-cpp", "", "")
}

func MainCppContent(botName, platform string) string {
	return Content("src/main.cpp", platform+"-cpp", botName, "")
}

func DotDockerIgnoreContent() string {
	return Content(".dockerignore", "discord-cpp", "", "")
}

func CmakeListsContent(botName, platform string) string {
	return Content("CMakeLists.txt", platform+"-cpp", botName, "")
}

func RunPsFileContent(platform string) string {
	return Content("run.ps1", "discord-cpp", "", "")
}

func CppTemplate(botName, platform, hostService string) {
	if platform == "discord" {
		if err := os.Mkdir(filepath.Join(botName, "cmake"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "include"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "include", "botway"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "include", botName), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	botwayHeader := os.WriteFile(filepath.Join(botName, "include", "botway", "botway.hpp"), []byte((BWCPPFileContent(botName))), 0644)
	dotDockerIgnoreFile := os.WriteFile(filepath.Join(botName, ".dockerignore"), []byte(DotDockerIgnoreContent()), 0644)
	cmakeListsFile := os.WriteFile(filepath.Join(botName, "CMakeLists.txt"), []byte(CmakeListsContent(botName, platform)), 0644)
	runPsFile := os.WriteFile(filepath.Join(botName, "run.ps1"), []byte(RunPsFileContent(platform)), 0644)
	mainFile := os.WriteFile(filepath.Join(botName, "src", "main.cpp"), []byte(MainCppContent(botName, platform)), 0644)
	dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService, "cmake"+platform, platform)), 0644)
	resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "cpp.md")), 0644)

	if platform == "discord" {
		findDPPCmakeFile := os.WriteFile(filepath.Join(botName, "cmake", "FindDPP.cmake"), []byte(FindDppCmakeContent()), 0644)
		mainIncludeFile := os.WriteFile(filepath.Join(botName, "include", botName, botName+".h"), []byte((MainIncludeFileContent())), 0644)

		if mainIncludeFile != nil {
			log.Fatal(mainIncludeFile)
		}

		if findDPPCmakeFile != nil {
			log.Fatal(findDPPCmakeFile)
		}
	}

	if botwayHeader != nil {
		log.Fatal(botwayHeader)
	}

	if runPsFile != nil {
		log.Fatal(runPsFile)
	}

	if dotDockerIgnoreFile != nil {
		log.Fatal(dotDockerIgnoreFile)
	}

	if cmakeListsFile != nil {
		log.Fatal(cmakeListsFile)
	}

	if mainFile != nil {
		log.Fatal(mainFile)
	}

	if dockerFile != nil {
		log.Fatal(dockerFile)
	}

	if resourcesFile != nil {
		log.Fatal(resourcesFile)
	}

	CheckProject(botName, platform)

}
