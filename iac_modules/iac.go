package iac_modules

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "strings"
)


// Current working directory
var workDir string

func init() {
    cwd, err := os.Getwd()
    if err != nil {
        panic("failed to get current directory")
    }
    workDir = cwd
    //fmt.Println("Work directory:", workDir)
}



// StackCheck returns the current stack name
func StackCheck() (string, error) {
    cmd := exec.Command("pulumi", "stack", "ls", "--json")
    output, err := cmd.Output()
    if err != nil {
        return "", errors.New("stack not found")
    }

    var stacks []map[string]interface{}
    err = json.Unmarshal(output, &stacks)
    if err != nil {
        return "", err
    }

    for _, stack := range stacks {
        if stack["current"] == true {
            name, ok := stack["name"].(string)
            if !ok {
                return "", errors.New("invalid stack name")
            }
            parts := strings.Split(name, "/")
            return parts[len(parts)-1], nil
        }
    }
    return "", errors.New("current stack not found")
}

// Dump extracts specific properties from stack configuration
func Dump(prop string, stackName string) (map[string]interface{}, error) {
    stackPath := fmt.Sprintf("%s/stacks/%s", workDir, stackName)
    DebugLog("Attempting to read directory: %s", stackPath)
    
    // Check if directory exists
    if _, err := os.Stat(stackPath); os.IsNotExist(err) {
        DebugLog("Warning: stacks/%s directory does not exist", stackName)
        return nil, fmt.Errorf("stack directory does not exist: %s", stackPath)
    }
    
    final := make(map[string]interface{})
    configs, err := LoadYaml(stackPath)
    if err != nil {
        DebugLog("Failed to read YAML: %v", err)
        return nil, err
    }
    
    DebugLog("Successfully read %d YAML files from %s", len(configs), stackPath)

    for _, config := range configs {
        if data, ok := config[prop]; ok {
            if subMap, ok := data.(map[string]interface{}); ok {
                for k, v := range subMap {
                    final[k] = v
                }
            }
        }
    }

    // Merge global variables for variables section
    if prop == "variables" {
        DebugLog("Merging global variables...")
        globalVars, err := GlobalVariables()
        if err != nil {
            DebugLog("Failed to read global variables: %v", err)
            return nil, err
        }
        DebugLog("Found %d global variables", len(globalVars))
        for k, v := range globalVars {
            final[k] = v
        }
    }

    return map[string]interface{}{prop: final}, nil
}

// GlobalVariables reads and returns global variables
func GlobalVariables() (map[string]interface{}, error) {
    globalVarsPath := fmt.Sprintf("%s/stack_share_variables", workDir)
    DebugLog("Attempting to read global variables: %s", globalVarsPath)
    
    // Check if file exists
    if _, err := os.Stat(globalVarsPath); os.IsNotExist(err) {
        DebugLog("Warning: global variables file does not exist: %s", globalVarsPath)
        return make(map[string]interface{}), nil // Return empty map instead of error
    }
    
    result := make(map[string]interface{})
    configs, err := LoadYaml(globalVarsPath)
    if err != nil {
        DebugLog("Failed to read global variables: %v", err)
        return nil, err
    }
    
    DebugLog("Successfully read %d YAML files from %s", len(configs), globalVarsPath)
    
    for _, config := range configs {
        if data, ok := config["variables"]; ok {
            if subMap, ok := data.(map[string]interface{}); ok {
                for k, v := range subMap {
                    result[k] = v
                }
            }
        }
    }
    return result, nil
}

// Join combines stack configuration sections into Pulumi.yaml
func Join(stackName string) error {
    DebugLog("Starting Join operation, merging stack: %s", stackName)
    sections := []string{"resources", "variables", "outputs"}
    merged := make(map[string]interface{})

    for _, section := range sections {
        DebugLog("Processing section: %s", section)
        data, err := Dump(section, stackName)
        if err != nil {
            DebugLog("Failed to read %s: %v", section, err)
            continue
        }
        DebugLog("Successfully read %s section", section)
        for k, v := range data {
            merged[k] = v
        }
    }

    DebugLog("Reading Pulumi.yaml base structure")
    projectData, err := SectionsRemove()
    if err != nil {
        DebugLog("Failed to read Pulumi.yaml: %v", err)
        return err
    }
    for k, v := range merged {
        projectData[k] = v
    }

    pulumiYamlPath := fmt.Sprintf("%s/Pulumi.yaml", workDir)
    DebugLog("Writing merged result to: %s", pulumiYamlPath)
    return WriteYamlFile(pulumiYamlPath, projectData)
}

// Recovery restores Pulumi.yaml to its original state
func Recovery() error {
    DebugLog("Executing Recovery operation")
    data, err := SectionsRemove()
    if err != nil {
        DebugLog("Failed to read Pulumi.yaml: %v", err)
        return err
    }
    pulumiYamlPath := fmt.Sprintf("%s/Pulumi.yaml", workDir)
    DebugLog("Restoring Pulumi.yaml to: %s", pulumiYamlPath)
    return WriteYamlFile(pulumiYamlPath, data)
}

// SectionsRemove removes specific sections from Pulumi.yaml
func SectionsRemove() (map[string]interface{}, error) {
    path := fmt.Sprintf("%s/Pulumi.yaml", workDir)
    DebugLog("Removing specific sections from %s", path)
    data, err := ReadYamlFile(path)  // 使用大寫開頭的新函數
    if err != nil {
        DebugLog("Failed to read Pulumi.yaml: %v", err)
        return nil, err
    }
    delete(data, "variables")
    delete(data, "resources")
    delete(data, "outputs")
    return data, nil
}



