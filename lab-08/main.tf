provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

resource "kubernetes_pod" "example" {
  metadata {
    name = "nginx-pod"
    labels = {
      app = "nginx"
    }
  }

  spec {
    container {
      name  = "nginx"
      image = "nginx"

      port {
        container_port = 80
      }
    }
  }
}