package handler

import (
	"errors"
	"log"
	"os/exec"
)

func (s *HTTPHandler) handleOnPush(repoName string) error {
	for _, repo := range s.repos.Repos {
		if repo.Name == repoName {
			cmd := exec.Command("git", "clone", repo.Url)
			cmd.Dir = repo.Path
			if err := cmd.Run(); err != nil {
				log.Println("[ERR] Error cloning repository: " + repoName + " Error: " + err.Error())
				return err
			}
			log.Println("Repository cloned successfully: " + repoName)
			length := len(repo.AfterScript)
			log.Println("Length of the slice:", length)
			if len(repo.AfterScript) > 0 {
				log.Println("Executing after script for repository: " + repoName)
				if err := s.executeAfterScript(repo.AfterScript); err != nil {
					log.Println("[ERR] Error executing after script: " + err.Error())
					return err
				}
			}
			return nil
		}
	}
	log.Println("[ERR] No matching repository found for: " + repoName)
	return errors.New("repository not found")
}

func (s *HTTPHandler) executeAfterScript(afterScript []AfterScript) error {
	for _, script := range afterScript {
		log.Println("Executing script: " + script.Command)
		// Here you would implement the logic to execute the script.
		cmd := exec.Command(script.Command, script.Args...)
		stdout, err := cmd.Output()
		if err != nil {
			log.Println("[ERR] Failed to execute script: " + script.Command + " Error: " + err.Error())
			return err
		}
		log.Println("Script executed successfully: " + string(stdout))
	}
	return nil
}
