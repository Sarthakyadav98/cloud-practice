## Clear old configs


```bash
chmod +x install.sh
sudo ./install.sh
newgrp docker
docker --version
```


```bash
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd
```

## Restart

```bash
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl restart containerd
sudo systemctl restart docker
```