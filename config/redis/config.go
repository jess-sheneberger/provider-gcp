package redis

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_redis_instance", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["host"].(string); ok {
				conn["host"] = []byte(a)
			}
			if certs, ok := attr["serverCaCerts"].([]map[string]any); ok {
				// need serialize serverCaCerts to []byte
				serverCaCerts["ca"] = certs
			}
			return conn, nil
		}
	})
}
