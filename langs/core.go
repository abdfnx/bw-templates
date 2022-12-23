package templates

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/resto/core/api"
	"github.com/tidwall/gjson"
)

func BotSecrets(platform string) string {
	if strings.Contains(platform, "discord") {
		return "DISCORD_TOKEN DISCORD_CLIENT_ID\n# You can add guild ids of your servers by adding ARG SERVER_NAME_GUILD_ID"
	} else if strings.Contains(platform, "telegram") {
		return "TELEGRAM_TOKEN"
	} else if strings.Contains(platform, "slack") {
		return "SLACK_TOKEN SLACK_APP_TOKEN SLACK_SIGNING_SECRET"
	} else if strings.Contains(platform, "twitch") {
		return "TWITCH_OAUTH_TOKEN TWITCH_CLIENT_ID TWITCH CLIENT_SECRET"
	}

	return "" + platform
}

func Content(arg, templateName, botName, platform string) string {
	org := "botwayorg"

	if templateName == "botway" {
		org = "abdfnx"
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", org, templateName, arg)
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404: " + url)
		os.Exit(0)
	}

	// TODO: fix botway c++ telegram template
	if strings.Contains(respone, "#include <{{.BotName}}/{{.BotName}}.h>") && strings.Contains(platform, "telegram") {
		respone = strings.ReplaceAll(respone, "#include <{{.BotName}}/{{.BotName}}.h>", "")
	} else if strings.Contains(respone, `#include "botway/botway.hpp"`) && strings.Contains(platform, "telegram") {
		respone = strings.ReplaceAll(respone, `#include "botway/botway.hpp"`, `#include "botway.hpp"`)
	} else if strings.Contains(arg, "pubspec.yaml") {
		respone = strings.ReplaceAll(respone, "{{.BotName}}", strings.ReplaceAll(botName, "-", ""))
	}

	respone = strings.ReplaceAll(respone, "{{.BotName}}", botName)

	author := gjson.Get(string(constants.BotwayConfig), "github.username").String()

	if author == "" {
		author = "botway"
	}

	respone = strings.ReplaceAll(respone, "{{.Author}}", author)

	respone = strings.ReplaceAll(respone, "{{.BotSecrets}}", BotSecrets(platform))

	return respone
}

func CheckProject(botName, botType string) {
	if _, err := os.Stat(botName); !os.IsNotExist(err) {
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("SUCCESS"))
		fmt.Println(constants.SUCCESS_FOREGROUND.Render(" " + botName + " Created successfully 🎉"))
		fmt.Print(constants.SUCCESS_BACKGROUND.Render("NEXT"))
		fmt.Println(" Now, run " + constants.COMMAND_FOREGROUND.Render("botway tokens set --"+botType+" "+botName) + " command to add tokens of your bot 🔑")
	}
}
