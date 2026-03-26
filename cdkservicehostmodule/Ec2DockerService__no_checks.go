//go:build no_runtime_type_checking

package cdkservicehostmodule

// Building without runtime type checking enabled, so all the below just return nil

func validateEc2DockerService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEc2DockerServiceParameters(scope constructs.Construct, id *string, props *Ec2DockerServiceProps) error {
	return nil
}

