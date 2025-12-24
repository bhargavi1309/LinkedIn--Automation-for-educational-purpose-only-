package main

import (
	"linkedin-automation/auth"
	"linkedin-automation/config"
	"linkedin-automation/logger"
	"linkedin-automation/search"
	"linkedin-automation/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Init()

	cfg := config.Load()

	// Force Rod to use EXISTING Chrome (NO leakless, NO download)
	u := launcher.New().
		Bin(`C:\Program Files\Google\Chrome\Application\chrome.exe`).
		Leakless(false).
		Headless(cfg.Headless).
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage()
	stealth.MaskFingerprint(page)

	auth.Login(page, cfg)

	profiles := search.SearchPeople(page)
	for _, p := range profiles {
		logger.Log.Info("Found profile:", p)
	}
}
