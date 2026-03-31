resource "null_resource" "deploy" {
  provisioner "local-exec" {
    command = "faas-cli up -f hello.yml"
  }
}