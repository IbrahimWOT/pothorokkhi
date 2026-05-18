# 🛡️ PothRokkhi (পথরক্ষী) - Secure Ride Sharing Platform

**PothRokkhi** is a secure, real-time, single-app ride-sharing backend ecosystem engineered with **Golang**, **Redis**, and **PostgreSQL (PostGIS)**. Designed and built from scratch using a learning-focused methodology, this platform serves as an academic cybersecurity evaluation model.

By implementing a responsive web interface with an instant **Rider/Driver Role-Switching** mechanism, the project demonstrates corporate-grade backend defenses without microservice overhead.

---

## 🚀 Core Architecture & Tech Stack
- **Backend Core:** Golang with Gin-Gonic (High-performance API Routing)
- **Frontend Layer:** Responsive HTML5, Tailwind CSS, Vanilla JavaScript
- **Geospatial Processing:** PostgreSQL + PostGIS Extension (Persistent Spatial Queries)
- **Live Telemetry Buffer:** Redis In-Memory Datastore (High-speed GPS coordination)
- **Deployment & Portability:** Docker Compose (Engineered for zero-config transfer between Windows & Kali Linux)

---

## 🎯 Security Features & Defense Mechanism
- **Cryptographic Access Control:** Password salting & hashing via `Bcrypt` with stateless context-injected `JWT` validation middlewares.
- **Geospatial Pipeline Protection:** Stateful WebSocket streams writing directly to memory via Redis `GEOADD` to shield persistent data stores from heavy I/O bottlenecks.
- **Infrastructure Hardening:** Anti brute-force protection using Redis Token Bucket Rate Limiters and implicit input sanitization against spatial telemetry spoofing.

---

## 📂 Project Structure Map
```text
ridesharing-project/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point & route registration
├── internal/
│   ├── auth/               # Bcrypt encryption & JWT handshake handlers
│   ├── db/                 # Postgres & Redis lifecycle initializers
│   ├── middleware/         # Security interception layers (Rate limiting, Auth)
│   ├── models/             # Shared application structures (Data layers)
│   └── tracking/           # WebSocket telemetry engines & spatial pipelines
├── web/                    # Responsive Frontend System
│   ├── index.html          # Phase 1: Interactive Landing Gateway
│   ├── auth.html           # Phase 2: Secure Access Panel
│   └── dashboard.html      # Phase 3 & 4: Live Spatial Tracking & Maps
├── .env                    # Hidden Environment Configurations (Not committed)
├── .gitignore              # Source control exclusion maps
├── docker-compose.yml      # Isolated stack database containers configuration
├── LICENSE                 # Open-source MIT Authorization Agreement
└── README.md               # Architecture definition & installation logs