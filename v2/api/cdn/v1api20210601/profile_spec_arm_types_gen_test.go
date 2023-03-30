// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

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

func Test_Profile_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfile_Spec_ARM, Profile_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfile_Spec_ARM runs a test to see if a specific instance of Profile_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfile_Spec_ARM(subject Profile_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_Spec_ARM
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

// Generator of Profile_Spec_ARM instances for property testing - lazily instantiated by Profile_Spec_ARMGenerator()
var profile_Spec_ARMGenerator gopter.Gen

// Profile_Spec_ARMGenerator returns a generator of Profile_Spec_ARM instances for property testing.
// We first initialize profile_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profile_Spec_ARMGenerator() gopter.Gen {
	if profile_Spec_ARMGenerator != nil {
		return profile_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_Spec_ARM(generators)
	profile_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profile_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfile_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProfile_Spec_ARM(generators)
	profile_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Profile_Spec_ARM{}), generators)

	return profile_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfile_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfile_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfile_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfile_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfileProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_ProfileProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileProperties_ARM, ProfileProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileProperties_ARM runs a test to see if a specific instance of ProfileProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileProperties_ARM(subject ProfileProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties_ARM
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

// Generator of ProfileProperties_ARM instances for property testing - lazily instantiated by
// ProfileProperties_ARMGenerator()
var profileProperties_ARMGenerator gopter.Gen

// ProfileProperties_ARMGenerator returns a generator of ProfileProperties_ARM instances for property testing.
func ProfileProperties_ARMGenerator() gopter.Gen {
	if profileProperties_ARMGenerator != nil {
		return profileProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileProperties_ARM(generators)
	profileProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_ARM{}), generators)

	return profileProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfileProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileProperties_ARM(gens map[string]gopter.Gen) {
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
}

func Test_Sku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_ARM, Sku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_ARM runs a test to see if a specific instance of Sku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_ARM(subject Sku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_ARM
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

// Generator of Sku_ARM instances for property testing - lazily instantiated by Sku_ARMGenerator()
var sku_ARMGenerator gopter.Gen

// Sku_ARMGenerator returns a generator of Sku_ARM instances for property testing.
func Sku_ARMGenerator() gopter.Gen {
	if sku_ARMGenerator != nil {
		return sku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_ARM(generators)
	sku_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_ARM{}), generators)

	return sku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_Custom_Verizon,
		Sku_Name_Premium_AzureFrontDoor,
		Sku_Name_Premium_Verizon,
		Sku_Name_StandardPlus_955BandWidth_ChinaCdn,
		Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn,
		Sku_Name_StandardPlus_ChinaCdn,
		Sku_Name_Standard_955BandWidth_ChinaCdn,
		Sku_Name_Standard_Akamai,
		Sku_Name_Standard_AvgBandWidth_ChinaCdn,
		Sku_Name_Standard_AzureFrontDoor,
		Sku_Name_Standard_ChinaCdn,
		Sku_Name_Standard_Microsoft,
		Sku_Name_Standard_Verizon))
}
