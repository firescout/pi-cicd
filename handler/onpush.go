package handler

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func (s *HTTPHandler) handleOnPush(repoName string) error {
	err := os.Chdir(s.repoCloneDir)
	if err != nil {
		log.Println("[ERR] Error changing directory to REPO_CLONE_DIR: " + err.Error())
		return err
	}
	for _, repo := range s.repos.Repos {
		if repo.Name == repoName {
			cmd := exec.Command("git", "clone", repo.Url)
			// cmd.Dir = repo.Path
			cmd.Dir = s.repoCloneDir
			if err := cmd.Run(); err != nil {
				log.Println("[ERR] Error cloning repository: " + repoName + " Error: " + err.Error())
				return err
			}
			log.Println("Repository cloned successfully: " + repoName)
			log.Println("Executing post-pull script for repository: " + repoName)
			if err := s.postPullClean(repo.Name); err != nil {
				log.Println("[ERR] Error executing post-pull script: " + err.Error())
				return err
			}
			if len(repo.AfterScript) > 0 {
				log.Println("Executing after script for repository: " + repoName)
				if err := s.executeAfterScript(repo.Name, repo.AfterScript); err != nil {
					log.Println("[ERR] Error executing after script: " + err.Error())
					return err
				}
			}
			err = s.cleanupCloneDir(repo.Name)
			if err != nil {
				log.Println("[ERR] Error executing cleanup script: " + err.Error())
				return err
			}
			return nil
		}
	}
	log.Println("[ERR] No matching repository found for: " + repoName)
	return errors.New("repository not found")
}

func (s *HTTPHandler) executeAfterScript(repoName string, afterScript []AfterScript) error {
	for _, script := range afterScript {
		log.Println("Executing script: " + script.Command)
		// Here you would implement the logic to execute the script.
		cmd := exec.Command(script.Command, script.Args...)
		cmd.Dir = s.repoCloneDir + s.pathSeparator + repoName
		stdout, err := cmd.Output()
		if err != nil {
			log.Println("[ERR] Failed to execute script: " + script.Command + " Error: " + err.Error())
			return err
		}
		log.Println("Script executed successfully: " + string(stdout))
	}
	return nil
}

func (s *HTTPHandler) postPullClean(repoName string) error {
	err := os.RemoveAll(s.repoCloneDir + s.pathSeparator + repoName + s.pathSeparator + ".git")
	if err != nil {
		log.Println("[ERR] Error removing .git directory: " + err.Error())
		return err
	}
	err = os.RemoveAll(s.repoCloneDir + s.pathSeparator + repoName + s.pathSeparator + ".gitea")
	if err != nil {
		log.Println("[ERR] Error removing .gitea directory: " + err.Error())
		return err
	}
	err = os.RemoveAll(s.repoCloneDir + s.pathSeparator + repoName + s.pathSeparator + ".github")
	if err != nil {
		log.Println("[ERR] Error removing .github directory: " + err.Error())
		return err
	}
	err = os.Remove(s.repoCloneDir + s.pathSeparator + repoName + s.pathSeparator + ".gitignore")
	if err != nil {
		log.Println("[ERR] Error removing .gitignore file: " + err.Error())
		return err
	}
	return nil
}

func (s *HTTPHandler) cleanupCloneDir(repoName string) error {
	log.Println("Executing cleanup script: " + repoName)
	cmd := exec.Command("rm", "-r", repoName)
	cmd.Dir = s.repoCloneDir
	err := os.RemoveAll(s.repoCloneDir + s.pathSeparator + repoName)
	if err != nil {
		return err
	}
	return nil
}
