
# Secure L2-to-L7 Networking Suite

## Overview
This project is a comprehensive networking application built in Go that spans multiple OSI layers:
- **Layer 2 (L2):** Data encapsulation and decapsulation.
- **Layer 3 (L3):** Mock IP packet generation and parsing.
- **Layer 4 (L4):** Reliable UDP implementation.
- **Layer 7 (L7):** Secure TLS-based communication.

### Why This Project Stands Out
This project demonstrates strong adherence to **Test-Driven Development (TDD)** and **Quality Assurance (QA)** practices. These methodologies ensure robust, maintainable, and high-quality code.

---

## Features
1. **Networking Functionality**:
   - **L2 Tunneling:** Encapsulation and decapsulation of data at Layer 2.
   - **L3 Packet Handling:** Generation and parsing of mock IP packets.
   - **L4 Reliable UDP:** Ensures data delivery with acknowledgment mechanisms.
   - **L7 Secure Communication:** TLS-based server-client communication for encrypted data transfer.

2. **TDD Practices**:
   - Unit tests written for each component **before** implementing core functionality.
   - Tests for edge cases, such as short or malformed packets.
   - High test coverage using `go test -cover`.

3. **QA Practices**:
   - **Integration Testing:** Ensures seamless interaction between components (e.g., UDP server-client and TLS communication).
   - **Static Analysis:** Linting with `golangci-lint` for code style and bug detection.
   - **Automated Testing Pipeline:** A `Makefile` simplifies testing, linting, and coverage reporting.

---

## Getting Started

### Prerequisites
- **Go**: Version 1.20 or later.
- **golangci-lint**: For linting (install via `brew install golangci-lint` or similar).
- **TLS Certificates**: `server.crt` and `server.key` for the TLS server.

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/networking-suite.git
   cd networking-suite
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   make build
   ```

---

## Usage
Run the application in various modes:

### Modes
- **TLS Server**:
  ```bash
  ./networking_suite server
  ```
- **TLS Client**:
  ```bash
  ./networking_suite client
  ```
- **Reliable UDP Server**:
  ```bash
  ./networking_suite udp-server
  ```
- **Reliable UDP Client**:
  ```bash
  ./networking_suite udp-client
  ```
- **Run Tests**:
  ```bash
  make test
  ```

---

## TDD Practices
1. **Unit Testing**:
   - Separate test files for each module (`l2_test.go`, `ip_test.go`, etc.).
   - Table-driven tests for edge cases (e.g., malformed packets).

2. **Integration Testing**:
   - Validates server-client interactions for both UDP and TLS communication.

3. **Coverage**:
   - Achieved high test coverage using:
     ```bash
     go test -cover ./...
     ```

4. **Test Automation**:
   - The `Makefile` provides an easy way to run all tests:
     ```bash
     make test
     ```

---

## QA Practices
1. **Static Analysis**:
   - Code is checked with `golangci-lint`:
     ```bash
     make lint
     ```

2. **Integration Tests**:
   - Validates end-to-end communication for reliable UDP and TLS.

3. **Logging**:
   - Logs are generated for debugging and operational insights:
     ```yaml
     log_file: "networking.log"
     ```

4. **Error Injection**:
   - Simulated packet corruption and loss for resilience testing.

