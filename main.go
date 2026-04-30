package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create a stats.wd3 object
	stats := Stats{}
	stats.Referrals = make(map[string]struct {
		From      string
		Visits    int
		APIVisits int
	})
	// create a fiber app
	app := fiber.New()

	// allow for base path - default to /
	var basePath string
	basePath = os.Getenv("BASE_PATH")
	group := app.Group("/" + basePath)

	// if set, add leading slash for replace in page endpoints
	if basePath != "" {
		basePath = "/" + basePath
	}

	// api endpoints
	group.Get("/api/wisdom/dispense", func(c *fiber.Ctx) error {
		stats.APIVisits++
		if c.Query("re") != "" {
			ref := stats.Referrals[c.Query("re")]
			ref.From = c.Query("re")
			ref.APIVisits++
			stats.Referrals[c.Query("re")] = ref
		}
		return c.SendString(generateWisdom())
	})
	group.Get("/api/wisdom/json", func(c *fiber.Ctx) error {
		stats.APIVisits++
		if c.Query("re") != "" {
			ref := stats.Referrals[c.Query("re")]
			ref.From = c.Query("re")
			ref.APIVisits++
			stats.Referrals[c.Query("re")] = ref
		}
		return c.SendString("{\"wisdom\":\"" + generateWisdom() + "\"}")
	})
	group.Get("/api/stats/json", func(c *fiber.Ctx) error {
		jsonBytes, _ := json.MarshalIndent(stats, "", "  ")
		return c.SendString(string(jsonBytes))
	})

	// page endpoints
	group.Get("/", func(c *fiber.Ctx) error {
		stats.Visits++
		if c.Query("re") != "" {
			ref := stats.Referrals[c.Query("re")]
			ref.From = c.Query("re")
			ref.Visits++
			stats.Referrals[c.Query("re")] = ref
		}
		fBytes, _ := os.ReadFile("./duck/index.html")
		fResp := strings.Replace(string(fBytes), "%WISDOM%", generateWisdom(), 1)
		fResp = strings.ReplaceAll(fResp, "%BASEPATH%", basePath)
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(fResp)
	})
	group.Get("/privacy", func(c *fiber.Ctx) error {
		fBytes, _ := os.ReadFile("./duck/privacy.html")
		fResp := strings.ReplaceAll(string(fBytes), "%BASEPATH%", basePath)
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(fResp)
	})
	group.Get("/stats", func(c *fiber.Ctx) error {
		fBytes, _ := os.ReadFile("./duck/stats.html")
		jsonBytes, _ := json.MarshalIndent(stats, "", "  ")
		fResp := strings.Replace(string(fBytes), "%STATS%", string(jsonBytes), 1)
		fResp = strings.ReplaceAll(fResp, "%BASEPATH%", basePath)
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(fResp)
	})
	group.Get("/wisdom", func(c *fiber.Ctx) error {
		stats.Visits++
		if c.Query("re") != "" {
			ref := stats.Referrals[c.Query("re")]
			ref.From = c.Query("re")
			ref.Visits++
			stats.Referrals[c.Query("re")] = ref
		}
		fBytes, _ := os.ReadFile("./duck/wisdom.html")
		fResp := strings.Replace(string(fBytes), "%WISDOM%", generateWisdom(), 1)
		fResp = strings.ReplaceAll(fResp, "%BASEPATH%", basePath)
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(fResp)
	})
	// page endpoints - static files
	group.Static("/assets", "./duck/assets", fiber.Static{})

	// start the server on port 5500
	log.Fatal(app.Listen(":5500"))
}
