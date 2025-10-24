# ============================================
# DJMCodeChain Makefile ⚙️
# Stack: Go + HTMX + Templ + SCSS + Fly.io
# ============================================

APP_NAME := djmcodechain
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./frontend/*")
SCSS_INPUT := ./frontend/scss/main.scss
CSS_OUTPUT := ./frontend/css/output.css
TEMPL_DIR := ./frontend/templates
BUILD_DIR := ./dist

.PHONY: all build run dev templ scss tidy clean fly-deploy fly-status fly-logs

# ============================================
# 🧱 Build Commands
# ============================================

# Build everything for production
all: tidy scss templ build

build:
	@echo "🔨 Building Go app..."
	go mod tidy
	go build -o app .

run:
	@echo "🚀 Starting app..."
	./app

# ============================================
# 🧩 Development Commands
# ============================================

# Watch Templ templates
templ:
	@echo "📦 Watching Templ templates..."
	templ generate --watch

# Watch SCSS -> CSS
scss:
	@echo "🎨 Watching SCSS for changes..."
	sass --watch $(SCSS_INPUT):$(CSS_OUTPUT)

# Start Air (live reload for Go)
dev:
	@echo "🧠 Starting live reload with Air..."
	air

# Run everything together (SCSS + Templ + Go)
dev-all:
	@echo "💻 Launching full GOTTH dev environment..."
	$(MAKE) -j3 scss templ dev

# ============================================
# 🧹 Utility Commands
# ============================================

tidy:
	@echo "🧾 Tidying modules..."
	go mod tidy

clean:
	@echo "🧼 Cleaning up build artifacts..."
	rm -f app
	rm -rf $(BUILD_DIR)

# ============================================
# ☁️ Deployment Commands
# ============================================

fly-deploy:
	@echo "🚀 Deploying to Fly.io..."
	flyctl deploy --no-cache

fly-status:
	@echo "📡 Checking Fly.io app status..."
	flyctl status

fly-logs:
	@echo "📜 Tailing Fly.io logs..."
	flyctl logs
