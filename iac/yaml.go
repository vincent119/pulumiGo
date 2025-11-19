package iac

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// RecursiveDirectoryWalk recursively walks through a directory and returns all files
func RecursiveDirectoryWalk(root string) ([]string, error) {
    var files []string
    err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}

func LoadYamlFileAll(path string) ([]map[string]interface{}, error) {
    DebugLog("Attempting to load YAML from: %s", path)
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    // 增加原始內容的日誌記錄 (限制長度避免垃圾信息)
    contentStr := string(content)
    if len(contentStr) > 500 {
        DebugLog("YAML content (first 500 chars): %s...", contentStr[:500])
    } else {
        DebugLog("YAML content: %s", contentStr)
    }

    var results []map[string]interface{}
    decoder := yaml.NewDecoder(strings.NewReader(string(content)))
    docIndex := 0

    for {
        var doc map[string]interface{}
        err := decoder.Decode(&doc)
        if err != nil {
            // 區分 EOF 和其他錯誤
            if err == io.EOF {
                break // 正常結束
            }

            // 記錄解析錯誤但繼續
            DebugLog("Error decoding document %d: %v", docIndex, err)
            return nil, fmt.Errorf("error parsing YAML document %d in %s: %w", docIndex, path, err)
        }

        DebugLog("Successfully decoded document %d", docIndex)
        results = append(results, doc)
        docIndex++
    }

    DebugLog("Successfully loaded %d documents from %s", len(results), path)
    return results, nil
}

// ReadYamlFile reads and parses a YAML file
func ReadYamlFile(path string) (map[string]interface{}, error) {
    DebugLog("Reading YAML file: %s", path)
    data := make(map[string]interface{})
    file, err := os.ReadFile(path)
    if err != nil {
        DebugLog("Failed to read file: %v", err)
        return nil, err
    }
    err = yaml.Unmarshal(file, &data)
    if err != nil {
        DebugLog("Failed to parse YAML: %v", err)
    }
    return data, err
}

// WriteYamlFile writes a map to a YAML file
func WriteYamlFile(path string, data map[string]interface{}) error {
    DebugLog("Writing YAML file: %s", path)
    file, err := os.Create(path)
    if err != nil {
        DebugLog("Failed to create file: %v", err)
        return err
    }
    defer file.Close()

    encoder := yaml.NewEncoder(file)
    defer encoder.Close()
    err = encoder.Encode(data)
    if err != nil {
        DebugLog("Failed to encode YAML: %v", err)
    }
    return err
}

// LoadYaml loads and parses YAML files from a path
func LoadYaml(path string) ([]map[string]interface{}, error) {
    DebugLog("Attempting to read YAML: %s", path)

    // Check if path is directory or file
    fileInfo, err := os.Stat(path)
    if err != nil {
        DebugLog("Failed to read path: %v", err)
        return nil, err
    }

    var files []string

    if fileInfo.IsDir() {
        // Use recursive method to read directory
        files, err = RecursiveDirectoryWalk(path)
        if err != nil {
            DebugLog("Failed to recursively read directory: %v", err)
            return nil, err
        }

        // Filter non-YAML files
        var yamlFiles []string
        for _, file := range files {
            if strings.HasSuffix(strings.ToLower(file), ".yaml") ||
               strings.HasSuffix(strings.ToLower(file), ".yml") {
                yamlFiles = append(yamlFiles, file)
            }
        }
        files = yamlFiles

        DebugLog("Found %d YAML files in directory and subdirectories", len(files))
    } else {
        // Single file
        files = append(files, path)
    }

    var result []map[string]interface{}

    for _, filePath := range files {
        DebugLog("Reading file: %s", filePath)
        data, err := LoadYamlFileAll(filePath)
        if err != nil {
            DebugLog("Failed to read YAML file %s: %v", filePath, err)
            content, readErr := os.ReadFile(filePath)
            if readErr == nil && len(content) > 0 {
                DebugLog("Problematic YAML content (truncated):\n%s", truncateString(string(content), 1000))
                errorLocation := LocateYAMLError(string(content), err)
                DebugLog("Error location: %s", errorLocation)
            }
            continue
        }

        // 檢查是否成功加載資源
        for i, doc := range data {
            DebugLog("Document %d contains keys: %v", i, getMapKeys(doc))

            // 呼叫獨立函數處理 WAF 規則
            processWAFRegexPatterns(doc, filePath)
        }

        result = append(result, data...)
    }

    DebugLog("Total YAML documents read: %d", len(result))
    return result, nil
}

// 處理 WAF 正則表達式模式的獨立函數
func processWAFRegexPatterns(doc map[string]interface{}, filePath string) {
    // 取得資源映射
    resources, ok := doc["resources"].(map[string]interface{})
    if !ok {
        DebugLog("No resources defined in document from %s", filePath)
        return
    }

    DebugLog("Resources defined in %s: %v", filePath, getMapKeys(resources))

    // 尋找並處理每個 WAF 正則表達式資源
    for resourceName, resourceDef := range resources {
        resourceMap, ok := resourceDef.(map[string]interface{})
        if !ok {
            continue
        }

        resourceType, ok := resourceMap["type"].(string)
        if !ok || resourceType != "aws:wafv2:RegexPatternSet" {
            continue
        }

        // 處理找到的 RegexPatternSet
        processRegexPatternSet(resourceName, resourceMap, resources, doc)
    }
}

// 處理單一 RegexPatternSet 資源
func processRegexPatternSet(resourceName string, resourceMap map[string]interface{}, resources map[string]interface{}, doc map[string]interface{}) {
    props, ok := resourceMap["properties"].(map[string]interface{})
    if !ok {
        return
    }

    regexes, ok := props["regularExpressions"].([]interface{})
    if !ok {
        return
    }

    DebugLog("Found RegexPatternSet: %s with %d regexes", resourceName, len(regexes))

    // 標準化和修正 WAF 正則表達式
    modified := standardizeRegexPatterns(regexes)

    // 如果有修改，更新文檔
    if modified {
        DebugLog("Updated regexes for %s", resourceName)
        props["regularExpressions"] = regexes
        resourceMap["properties"] = props
        resources[resourceName] = resourceMap
        doc["resources"] = resources
    }
}



func getMapKeys(m map[string]interface{}) []string {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}




// SanitizeRegexForYAML 處理正則表達式中的特殊字符，確保它們在 YAML 中正確表示
func SanitizeRegexForYAML(regex string) string {
    // 檢查是否已經有足夠的轉義
    if isProperlyEscaped(regex) {
        return regex
    }

    // 先將可能已存在的轉義還原，以避免重複轉義
    normalizedRegex := RestoreRegexFromYAML(regex)

    // 步驟 1: 轉義反斜線 (\ → \\)
    result := ""
    for i := 0; i < len(normalizedRegex); i++ {
        if normalizedRegex[i] == '\\' {
            result += "\\\\"
        } else if normalizedRegex[i] == '/' {
            // 步驟 2: 轉義斜線 (/ → \/)
            result += "\\/"
        } else {
            result += string(normalizedRegex[i])
        }
    }

    return result
}

// isProperlyEscaped 檢查正則表達式是否已經正確轉義
func isProperlyEscaped(regex string) bool {
    // 檢查是否有未轉義的斜線
    for i := 0; i < len(regex); i++ {
        if regex[i] == '/' {
            // 確認前方沒有單個反斜線
            if i == 0 || regex[i-1] != '\\' {
                return false
            }
        }
    }

    return true
}


// RestoreRegexFromYAML 將 YAML 中的正則表達式恢復為有效格式
func RestoreRegexFromYAML(yamlRegex string) string {
    // 移除 YAML 和正則表達式雙重轉義的影響
    result := strings.ReplaceAll(yamlRegex, "\\/", "/")

    // 步驟 2: 將剩餘的雙反斜線恢復為單反斜線 (\\ → \)
    result = strings.ReplaceAll(result, "\\\\", "\\")

    return result
}



// LocateYAMLError 嘗試定位 YAML 錯誤的具體位置和行號
func LocateYAMLError(content string, err error) string {
    if err == nil {
        return "No error"
    }

    // 查看錯誤是否包含行號信息
    // 這是基於 yaml.v3 錯誤格式的一個簡單示例
    errMsg := err.Error()

    // 尋找可能的行號
    lineMatches := regexp.MustCompile(`line (\d+)`).FindStringSubmatch(errMsg)

    if len(lineMatches) > 1 {
        lineNumber, parseErr := strconv.Atoi(lineMatches[1])
        if parseErr == nil {
            // 找到了行號，嘗試顯示問題行及其上下文
            lines := strings.Split(content, "\n")
            if lineNumber > 0 && lineNumber <= len(lines) {
                context := ""

                // 顯示問題行的前後幾行
                start := max(0, lineNumber-3)
                end := min(len(lines), lineNumber+2)

                for i := start; i < end; i++ {
                    prefix := "  "
                    if i+1 == lineNumber {
                        prefix = "> " // 標記問題行
                    }
                    context += fmt.Sprintf("%s%4d: %s\n", prefix, i+1, lines[i])
                }

                return fmt.Sprintf("Error at line %d:\n%s\nError: %s", lineNumber, context, errMsg)
            }
        }
    }

    return fmt.Sprintf("YAML Error: %v", err)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


// 輔助函數截斷字符串
func truncateString(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen] + "..."
}


// StandardizeWAFRegex 將 WAF 正則表達式保持為 Pulumi 期望的格式
func StandardizeWAFRegex(regex string) string {
    // 移除所有引號
    regex = strings.Trim(regex, "'\"")

    // 檢查是否已經是一個簡單的格式
    // 如果是簡單格式，例如 Pulumi 中看到的格式，直接返回
    if !needsStandardization(regex) {
        return regex
    }

    // 只有在真正需要轉換時才進行標準化處理
    normalizedRegex := RestoreRegexFromYAML(regex)
    hasStartAnchor := strings.HasPrefix(normalizedRegex, "^")

    // 基於當前格式偏好進行轉換
    // 檢查截圖中的格式，它似乎偏好簡單格式，不使用額外的斜線轉義
    simplifiedRegex := simplifyRegex(normalizedRegex)

    // 恢復開頭錨點
    if hasStartAnchor && !strings.HasPrefix(simplifiedRegex, "^") {
        simplifiedRegex = "^" + simplifiedRegex
    }

    return simplifiedRegex
}

// needsStandardization 檢查正則表達式是否需要標準化
func needsStandardization(regex string) bool {
    // 檢查常見的格式問題

    // 1. 有尾部單引號的情況
    if strings.HasSuffix(regex, "'") {
        return true
    }

    // 2. 有過度轉義的斜線
    if strings.Contains(regex, "\\/") &&
       (strings.HasPrefix(regex, "^/") || strings.HasPrefix(regex, "/")) {
        return true
    }

    // 3. 檢查轉義的一致性
    // 如果既有 / 又有 \/ 混合使用，需要標準化
    slashCount := strings.Count(regex, "/")
    escapedSlashCount := strings.Count(regex, "\\/")

    if slashCount > 0 && escapedSlashCount > 0 &&
       slashCount != escapedSlashCount {
        return true
    }

    return false
}

// simplifyRegex 簡化正則表達式，使其符合 Pulumi 偏好的格式
// 基於截圖中觀察到的格式
func simplifyRegex(regex string) string {
    // 截圖顯示 Pulumi 偏好簡單的格式:
    // 1. 路徑分隔符使用普通的斜線 /，而非轉義的 \/
    result := strings.ReplaceAll(regex, "\\/", "/")

    // 2. 但對於特殊字符，仍然需要保留正確的轉義
    // 例如 "\." 應該保持為 "\."

    // 先檢查結果是否是有效的正則表達式
    if _, err := regexp.Compile(result); err != nil {
        // 如果轉換後無效，嘗試修復
        result = regex // 回退到原始版本
        DebugLog("簡化後的正則表達式無效，保持原樣: %s", regex)
    }

    return result
}


// DetectRegexFormatStyle 檢測正則表達式集合的格式風格
func DetectRegexFormatStyle(regexes []interface{}) (bool, bool) {
    // 分析正則表達式集合，確定主流格式

    // 追蹤格式屬性
    useStartAnchor := 0  // 使用 ^ 前綴的數量
    useEscapedSlash := 0 // 使用 \/ 轉義斜線的數量

    for _, regex := range regexes {
        regexMap, ok := regex.(map[string]interface{})
        if !ok {
            continue
        }

        pattern, ok := regexMap["regexString"].(string)
        if !ok {
            continue
        }

        // 移除引號
        pattern = strings.Trim(pattern, "'\"")

        // 檢查是否有前綴 ^
        if strings.HasPrefix(pattern, "^") {
            useStartAnchor++
        }

        // 檢查是否使用轉義斜線 \/
        if strings.Contains(pattern, "\\/") {
            useEscapedSlash++
        } else if strings.Contains(pattern, "/") &&
                !strings.Contains(pattern, "\\/") {
            // 明確使用非轉義斜線
            useEscapedSlash--
        }
    }

    // 根據多數規則決定格式風格
    preferStartAnchor := useStartAnchor > len(regexes)/2
    preferEscapedSlash := useEscapedSlash > 0

    return preferStartAnchor, preferEscapedSlash
}

// 修改 standardizeRegexPatterns 函數，使用檢測到的格式風格
func standardizeRegexPatterns(regexes []interface{}) bool {
    if len(regexes) == 0 {
        return false
    }

    // 檢測集合的格式風格
    preferStartAnchor, preferEscapedSlash := DetectRegexFormatStyle(regexes)
    DebugLog("Detected pattern style: preferStartAnchor=%v, preferEscapedSlash=%v",
             preferStartAnchor, preferEscapedSlash)

    modified := false

    for j, regex := range regexes {
        regexMap, ok := regex.(map[string]interface{})
        if !ok {
            continue
        }

        pattern, ok := regexMap["regexString"].(string)
        if !ok {
            continue
        }

        DebugLog("  Original Regex %d: %s", j, pattern)

        // 清除格式問題
        cleanPattern := strings.Trim(pattern, "'\"")

        // 檢查是否需要修復明顯錯誤
        if needsStandardization(cleanPattern) {
            // 取出核心正則表達式內容
            normalizedRegex := RestoreRegexFromYAML(cleanPattern)
            hasStartAnchor := strings.HasPrefix(normalizedRegex, "^")

            if hasStartAnchor && !preferStartAnchor {
                normalizedRegex = normalizedRegex[1:] // 移除 ^ 前綴
            } else if !hasStartAnchor && preferStartAnchor {
                normalizedRegex = "^" + normalizedRegex // 添加 ^ 前綴
            }

            // 根據偏好轉義或不轉義斜線
            if preferEscapedSlash {
                normalizedRegex = strings.ReplaceAll(normalizedRegex, "/", "\\/")
            } else {
                normalizedRegex = strings.ReplaceAll(normalizedRegex, "\\/", "/")
            }

            if normalizedRegex != cleanPattern {
                DebugLog("  Standardized Regex %d: %s -> %s", j, cleanPattern, normalizedRegex)
                regexMap["regexString"] = normalizedRegex
                regexes[j] = regexMap
                modified = true
            }
        }
    }

    return modified
}
