package converter

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
)

func BuildConvertRelSamePackageTestCase() ConvertRelTestCase {
	var gops = genopts.New()
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath("github.com/mus-format/mus-gen-go/test/converter"),
		genopts.WithPackage("converter"),
	}, &gops)

	return ConvertRelTestCase{
		Setup: ConvertRelSetup{
			Gops: gops,
		},
		Params: ConvertRelParams{
			FName: "converter.Type",
		},
		Want: ConvertRelWant{
			RName: "Type",
		},
	}
}

func BuildConvertRelOtherPackageTestCase() ConvertRelTestCase {
	var gops = genopts.New()
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath("github.com/mus-format/mus-gen-go/test/converter"),
		genopts.WithPackage("converter"),
	}, &gops)

	return ConvertRelTestCase{
		Setup: ConvertRelSetup{
			Gops: gops,
		},
		Params: ConvertRelParams{
			FName: "other.Type",
		},
		Want: ConvertRelWant{
			RName: "other.Type",
		},
	}
}

func BuildConvertRelPrimTestCase() ConvertRelTestCase {
	return ConvertRelTestCase{
		Params: ConvertRelParams{
			FName: "int",
		},
		Want: ConvertRelWant{
			RName: "int",
		},
	}
}

func BuildConvertRelComplexTestCase() ConvertRelTestCase {
	var gops = genopts.New()
	genopts.Apply([]genopts.SetOption{
		genopts.WithPkgPath("github.com/mus-format/mus-gen-go/test/converter"),
		genopts.WithPackage("converter"),
	}, &gops)

	return ConvertRelTestCase{
		Setup: ConvertRelSetup{
			Gops: gops,
		},
		Params: ConvertRelParams{
			FName: "[]map[other.Type]converter.Type[other.Param1,converter.Param2]",
		},
		Want: ConvertRelWant{
			RName: "[]map[other.Type]Type[other.Param1,Param2]",
		},
	}
}
