package latest

import (
	"fmt"
	"testing"
)

func TestLatestVersionSimple(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.1.0"] = 100
	versions["1.2.0"] = 60
	versions["1.2.1"] = 10

	if err := versionSimpleHelper(versions, 60.0, "1.1.0"); err != nil {
		t.Fatal(err)
	}

	if err := versionSimpleHelper(versions, 30.0, "1.2.0"); err != nil {
		t.Fatal(err)
	}

	if err := versionSimpleHelper(versions, 5.0, "1.2.1"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestVersion(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.1.0"] = 100
	versions["1.2.0"] = 50
	versions["1.2.1"] = 10
	versions["invalid.version"] = 1
	versions["1.0.0-alpha"] = 3

	if err := versionHelper(versions, 30.0, 15.0, "1.2.0"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestStringVersion(t *testing.T) {
	versions := make(map[string]uint32)
	versions["a"] = 100
	versions["b"] = 50
	versions["c"] = 50

	if err := versionHelper(versions, 30.0, 15.0, "a"); err != nil {
		t.Fatal(err)
	}

	versions["b"] = 100

	if err := versionHelper(versions, 30.0, 15.0, "b"); err != nil {
		t.Fatal(err)
	}

	versions["c"] = 100

	if err := versionHelper(versions, 30.0, 15.0, "c"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestUnicodeVersion(t *testing.T) {
	versions := make(map[string]uint32)
	versions["вода"] = 100
	versions["крот"] = 90
	versions["хвоя"] = 30

	if err := versionHelper(versions, 30.0, 15.0, "крот"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestStrangeVersion(t *testing.T) {
	versions := make(map[string]uint32)
	versions["---"] = 100
	versions["."] = 100
	versions[". "] = 100
	versions[""] = 100

	if err := versionHelper(versions, 0.0, 0.0, "---"); err != nil {
		t.Fatal(err)
	}

	versions["💁"] = 100
	if err := versionHelper(versions, 0.0, 0.0, "💁"); err != nil {
		t.Fatal(err)
	}

	versions["1. 0.0"] = 100
	if err := versionHelper(versions, 0.0, 0.0, "1. 0.0"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestMixedVersion(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1"] = 100
	versions["some"] = 50

	if err := versionHelper(versions, 30.0, 15.0, "1"); err != nil {
		t.Fatal(err)
	}

	versions["0.5.1"] = 100
	if err := versionHelper(versions, 30.0, 15.0, "1"); err != nil {
		t.Fatal(err)
	}

	versions["1"] = 0
	if err := versionHelper(versions, 30.0, 15.0, "0.5.1"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestNotSemver(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.0.0.123.9"] = 100
	versions["1.0.0.345.8"] = 50
	versions["1.0.0.678.7"] = 10

	if err := versionHelper(versions, 30.0, 15.0, "1.0.0.345.8"); err != nil {
		t.Fatal(err)
	}
}

func TestLatestDifficult(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1"] = 100
	versions["1.2"] = 50
	versions["1.2.3"] = 25
	versions["1.2.3.4"] = 12

	if err := versionHelper(versions, 0, 0, "1.2.3.4"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 5.0, 1.0, "1.2.3.4"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 10.0, 3.0, "1.2.3"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 20.0, 5.0, "1.2"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 30.0, 20.0, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestThreshold(t *testing.T) {
	versions := map[string]uint32{
		"1": 10, "2": 10, "3": 10, "4": 10, "5": 10, "6": 10, "7": 10, "8": 10, "9": 10, "10": 10, "11": 10,
	}

	if err := versionHelper(versions, 50.0, 25.0, "11"); err != nil {
		t.Fatal(err)
	}

	versions["11"] = 1

	if err := versionHelper(versions, 50.0, 25.0, "11"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 10.0, 5.0, "10"); err != nil {
		t.Fatal(err)
	}

	versions["10"] = 0

	if err := versionHelper(versions, 10.0, 5.0, "9"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 1.0, 0.5, "11"); err != nil {
		t.Fatal(err)
	}

	versions["4"] = 20

	if err := versionHelper(versions, 10.0, 5.0, "4"); err != nil {
		t.Fatal(err)
	}

	if err := versionHelper(versions, 0.0, 0.0, "11"); err != nil {
		t.Fatal(err)
	}
}

func TestFallback(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1"] = 45
	versions["2"] = 35
	versions["3"] = 20

	// low fallback
	if err := versionHelper(versions, 50.0, 15.0, "3"); err != nil {
		t.Fatal(err)
	}

	// middle fallback
	if err := versionHelper(versions, 50.0, 30.0, "2"); err != nil {
		t.Fatal(err)
	}

	// big fallback
	if err := versionHelper(versions, 50.0, 45.0, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestSingleVersions(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.0.0"] = 100

	if err := versionHelper(versions, 30.0, 15.0, "1.0.0"); err != nil {
		t.Fatal(err)
	}
}

func TestEmptyVersions(t *testing.T) {
	_, err := FindVersionSafe(make(map[string]uint32), 70.0, 30.0)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	fmt.Printf("Passed with error: %v\n", err)
}

func TestWrongThreshold(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.0.0"] = 100

	_, err := FindVersionSafe(versions, 99.9, 30.0)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	fmt.Printf("Passed with error: %v\n", err)
}

func TestWrongFallback(t *testing.T) {
	versions := make(map[string]uint32)
	versions["1.0.0"] = 100

	_, err := FindVersionSafe(versions, 50.0, 46.0)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	fmt.Printf("Passed with error: %v\n", err)
}

func TestCompareEquals(t *testing.T) {
	result := CompareVersions("0", "0")

	if result != 0 {
		t.Error("Got not expected value")
	}
}

func versionHelper(versions map[string]uint32, threshold, fallback float64, expected string) error {
	result, err := FindVersionSafe(versions, threshold, fallback)
	if err != nil {
		return fmt.Errorf("Expected no error, got %s: %v", result, err)
	}
	if result != expected {
		return fmt.Errorf("Expected result %s, got %s", expected, result)
	}

	fmt.Printf("result: %s\n", result)
	return nil
}

func versionSimpleHelper(versions map[string]uint32, threshold float64, expected string) error {
	result, err := FindVersion(versions, threshold)
	if err != nil {
		return fmt.Errorf("Expected no error, got %s: %v", result, err)
	}
	if result != expected {
		return fmt.Errorf("Expected result %s, got %s", expected, result)
	}

	fmt.Printf("result: %s\n", result)
	return nil
}
