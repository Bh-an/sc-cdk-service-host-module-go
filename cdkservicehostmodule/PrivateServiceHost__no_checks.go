//go:build no_runtime_type_checking

package cdkservicehostmodule

// Building without runtime type checking enabled, so all the below just return nil

func validatePrivateServiceHost_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPrivateServiceHostParameters(scope constructs.Construct, id *string, props *PrivateServiceHostProps) error {
	return nil
}

