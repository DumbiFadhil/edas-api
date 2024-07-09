package main

import (
    "DumbiFadhil/edas-api/config"
    "DumbiFadhil/edas-api/routes"
)

func main() {
    router := config.SetupRouter()
    routes.SetupRoutes(router)
    router.Run(":8080")
}

