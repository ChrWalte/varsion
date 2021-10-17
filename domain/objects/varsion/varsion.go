package varsion

import "fmt"

type Varsion struct {
	Prefix    string
	Major     uint64
	Minor     uint64
	Patch     uint64
	Suffix    string
	Delimiter string
}

func InitializeVarsion(prefix string, suffix string, delmiter string) Varsion {
	return Varsion{
		Prefix:    prefix,
		Major:     0,
		Minor:     1,
		Patch:     0,
		Suffix:    suffix,
		Delimiter: delmiter,
	}
}

func (v Varsion) IncrementMajor() Varsion {
	return Varsion{
		Prefix:    v.Prefix,
		Major:     v.Major + 1,
		Minor:     0,
		Patch:     0,
		Suffix:    v.Suffix,
		Delimiter: v.Delimiter,
	}
}

func (v Varsion) IncrementMinor() Varsion {
	return Varsion{
		Prefix:    v.Prefix,
		Major:     v.Major,
		Minor:     v.Minor + 1,
		Patch:     0,
		Suffix:    v.Suffix,
		Delimiter: v.Delimiter,
	}
}

func (v Varsion) IncrementPatch() Varsion {
	return Varsion{
		Prefix:    v.Prefix,
		Major:     v.Major,
		Minor:     v.Minor,
		Patch:     v.Patch + 1,
		Suffix:    v.Suffix,
		Delimiter: v.Delimiter,
	}
}

func (v Varsion) ToString() string {
	return fmt.Sprintf(
		"%s%v%s%v%s%v%s",
		v.Prefix,
		v.Major,
		v.Delimiter,
		v.Minor,
		v.Delimiter,
		v.Patch,
		v.Suffix)
}
