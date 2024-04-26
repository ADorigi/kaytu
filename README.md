## Overview

The Kaytu CLI helps you save on cloud costs by finding the perfect server sizes. Kaytu analyzes historical usage and provides tailored recommendations, ensuring you only pay for the resources you **actually need**.

- **Ease of use**: Doesn't touch the workload or require configuration changes to be made
- **Historical Usage**: Analyzes the past seven days of usage from Cloud native monitoring (CloudWatch), including advanced AWS CloudWatch metrics (where available).
- **Optimize as needed**: Optimize based on Region, CPU, memory, network performance, storage, and licenses.
- **Secure** - no credentials to share; extracts required metrics from the client side
- **Open-core philosophy** Use without fear of lock-in. The CLI is open-sourced, and the Server side will be open-sourced soon.
- **Coming Soon**: Non-Interactive mode, Azure support, GPU Optimization, Credit utilization for Burst instances, and Observability data from Prometheus

## Getting Started

### 1. Install Kaytu CLI

**MacOS**
```shell
brew tap kaytu-io/cli-tap && brew install kaytu
```

**Windows w/Chocolatey**
```shell
choco install infracost
```

**Linux**
```shell
curl -fsSL https://raw.githubusercontent.com/kaytu-io/kaytu/main/scripts/install.sh | sh
```
***Binary Download***
Download and install Windows, MacOS, and Linux binaries manually from [releases](https://github.com/kaytu-io/kaytu/releases) 

### 2. Ensure you have AWS CLI

Ensure you are logged in to AWS CLI

### 3. Run

```shell
kaytu
```
