package utils

const (
	//LicenseTxt const
	LicenseTxt = "license.txt"
	//LicenseDot const
	LicenseDot = "license."
	//LicenseMd const
	LicenseMd = "license.md"
	//License const
	License = "license"
	//Bsdl const
	Bsdl = "bsdl"
	//Copying const
	Copying = "copying"
	//COPYINGv3 const
	COPYINGv3 = "COPYINGv3"
	//CopyingDot const
	CopyingDot = "copying."
	//CopyingDash const
	CopyingDash = "copying-"
	//Legal const
	Legal = "legal"
	//Readme const
	Readme = "readme"
	//Copyright const
	Copyright = "copyright"
	//Ftl const
	Ftl = "ftl.txt"
	//GPLv2 const
	GPLv2 = "gplv2.TXT"
	//gpl20 const
	gpl20 = "gpl-2.0"
	//MITtxt const
	MITtxt = "mit.txt"
	//LisenseRst const
	LisenseRst = "lisense.rst"
	//LisenceHTML const
	LisenceHTML = "license.html"
	//Licenses2 const
	Licenses2 = "LICENSE-2.0.txt"
	//LicenseUnknown const
	LicenseUnknown = "Unknown"
	//Licenses const
	Licenses = "/licenses/"
	//EmptyString const
	EmptyString = ""
)

//GetLicensesFiles get license file type
func GetLicensesFiles() []string {
	return []string{LicenseTxt, LicenseMd, License, Copying, Legal, COPYINGv3,
		Readme, Ftl, GPLv2, gpl20, Bsdl, Copyright, MITtxt, LisenseRst,
		LisenceHTML, Licenses2}
}

//GetLicensesFilesPrefix get license by prefix
func GetLicensesFilesPrefix() []string {
	return []string{LicenseDot, CopyingDot, Copyright, CopyingDash}
}
