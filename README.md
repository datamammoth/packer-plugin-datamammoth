# packer-plugin-datamammoth

Packer builder plugin for creating server images on the DataMammoth platform.

## Overview

This plugin provides a `datamammoth-server` builder that:
1. Provisions a temporary VPS via the DataMammoth API v2
2. Runs your Packer provisioners (shell, Ansible, Chef, etc.)
3. Creates a snapshot of the configured server
4. Terminates the temporary server
5. Returns the snapshot ID as a build artifact

## Usage

```hcl
packer {
  required_plugins {
    datamammoth = {
      version = ">= 0.1.0"
      source  = "github.com/datamammoth/datamammoth"
    }
  }
}

source "datamammoth-server" "ubuntu" {
  api_key       = var.dm_api_key
  product_id    = "prod_vps_4core"
  region        = "EU"
  image_id      = "ubuntu-24.04"
  snapshot_name = "golden-ubuntu-{{timestamp}}"
  ssh_username  = "root"
}

build {
  sources = ["source.datamammoth-server.ubuntu"]

  provisioner "shell" {
    inline = [
      "apt-get update",
      "apt-get install -y nginx",
    ]
  }
}
```

## Configuration Reference

| Key | Required | Description |
|---|---|---|
| `api_key` | Yes | DataMammoth API key (`dm_live_...`) |
| `base_url` | No | Override API URL (self-hosted) |
| `product_id` | Yes | Hosting product ID for the temp server |
| `region` | Yes | Datacenter region |
| `image_id` | Yes | Base OS image ID |
| `snapshot_name` | No | Name for the resulting snapshot |
| `ssh_username` | No | SSH user (default: `root`) |

## Building

```bash
go build -o packer-plugin-datamammoth
```

## License

MIT
