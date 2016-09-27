// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/routing-api/migration"
	"github.com/jinzhu/gorm"
)

type FakeMigration struct {
	RunMigrationStub        func(db *gorm.DB) error
	runMigrationMutex       sync.RWMutex
	runMigrationArgsForCall []struct {
		db *gorm.DB
	}
	runMigrationReturns struct {
		result1 error
	}
	VersionStub        func() int
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns struct {
		result1 int
	}
}

func (fake *FakeMigration) RunMigration(db *gorm.DB) error {
	fake.runMigrationMutex.Lock()
	fake.runMigrationArgsForCall = append(fake.runMigrationArgsForCall, struct {
		db *gorm.DB
	}{db})
	fake.runMigrationMutex.Unlock()
	if fake.RunMigrationStub != nil {
		return fake.RunMigrationStub(db)
	} else {
		return fake.runMigrationReturns.result1
	}
}

func (fake *FakeMigration) RunMigrationCallCount() int {
	fake.runMigrationMutex.RLock()
	defer fake.runMigrationMutex.RUnlock()
	return len(fake.runMigrationArgsForCall)
}

func (fake *FakeMigration) RunMigrationArgsForCall(i int) *gorm.DB {
	fake.runMigrationMutex.RLock()
	defer fake.runMigrationMutex.RUnlock()
	return fake.runMigrationArgsForCall[i].db
}

func (fake *FakeMigration) RunMigrationReturns(result1 error) {
	fake.RunMigrationStub = nil
	fake.runMigrationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) Version() int {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeMigration) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeMigration) VersionReturns(result1 int) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 int
	}{result1}
}

var _ migration.Migration = new(FakeMigration)
