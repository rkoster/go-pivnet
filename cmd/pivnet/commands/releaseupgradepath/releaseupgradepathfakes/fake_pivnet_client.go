// This file was generated by counterfeiter
package releaseupgradepathfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/go-pivnet/cmd/pivnet/commands/releaseupgradepath"
)

type FakePivnetClient struct {
	ReleaseForProductVersionStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	releaseForProductVersionMutex       sync.RWMutex
	releaseForProductVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForProductVersionReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	ReleaseUpgradePathsStub        func(productSlug string, releaseID int) ([]go_pivnet.ReleaseUpgradePath, error)
	releaseUpgradePathsMutex       sync.RWMutex
	releaseUpgradePathsArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseUpgradePathsReturns struct {
		result1 []go_pivnet.ReleaseUpgradePath
		result2 error
	}
	AddReleaseUpgradePathStub        func(productSlug string, releaseID int, upgradeFromReleaseID int) error
	addReleaseUpgradePathMutex       sync.RWMutex
	addReleaseUpgradePathArgsForCall []struct {
		productSlug          string
		releaseID            int
		upgradeFromReleaseID int
	}
	addReleaseUpgradePathReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseForProductVersion(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.releaseForProductVersionMutex.Lock()
	fake.releaseForProductVersionArgsForCall = append(fake.releaseForProductVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForProductVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForProductVersionMutex.Unlock()
	if fake.ReleaseForProductVersionStub != nil {
		return fake.ReleaseForProductVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForProductVersionReturns.result1, fake.releaseForProductVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForProductVersionCallCount() int {
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	return len(fake.releaseForProductVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForProductVersionArgsForCall(i int) (string, string) {
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	return fake.releaseForProductVersionArgsForCall[i].productSlug, fake.releaseForProductVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForProductVersionReturns(result1 go_pivnet.Release, result2 error) {
	fake.ReleaseForProductVersionStub = nil
	fake.releaseForProductVersionReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseUpgradePaths(productSlug string, releaseID int) ([]go_pivnet.ReleaseUpgradePath, error) {
	fake.releaseUpgradePathsMutex.Lock()
	fake.releaseUpgradePathsArgsForCall = append(fake.releaseUpgradePathsArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseUpgradePaths", []interface{}{productSlug, releaseID})
	fake.releaseUpgradePathsMutex.Unlock()
	if fake.ReleaseUpgradePathsStub != nil {
		return fake.ReleaseUpgradePathsStub(productSlug, releaseID)
	} else {
		return fake.releaseUpgradePathsReturns.result1, fake.releaseUpgradePathsReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseUpgradePathsCallCount() int {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return len(fake.releaseUpgradePathsArgsForCall)
}

func (fake *FakePivnetClient) ReleaseUpgradePathsArgsForCall(i int) (string, int) {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return fake.releaseUpgradePathsArgsForCall[i].productSlug, fake.releaseUpgradePathsArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseUpgradePathsReturns(result1 []go_pivnet.ReleaseUpgradePath, result2 error) {
	fake.ReleaseUpgradePathsStub = nil
	fake.releaseUpgradePathsReturns = struct {
		result1 []go_pivnet.ReleaseUpgradePath
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddReleaseUpgradePath(productSlug string, releaseID int, upgradeFromReleaseID int) error {
	fake.addReleaseUpgradePathMutex.Lock()
	fake.addReleaseUpgradePathArgsForCall = append(fake.addReleaseUpgradePathArgsForCall, struct {
		productSlug          string
		releaseID            int
		upgradeFromReleaseID int
	}{productSlug, releaseID, upgradeFromReleaseID})
	fake.recordInvocation("AddReleaseUpgradePath", []interface{}{productSlug, releaseID, upgradeFromReleaseID})
	fake.addReleaseUpgradePathMutex.Unlock()
	if fake.AddReleaseUpgradePathStub != nil {
		return fake.AddReleaseUpgradePathStub(productSlug, releaseID, upgradeFromReleaseID)
	} else {
		return fake.addReleaseUpgradePathReturns.result1
	}
}

func (fake *FakePivnetClient) AddReleaseUpgradePathCallCount() int {
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	return len(fake.addReleaseUpgradePathArgsForCall)
}

func (fake *FakePivnetClient) AddReleaseUpgradePathArgsForCall(i int) (string, int, int) {
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	return fake.addReleaseUpgradePathArgsForCall[i].productSlug, fake.addReleaseUpgradePathArgsForCall[i].releaseID, fake.addReleaseUpgradePathArgsForCall[i].upgradeFromReleaseID
}

func (fake *FakePivnetClient) AddReleaseUpgradePathReturns(result1 error) {
	fake.AddReleaseUpgradePathStub = nil
	fake.addReleaseUpgradePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ releaseupgradepath.PivnetClient = new(FakePivnetClient)
