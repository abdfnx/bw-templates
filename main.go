package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	bwt "github.com/abdfnx/create-botway-bot/langs"
)

func main() {
	name := os.Args[1]
	platform := os.Args[2]
	lang := os.Args[3]
	packageManager := os.Args[4]
	hostService := os.Args[5]
	authToken := ""

	if len(os.Args[1:]) == 6 {
		authToken = os.Args[6]
	}

	bwt.CreateBot(name, platform, lang, packageManager, hostService)

	if lang == "c" {
		bwt.CTemplate(name, hostService)
	} else if lang == "cpp" {
		bwt.CppTemplate(name, platform, hostService)
	} else if lang == "crystal" {
		bwt.CrystalTemplate(name, hostService)
	} else if lang == "csharp" {
		bwt.CsharpTemplate(name, platform, hostService)
	} else if lang == "dart" {
		bwt.DartTemplate(name, platform, hostService)
	} else if lang == "deno" {
		bwt.DenoTemplate(name, platform, hostService)
	} else if lang == "go" {
		bwt.GoTemplate(name, platform, hostService)
	} else if lang == "java" {
		bwt.JavaTemplate(name, platform, hostService)
	} else if lang == "kotlin" {
		bwt.KotlinTemplate(name, platform, hostService)
	} else if lang == "nim" {
		bwt.NimTemplate(name, platform, hostService)
	} else if lang == "nodejs" {
		bwt.NodejsTemplate(name, packageManager, platform, hostService, false)
	} else if lang == "typescript" {
		bwt.NodejsTemplate(name, packageManager, platform, hostService, true)
	} else if lang == "php" {
		bwt.PHPTemplate(name, platform, hostService)
	} else if lang == "python" {
		bwt.PythonTemplate(name, platform, packageManager, hostService)
	} else if lang == "ruby" {
		bwt.RubyTemplate(name, platform, hostService)
	} else if lang == "rust" {
		bwt.RustTemplate(name, platform, packageManager, hostService)
	} else if lang == "swift" {
		bwt.SwiftTemplate(name, platform, hostService)
	}

	if authToken != "" {
		username := os.Args[7]
		email := os.Args[8]

		createRepo := exec.Command("bash", "-c", fmt.Sprintf(`
			echo %s | gh auth login --with-token
			gh auth setup-git
			git init
			git config --global user.name "%s"
			git config --global user.email "%s"
			git add .
			git commit -m "new botway bot project"
			git branch -M main
			git remote add origin https://github.com/%s/%s.git
			git push -u origin main
		`, authToken, username, email, username, name))

		createRepo.Dir = name
		createRepo.Stdin = os.Stdin
		createRepo.Stdout = os.Stdout
		createRepo.Stderr = os.Stderr

		err := createRepo.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	// os.RemoveAll(name)
}
