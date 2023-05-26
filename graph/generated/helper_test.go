package generated_test

import (
	"encoding/json"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jbactad/loop/application"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/resolvers"
	"github.com/jbactad/loop/infrastructure"
	"github.com/samber/do"
)

type snapshot struct {
	*cupaloy.Config
}

func (*snapshot) marshalIndent(v interface{}) ([]byte, error) {
	m, err := json.MarshalIndent(v, "", "  ")
	return m, err
}

func (s *snapshot) SnapshotJSON(v interface{}) error {
	m, err := s.marshalIndent(v)
	if err != nil {
		return err
	}

	s.Config.Snapshot(m)

	return nil
}

func (s *snapshot) SnapshotJSONT(t *testing.T, v interface{}) {
	m, err := s.marshalIndent(v)
	if err != nil {
		t.Fatal(err)
	}

	s.Config.SnapshotT(t, m)
}

func (s *snapshot) SnapshotJSONMulti(snapshotID string, v ...interface{}) {
	m, err := s.marshalIndent(v)
	if err != nil {
		panic(err)
	}

	s.Config.SnapshotMulti(snapshotID, m)
}

func (s *snapshot) Snapshot(v interface{}) {
	s.Config.Snapshot(v)
}

func (s *snapshot) SnapshotT(t *testing.T, v interface{}) {
	s.Config.SnapshotT(t, v)
}

func (s *snapshot) SnapshotMulti(snapshotID string, v ...interface{}) {
	s.Config.SnapshotMulti(snapshotID, v...)
}

func (s *snapshot) WithOptions(configurator ...cupaloy.Configurator) {
	s.Config.WithOptions(configurator...)
}

func NewSnapshoter() *snapshot {
	return &snapshot{
		Config: cupaloy.New(cupaloy.SnapshotFileExtension(".json")),
	}
}

var Snapshoter = NewSnapshoter()

func NewTestClient(t *testing.T) *client.Client {
	t.Helper()
	injector := do.NewWithOpts(&do.InjectorOpts{})

	err := infrastructure.ProvideDatabaseConnection(injector)
	if err != nil {
		t.Fatal(err)
	}

	infrastructure.ProvideRepositories(injector)
	application.ProvideQueryUseCases(injector)
	resolvers.ProvideResolver(injector)
	if err != nil {
		t.Fatal(err)
	}

	r := do.MustInvoke[*resolvers.Resolver](injector)

	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(*generated.NewConfig(r))))
}
