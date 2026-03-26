//go:build no_runtime_type_checking

package cdkservicehostmodule

// Building without runtime type checking enabled, so all the below just return nil

func validatePublicServiceHost_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPublicServiceHostParameters(scope constructs.Construct, id *string, props *PublicServiceHostProps) error {
	return nil
}

