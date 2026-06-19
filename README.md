# 🏍️ Dinz RentBike

Rental motor & mobil — booking cepat, bayar mudah, langsung jalan.

**Deployment:** [https://dinz-rentbike.up.railway.app/](https://dinz-rentbike.up.railway.app/)

---

## Tech Stack

| Layer | Tech |
|---|---|
| Language | Go 1.23+ |
| HTTP Router | [Echo v4](https://echo.labstack.com/) |
| ORM | [GORM](https://gorm.io/) |
| Database | PostgreSQL |
| Auth | JWT (golang-jwt) |
| Payment | [Xendit Payment Sessions](https://docs.xendit.co/) |
| Email | [Mailjet v3.1](https://dev.mailjet.com/email/guides/) |
| Test | testify + mockery v3 |

---

## Project Structure

```
.
├── cmd/app/                  # Entry point
├── internal/
│   ├── bootstrap/            # App init & server wiring
│   ├── config/               # Viper config (.env + env vars)
│   ├── delivery/http/
│   │   ├── handler/          # HTTP handlers (auth, user, vehicle, rental, payment, review, admin, webhook)
│   │   └── middleware/       # Auth, Role, Logger middleware
│   ├── domain/
│   │   ├── constants/        # Status enums
│   │   ├── contract/         # Interfaces (repository & usecase)
│   │   ├── dto/              # Request/Response DTOs
│   │   └── entity/           # GORM entities
│   ├── external/
│   │   ├── mailjet/          # Mailjet email client
│   │   └── xendit/           # Xendit payment client
│   ├── infrastructure/       # Database connection
│   ├── repository/           # Repository implementations
│   └── usecase/              # Business logic
├── pkg/
│   ├── jwt/                  # JWT auth manager
│   ├── logger/               # Zerolog logger
│   ├── response/             # API response helpers
│   └── utils/                # Password hashing, random string
├── sql/                      # DDL & seed data
├── docs/                     # OpenAPI spec & Postman collection
├── tests/                    # Unit tests
│   └── mock/                 # Auto-generated mocks (mockery v3)
└── .mockery.yml              # Mockery config
```

---

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL
- [mockery v3](https://vektra.github.io/mockery/) (for generating mocks)

### Setup

```bash
# Clone
git clone <repo-url> && cd dinz-rentbike

# Install deps
go mod tidy

# Copy env
cp .env.example .env
# Edit .env with your credentials

# Run DDL + Seed
psql -h localhost -U postgres -d rentbike_db -f sql/ddl.sql
psql -h localhost -U postgres -d rentbike_db -f sql/seed.sql

# Run
go run cmd/app/main.go
```

### Environment Variables

```env
# Server
APP_NAME=P02LC3
APP_HOST=localhost
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=rentbike_db
DB_SCHEMA=public

# JWT
JWT_SECRET=your-jwt-secret-key-here

# Xendit
XENDIT_BASE_URL=https://api.xendit.co/
XENDIT_PUBLIC_KEY=
XENDIT_SECRET_KEY=
XENDIT_WEBHOOK_TOKEN=your-webhook-token

# Mailjet
MAILJET_BASE_URL=https://api.mailjet.com/v3.1/send
MAILJET_API_KEY=
MAILJET_SECRET_KEY=
MAILJET_FROM_EMAIL=noreply@dinzrentbike.com
```

---

## API Endpoints

### Auth
| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/v1/auth/register` | Register new user |
| POST | `/api/v1/auth/login` | Login, returns JWT |

### User (🔒)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/users/me` | Get my profile |
| PUT | `/api/v1/users/update` | Update name & phone |
| PATCH | `/api/v1/users/change-password` | Change password |

### Vehicles (🔒)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/vehicles` | List all vehicles |
| GET | `/api/v1/vehicles/:id` | Detail + reviews |

### Rentals (🔒)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/rentals` | My rental history |
| GET | `/api/v1/rentals/:id` | Detail + payment + review |
| POST | `/api/v1/rentals/create` | Create rental + auto payment |
| POST | `/api/v1/rentals/cancel` | Cancel rental |

### Payments (🔒)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/payments` | My payment history |
| GET | `/api/v1/payments/:id` | Payment detail |
| POST | `/api/v1/payments/cancel` | Cancel payment |

### Reviews (🔒)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/reviews` | My reviews |
| POST | `/api/v1/reviews/create` | Create review |
| POST | `/api/v1/reviews/edit` | Edit review |
| POST | `/api/v1/reviews/delete` | Delete review |

### Admin (🔒 admin role)
| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/v1/admin/vehicles` | List vehicles (admin) |
| POST | `/api/v1/admin/vehicles` | Create vehicle |
| PUT | `/api/v1/admin/vehicles/:id` | Update vehicle |
| DELETE | `/api/v1/admin/vehicles/:id` | Delete vehicle |
| GET | `/api/v1/admin/rentals` | List all rentals |
| GET | `/api/v1/admin/rentals/:id` | Rental detail |
| PATCH | `/api/v1/admin/rentals/:id/status` | Update rental status |
| GET | `/api/v1/admin/payments` | List all payments |
| GET | `/api/v1/admin/payments/:id` | Payment detail |

### Webhook (Xendit)
| Method | Endpoint | Description |
|---|---|---|
| POST | `/webhook` | Receive payment session updates |

---

## Testing

```bash
# Run all tests
go test ./tests/ -v

# Run specific tests
go test ./tests/ -v -run TestRegister
go test ./tests/ -v -run TestLogin

# Generate mocks (after contract changes)
mockery
```

---

## Flow

```
1. Register / Login → JWT token
2. Browse vehicles → GET /vehicles
3. Create rental → POST /rentals/create
   ├── Rental created (vehicle → rented)
   ├── Payment auto-created via Xendit
   └── Response includes xendit_payment_url
4. Pay via Xendit hosted page
5. Xendit webhook → POST /webhook
   ├── Payment status → paid
   └── Email notification sent to user
6. Review after rental completed → POST /reviews
```

---

## Import to Postman

1. Open Postman → Import
2. Select `docs/postman.json`
3. Set `base_url` variable to `https://dinz-rentbike.up.railway.app/api/v1`
4. Login via Auth → Login (auto-saves token)

---

## Seed Users

| Email | Password | Role |
|---|---|---|
| admin@rentbike.com | password123 | admin |
| budi@rentbike.com | password123 | customer |
| siti@rentbike.com | password123 | customer |
| ahmad@rentbike.com | password123 | customer |
