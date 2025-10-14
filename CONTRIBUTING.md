# 🧭 Contributing Guide (a.k.a. Notes to Future Me)

Hey, me 👋 — welcome back to the codebase.  
This document keeps me (and maybe future collaborators) consistent when making changes, opening issues, or submitting pull requests.  
Think of it as my personal dev manual for this project.

---

## 🧩 Workflow Overview

**Main Branch:**  
- `main` (production-ready only — no experiments here)  

**Feature Branches:**  
- Always branch off `main`  
- Naming convention: `camelCase`
- feature/<short-description>
- bugfix/<short-description>
- refactor/<short-description>

**Example:** 
- feature/create-handlers
- bugfix/template-rendering


---

## 🧠 Issues

All issues use one of the pre-defined templates in `.github/ISSUE_TEMPLATE`:

- 🪄 **Feature Request** → for new functionality or enhancements  
- ⚙️ **Implementation Task** → for structured development work (like handlers, models, etc.)  
- 🐞 **Bug Report** → for broken things  

Each issue includes:
- Summary (what + why)
- Details/Objectives (how)
- Acceptance Criteria (definition of done)
- Optional notes or next steps  

This keeps the project self-documented and my future self grateful.

---

## 🔀 Pull Requests

All PRs should use the [pull_request_template.md](.github/pull_request_template.md).

### PR naming convention:
- [feature] Create Handlers
- [bugfix] Fix Template Crash
- [refactor] Clean Up Routing


### PR rules:
- Each PR should close or reference its issue (e.g., `Closes #10`)
- Keep PRs small and focused — one logical change per PR  
- Include testing steps (manual or automated)
- Make sure it builds and runs before pushing 🙃  

---

## 🧹 Code Style

- Use **Go fmt** and **go vet** before committing  
- Keep functions small, purposeful, and readable  
- Comments should explain *why*, not *what*  
- Directory naming: lowercase, no underscores (Go idiomatic)
- Avoid long god structs — compose smaller models

---

## 🧪 Testing & Running
```bash
go run ./cmd/server
```
go test ./...


Keep tests simple and isolated — no fancy frameworks unless necessary.

---

## 💡 Notes to Future Me

- Use issues even for solo work — they make progress visible and organized.  
- Keep templates up to date as structure evolves.  
- Refactor early, document always.  
- If something feels hacky, it probably is — fix it next commit.  
- Future me, you’re awesome. Don’t forget coffee ☕️

---

**Last Updated:** 14 October, 2025
*(If you’re reading this in the future, please maintain the same tone — professional but chill.)*



