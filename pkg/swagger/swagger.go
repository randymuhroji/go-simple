package swagger

import "kumparan/pkg/swagger/docs"

// @termsOfService http://swagger.io/terms/

func Init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Kumparan Article Management"
	docs.SwaggerInfo.Description = "KUMPARAN"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
