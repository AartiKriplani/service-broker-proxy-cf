package cf

import (
	"github.com/Peripli/service-manager/pkg/env"
	"github.com/cloudfoundry-community/go-cfenv"
	"os"
	"fmt"
)

// SetCFOverrides overrides some SM environment with values from CF's VCAP environment variables
func SetCFOverrides(env env.Environment) error {
	if _, exists := os.LookupEnv("VCAP_APPLICATION"); exists {
		cfEnv, err := cfenv.Current()
		if err != nil {
			return fmt.Errorf("could not load VCAP environment: %s", err)
		}

		env.Set("server.host", "https://"+cfEnv.ApplicationURIs[0])
		env.Set("server.port", cfEnv.Port)
		env.Set("cf.api", cfEnv.CFAPI)
	}
	return nil
}