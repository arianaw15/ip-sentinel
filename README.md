# IP Sentinel — IP Geolocation Country Whitelist Validator

## Overview

**IP Sentinel** is a Go-based HTTP API that determines whether a given IP address belongs to one of a customer-defined list of allowed countries. This service is intended to be deployed behind an API gateway, which captures the requester's IP and the target customer's country whitelist.

The service leverages the **MaxMind GeoLite2** IP-to-country dataset to perform accurate and efficient geolocation checks.

This project was developed as part of a technical assessment and showcases production-ready API development using Go. While it currently only accepts CIDR IPv4 addresses, it can easily be adapted to include other types.

---

## Features

- Accepts a request containing:
  - An CIDR IPv4 address
  - A list of country names (e.g. `"United States"`, `"Jordan"`, `"Peru"`)
- Returns a boolean indicator of whether the IP belongs to one of the allowed countries
- Uses the MaxMind GeoLite2 Country database
- Easily extendable to support future features

---

## Tech Stack

- **Language**: Go (Golang)
- **API**: HTTP + gRPC (optional gateway support)
- **IP Geolocation**: MaxMind GeoLite2
- **Testing**: Go’s native testing framework

---

## Setup

### 1. Prerequisites

- Go 1.24.4+
- MaxMind GeoLite2 Country database (`GeoIP2-Country-CSV_Example.zip`)
- [github.com/googleapis/googleapis annotation.proto file](https://github.com/googleapis/googleapis)

Register and download the GeoLite2 database from:  
[GeoIP2-Country-CSV_Example.zip](https://dev.maxmind.com/static/GeoIP2-Country-CSV_Example.zip)

---

### 2. Run Locally

```bash
go run ./cmd/main.go

## Developer Notes

As part of the submission, I’ve also attempted to integrate:

- A `Dockerfile` for containerized builds and deployment (in progress)
- A `Kubernetes` deployment and service YAML configuration (in progress)
- Initial `gRPC` support and `.proto` definitions (in progress)

## AI & Tooling Disclosure

During development, I used GitHub Copilot to assist with code completions and suggestions, particularly for routine scaffolding and standard patterns. While Copilot provided helpful starting points, all generated code was reviewed and adjusted manually to meet the specific functional and architectural requirements of the project. I also used Copilot to answer a few targeted questions related to the Dockerfile and Kubernetes configuration files, as I have limited hands-on experience with container orchestration and deployment tooling. This allowed me to explore those areas more effectively while ensuring the final implementation aligned with the project goals.