package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string
func doNew(appName string) {
	appName = strings.ToLower(appName)
	appURL = appName
	//sanitize app name(convert url to single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}
	log.Println("AppName is:", appName)

	//git clone the skeleton application
	color.Green("\tCloning repository...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/mirei965/framingo-project.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		exitGracefully(err)
	}
	//remove the .git folder
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	//create a ready to go .env file
	color.Yellow("\tCreating .env file...")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}
	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${ENCRYPTION_KEY}", fra.RandomString(32))
	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}

	//create a make file
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	}

	_ = os.Remove(fmt.Sprintf("./%s/Makefile.mac", appName))
	_ = os.Remove(fmt.Sprintf("./%s/Makefile.windows", appName))

	//update the go.mod file
	color.Yellow("\tCreating go.mod file...")
	_ = os.Remove(fmt.Sprintf("./"+appName+"/go.mod"))
  data, err = templateFS.ReadFile("templates/go.mod.txt")
  if err != nil {
    exitGracefully(err)
  }

	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)
	err = copyDataToFile([]byte(mod), "./"+appName+"/go.mod")
	if err != nil {
		exitGracefully(err)
	}

	//update existing  .go file with correct name/imports
  color.Yellow("\tUpdating source files...")
  os.Chdir("./"+appName)
  updateSource()

	//run mod tidy in the project folder
	color.Yellow("\tRunning go mod tidy...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Start()
	if err != nil {
		exitGracefully(err)
	}
	color.Green("\tSuccessfully created new Framingo project! with: " +appURL)
	color.Green("\tTo start your awsome project")
}