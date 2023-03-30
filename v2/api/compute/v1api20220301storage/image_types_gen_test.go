// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220301storage

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

func Test_Image_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImage, ImageGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImage runs a test to see if a specific instance of Image round trips to JSON and back losslessly
func RunJSONSerializationTestForImage(subject Image) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image
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

// Generator of Image instances for property testing - lazily instantiated by ImageGenerator()
var imageGenerator gopter.Gen

// ImageGenerator returns a generator of Image instances for property testing.
func ImageGenerator() gopter.Gen {
	if imageGenerator != nil {
		return imageGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForImage(generators)
	imageGenerator = gen.Struct(reflect.TypeOf(Image{}), generators)

	return imageGenerator
}

// AddRelatedPropertyGeneratorsForImage is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImage(gens map[string]gopter.Gen) {
	gens["Spec"] = Image_SpecGenerator()
	gens["Status"] = Image_STATUSGenerator()
}

func Test_Image_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImage_Spec, Image_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImage_Spec runs a test to see if a specific instance of Image_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForImage_Spec(subject Image_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_Spec
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

// Generator of Image_Spec instances for property testing - lazily instantiated by Image_SpecGenerator()
var image_SpecGenerator gopter.Gen

// Image_SpecGenerator returns a generator of Image_Spec instances for property testing.
// We first initialize image_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Image_SpecGenerator() gopter.Gen {
	if image_SpecGenerator != nil {
		return image_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_Spec(generators)
	image_SpecGenerator = gen.Struct(reflect.TypeOf(Image_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_Spec(generators)
	AddRelatedPropertyGeneratorsForImage_Spec(generators)
	image_SpecGenerator = gen.Struct(reflect.TypeOf(Image_Spec{}), generators)

	return image_SpecGenerator
}

// AddIndependentPropertyGeneratorsForImage_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImage_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImage_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImage_Spec(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResourceGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfileGenerator())
}

func Test_Image_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImage_STATUS, Image_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImage_STATUS runs a test to see if a specific instance of Image_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImage_STATUS(subject Image_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_STATUS
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

// Generator of Image_STATUS instances for property testing - lazily instantiated by Image_STATUSGenerator()
var image_STATUSGenerator gopter.Gen

// Image_STATUSGenerator returns a generator of Image_STATUS instances for property testing.
// We first initialize image_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Image_STATUSGenerator() gopter.Gen {
	if image_STATUSGenerator != nil {
		return image_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_STATUS(generators)
	image_STATUSGenerator = gen.Struct(reflect.TypeOf(Image_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_STATUS(generators)
	AddRelatedPropertyGeneratorsForImage_STATUS(generators)
	image_STATUSGenerator = gen.Struct(reflect.TypeOf(Image_STATUS{}), generators)

	return image_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImage_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImage_STATUS(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImage_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImage_STATUS(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfile_STATUSGenerator())
}

func Test_ExtendedLocation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation, ExtendedLocationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation runs a test to see if a specific instance of ExtendedLocation round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation(subject ExtendedLocation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation
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

// Generator of ExtendedLocation instances for property testing - lazily instantiated by ExtendedLocationGenerator()
var extendedLocationGenerator gopter.Gen

// ExtendedLocationGenerator returns a generator of ExtendedLocation instances for property testing.
func ExtendedLocationGenerator() gopter.Gen {
	if extendedLocationGenerator != nil {
		return extendedLocationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation(generators)
	extendedLocationGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation{}), generators)

	return extendedLocationGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExtendedLocation_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation_STATUS, ExtendedLocation_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation_STATUS runs a test to see if a specific instance of ExtendedLocation_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation_STATUS(subject ExtendedLocation_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_STATUS
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

// Generator of ExtendedLocation_STATUS instances for property testing - lazily instantiated by
// ExtendedLocation_STATUSGenerator()
var extendedLocation_STATUSGenerator gopter.Gen

// ExtendedLocation_STATUSGenerator returns a generator of ExtendedLocation_STATUS instances for property testing.
func ExtendedLocation_STATUSGenerator() gopter.Gen {
	if extendedLocation_STATUSGenerator != nil {
		return extendedLocation_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation_STATUS(generators)
	extendedLocation_STATUSGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_STATUS{}), generators)

	return extendedLocation_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageStorageProfile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfile, ImageStorageProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfile runs a test to see if a specific instance of ImageStorageProfile round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfile(subject ImageStorageProfile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile
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

// Generator of ImageStorageProfile instances for property testing - lazily instantiated by
// ImageStorageProfileGenerator()
var imageStorageProfileGenerator gopter.Gen

// ImageStorageProfileGenerator returns a generator of ImageStorageProfile instances for property testing.
// We first initialize imageStorageProfileGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfileGenerator() gopter.Gen {
	if imageStorageProfileGenerator != nil {
		return imageStorageProfileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile(generators)
	imageStorageProfileGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfile(generators)
	imageStorageProfileGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile{}), generators)

	return imageStorageProfileGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfile is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfile(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfile is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfile(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDiskGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDiskGenerator())
}

func Test_ImageStorageProfile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfile_STATUS, ImageStorageProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfile_STATUS runs a test to see if a specific instance of ImageStorageProfile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfile_STATUS(subject ImageStorageProfile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile_STATUS
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

// Generator of ImageStorageProfile_STATUS instances for property testing - lazily instantiated by
// ImageStorageProfile_STATUSGenerator()
var imageStorageProfile_STATUSGenerator gopter.Gen

// ImageStorageProfile_STATUSGenerator returns a generator of ImageStorageProfile_STATUS instances for property testing.
// We first initialize imageStorageProfile_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfile_STATUSGenerator() gopter.Gen {
	if imageStorageProfile_STATUSGenerator != nil {
		return imageStorageProfile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS(generators)
	imageStorageProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS(generators)
	imageStorageProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_STATUS{}), generators)

	return imageStorageProfile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfile_STATUS(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfile_STATUS(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDisk_STATUSGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDisk_STATUSGenerator())
}

func Test_SubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource runs a test to see if a specific instance of SubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource(subject SubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource
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

// Generator of SubResource instances for property testing - lazily instantiated by SubResourceGenerator()
var subResourceGenerator gopter.Gen

// SubResourceGenerator returns a generator of SubResource instances for property testing.
func SubResourceGenerator() gopter.Gen {
	if subResourceGenerator != nil {
		return subResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

func Test_SubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS runs a test to see if a specific instance of SubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS
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

// Generator of SubResource_STATUS instances for property testing - lazily instantiated by SubResource_STATUSGenerator()
var subResource_STATUSGenerator gopter.Gen

// SubResource_STATUSGenerator returns a generator of SubResource_STATUS instances for property testing.
func SubResource_STATUSGenerator() gopter.Gen {
	if subResource_STATUSGenerator != nil {
		return subResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS(generators)
	subResource_STATUSGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS{}), generators)

	return subResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageDataDisk_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDisk, ImageDataDiskGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDisk runs a test to see if a specific instance of ImageDataDisk round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDisk(subject ImageDataDisk) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk
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

// Generator of ImageDataDisk instances for property testing - lazily instantiated by ImageDataDiskGenerator()
var imageDataDiskGenerator gopter.Gen

// ImageDataDiskGenerator returns a generator of ImageDataDisk instances for property testing.
// We first initialize imageDataDiskGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDiskGenerator() gopter.Gen {
	if imageDataDiskGenerator != nil {
		return imageDataDiskGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk(generators)
	imageDataDiskGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk(generators)
	AddRelatedPropertyGeneratorsForImageDataDisk(generators)
	imageDataDiskGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk{}), generators)

	return imageDataDiskGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDisk is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDisk(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageDataDisk is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDisk(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceGenerator())
}

func Test_ImageDataDisk_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDisk_STATUS, ImageDataDisk_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDisk_STATUS runs a test to see if a specific instance of ImageDataDisk_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDisk_STATUS(subject ImageDataDisk_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk_STATUS
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

// Generator of ImageDataDisk_STATUS instances for property testing - lazily instantiated by
// ImageDataDisk_STATUSGenerator()
var imageDataDisk_STATUSGenerator gopter.Gen

// ImageDataDisk_STATUSGenerator returns a generator of ImageDataDisk_STATUS instances for property testing.
// We first initialize imageDataDisk_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDisk_STATUSGenerator() gopter.Gen {
	if imageDataDisk_STATUSGenerator != nil {
		return imageDataDisk_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_STATUS(generators)
	imageDataDisk_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_STATUS(generators)
	AddRelatedPropertyGeneratorsForImageDataDisk_STATUS(generators)
	imageDataDisk_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_STATUS{}), generators)

	return imageDataDisk_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDisk_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDisk_STATUS(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageDataDisk_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDisk_STATUS(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_ImageOSDisk_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDisk, ImageOSDiskGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDisk runs a test to see if a specific instance of ImageOSDisk round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDisk(subject ImageOSDisk) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk
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

// Generator of ImageOSDisk instances for property testing - lazily instantiated by ImageOSDiskGenerator()
var imageOSDiskGenerator gopter.Gen

// ImageOSDiskGenerator returns a generator of ImageOSDisk instances for property testing.
// We first initialize imageOSDiskGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDiskGenerator() gopter.Gen {
	if imageOSDiskGenerator != nil {
		return imageOSDiskGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk(generators)
	imageOSDiskGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk(generators)
	AddRelatedPropertyGeneratorsForImageOSDisk(generators)
	imageOSDiskGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk{}), generators)

	return imageOSDiskGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDisk is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDisk(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageOSDisk is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDisk(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResourceGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceGenerator())
}

func Test_ImageOSDisk_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDisk_STATUS, ImageOSDisk_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDisk_STATUS runs a test to see if a specific instance of ImageOSDisk_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDisk_STATUS(subject ImageOSDisk_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk_STATUS
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

// Generator of ImageOSDisk_STATUS instances for property testing - lazily instantiated by ImageOSDisk_STATUSGenerator()
var imageOSDisk_STATUSGenerator gopter.Gen

// ImageOSDisk_STATUSGenerator returns a generator of ImageOSDisk_STATUS instances for property testing.
// We first initialize imageOSDisk_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDisk_STATUSGenerator() gopter.Gen {
	if imageOSDisk_STATUSGenerator != nil {
		return imageOSDisk_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_STATUS(generators)
	imageOSDisk_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_STATUS(generators)
	AddRelatedPropertyGeneratorsForImageOSDisk_STATUS(generators)
	imageOSDisk_STATUSGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_STATUS{}), generators)

	return imageOSDisk_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDisk_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDisk_STATUS(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountType"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImageOSDisk_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDisk_STATUS(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_STATUSGenerator())
}
