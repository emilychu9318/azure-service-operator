// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_MongoDBCollectionGetResults_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetResults_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetResultsSTATUSARM, MongoDBCollectionGetResultsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetResultsSTATUSARM runs a test to see if a specific instance of MongoDBCollectionGetResults_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetResultsSTATUSARM(subject MongoDBCollectionGetResults_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetResults_STATUSARM
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

// Generator of MongoDBCollectionGetResults_STATUSARM instances for property testing - lazily instantiated by
// MongoDBCollectionGetResultsSTATUSARMGenerator()
var mongoDBCollectionGetResultsSTATUSARMGenerator gopter.Gen

// MongoDBCollectionGetResultsSTATUSARMGenerator returns a generator of MongoDBCollectionGetResults_STATUSARM instances for property testing.
// We first initialize mongoDBCollectionGetResultsSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetResultsSTATUSARMGenerator() gopter.Gen {
	if mongoDBCollectionGetResultsSTATUSARMGenerator != nil {
		return mongoDBCollectionGetResultsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM(generators)
	mongoDBCollectionGetResultsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM(generators)
	mongoDBCollectionGetResultsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_STATUSARM{}), generators)

	return mongoDBCollectionGetResultsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBCollectionGetPropertiesSTATUSARMGenerator())
}

func Test_MongoDBCollectionGetProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSARM, MongoDBCollectionGetPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSARM runs a test to see if a specific instance of MongoDBCollectionGetProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSARM(subject MongoDBCollectionGetProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_STATUSARM
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

// Generator of MongoDBCollectionGetProperties_STATUSARM instances for property testing - lazily instantiated by
// MongoDBCollectionGetPropertiesSTATUSARMGenerator()
var mongoDBCollectionGetPropertiesSTATUSARMGenerator gopter.Gen

// MongoDBCollectionGetPropertiesSTATUSARMGenerator returns a generator of MongoDBCollectionGetProperties_STATUSARM instances for property testing.
func MongoDBCollectionGetPropertiesSTATUSARMGenerator() gopter.Gen {
	if mongoDBCollectionGetPropertiesSTATUSARMGenerator != nil {
		return mongoDBCollectionGetPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSARM(generators)
	mongoDBCollectionGetPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_STATUSARM{}), generators)

	return mongoDBCollectionGetPropertiesSTATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceSTATUSARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionGetPropertiesSTATUSResourceARMGenerator())
}

func Test_MongoDBCollectionGetProperties_STATUS_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_STATUS_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSResourceARM, MongoDBCollectionGetPropertiesSTATUSResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSResourceARM runs a test to see if a specific instance of MongoDBCollectionGetProperties_STATUS_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetPropertiesSTATUSResourceARM(subject MongoDBCollectionGetProperties_STATUS_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_STATUS_ResourceARM
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

// Generator of MongoDBCollectionGetProperties_STATUS_ResourceARM instances for property testing - lazily instantiated
// by MongoDBCollectionGetPropertiesSTATUSResourceARMGenerator()
var mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator gopter.Gen

// MongoDBCollectionGetPropertiesSTATUSResourceARMGenerator returns a generator of MongoDBCollectionGetProperties_STATUS_ResourceARM instances for property testing.
// We first initialize mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetPropertiesSTATUSResourceARMGenerator() gopter.Gen {
	if mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator != nil {
		return mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM(generators)
	mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_STATUS_ResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM(generators)
	mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_STATUS_ResourceARM{}), generators)

	return mongoDBCollectionGetPropertiesSTATUSResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesSTATUSResourceARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexSTATUSARMGenerator())
}

func Test_OptionsResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResourceSTATUSARM, OptionsResourceSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResourceSTATUSARM runs a test to see if a specific instance of OptionsResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResourceSTATUSARM(subject OptionsResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_STATUSARM
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

// Generator of OptionsResource_STATUSARM instances for property testing - lazily instantiated by
// OptionsResourceSTATUSARMGenerator()
var optionsResourceSTATUSARMGenerator gopter.Gen

// OptionsResourceSTATUSARMGenerator returns a generator of OptionsResource_STATUSARM instances for property testing.
// We first initialize optionsResourceSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResourceSTATUSARMGenerator() gopter.Gen {
	if optionsResourceSTATUSARMGenerator != nil {
		return optionsResourceSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceSTATUSARM(generators)
	optionsResourceSTATUSARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResourceSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForOptionsResourceSTATUSARM(generators)
	optionsResourceSTATUSARMGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUSARM{}), generators)

	return optionsResourceSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResourceSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResourceSTATUSARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResourceSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResourceSTATUSARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsSTATUSARMGenerator())
}

func Test_AutoscaleSettings_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsSTATUSARM, AutoscaleSettingsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsSTATUSARM runs a test to see if a specific instance of AutoscaleSettings_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsSTATUSARM(subject AutoscaleSettings_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_STATUSARM
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

// Generator of AutoscaleSettings_STATUSARM instances for property testing - lazily instantiated by
// AutoscaleSettingsSTATUSARMGenerator()
var autoscaleSettingsSTATUSARMGenerator gopter.Gen

// AutoscaleSettingsSTATUSARMGenerator returns a generator of AutoscaleSettings_STATUSARM instances for property testing.
func AutoscaleSettingsSTATUSARMGenerator() gopter.Gen {
	if autoscaleSettingsSTATUSARMGenerator != nil {
		return autoscaleSettingsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsSTATUSARM(generators)
	autoscaleSettingsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_STATUSARM{}), generators)

	return autoscaleSettingsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsSTATUSARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_MongoIndex_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexSTATUSARM, MongoIndexSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexSTATUSARM runs a test to see if a specific instance of MongoIndex_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexSTATUSARM(subject MongoIndex_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_STATUSARM
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

// Generator of MongoIndex_STATUSARM instances for property testing - lazily instantiated by
// MongoIndexSTATUSARMGenerator()
var mongoIndexSTATUSARMGenerator gopter.Gen

// MongoIndexSTATUSARMGenerator returns a generator of MongoIndex_STATUSARM instances for property testing.
func MongoIndexSTATUSARMGenerator() gopter.Gen {
	if mongoIndexSTATUSARMGenerator != nil {
		return mongoIndexSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexSTATUSARM(generators)
	mongoIndexSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoIndex_STATUSARM{}), generators)

	return mongoIndexSTATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexSTATUSARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysSTATUSARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsSTATUSARMGenerator())
}

func Test_MongoIndexKeys_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysSTATUSARM, MongoIndexKeysSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysSTATUSARM runs a test to see if a specific instance of MongoIndexKeys_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysSTATUSARM(subject MongoIndexKeys_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_STATUSARM
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

// Generator of MongoIndexKeys_STATUSARM instances for property testing - lazily instantiated by
// MongoIndexKeysSTATUSARMGenerator()
var mongoIndexKeysSTATUSARMGenerator gopter.Gen

// MongoIndexKeysSTATUSARMGenerator returns a generator of MongoIndexKeys_STATUSARM instances for property testing.
func MongoIndexKeysSTATUSARMGenerator() gopter.Gen {
	if mongoIndexKeysSTATUSARMGenerator != nil {
		return mongoIndexKeysSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysSTATUSARM(generators)
	mongoIndexKeysSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_STATUSARM{}), generators)

	return mongoIndexKeysSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysSTATUSARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsSTATUSARM, MongoIndexOptionsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsSTATUSARM runs a test to see if a specific instance of MongoIndexOptions_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsSTATUSARM(subject MongoIndexOptions_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_STATUSARM
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

// Generator of MongoIndexOptions_STATUSARM instances for property testing - lazily instantiated by
// MongoIndexOptionsSTATUSARMGenerator()
var mongoIndexOptionsSTATUSARMGenerator gopter.Gen

// MongoIndexOptionsSTATUSARMGenerator returns a generator of MongoIndexOptions_STATUSARM instances for property testing.
func MongoIndexOptionsSTATUSARMGenerator() gopter.Gen {
	if mongoIndexOptionsSTATUSARMGenerator != nil {
		return mongoIndexOptionsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsSTATUSARM(generators)
	mongoIndexOptionsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_STATUSARM{}), generators)

	return mongoIndexOptionsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsSTATUSARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}