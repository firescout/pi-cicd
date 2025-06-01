package handler

import (
	"errors"
	"log"
)

func (s *HTTPHandler) handleOnPush(repoName string) error {
	for _, repo := range s.repos.Repos {
		if repo.Name == repoName {
			if repo.AfterScript != "" {
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

func (s *HTTPHandler) executeAfterScript(afterScript string) error {
	return nil
}
