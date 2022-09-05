# The name of your project. A project typically maps 1:1 to a VCS repository.
# This name must be unique for your Waypoint server. If you're running in
# local mode, this must be unique to your machine.
project = "api-goLang"

# Labels can be specified for organizational purposes.
# labels = { "foo" = "bar" }

# An application to deploy.
app "web" {
    # Build specifies how an application should be deployed. In this case,
    # we'll build using a Dockerfile and keeping it in a local registry.
    build {
        use "pack" {}
        
        #Uncomment below to use a remote docker registry to push your built images.
       
        registry {
        use "docker" {
           image = "symonnascimento/apigolangwaypoint"
            tag  = "latest"
          }
        }
    }

    # Deploy to Docker
    deploy {
        use "docker" {
             service_port=8090
        command = ["-./api", ]
        }
    }
}
