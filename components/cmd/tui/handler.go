package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/autoai-org/aid/internal/initialization"

	"github.com/autoai-org/aid/internal/configuration"

	markdown "github.com/MichaelMure/go-term-markdown"
	entRepository "github.com/autoai-org/aid/ent/generated/repository"
	"github.com/autoai-org/aid/internal/daemon"
	"github.com/autoai-org/aid/internal/database"

	"github.com/autoai-org/aid/internal/runtime/cargo"
	"github.com/autoai-org/aid/internal/runtime/docker"
	"github.com/autoai-org/aid/internal/runtime/requests"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/autoai-org/aid/internal/workflow"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func installPackage(remoteURL string) {
	workflow.PullPackageSource(remoteURL)
}

// listObject
func listObject(objectName string) {
	listEntity(objectName)
}

// startServer
func startServer(port string) {
	daemon.RunServer(port)
}

func buildImage(buildContext string, gpu bool) {
	buildInfo := strings.Split(buildContext, "/")
	workflow.BuildDockerImage(buildInfo[0], buildInfo[1], buildInfo[2], gpu, false)
}

func buildImageByPath(path string, solver string, autoRemove bool) {
	workflow.BuildDockerImageWithPath(path, solver, autoRemove)
}

func createContainer(imageID string, hostPort string, gpu bool) {
	if hostPort == "" {
		utilities.Formatter.Error("Hostport is not given... Aborted")
		os.Exit(4)
	}
	workflow.CreateContainer(imageID, hostPort, gpu)
}

func startContainer(containerID string) {
	workflow.StartContainer(containerID)
}

func infer(containerID string, args cli.Args) {
	params := make(map[string]string)
	for _, param := range args.Tail() {
		kv := strings.Split(param, "=")
		params[kv[0]] = kv[1]
	}
	resp, err := requests.NewHTTPClient().Infer(containerID, params)
	if err != nil {
		utilities.ReportError(err, "Cannot handle requests from "+containerID)
		os.Exit(6)
	}
	utilities.Formatter.Info("Inference successfully returned:")
	fmt.Println(resp.String())
}

func help(packageID string) {
	repo, err := database.NewDefaultDB().Repository.Query().Where(entRepository.UID(packageID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch repository " + fmt.Sprint(packageID) + ", Aborted")
		os.Exit(3)
	}
	readmeFilePath := filepath.Join(repo.Localpath, "README.md")
	readmeFile, err := ioutil.ReadFile(readmeFilePath)
	if err != nil {
		utilities.Formatter.Error("Cannot read the readme file from " + readmeFilePath)
		os.Exit(6)
	}
	result := markdown.Render(string(readmeFile), 80, 6)
	separator := markdown.Render(string("---"), 80, 6)
	fmt.Println(string(result))
	fmt.Println(string(separator))
	for _, solver := range repo.Edges.Solvers {
		fmt.Println(solver.Name)
	}
}

func remove(entity string, identifier string) error {
	var err error
	switch entity {
	case "package":
		err = cargo.RemovePackage(identifier)
	case "container":
		err = docker.RemoveContainer(identifier)
	case "image":
		err = docker.RemoveImage(identifier)
	}
	if err != nil {
		utilities.Formatter.Error("Cannot remove " + entity + " with the id " + identifier + ": " + err.Error())
		os.Exit(3)
	}
	return err
}

func initRepo(c *cli.Context) {
	prompt := promptui.Prompt{
		Label: "Your Name (Please use the same username with your GitHub account for now)",
	}
	vendorName, err := prompt.Run()
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	prompt = promptui.Prompt{
		Label: "Your Package Name",
	}
	repoName, err := prompt.Run()
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	prompt = promptui.Prompt{
		Label: "One-line Description (You can change it in README.md file later as well)",
	}
	description, err := prompt.Run()

	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	workflow.InitNewRepo(vendorName, repoName, description)
	utilities.Formatter.Info(repoName + " created!")
	utilities.Formatter.Info("Check https://aid.autoai.org/docs/usage/manual/publish/prepare_locally for what to do next.")
}

func headlessDaemon(c *cli.Context) {
	Port := os.Getenv("AID_PORT")
	remoteURL := os.Getenv("AID_MODEL")
	installPackage(remoteURL)
	targetPath := filepath.Join(utilities.GetBasePath(), "models")
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	targetSubFolder := filepath.Join(targetPath, vendorName, repoName)

	absTargetSubFolder, _ := filepath.Abs(targetSubFolder)

	tomlString, err := utilities.ReadFileContent(filepath.Join(absTargetSubFolder, "aid.toml"))

	utilities.ReportError(err, "cannot parse aid.toml file")
	packageConfig := configuration.LoadPackageFromConfig(tomlString)
	solver := packageConfig.Solvers[0]
	workflow.BuildDockerImage(vendorName, repoName, solver.Name, false, true)
	image, err := database.DefaultDB.Image.Query().First(context.Background())
	if err != nil {
		utilities.ReportError(err, "cannot find container")
	}
	workflow.CreateContainer(image.UID, Port, false)
	container, err := database.DefaultDB.Container.Query().First(context.Background())
	if err != nil {
		utilities.ReportError(err, "cannot find container")
	}
	workflow.StartContainer(container.UID)
}

func generate(c *cli.Context) {
	remoteURL := os.Getenv("AID_MODEL")
	installPackage(remoteURL)
	targetPath := filepath.Join(utilities.GetBasePath(), "models")
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	targetSubFolder := filepath.Join(targetPath, vendorName, repoName)
	absTargetSubFolder, _ := filepath.Abs(targetSubFolder)

	tomlString, err := utilities.ReadFileContent(filepath.Join(absTargetSubFolder, "aid.toml"))

	utilities.ReportError(err, "cannot parse aid.toml file")
	solvers := configuration.LoadSolversFromConfig(tomlString)
	docker.RenderRunnerTpl(absTargetSubFolder, solvers.Solvers)
	utilities.Formatter.Info("Generated runner_sentimentSolver.py")
}
