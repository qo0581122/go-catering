package initialize

import "catering/pkg"

func init() {
	_ = pkg.RegisterRule("PageVerify",
		pkg.Rules{
			"Page":     {pkg.NotEmpty()},
			"PageSize": {pkg.NotEmpty()},
		},
	)
	_ = pkg.RegisterRule("IdVerify",
		pkg.Rules{
			"Id": {pkg.NotEmpty()},
		},
	)
	_ = pkg.RegisterRule("AuthorityIdVerify",
		pkg.Rules{
			"AuthorityId": {pkg.NotEmpty()},
		},
	)
}
