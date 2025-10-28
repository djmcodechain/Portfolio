# ============================================
# ⚙️ DJMCodeChain Makefile
# Stack: Go + HTMX + Templ + CSS + Fly.io
# ============================================
# Author: Daniel J. Manning
# Website: https://djmcodechain.dev
# ============================================

APP_NAME := djmcodechain

# ─────────────────────────────────────────────
# 🧱 Paths
# ─────────────────────────────────────────────
GO_SRC := ./cmd
GO_BIN := ./bin/app
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./frontend/*")

CSS_DIR := frontend/styles/partials
CSS_OUTPUT := frontend/styles/style.css
CSS_MIN := frontend/styles/min.style.css

JS_DIR := frontend/scripts/partials
JS_OUTPUT := frontend/scripts/script.js
JS_MIN := frontend/scripts/min.script.js

HTML_DIR := frontend/public
HTML_OUTPUT := frontend/public/index.html
HTML_MIN := frontend/public/min.index.html

TEMPL_DIR := ./frontend/templates
BUILD_DIR := ./dist

.PHONY: all build run dev dev-all templ scss tidy clean \
        build-css minify-css build-js minify-js minify-html \
        build-go build-all minify-all fly-deploy fly-status fly-logs

# ============================================
# 🧩 Core Build Commands
# ============================================
all: tidy build-all minify-all

build:
	@echo "🔨 Building Go app..."
	go mod tidy
	go build -o app .

run:
	@echo "🚀 Starting app..."
	./app

tidy:
	@echo "🧾 Tidying Go modules..."
	go mod tidy

# ============================================
# 🎨 Frontend Tasks (CSS, JS, HTML)
# ============================================

# 🧱 CSS
build-css:
	@echo "🧩 Building readable main.css..."
	@cat $(CSS_DIR)/header.css \
	    $(filter-out $(CSS_DIR)/header.css $(CSS_DIR)/footer.css, $(wildcard $(CSS_DIR)/*.css)) \
	    $(CSS_DIR)/footer.css \
	    > $(CSS_OUTPUT)
	@echo "✅ main.css built in correct order!"

minify-css: build-css
	@echo "⚙️  Minifying CSS → $(CSS_MIN)..."
	@sed '/\/\*/{:a; /\*\//!{N;ba}; s/\/\*.*\*\///g}' $(CSS_OUTPUT) \
	| tr -d '\n\t' \
	| sed 's/  */ /g; s/ *{/{/g; s/ *}/}/g; s/ *:/:/g; s/ *,/,/g; s/ *; */;/g' \
	> $(CSS_MIN)
	@echo "✅ min.main.css ready!"

# 🧠 JS
build-js:
	@echo "🧩 Building readable main.js..."
	@cat $(JS_DIR)/*.js > $(JS_OUTPUT)
	@echo "✅ main.js built!"

minify-js: build-js
	@echo "⚙️  Minifying JS → $(JS_MIN)..."
	@sed '/\/\//d;/\/\*/,/\*\//d' $(JS_OUTPUT) \
	| tr -d '\n\t' \
	| sed 's/  */ /g' \
	> $(JS_MIN)
	@echo "✅ min.main.js ready!"

# 💎 HTML
minify-html:
	@echo "⚙️  Minifying HTML → $(HTML_MIN)..."
	@sed '/<!--/,/-->/d' $(HTML_OUTPUT) \
	| tr -d '\n\t' \
	| sed 's/  */ /g' \
	> $(HTML_MIN)
	@echo "✅ min.index.html ready!"

# ============================================
# 🔨 Go Build + Templ
# ============================================
templ:
	@echo "📦 Watching Templ templates..."
	templ generate --watch

build-go:
	@echo "⚙️ Generating templ files..."
	@templ generate
	@echo "🚀 Building Go binaries..."
	@go mod tidy
	@go build -o $(GO_BIN) $(GO_SRC)
	@echo "✅ Go build complete!"

# ============================================
# 💻 Development Environment
# ============================================
dev:
	@echo "🧠 Starting live reload with Air..."
	air

dev-all:
	@echo "💻 Launching full GOTTH dev environment..."
	$(MAKE) -j3 scss templ dev

# ============================================
# 🧹 Clean + Bundle Tasks
# ============================================
build-all: build-go build-css build-js
	@echo "🏗️  All assets built!"

minify-all: minify-css minify-js minify-html
	@echo "✨ All files minified!"

clean:
	@echo "🧼 Cleaning up build outputs..."
	@rm -f app $(CSS_OUTPUT) $(CSS_MIN) $(JS_OUTPUT) $(JS_MIN) $(HTML_MIN) $(GO_BIN)
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete!"

# ============================================
# ☁️ Fly.io Deployment
# ============================================
fly-deploy:
	@echo "🚀 Deploying $(APP_NAME) to Fly.io..."
	flyctl deploy --no-cache

fly-status:
	@echo "📡 Checking Fly.io app status..."
	flyctl status

fly-logs:
	@echo "📜 Tailing Fly.io logs..."
	flyctl logs
