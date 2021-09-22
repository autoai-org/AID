// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	ent "github.com/autoai-org/aid/ent/generated"
	entImage "github.com/autoai-org/aid/ent/generated/image"
	"github.com/autoai-org/aid/ent/generated/repository"
	"github.com/autoai-org/aid/internal/configuration"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"github.com/sirupsen/logrus"
)

// realBuild calls the Docker daemon to actually perform the build operation.
func realBuild(dockerfile string, imageName string, buildLogger *logrus.Logger, solver ent.Solver) {
	buildResponse, err := NewDockerRuntime().ImageBuild(context.Background(), getBuildCtx(path.Dir(dockerfile)), types.ImageBuildOptions{
		Tags:       []string{strings.ToLower(imageName)},
		Dockerfile: filepath.Base(dockerfile),
		Remove:     true,
	})
	utilities.ReportError(err, "Cannot build image "+imageName)
	if err != nil {
		buildLogger.Error("Cannot build image " + imageName)
		buildLogger.Error(err.Error())
	}
	// set logs to build logs
	reader := buildResponse.Body
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	// then scanning the build response and stream them to build logger.
	for scanner.Scan() {
		var buildLog BuildLog
		json.Unmarshal(scanner.Bytes(), &buildLog)
		buildLogger.Info(buildLog.Stream)
		if utilities.Verbose {
			fmt.Printf(buildLog.Stream)
		}
	}
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		buildLogger.Error("Cannot show log from " + imageName)
		buildLogger.Error(err.Error())
	}
	inspect, _, err := NewDockerRuntime().ImageInspectWithRaw(context.Background(), strings.ToLower(imageName))
	if err != nil {
		utilities.ReportError(err, "Cannot fetch image detail from docker host")
	}
	utilities.ReportError(err, "Cannot build image")
	if err == nil {
		utilities.Formatter.Info("Finishing building" + solver.Name + " ...")
		_, err = database.NewDefaultDB().Image.Create().SetUID(inspect.ID[7:15]).SetTitle(imageName).SetSolver(&solver).Save(context.Background())
		utilities.ReportError(err, "cannot save image to db")
		utilities.Formatter.Info("Please use " + inspect.ID[7:15] + " as the reference of the image.")
	}
}

// prepareBuild will prepare everything for the building process.
func prepareBuild(solver ent.Solver, gpu bool, block bool) (*ent.SystemLog, error) {
	utilities.Formatter.Info("Building Image for " + solver.Name + " ...")
	logUID := utilities.GenerateUUIDv4()[0:8]
	logPath := filepath.Join(utilities.GetBasePath(), "logs", "builds", logUID[0:8])
	log, err := database.NewDefaultDB().SystemLog.Create().SetFilepath(logPath).SetUID(logUID[0:8]).SetTitle(logUID[0:8]).SetSource("build").Save(context.Background())
	utilities.ReportError(err, "Cannot save to database")
	utilities.Formatter.Info("Building in progress, view full log at " + logPath)
	if utilities.Verbose {
		utilities.Formatter.Info("Verbose mode is on, detailed logs will be shown below.")
	}
	buildLogger := utilities.NewLogger(logPath)
	repo, err := solver.QueryRepository().First(context.Background())
	utilities.ReportError(err, "Cannot query repository of "+solver.Name)
	var dockerfile string
	if gpu {
		utilities.Formatter.Info("Building with GPU Enabled")
		dockerfile = filepath.Join(repo.Localpath, "docker_gpu_"+solver.Name)
	} else {
		utilities.Formatter.Info("Building with GPU Disabled")
		dockerfile = filepath.Join(repo.Localpath, "docker_"+solver.Name)
	}
	// if dockerfile does not exists, generate a default one.
	if !utilities.IsFileExists(dockerfile) {
		utilities.Formatter.Warn("Dockerfile not found, AID will generate a default version.")
		GenerateDockerFiles(filepath.Dir(dockerfile))
		tomlFilePath := filepath.Join(filepath.Dir(dockerfile), "aid.toml")
		aidToml, err := utilities.ReadFileContent(tomlFilePath)
		utilities.ReportError(err, "Cannot open file "+tomlFilePath)
		solvers := configuration.LoadSolversFromConfig(aidToml)
		RenderRunnerTpl(filepath.Dir(dockerfile), solvers.Solvers)
	}
	title := "aid/" + repo.Vendor + "/" + repo.Name + "/" + solver.Name
	if block {
		realBuild(dockerfile, title, buildLogger, solver)
	} else {
		go realBuild(dockerfile, title, buildLogger, solver)
	}
	return log, err
}

// BuildImage builds the image
func BuildImage(vendor string, packageName string, solverName string, gpu bool, block bool) string {
	var logID string
	repos, err := database.NewDefaultDB().Repository.Query().Where(repository.And(repository.Name(packageName), repository.Vendor(vendor))).First(context.Background())
	utilities.ReportError(err, "cannot find repos "+packageName)
	solvers, err := repos.QuerySolvers().All(context.Background())
	utilities.ReportError(err, "cannot find solvers of "+packageName)
	for _, solver := range solvers {
		if solver.Name == solverName {
			log, err := prepareBuild(*solver, gpu, block)
			if err != nil {
				utilities.Formatter.Error(err.Error())
			}
			logID = log.UID
		}
	}
	return logID
}

// RemoveImage deletes the image
func RemoveImage(imageUID string) error {
	imageEnt, err := database.NewDefaultDB().Image.Query().Where(entImage.UID(imageUID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch image from database: " + err.Error())
		os.Exit(3)
	}
	if _, err := NewDockerRuntime().ImageRemove(context.Background(), imageUID, types.ImageRemoveOptions{}); err != nil {
		utilities.Formatter.Error("Cannot remove image from docker: " + err.Error())
		os.Exit(4)
	}
	utilities.Formatter.Info("Image " + imageEnt.Title + "(" + imageUID + ") removed from Docker.")
	_, err = database.NewDefaultDB().Image.Delete().Where(entImage.UID(imageUID)).Exec(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot remove image from database: " + err.Error())
		os.Exit(4)
	}
	utilities.Formatter.Info("Image " + imageEnt.Title + "(" + imageUID + ") removed from database.")
	return err
}

// BuildWithPath builds the image with a path parameter.
// It should be used for testing, especially for ci-purpose.
// At this time, no database entries have been added.
// Hence the log will be printed to the console, rather than a file.
// If the solver in parameter is "all", the all solvers will be built.
func BuildWithPath(path string, solver string, removeAfterBuild bool) error {
	// first determine the docker file
	utilities.Formatter.Info("Building " + path + " in progess...")
	utilities.Formatter.Warn("Building with path mode, no log file will be created.")
	tomlFilePath := filepath.Join(path, "aid.toml")
	aidToml, err := utilities.ReadFileContent(tomlFilePath)
	if err != nil {
		utilities.Formatter.Error("Cannot read the toml file " + tomlFilePath + ": " + err.Error())

	}
	solvers := configuration.LoadSolversFromConfig(aidToml)
	if solver == "all" {
		utilities.Formatter.Warn("No solver name specified, will build all solvers under this folder")
	}
	var existingSolver []string
	for _, sol := range solvers.Solvers {
		existingSolver = append(existingSolver, sol.Name)
	}
	if solver != "all" {
		if !utilities.StringInArray(solver, existingSolver) {
			utilities.Formatter.Error("Cannot find the solver " + solver)
			os.Exit(4)
		}
		dockerFilepath := filepath.Join(path, "docker_"+solver)
		buildWithDockerfile(dockerFilepath, solver)
	} else {
		for _, sol := range existingSolver {
			dockerFilepath := filepath.Join(path, "docker_"+sol)
			buildWithDockerfile(dockerFilepath, sol)
		}
	}
	return err
}

func buildWithDockerfile(dockerfile string, solName string) {
	utilities.Formatter.Info(dockerfile)
	utilities.Formatter.Info(solName)
	if !utilities.IsFileExists(dockerfile) {
		utilities.Formatter.Warn("Dockerfile not found, AID will generate a default version.")
		GenerateDockerFiles(filepath.Dir(dockerfile))
		tomlFilePath := filepath.Join(filepath.Dir(dockerfile), "aid.toml")
		aidToml, err := utilities.ReadFileContent(tomlFilePath)
		utilities.ReportError(err, "Cannot open file "+tomlFilePath)
		solvers := configuration.LoadSolversFromConfig(aidToml)
		RenderRunnerTpl(filepath.Dir(dockerfile), solvers.Solvers)
	}
	imageName := "aid-standalone-" + solName
	buildResponse, err := NewDockerRuntime().ImageBuild(context.Background(), getBuildCtx(path.Dir(dockerfile)), types.ImageBuildOptions{
		Tags:       []string{strings.ToLower(imageName)},
		Dockerfile: filepath.Base(dockerfile),
		Remove:     true,
	})
	reader := buildResponse.Body
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var buildLog BuildLog
		json.Unmarshal(scanner.Bytes(), &buildLog)
		fmt.Printf(buildLog.Stream)
	}
	if err != nil {
		utilities.Formatter.Error(err.Error())
		os.Exit(5)
	}
}
