# Aviator: VM Provisioning Kubernetes Controller

A custom Kubernetes controller for automating OS provisioning in cloud environments, with integrated support for the Naver Cloud Platform (NCP) through the Aviator service SDK.
The [Aviator-service](https://github.com/cloud-club/Aviator-service) repository contains the core SDK implementation for cloud infrastructure automation.

## Overview

This Kubernetes controller automates the provisioning and management of virtual machines through custom resources. It integrates with NCP APIs to handle VM lifecycle operations including creation, updates, deletion, and status monitoring.

## Features

- **Custom Resource Definitions (CRDs)**
  - Group: `vm.cloudclub.io`
  - Version: `v1`
  - Resources:
    - `Provision`: Manages VM provisioning lifecycle
    - `Plan`: Handles VM planning and configuration

- **Integrated NCP Operations**
  - Server instance creation and management
  - Network interface configuration
  - Block storage management
  - Access control group handling
  - Status monitoring and updates

## Architecture

### Component Structure

1. **Controller Core**
   - Custom Resource Definitions
   - Reconciliation Logic
   - Status Management

2. **Aviator Service SDK Integration**
   - ServerService: Handles VM operations
   - KeyService: Manages authentication
   - Error handling and status checking
   - NCP API integration

### Server Operations Interface
```go
type ServerInterface interface {
    List(url string, request *types.ListServerRequest) (*types.ListServerResponse, error)
    Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error)
    Update(url string, request *types.UpdateServerRequest) (*types.UpdateServerResponse, error)
    Start(url string, request *types.StartServerRequest) (*types.StartServerResponse, error)
    Stop(url string, request *types.StopServerRequest) (*types.StopServerResponse, error)
    Delete(url string, request *types.DeleteServerRequest) (*types.DeleteServerResponse, error)
}
```

### Custom Resources

#### Provision Resource Example
```yaml
apiVersion: vm.cloudclub.io/v1
kind: Provision
spec:
  regionCode: KR
  server:
    imageProductCode: UBUNTU20
    productCode: SPSVRSTAND000056
  vpcNo: "vpc-123"
  subnetNo: "subnet-456"
  phase: Create
```

## Development

### Prerequisites
- Go 1.16+
- Kubernetes cluster (1.19+)
- Docker or compatible container runtime
- NCP Account and API Credentials
- Aviator Service SDK

### Setting Up Development Environment

1. Clone the repositories:
```bash
# Main controller
git clone https://github.com/your-org/vm-controller
# Aviator Service SDK
git clone https://github.com/cloud-club/Aviator-service
```

2. Configure NCP credentials:
```bash
export NCP_ACCESS_KEY="your-access-key"
export NCP_SECRET_KEY="your-secret-key"
```

### Build and Test

```bash
# Build the controller
make build

# Run tests
make test

# Run linter
make lint
```

## CI/CD Pipeline

The project uses GitHub Actions for CI/CD with the following stages:

1. **Build & Test**
   - Code compilation
   - Unit tests
   - Integration tests
   - Code coverage reporting

2. **Quality Gates**
   - Linting
   - Security scanning
   - Dependency checking

3. **Deployment**
   - Container image building
   - Image scanning
   - Registry push
   - Kubernetes deployment

### Pipeline Configuration

```yaml
# Example workflow stages
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Build
        run: make build
      - name: Test
        run: make test
```

## Usage

### 1. Installation

```bash
# Install CRDs
make install

# Deploy controller
make deploy IMG=<your-registry>/controller:tag
```

### 2. Create a VM Instance

```yaml
apiVersion: vm.cloudclub.io/v1
kind: Provision
metadata:
  name: example-vm
spec:
  phase: Create
  server:
    imageProductCode: "UBUNTU20"
    productCode: "SPSVRSTAND000056"
  vpcNo: "vpc-123"
  subnetNo: "subnet-456"
```

### 3. Monitor Status

```bash
kubectl get provisions
kubectl describe provision example-vm
```

## Contributing

1. Fork both repositories (controller and [Aviator-service](https://github.com/cloud-club/Aviator-service))
2. Create feature branches
3. Run tests and linting
4. Submit Pull Requests to both repos if needed
5. Ensure CI pipeline passes