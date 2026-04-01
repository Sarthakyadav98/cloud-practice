# Lab 2 — Virtual Machines, MPI & AWS

---

## Objective

To create Ubuntu-based virtual machines, execute an MPI Hello World program across multiple VMs, and deploy an Ubuntu VM on AWS.

---

# TASK 1 — Create VMs & Run MPI

---

## Step 1: Create Virtual Machines

* Install VirtualBox / VMware
* Create **2 Ubuntu VMs**
* Configure network:
    Go to VM Settings → Network
    Adapter 1:
    Enable Network Adapter
    Attached to: Host-only Adapter

    Do this for BOTH VMs.

  * Bridged / Host-only network

---

## Step 2: Install MPI (on BOTH VMs)

```bash
sudo apt update
sudo apt install -y openmpi-bin openmpi-common libopenmpi-dev
```

---

## Step 3: Verify MPI

```bash
mpirun --version
```

---

## Step 4: Setup SSH

```bash
sudo apt install -y openssh-server
sudo systemctl enable ssh
sudo systemctl start ssh
ssh localhost
```

(Optional: passwordless SSH for multi-node execution)

---

## Step 5: Write MPI Program

Create file: `hello.c`

```c
#include <mpi.h>
#include <stdio.h>

int main(int argc, char** argv) {
    MPI_Init(&argc, &argv);

    int rank, size;
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    printf("Hello from process %d of %d\n", rank, size);

    MPI_Finalize();
    return 0;
}
```

---

## Step 6: Compile

```bash
mpicc hello.c -o hello
```

---

## Step 7: Run Program

```bash
mpirun -np 2 ./hello
```

---

## Expected Output

```text
Hello from process 0 of 2
Hello from process 1 of 2
```



# TASK 2 — AWS Ubuntu VM

---

## Step 1: Login to AWS

* Go to AWS Console
* Open EC2 Dashboard

---

## Step 2: Launch Instance

* Choose Ubuntu AMI
* Instance type: `t2.micro`
* Create / select key pair
* Enable SSH (port 22)

---

## Step 3: Connect to Instance

```bash
ssh -i key.pem ubuntu@<public-ip>
```

---

## Fix Permission Error

```bash
chmod 400 key.pem
```

---

## Expected Output

* Ubuntu terminal opens via SSH
* Instance is accessible

---

# Key Concepts

* VM → Virtual computer
* MPI → Parallel processing
* AWS EC2 → Cloud VM

---

# Quick Revision

```text
VM → Install MPI → Write code → mpicc → mpirun
AWS → Launch EC2 → SSH connect
```
