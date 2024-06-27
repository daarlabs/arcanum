package migrator

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
	
	"github.com/daarlabs/arcanum/quirk"
)

type Migrator interface {
	IgnoreDatabase(names ...string) Migrator
	Run() error
	Init() error
	New() error
	Up() error
	Down() error
	
	MustRun()
	MustInit()
	MustNew()
	MustUp()
	MustDown()
}

type migrator struct {
	dir            string
	databases      map[string]*quirk.DB
	ignoreDatabase []string
	migrations     []*Migration
	init           *bool
	new            *bool
	up             *bool
	down           *bool
}

const (
	migrationsTable string = "quirk_migrations"
)

const (
	migrationFileContent = `package main

import "github.com/daarlabs/arcanum/quirk/migrator"

func init() {
	manager.Add().
		Up(
			func(c migrator.Control) {

			},
		).
		Down(
			func(c migrator.Control) {
			
			},
		)
}
`
)

func New(dir string, databases map[string]*quirk.DB, migrations []*Migration) Migrator {
	m := &migrator{dir: dir, databases: databases, migrations: migrations, ignoreDatabase: make([]string, 0)}
	return m
}

func (m *migrator) IgnoreDatabase(names ...string) Migrator {
	m.ignoreDatabase = append(m.ignoreDatabase, names...)
	return m
}

func (m *migrator) Run() error {
	m.init = flag.Bool("init", false, "Init migrations")
	m.new = flag.Bool("new", false, "New migration")
	m.up = flag.Bool("up", false, "Up migrations")
	m.down = flag.Bool("down", false, "Down migration")
	flag.Parse()
	flag.Parse()
	if *m.init {
		return m.Init()
	}
	if *m.new {
		return m.New()
	}
	if *m.up {
		return m.Up()
	}
	if *m.down {
		return m.Down()
	}
	return nil
}

func (m *migrator) MustRun() {
	if err := m.Run(); err != nil {
		panic(err)
	}
}

func (m *migrator) Init() error {
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		exists, err := m.migrationsTableExists(db)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		if err := quirk.New(db).Q(
			fmt.Sprintf(
				`CREATE TABLE IF NOT EXISTS %s (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
    )`, migrationsTable,
			),
		).Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (m *migrator) MustInit() {
	if err := m.Init(); err != nil {
		panic(err)
	}
}

func (m *migrator) New() error {
	if len(m.dir) == 0 {
		return ErrorInvalidDir
	}
	m.check(os.MkdirAll(m.dir, os.ModePerm))
	filepath := fmt.Sprintf("%s/%d.go", m.dir, time.Now().UnixNano())
	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(migrationFileContent)
	if err != nil {
		return err
	}
	return nil
}

func (m *migrator) MustNew() {
	if err := m.New(); err != nil {
		panic(err)
	}
}

func (m *migrator) Up() error {
	existingMigrationsNames, err := m.getExistingMigrationsNames()
	if err != nil {
		return err
	}
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		if _, err := db.Begin(); err != nil {
			return err
		}
	}
	for _, item := range m.migrations {
		if slices.Contains(existingMigrationsNames, item.name) {
			continue
		}
		fmt.Printf("Up [%s]...\n", item.name)
		item.up(&control{m})
		if err := m.insertMigration(item.name); err != nil {
			return err
		}
	}
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		if err := db.Commit(); err != nil {
			return err
		}
	}
	for _, item := range m.migrations {
		if slices.Contains(existingMigrationsNames, item.name) {
			continue
		}
		fmt.Printf("Up successful [%s]!\n", item.name)
	}
	return nil
}

func (m *migrator) MustUp() {
	if err := m.Up(); err != nil {
		panic(err)
	}
}

func (m *migrator) Down() error {
	lastMigrationName, err := m.getLastMigrationName()
	if err != nil {
		return err
	}
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		if _, err := db.Begin(); err != nil {
			return err
		}
	}
	for _, item := range m.migrations {
		if !slices.Contains(lastMigrationName, item.name) {
			continue
		}
		fmt.Printf("Down [%s]...\n", item.name)
		item.down(&control{m})
		if err := m.deleteMigration(item.name); err != nil {
			return err
		}
	}
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		if err := db.Commit(); err != nil {
			return err
		}
	}
	for _, item := range m.migrations {
		if !slices.Contains(lastMigrationName, item.name) {
			continue
		}
		fmt.Printf("Down successful [%s]!\n", item.name)
	}
	return nil
}

func (m *migrator) MustDown() {
	if err := m.Down(); err != nil {
		panic(err)
	}
}

func (m *migrator) check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func (m *migrator) getLastMigrationName() ([]string, error) {
	result := make([]string, 0)
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		exists, err := m.migrationsTableExists(db)
		if err != nil {
			return result, err
		}
		if !exists {
			continue
		}
		var r string
		if err := quirk.New(db).Q(
			fmt.Sprintf(
				`SELECT name FROM %s ORDER BY created_at DESC LIMIT 1`, migrationsTable,
			),
		).Exec(&r); err != nil {
			return result, err
		}
		if !slices.Contains(result, r) {
			result = append(result, r)
		}
	}
	return result, nil
}

func (m *migrator) migrationsTableExists(db *quirk.DB) (bool, error) {
	var r bool
	if err := db.Q(
		fmt.Sprintf(
			`SELECT EXISTS (
SELECT 1
FROM pg_tables
WHERE tablename = '%s'
) AS table_existence`, migrationsTable,
		),
	).Exec(&r); err != nil {
		return r, err
	}
	return r, nil
}

func (m *migrator) getExistingMigrationsNames() ([]string, error) {
	result := make([]string, 0)
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		exists, err := m.migrationsTableExists(db)
		if err != nil {
			return result, err
		}
		if !exists {
			continue
		}
		r := make([]string, 0)
		if err := quirk.New(db).Q(
			fmt.Sprintf(
				`SELECT name FROM %s ORDER BY created_at ASC`, migrationsTable,
			),
		).Exec(&r); err != nil {
			return result, err
		}
		for _, name := range r {
			if !slices.Contains(result, name) {
				result = append(result, name)
			}
		}
	}
	return result, nil
}

func (m *migrator) insertMigration(name string) error {
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		exists, err := m.migrationsTableExists(db)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		if err := quirk.New(db).
			Q(
				fmt.Sprintf(
					`INSERT INTO %s (id, name, created_at, updated_at) VALUES (DEFAULT, @name, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`,
					migrationsTable,
				),
				quirk.Map{"name": name},
			).
			Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (m *migrator) deleteMigration(name string) error {
	for name, db := range m.databases {
		if slices.Contains(m.ignoreDatabase, name) {
			continue
		}
		exists, err := m.migrationsTableExists(db)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		if err := quirk.New(db).Q(
			fmt.Sprintf(`DELETE FROM %s WHERE name = @name`, migrationsTable), quirk.Map{"name": name},
		).Exec(); err != nil {
			return err
		}
	}
	return nil
}
