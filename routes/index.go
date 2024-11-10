package routes

import (
	"an-2024/models"
	"an-2024/views/index"

	"github.com/labstack/echo/v4"
)

func IndexView(ctx echo.Context) error {
	var sideProjects []models.SideProject
	htmlHighlighter := models.SideProject{
		Title:               "HTML Highlighter for template strings in JS/TS",
		Body:                "A VSCode plugin for syntax highlighting HTML in JS/TS template strings. Based on my other project \"SQL Syntax Higlighter for Go\".",
		Icon:                "/assets/images/side-projects/brackets.png",
		IsAppStoreIconStyle: true,
		Link:                "https://marketplace.visualstudio.com/items?itemName=antonnyman.vscode-js-html-highlight",
		LinkText:            "Open in VSCode Marketplace",
	}
	bevattningsforbud := models.SideProject{
		Title:               "Bevattningsförbud",
		Body:                "An iOS app for current irrigation bans in Sweden per municipality. Written in Swift, backend using Cloudflare Workers.",
		Icon:                "/assets/images/side-projects/forbud.png",
		IsAppStoreIconStyle: true,
		Link:                "https://apps.apple.com/se/app/bevattningsf%C3%B6rbud/id6450721436",
		LinkText:            "Bevattningsforbud link to App Store",
		IsAppStoreLink:      true,
	}
	sqlHighlighter := models.SideProject{
		Title:    "SQL Syntax Highlighter for Go",
		Body:     "A VSCode plugin for syntax highlighting raw SQL string literals in Go. Similar syntax to what the Apollo GraphQL plugin uses.",
		Comment:  "5/5 ★, 2100+ active installs",
		Icon:     "/assets/images/side-projects/gopher.png",
		Link:     "https://marketplace.visualstudio.com/items?itemName=antonnyman.go-sql-highlight",
		LinkText: "Open in VSCode Marketplace",
	}
	quickTunnel := models.SideProject{
		Title:    "Quick Tunnel",
		Body:     "A macOS app to forward your localhost server to the world using Cloudflare Tunnels. Written in Swift.",
		Comment:  "For now, this app is unsigned, therefore you need to unzip, drag to applications, right click and open.",
		Icon:     "/assets/images/side-projects/quick-tunnel.png",
		Link:     "/downloads/QuickTunnel-1.0.zip",
		LinkText: "Download for macOS, Apple Silicon (v1.0)",
	}
	compris := models.SideProject{
		Title:               "Compris",
		Body:                "An iOS app to compress videos. Written in Swift.",
		Icon:                "/assets/images/side-projects/compris.png",
		IsAppStoreIconStyle: true,
		Link:                "https://apps.apple.com/se/app/compris/id1632895753",
		LinkText:            "Bevattningsforbud link to App Store",
		IsAppStoreLink:      true,
	}
	embroidery := models.SideProject{
		Title:               "Embroidery",
		Body:                "A toy JavaScript framework for modifying the DOM.",
		Icon:                "/assets/images/side-projects/code-slash.png",
		IsAppStoreIconStyle: true,
		Link:                "/embroidery",
		LinkText:            "Read more",
	}

	sideProjects = append(sideProjects, htmlHighlighter, bevattningsforbud, sqlHighlighter, quickTunnel, compris, embroidery)
	return html(ctx, index.IndexView(sideProjects))
}
