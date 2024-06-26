package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Bitte geben Sie das Root-Verzeichnis ein: ")
    root, err := reader.ReadString('\n')
    if err != nil {
        fmt.Printf("Fehler beim Einlesen des Root-Verzeichnisses: %v\n", err)
        return
    }
    root = strings.TrimSpace(root)

    err = filepath.Walk(root, visit)
    if err != nil {
        fmt.Printf("Fehler beim Durchlaufen der Verzeichnisse: %v\n", err)
    }
}

func visit(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    if info.IsDir() {
        gitIgnorePath := filepath.Join(path, ".gitignore")
        if _, err := os.Stat(gitIgnorePath); err == nil {
            fmt.Printf("Git-ähnliches Verzeichnis gefunden: %s\n", path)
            pullLocalBranches(path)
            return filepath.SkipDir
        }
    }
    return nil
}

func pullLocalBranches(repoPath string) {
    cmd := exec.Command("git", "branch")
    cmd.Dir = repoPath
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Fehler beim Abrufen der lokalen Branches in %s: %v\n", repoPath, err)
        return
    }

    branches := strings.Split(string(output), "\n")
    for _, branch := range branches {
        branch = strings.TrimSpace(branch)
        if branch != "" {
            branch = strings.TrimPrefix(branch, "* ")
            if remoteBranchExists(repoPath, branch) {
                checkoutCmd := exec.Command("git", "checkout", branch)
                checkoutCmd.Dir = repoPath
                if err := checkoutCmd.Run(); err != nil {
                    fmt.Printf("Fehler beim Checkout des Branches %s in %s: %v\n", branch, repoPath, err)
                    continue
                }

                pullCmd := exec.Command("git", "pull")
                pullCmd.Dir = repoPath
                if err := pullCmd.Run(); err != nil {
                    fmt.Printf("Fehler beim Pullen des Branches %s in %s: %v\n", branch, repoPath, err)
                    continue
                }

                fmt.Printf("Branch %s in %s erfolgreich gepullt.\n", branch, repoPath)
            } else {
                fmt.Printf("Remote-Branch %s existiert nicht, Überspringen des Pulls in %s.\n", branch, repoPath)
            }
        }
    }
}

func remoteBranchExists(repoPath, branch string) bool {
    cmd := exec.Command("git", "ls-remote", "--heads", "origin", branch)
    cmd.Dir = repoPath
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Fehler beim Überprüfen des Remote-Branches %s in %s: %v\n", branch, repoPath, err)
        return false
    }
    return strings.TrimSpace(string(output)) != ""
}
