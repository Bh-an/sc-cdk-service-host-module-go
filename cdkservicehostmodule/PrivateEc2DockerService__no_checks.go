//go:build no_runtime_type_checking

package cdkservicehostmodule

// Building without runtime type checking enabled, so all the below just return nil

func validatePrivateEc2DockerService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPrivateEc2DockerServiceParameters(scope constructs.Construct, id *string, props *PrivateEc2DockerServiceProps) error {
	return nil
}

