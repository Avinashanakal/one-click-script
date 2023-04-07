package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TemplateRepo struct {
	Owner              string `json:"owner"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	IncludeAllBranches bool   `json:"include_all_branches"`
	Private            bool   `json:"private"`
}

func main() {
	token := "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"

	templateOwner := "Avinashanakal"
	templateRepo := "LoggerLLD"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the owner name of the new repository: ")
	newOwner, _ := reader.ReadString('\n')
	newOwner = newOwner[:len(newOwner)-1]

	fmt.Print("Enter the description of the new repository: ")
	newDescription, _ := reader.ReadString('\n')
	newDescription = newDescription[:len(newDescription)-1]

	fmt.Print("Enter the new name of the repository: ")
	newName, _ := reader.ReadString('\n')
	newName = newName[:len(newName)-1]

	includeAllBranches := false
	private := true

	repo := TemplateRepo{
		Owner:              newOwner,
		Name:               newName,
		Description:        newDescription,
		IncludeAllBranches: includeAllBranches,
		Private:            private,
	}
	payload, err := json.Marshal(repo)
	if err != nil {
		fmt.Println("Failed to encode payload:", err)
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/generate", templateOwner, templateRepo)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Failed to create repository: %s\n", resp.Status)
		return
	}

	fmt.Println("Repository created successfully!")
}
