package matcher

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	NoLicense = "no licenes"
)

func TestInitLicenseMatcherWithNoLicensesDB(t *testing.T) {
	err := InitLicenseMatcher("unknown folder")
	if err == nil {
		t.Fatal(err)
	}
}

func TestInitLicenseMatcherWithLicensesDB(t *testing.T) {
	err := InitLicenseMatcher(".")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMatchLicenseTxtNoLicenseFile(t *testing.T) {
	err := InitLicenseMatcher(".")
	if err != nil {
		t.Fatal(err)
	}
	lics := MatchLicenseTxt(NoLicense)
	assert.True(t, len(lics) == 0 || lics[0] == "Unknown")

}

func TestMatchLicenseTxtWithClassifier(t *testing.T) {
	err := InitLicenseMatcher(".")
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open("./fixtures/Multi_LICENSE.txt")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	lics := MatchLicenseTxt(string(data))
	assert.True(t, len(lics) == 6)
}

func TestMatchLicenseTxtWithDetector(t *testing.T) {
	err := InitLicenseMatcher(".")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ReadLicense("./fixtures/Partial_LICENSE.txt")
	if err != nil {
		t.Error(err)
	}
	lics := MatchLicenseTxt(data)
	assert.True(t, len(lics) == 1)
}

func TestMatchLicenseTxtWithPom(t *testing.T) {
	err := InitLicenseMatcher(".")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ReadLicense("./fixtures/PomWithLicenseAsComment.xml")
	if err != nil {
		t.Error(err)
	}
	lic := GetPomCommentLicense(data)
	assert.True(t, lic == "Apache-2.0")
}

func ReadLicense(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
