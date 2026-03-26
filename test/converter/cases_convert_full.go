package converter

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
)

func BuildConvertFullPrimTestCase() ConvertFullTestCase {
	return ConvertFullTestCase{
		Params: ConvertFullParams{
			CName: "int",
		},
		Want: ConvertFullWant{
			FName: "int",
		},
	}
}

func BuildConvertFullSamePackageTestCase() ConvertFullTestCase {
	var (
		pkgPath = "github.com/mus-format/mus-gen-go/test/converter"
		pkg     = "converter"
		gops    = genopts.New()
	)
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath(pkgPath),
		genopts.WithPackage(pkg),
	}, &gops)

	return ConvertFullTestCase{
		Setup: ConvertFullSetup{
			Gops: gops,
		},
		Params: ConvertFullParams{
			CName: "github.com/mus-format/mus-gen-go/test/converter/converter.Type",
		},
		Want: ConvertFullWant{
			FName: "converter.Type",
		},
	}
}

func BuildConvertFullOtherPackageTestCase() ConvertFullTestCase {
	var (
		pkgPath = "github.com/mus-format/mus-gen-go/test/converter"
		pkg     = "converter"
		gops    = genopts.New()
	)
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath(pkgPath),
		genopts.WithPackage(pkg),
	}, &gops)

	return ConvertFullTestCase{
		Setup: ConvertFullSetup{
			Gops: gops,
		},
		Params: ConvertFullParams{
			CName: "github.com/mus-format/mus-gen-go/other/other.Type",
		},
		Want: ConvertFullWant{
			FName: "other.Type",
		},
	}
}

func BuildConvertFullAliasedPackageTestCase() ConvertFullTestCase {
	var (
		pkgPath   = "github.com/mus-format/mus-gen-go/test/converter"
		pkg       = "converter"
		otherPath = "github.com/mus-format/mus-gen-go/other"
		alias     = "alias"
		gops      = genopts.New()
	)
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath(pkgPath),
		genopts.WithPackage(pkg),
		genopts.WithImportAlias(otherPath, alias),
	}, &gops)

	return ConvertFullTestCase{
		Setup: ConvertFullSetup{
			Gops: gops,
		},
		Params: ConvertFullParams{
			CName: "github.com/mus-format/mus-gen-go/other/other.Type",
		},
		Want: ConvertFullWant{
			FName: "alias.Type",
		},
	}
}

func BuildConvertFullSliceTestCase() ConvertFullTestCase {
	return ConvertFullTestCase{
		Params: ConvertFullParams{
			CName: "[]int",
		},
		Want: ConvertFullWant{
			FName: "[]int",
		},
	}
}

func BuildConvertFullMapTestCase() ConvertFullTestCase {
	return ConvertFullTestCase{
		Params: ConvertFullParams{
			CName: "map[string]int",
		},
		Want: ConvertFullWant{
			FName: "map[string]int",
		},
	}
}

func BuildConvertFullTypeParamTestCase() ConvertFullTestCase {
	var (
		pkgPath   = "github.com/mus-format/mus-gen-go/test/converter"
		pkg       = "converter"
		otherPath = "github.com/mus-format/mus-gen-go/other"
		alias     = "alias"
		gops      = genopts.New()
	)
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath(pkgPath),
		genopts.WithPackage(pkg),
		genopts.WithImportAlias(otherPath, alias),
	}, &gops)

	return ConvertFullTestCase{
		Setup: ConvertFullSetup{
			Gops: gops,
		},
		Params: ConvertFullParams{
			CName: "github.com/mus-format/mus-gen-go/test/converter/converter.Type[github.com/mus-format/mus-gen-go/other.Type]",
		},
		Want: ConvertFullWant{
			FName: "converter.Type[alias.Type]",
		},
	}
}

func BuildConvertFullComplexTestCase() ConvertFullTestCase {
	var (
		pkgPath   = "github.com/mus-format/mus-gen-go/test/converter"
		pkg       = "converter"
		firstPath = "github.com/mus-format/mus-gen-go/other/first"
		thirdPath = "github.com/mus-format/mus-gen-go/other/third"
		first     = "first_alias"
		third     = "third_alias"
		gops      = genopts.New()
	)
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath(pkgPath),
		genopts.WithPackage(pkg),
		genopts.WithImportAlias(firstPath, first),
		genopts.WithImportAlias(thirdPath, third),
	}, &gops)

	return ConvertFullTestCase{
		Setup: ConvertFullSetup{
			Gops: gops,
		},
		Params: ConvertFullParams{
			CName: "[]map[github.com/mus-format/mus-gen-go/other/third/third.Type]github.com/mus-format/mus-gen-go/test/converter/converter.Type[github.com/mus-format/mus-gen-go/other/first.Type, github.com/mus-format/mus-gen-go/other/second.Type]",
		},
		Want: ConvertFullWant{
			FName: "[]map[third_alias.Type]converter.Type[first_alias.Type,second.Type]",
		},
	}
}
