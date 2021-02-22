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

func realBuild(dockerfile string, imageName string, buildLogger *logrus.Logger) (types.ImageInspect, error) {
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
	imageInspect, _, err := NewDockerRuntime().ImageInspectWithRaw(context.Background(), strings.ToLower(imageName))
	if err != nil {
		utilities.ReportError(err, "Cannot fetch image detail from docker host")
	}
	return imageInspect, err
}

// prepareBuild will prepare everything for the building process.
func prepareBuild(solver ent.Solver) (*ent.SystemLog, error) {
	utilities.Formatter.Info("Building Image for " + solver.Name + " ...")
	logUID := utilities.GenerateUUIDv4()
	logPath := filepath.Join(utilities.GetBasePath(), "logs", "builds", logUID[0:10])
	log, err := database.NewDefaultDB().SystemLog.Create().SetFilepath(logPath).SetTitle(logUID[0:10]).SetSource("build").Save(context.Background())
	utilities.ReportError(err, "Cannot save to database")
	utilities.Formatter.Info("Building in progress, view full log at " + logPath)
	if utilities.Verbose {
		utilities.Formatter.Info("Verbose mode is on, detailed logs will be shown below.")
	}
	buildLogger := utilities.NewLogger(logPath)
	repo, err := solver.QueryRepository().First(context.Background())
	utilities.ReportError(err, "Cannot query repository of "+solver.Name)
	dockerfile := filepath.Join(repo.Localpath, "docker_"+solver.Name)
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
	inspect, err := realBuild(dockerfile, title, buildLogger)
	utilities.ReportError(err, "Cannot build image")
	if err == nil {
		utilities.Formatter.Info("Finishing building" + solver.Name + " ...")
		_, err = database.NewDefaultDB().Image.Create().SetUID(inspect.ID[7:17]).SetTitle(title).SetSolver(&solver).Save(context.Background())
		utilities.ReportError(err, "cannot save image to db")
		utilities.Formatter.Info("Please use " + inspect.ID[7:17] + " as the reference of the image.")
	}
	return log, err
}

// BuildImage builds the image
func BuildImage(vendor string, packageName string, solverName string) {
	repos, err := database.NewDefaultDB().Repository.Query().Where(repository.And(repository.Name(packageName), repository.Vendor(vendor))).First(context.Background())
	utilities.ReportError(err, "cannot find repos "+packageName)
	solvers, err := repos.QuerySolvers().All(context.Background())
	utilities.ReportError(err, "cannot find solvers of "+packageName)
	for _, solver := range solvers {
		if solver.Name == solverName {
			prepareBuild(*solver)
		}
	}
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
