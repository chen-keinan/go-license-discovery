package matcher

import (
	"github.com/chen-keinan/go-license-detector/licensedb"
	"github.com/chen-keinan/go-license-detector/licensedb/filer"
	"github.com/chen-keinan/go-license-discovery/utils"
	"github.com/jfrog/gofrog/lru"
	"github.com/jfrog/licenseclassifier/licenseclassifier"
	"strings"
	"time"
)

/**
 * @author chenk on 29/08/2017.
 */
var lc *licenseclassifier.License

var extractedLicenseCache *lru.Cache = lru.New(1000, lru.WithExpiry(168*time.Hour))

//InitLicenseMatcher read licenses db and instantiate the classifier
func InitLicenseMatcher(licensesFolder string) error {
	licenseclassifier.LicensesDir(licensesFolder + utils.Licenses)
	var err error
	lc, err = licenseclassifier.New(0.8)
	if err != nil {
		return err
	}
	return nil
}

//MatchLicenseTxt try to match license txt to known spdx license
// 1st try to match to license classifier and then to license detector
func MatchLicenseTxt(licenseTxt string) []string {
	licSha := utils.LicenseToSha256(licenseTxt)
	if val, ok := extractedLicenseCache.Get(licSha); ok {
		return val.([]string)
	}
	n := lc.MultipleMatch(licenseTxt, true)
	set := utils.NewSet()
	if len(n) > 0 {
		for _, m := range n {
			if m.Confidence > 0.8 {
				set.AddString(m.Name)
			}
		}
		if set.Size() > 0 {
			licenses := set.StringValues()
			extractedLicenseCache.Add(licSha, licenses)
			return licenses
		}
	}
	licenses := GetLicenseFromDetector(licenseTxt, licSha)
	if len(licenses) > 0 {
		return licenses
	}

	return []string{utils.LicenseUnknown}
}

//DiscoveryFiler license object
type DiscoveryFiler struct {
	LicenseContent string
}

//ReadFile read license file
func (df DiscoveryFiler) ReadFile(path string) (content []byte, err error) {
	return []byte(df.LicenseContent), nil
}

//ReadDir read license directory
func (df DiscoveryFiler) ReadDir(path string) ([]filer.File, error) {
	return []filer.File{{Name: "license.txt"}}, nil
}

//Close close license file
func (df DiscoveryFiler) Close() {

}

//PathsAreAlwaysSlash check if path has slash
func (df DiscoveryFiler) PathsAreAlwaysSlash() bool {
	return false
}

//GetLicenseFromDetector check license from detector
func GetLicenseFromDetector(licenseTxt string, licSha string) []string {
	set := utils.NewSet()
	if len(licSha) == 0 {
		licSha = utils.LicenseToSha256(licenseTxt)
	}
	if val, ok := extractedLicenseCache.Get(licSha); ok {
		return val.([]string)
	}
	df := DiscoveryFiler{LicenseContent: licenseTxt}
	licMap, err := licensedb.Detect(df)
	if err != nil {
		return set.StringValues()
	}
	var maxScore float32
	var lic string
	if len(licMap) > 0 {
		for key, value := range licMap {
			if maxScore < value.Confidence && value.Confidence > 0.8 {
				maxScore = value.Confidence
				lic = key
			}
		}
		if len(lic) > 0 {
			extractedLicenseCache.Add(licSha, []string{lic})
			set.AddString(lic)
		}
	}
	return set.StringValues()
}

//GetPomCommentLicense read licenses which presented as a comment in the pom file
func GetPomCommentLicense(pomTxt string) string {
	pomComments := utils.ReadPomComments(pomTxt)
	if len(pomComments) == 0 {
		return utils.LicenseUnknown
	}
	tLicense := strings.TrimSpace(pomComments)
	lic := GetLicenseFromDetector(tLicense, utils.EmptyString)
	if len(lic) > 0 {
		return lic[0]
	}

	return utils.LicenseUnknown
}
