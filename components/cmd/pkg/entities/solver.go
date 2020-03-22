// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

// Solver defines the struct of a solver, the minimal struct of a inference program
type Solver struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Class     string    `db:"solverpath"`
	Vendor    string    `db:"vendor"`
	Package   string    `db:"package"`
	Status    string    `db:"status"`
	ImageName string    `db:"imagename"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Image saves the built images from docker client
type Image struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Container saves the containers that are created from images
type Container struct {
	ID        string    `db:"id"`
	ImageID   string    `db:"imageid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// RunningSolver wraps the docker container that runs as solver
type RunningSolver struct {
	ID         string    `db:"id"`
	ImageID    string    `db:"imageId"`
	Status     string    `db:"status"`
	ImageName  string    `db:"imagename"`
	EntryPoint string    `db:"entrypoint"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Solvers defines a list of solvers
type Solvers struct {
	Solvers []Solver
}

// TableName defines the tablename of solver in database
func (s *Solver) TableName() string {
	return "solver"
}

// TableName defines the tablename of image in database
func (i *Image) TableName() string {
	return "image"
}

// TableName defines the tablename of container in database
func (c *Container) TableName() string {
	return "container"
}

// TableName defines the tablename of running solver in database
func (rs *RunningSolver) TableName() string {
	return "runningsolver"
}

// PK defines the primary key of Solver
func (s *Solver) PK() string {
	return "id"
}

// PK defines the primary key of Image
func (i *Image) PK() string {
	return "id"
}

// PK defines the primary key of Container
func (c *Container) PK() string {
	return "id"
}

// PK defines the primary key of RunningSolver
func (rs *RunningSolver) PK() string {
	return "id"
}

// Save stores package into database
func (s *Solver) Save() error {
	s.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(s)
}

// Save stores image into database
func (i *Image) Save() error {
	i.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(i)
}

// Save runningSolver into database
func (rs *RunningSolver) Save() error {
	rs.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(rs)
}

// Save stores container into database
func (c *Container) Save() error {
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(c)
}

// GetImage returns the image object by image id
func GetImage(id string) Image {
	image := Image{ID: id}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&image)
	utilities.CheckError(err, "Cannot fetch image object with id "+id)
	return image
}

// GetImagebyName returns the image object by image name
func GetImagebyName(name string) Image {
	image := Image{Name: name}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&image)
	utilities.CheckError(err, "Cannot fetch image object with name "+name)
	return image
}

// GetContainer returns the container object by container id
func GetContainer(id string) Container {
	container := Container{ID: id}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&container)
	utilities.CheckError(err, "Cannot fetch container object with id "+id)
	return container
}

// GetNextImageNumber returns the count+1 of current images in string
func GetNextImageNumber() string {
	image := Image{}
	db := storage.GetDefaultDB()
	count, err := db.Count(&image)
	utilities.CheckError(err, "Cannot count image object")
	nextCount := count + 1
	return strconv.FormatInt(int64(nextCount), 10)
}

// FetchSolvers returns all solvers in a single call
func FetchSolvers() []Solver {
	solversPointers := make([]*Solver, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&solversPointers)
	solvers := make([]Solver, len(solversPointers))
	for i := range solversPointers {
		solvers[i] = *solversPointers[i]
	}
	return solvers
}

// FetchImages returns all images in a single call
func FetchImages() []Image {
	imagesPointers := make([]*Image, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&imagesPointers)
	images := make([]Image, len(imagesPointers))
	for i := range imagesPointers {
		images[i] = *imagesPointers[i]
	}
	return images
}

// FetchContainers returns all containers in a single call
func FetchContainers() []Container {
	containersPointers := make([]*Container, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&containersPointers)
	containers := make([]Container, len(containersPointers))
	for i := range containersPointers {
		containers[i] = *containersPointers[i]
	}
	return containers
}

// LoadSolversFromConfig reads the config string and returns the objects
func LoadSolversFromConfig(tomlString string) Solvers {
	var solvers Solvers
	_, err := toml.Decode(tomlString, &solvers)
	utilities.CheckError(err, "Cannot Load Solvers")
	return solvers
}

// GetRunningSolver returns the current solver that is marked as running.
func GetRunningSolver(ID string) RunningSolver {
	var targetSolver = RunningSolver{ID: ID}
	db := storage.GetDefaultDB()
	db.FetchOne(&targetSolver)
	return targetSolver
}

// GetAllRunningSolvers returns the array of all running solvers
// which belongs the a single solver.
func GetAllRunningSolvers(ID string) []RunningSolver {
	runningSolverPointers := make([]*RunningSolver, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&runningSolverPointers)
	var runningSolvers []RunningSolver
	for i := range runningSolverPointers {
		if runningSolverPointers[i].ImageID == ID {
			runningSolvers = append(runningSolvers, *runningSolverPointers[i])
		}
	}
	return runningSolvers
}
