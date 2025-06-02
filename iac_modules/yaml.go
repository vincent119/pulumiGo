package iac_modules

import (
    "os"
    "path/filepath"
    "strings"
    "io/fs"
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

// LoadYamlFileAll loads and parses all YAML documents from a file
func LoadYamlFileAll(path string) ([]map[string]interface{}, error) {
		DebugLog("Attempting to load YAML from: %s", path) 
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var results []map[string]interface{}
    decoder := yaml.NewDecoder(strings.NewReader(string(content)))
    for {
        var doc map[string]interface{}
        err := decoder.Decode(&doc)
        if err != nil {
            break
        }
        results = append(results, doc)
    }
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
            DebugLog("Failed to read YAML file: %v", err)
            continue
        }
        result = append(result, data...)
    }
    
    DebugLog("Total YAML documents read: %d", len(result))
    return result, nil
}