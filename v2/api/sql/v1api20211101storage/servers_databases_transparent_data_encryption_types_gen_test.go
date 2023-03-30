// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ServersDatabasesTransparentDataEncryption_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesTransparentDataEncryption via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesTransparentDataEncryption, ServersDatabasesTransparentDataEncryptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesTransparentDataEncryption runs a test to see if a specific instance of ServersDatabasesTransparentDataEncryption round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesTransparentDataEncryption(subject ServersDatabasesTransparentDataEncryption) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesTransparentDataEncryption
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ServersDatabasesTransparentDataEncryption instances for property testing - lazily instantiated by
// ServersDatabasesTransparentDataEncryptionGenerator()
var serversDatabasesTransparentDataEncryptionGenerator gopter.Gen

// ServersDatabasesTransparentDataEncryptionGenerator returns a generator of ServersDatabasesTransparentDataEncryption instances for property testing.
func ServersDatabasesTransparentDataEncryptionGenerator() gopter.Gen {
	if serversDatabasesTransparentDataEncryptionGenerator != nil {
		return serversDatabasesTransparentDataEncryptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesTransparentDataEncryption(generators)
	serversDatabasesTransparentDataEncryptionGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesTransparentDataEncryption{}), generators)

	return serversDatabasesTransparentDataEncryptionGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesTransparentDataEncryption is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesTransparentDataEncryption(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Databases_TransparentDataEncryption_SpecGenerator()
	gens["Status"] = Servers_Databases_TransparentDataEncryption_STATUSGenerator()
}

func Test_Servers_Databases_TransparentDataEncryption_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_TransparentDataEncryption_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_Spec, Servers_Databases_TransparentDataEncryption_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_Spec runs a test to see if a specific instance of Servers_Databases_TransparentDataEncryption_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_Spec(subject Servers_Databases_TransparentDataEncryption_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_TransparentDataEncryption_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Servers_Databases_TransparentDataEncryption_Spec instances for property testing - lazily instantiated by
// Servers_Databases_TransparentDataEncryption_SpecGenerator()
var servers_Databases_TransparentDataEncryption_SpecGenerator gopter.Gen

// Servers_Databases_TransparentDataEncryption_SpecGenerator returns a generator of Servers_Databases_TransparentDataEncryption_Spec instances for property testing.
func Servers_Databases_TransparentDataEncryption_SpecGenerator() gopter.Gen {
	if servers_Databases_TransparentDataEncryption_SpecGenerator != nil {
		return servers_Databases_TransparentDataEncryption_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_Spec(generators)
	servers_Databases_TransparentDataEncryption_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_TransparentDataEncryption_Spec{}), generators)

	return servers_Databases_TransparentDataEncryption_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_Spec(gens map[string]gopter.Gen) {
	gens["OriginalVersion"] = gen.AlphaString()
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Databases_TransparentDataEncryption_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_TransparentDataEncryption_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_STATUS, Servers_Databases_TransparentDataEncryption_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_STATUS runs a test to see if a specific instance of Servers_Databases_TransparentDataEncryption_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_TransparentDataEncryption_STATUS(subject Servers_Databases_TransparentDataEncryption_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_TransparentDataEncryption_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Servers_Databases_TransparentDataEncryption_STATUS instances for property testing - lazily instantiated
// by Servers_Databases_TransparentDataEncryption_STATUSGenerator()
var servers_Databases_TransparentDataEncryption_STATUSGenerator gopter.Gen

// Servers_Databases_TransparentDataEncryption_STATUSGenerator returns a generator of Servers_Databases_TransparentDataEncryption_STATUS instances for property testing.
func Servers_Databases_TransparentDataEncryption_STATUSGenerator() gopter.Gen {
	if servers_Databases_TransparentDataEncryption_STATUSGenerator != nil {
		return servers_Databases_TransparentDataEncryption_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_STATUS(generators)
	servers_Databases_TransparentDataEncryption_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_TransparentDataEncryption_STATUS{}), generators)

	return servers_Databases_TransparentDataEncryption_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_TransparentDataEncryption_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
