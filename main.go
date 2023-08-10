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
	authToken := ""

	if len(os.Args[1:]) >= 5 {
		authToken = os.Args[5]
	}

	bwt.CreateBot(name, platform, lang, packageManager)

	if lang == "c" {
		bwt.CTemplate(name)
	} else if lang == "cpp" {
		bwt.CppTemplate(name, platform)
	} else if lang == "crystal" {
		bwt.CrystalTemplate(name)
	} else if lang == "csharp" {
		bwt.CsharpTemplate(name, platform)
	} else if lang == "dart" {
		bwt.DartTemplate(name, platform)
	} else if lang == "deno" {
		bwt.DenoTemplate(name, platform)
	} else if lang == "go" {
		bwt.GoTemplate(name, platform)
	} else if lang == "java" {
		bwt.JavaTemplate(name, platform)
	} else if lang == "kotlin" {
		bwt.KotlinTemplate(name, platform)
	} else if lang == "nim" {
		bwt.NimTemplate(name, platform)
	} else if lang == "nodejs" {
		bwt.NodejsTemplate(name, packageManager, platform, false)
	} else if lang == "typescript" {
		bwt.NodejsTemplate(name, packageManager, platform, true)
	} else if lang == "php" {
		bwt.PHPTemplate(name, platform)
	} else if lang == "python" {
		bwt.PythonTemplate(name, platform, packageManager)
	} else if lang == "ruby" {
		bwt.RubyTemplate(name, platform)
	} else if lang == "rust" {
		bwt.RustTemplate(name, platform, packageManager)
	} else if lang == "swift" {
		bwt.SwiftTemplate(name, platform)
	}

	if authToken != "" {
		username := os.Args[6]
		email := os.Args[7]

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
