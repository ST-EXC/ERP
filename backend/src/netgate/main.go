package main

import "EDA.Project.ERP.backend.netgate/routers"

func main() {
	r := routers.Router()
	r.Run(":9999")
}
