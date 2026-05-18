# Changelog

All notable changes to this project will be documented in this file.

---

## [v0.1.4] - 2026-05-18

### ✨ Added

- **Scoop 支援（Windows）**：可透過 Scoop 安裝 pulumiGo，Windows 使用者不再需要手動下載
- **Homebrew Cask 支援（macOS）**：改用 Cask 發佈，安裝後自動移除 macOS 隔離屬性，不再需要手動執行 `xattr`

### 🔧 Changed

- Example 專案的 AWS Provider 命名從 `Uat` 統一改為 `Dev`，與 dev stack 環境一致

---

## [v0.1.3] - 2026-05-15

### 🔧 Changed

- Homebrew formula 名稱改為小寫 `pulumigo`，符合 Homebrew 命名規範

---

## [v0.1.2] - 2026-05-15

### Refactor / 重構

- **WAF regex logic decoupled from `LoadYaml`**
  Moved WAF regex normalization out of the generic YAML loader. The loader is now a pure I/O function with no business logic side effects.
  將 WAF 正規表達式標準化邏輯從通用 YAML 載入函式中分離，載入器現在是純 I/O 函式，不含業務邏輯副作用。

- **Unified flag-forwarding helpers across all handlers**
  `refresh`, `import`, `state` sub-commands, `config`, and `login` handlers now use the shared `forwardStringFlag` / `forwardStringArrayFlag` / `forwardBoolFlag` / `forwardInt32Flag` helpers, consistent with `up` and `preview`.
  所有 handler（`refresh`、`import`、`state` 子命令、`config`、`login`）統一改用共用的旗標轉發 helper，與 `up`、`preview` 風格一致。

- **Removed per-command flag handling from executor**
  `executor.go` no longer contains special-case flag logic for `config` or `login`. Each handler is now fully responsible for its own flag forwarding.
  `executor.go` 不再包含 `config`、`login` 的特殊旗標判斷，各 handler 完全負責自身的旗標轉發。

- **`workDir` lazy initialization replaces `init()` panic**
  `iac/iac.go` now uses `sync.Once` to initialize the working directory on first use, returning an error instead of panicking.
  `iac/iac.go` 改用 `sync.Once` 在首次使用時初始化工作目錄，失敗時回傳 error 而非 panic。

- **Replaced `log` with `zlogger` in executor**
  All log output in `executor.go` now goes through `zlogger` (via `iac.DebugLog`), consistent with the rest of the codebase.
  `executor.go` 的日誌輸出統一改用 `zlogger`（透過 `iac.DebugLog`），與其他模組一致。

- **Removed redundant `max`/`min` functions**
  Dropped hand-rolled `max`/`min` in `iac/yaml.go`; Go 1.21+ built-ins are used directly.
  移除 `iac/yaml.go` 中自訂的 `max`/`min` 函式，直接使用 Go 1.21+ 內建版本。

### Tests / 測試

- Added unit tests for `iac` package: `RestoreRegexFromYAML`, `SanitizeRegexForYAML`, `truncateString`, `LocateYAMLError`, `ReadYamlFile`/`WriteYamlFile`, and stack JSON parsing.
  新增 `iac` package 單元測試：regex 轉換、字串截斷、YAML 讀寫、stack JSON 解析。

- Added unit tests for `command/handlers`: all four `forwardXxxFlag` helpers.
  新增 `command/handlers` 單元測試：四種旗標轉發 helper。

- Added unit tests for `command`: stackless command detection logic.
  新增 `command` 單元測試：stackless 命令判斷邏輯。

- Updated `Makefile`: `test` target now uses `-count=1`; added `cover` target generating `coverage.out` and `coverage.html`.
  更新 `Makefile`：`test` 加入 `-count=1`；新增 `cover` 目標產生覆蓋率報告。

---

## [v0.1.1] - 2024-03-13

### Added / 新增

- Initial release with Pulumi CLI wrapper support.
  初始版本，支援 Pulumi CLI 包裝器。
- Commands: `up`, `preview`, `stack`, `config`, `login`, `logout`, `whoami`, `refresh`, `import`, `state`, `plugin`, `org`, `about`, `version`.
- Homebrew tap support.
  支援 Homebrew tap 安裝。
- Debug mode via `--debug` / `-d` flag.
  透過 `--debug` / `-d` 旗標啟用除錯模式。

---

## [v0.1.0] - 2024-03-13

### Added / 新增

- Project bootstrap.
  專案初始化。
