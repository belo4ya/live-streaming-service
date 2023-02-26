package v1

import _ "embed"

//go:embed gateway.swagger.json
var OpenAPISpec []byte
