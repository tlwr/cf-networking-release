// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"policy-server/store/migrations"
	"sync"
)

type Migrator struct {
	PerformMigrationsStub        func(driverName string, migrationDb migrations.MigrationDb, migrateExecutor migrations.MigrateExecutor, maxNumMigrations int) (int, error)
	performMigrationsMutex       sync.RWMutex
	performMigrationsArgsForCall []struct {
		driverName       string
		migrationDb      migrations.MigrationDb
		migrateExecutor  migrations.MigrateExecutor
		maxNumMigrations int
	}
	performMigrationsReturns struct {
		result1 int
		result2 error
	}
	performMigrationsReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Migrator) PerformMigrations(driverName string, migrationDb migrations.MigrationDb, migrateExecutor migrations.MigrateExecutor, maxNumMigrations int) (int, error) {
	fake.performMigrationsMutex.Lock()
	ret, specificReturn := fake.performMigrationsReturnsOnCall[len(fake.performMigrationsArgsForCall)]
	fake.performMigrationsArgsForCall = append(fake.performMigrationsArgsForCall, struct {
		driverName       string
		migrationDb      migrations.MigrationDb
		migrateExecutor  migrations.MigrateExecutor
		maxNumMigrations int
	}{driverName, migrationDb, migrateExecutor, maxNumMigrations})
	fake.recordInvocation("PerformMigrations", []interface{}{driverName, migrationDb, migrateExecutor, maxNumMigrations})
	fake.performMigrationsMutex.Unlock()
	if fake.PerformMigrationsStub != nil {
		return fake.PerformMigrationsStub(driverName, migrationDb, migrateExecutor, maxNumMigrations)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.performMigrationsReturns.result1, fake.performMigrationsReturns.result2
}

func (fake *Migrator) PerformMigrationsCallCount() int {
	fake.performMigrationsMutex.RLock()
	defer fake.performMigrationsMutex.RUnlock()
	return len(fake.performMigrationsArgsForCall)
}

func (fake *Migrator) PerformMigrationsArgsForCall(i int) (string, migrations.MigrationDb, migrations.MigrateExecutor, int) {
	fake.performMigrationsMutex.RLock()
	defer fake.performMigrationsMutex.RUnlock()
	return fake.performMigrationsArgsForCall[i].driverName, fake.performMigrationsArgsForCall[i].migrationDb, fake.performMigrationsArgsForCall[i].migrateExecutor, fake.performMigrationsArgsForCall[i].maxNumMigrations
}

func (fake *Migrator) PerformMigrationsReturns(result1 int, result2 error) {
	fake.PerformMigrationsStub = nil
	fake.performMigrationsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Migrator) PerformMigrationsReturnsOnCall(i int, result1 int, result2 error) {
	fake.PerformMigrationsStub = nil
	if fake.performMigrationsReturnsOnCall == nil {
		fake.performMigrationsReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.performMigrationsReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Migrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.performMigrationsMutex.RLock()
	defer fake.performMigrationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Migrator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ store.Migrator = new(Migrator)
