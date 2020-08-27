package sqlite

import (
	"context"
	"database/sql"

	"github.com/aclivo/olap"
)

type storage struct {
	db *sql.DB
}

func NewSqliteStorage(db *sql.DB) olap.Storage {
	return &storage{db: db}
}

func (s *storage) AddCube(ctx context.Context, cube olap.Cube) error {
	return nil
}

func (s *storage) GetCube(ctx context.Context, name string) (olap.Cube, error) {
	return olap.Cube{}, nil
}

func (s *storage) CubeExists(ctx context.Context, name string) (bool, error) {
	return false, nil
}

func (s *storage) AddDimension(ctx context.Context, dim olap.Dimension) error {
	return nil
}

func (s *storage) GetDimension(ctx context.Context, name string) (olap.Dimension, error) {
	return olap.Dimension{}, nil
}

func (s *storage) DimensionExists(ctx context.Context, name string) (bool, error) {
	return false, nil
}

func (s *storage) AddElement(ctx context.Context, el olap.Element) error {
	return nil
}

func (s *storage) GetElement(ctx context.Context, dim, el string) (olap.Element, error) {
	return olap.Element{}, nil
}

func (s *storage) ElementExists(ctx context.Context, dim, name string) (bool, error) {
	return false, nil
}

func (s *storage) AddComponent(ctx context.Context, tot, el olap.Element) error {
	return nil
}

func (s *storage) GetComponent(ctx context.Context, dim, name string) (olap.Element, error) {
	return olap.Element{}, nil
}

func (s *storage) ComponentExists(ctx context.Context, dim, name string) (bool, error) {
	return false, nil
}

func (s *storage) Children(ctx context.Context, dim, name string) ([]olap.Element, error) {
	return []olap.Element{}, nil
}

func (s *storage) AddCell(ctx context.Context, cell olap.Cell) error {
	return nil
}

func (s *storage) GetCell(ctx context.Context, cube string, elements ...string) (olap.Cell, error) {
	return olap.Cell{}, nil
}

func (s *storage) AddCubeRules(ctx context.Context, cube string, rules olap.CubeRules) error {
	return nil
}

func (s *storage) GetCubeRules(ctx context.Context, cube string) (olap.CubeRules, error) {
	return olap.CubeRules{}, nil
}

func (s *storage) AddProcess(ctx context.Context, process olap.Process) error {
	return nil
}

func (s *storage) GetProcess(ctx context.Context, process string) (olap.Process, error) {
	return olap.Process{}, nil
}

func initDB(db *sql.DB) error {
	return nil
}
