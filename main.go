package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	bwt "github.com/abdfnx/bw-templates/langs"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok 👍")
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		body := func(value string) string {
			return gjson.Get(string(c.Body()), value).String()
		}

		bwt.CreateBot(body("name"), body("platform"), body("lang"), body("packageManager"), body("hostService"))

		if body("lang") == "c" {
			bwt.CTemplate(body("name"), body("hostService"))
		} else if body("lang") == "cpp" {
			bwt.CppTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "crystal" {
			bwt.CrystalTemplate(body("name"), body("hostService"))
		} else if body("lang") == "csharp" {
			bwt.CsharpTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "dart" {
			bwt.DartTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "deno" {
			bwt.DenoTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "go" {
			bwt.GoTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "java" {
			bwt.JavaTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "kotlin" {
			bwt.KotlinTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "nim" {
			bwt.NimTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "nodejs" {
			bwt.NodejsTemplate(body("name"), body("packageManager"), body("platform"), body("hostService"), false)
		} else if body("lang") == "typescript" {
			bwt.NodejsTemplate(body("name"), body("packageManager"), body("platform"), body("hostService"), true)
		} else if body("lang") == "php" {
			bwt.PHPTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "python" {
			bwt.PythonTemplate(body("name"), body("platform"), body("packageManager"), body("hostService"))
		} else if body("lang") == "ruby" {
			bwt.RubyTemplate(body("name"), body("platform"), body("hostService"))
		} else if body("lang") == "rust" {
			bwt.RustTemplate(body("name"), body("platform"), body("packageManager"), body("hostService"))
		} else if body("lang") == "swift" {
			bwt.SwiftTemplate(body("name"), body("platform"), body("hostService"))
		}

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
			gh auth logout --hostname github.com
		`, c.GetReqHeaders()["Authorization"], body("username"), body("email"), body("username"), body("name")))

		createRepo.Dir = body("name")
		createRepo.Stdin = os.Stdin
		createRepo.Stdout = os.Stdout
		createRepo.Stderr = os.Stderr

		err := createRepo.Run()

		if err != nil {
			log.Printf("error: %v\n", err)

			return c.SendString("error: " + err.Error())
		}

		os.RemoveAll(body("name"))

		return c.SendString("created successfully 📦")
	})

	log.Fatal(app.Listen(":7050"))
}
