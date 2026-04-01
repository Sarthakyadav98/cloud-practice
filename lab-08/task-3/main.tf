provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_pod" "mongodb" {
  metadata {
    name = "mongodb-pod"
    labels = {
      app = "mongodb"
    }
  }

  spec {
    container {
      name  = "mongodb"
      image = "mongo:6.0"

      port {
        container_port = 27017
      }
    }
  }
}

resource "kubernetes_service" "mongodb_service" {
  metadata {
    name = "mongodb-service"
  }

  spec {
    selector = {
      app = "mongodb"
    }

    port {
      port        = 27017
      target_port = 27017
    }

    type = "NodePort"
  }
}