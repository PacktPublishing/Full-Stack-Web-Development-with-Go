#script to pull chapter12 image and run it locally
#it also store the image locally
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }

  required_version = ">= 1.2.0"
}


data "docker_registry_image" "github" {
  name = "ghcr.io/nanikjava/golangci/chapter12:latest"
}

resource "docker_image" "embed" {
  keep_locally = true
  name         = "${data.docker_registry_image.github.name}"
}

resource "docker_container" "embed" {
  image = "ghcr.io/nanikjava/golangci/chapter12"
  name  = "chapter12"
  ports {
    internal = 3333
    external = 3333
  }
}

