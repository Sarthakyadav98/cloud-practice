provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_pod" "example" {
  metadata {
    name = "nginx-pod"
  }

  spec {
    container {
      name  = "nginx"
      image = "nginx"
    }
  }
}