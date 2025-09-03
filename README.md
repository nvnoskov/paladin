# Paladin - Caching Service for Mobile Applications

Paladin is a high-performance caching service designed for mobile applications, built with Go. It provides a robust API for storing, retrieving, and managing currency exchange rate data with a focus on performance and reliability.

## Features

- **High-Performance Caching**: In-memory cache with statistics tracking
- **Persistent Storage**: Uses BadgerDB for durable storage with automatic garbage collection
- **RESTful API**: Simple and efficient API for data operations
- **Docker Support**: Containerized deployment with minimal image size
- **Currency Exchange Data**: Specialized for storing and retrieving currency exchange rates
- **TTL Support**: Automatic expiration of cached data
- **Batch Operations**: Efficient bulk data synchronization

## Architecture

Paladin is built with a clean architecture that separates concerns:

- **Models**: Data structures for currency exchange points and rates
- **Controllers**: API endpoints for data operations
- **Cache**: In-memory caching layer for fast access
- **Database**: Persistent storage using BadgerDB

## API Endpoints

### Punkt (Exchange Point) Operations

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/punkts/` | Get all exchange points |
| GET | `/punkts/{id}` | Get a specific exchange point by ID |
| POST | `/punkts/` | Synchronize multiple exchange points |
| PUT | `/punkts/{id}` | Update a specific exchange point |
| POST | `/punkts/{id}` | Update exchange rate data for a specific point |

### Statistics

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/stats` | Get cache statistics (hits, misses, sets, etc.) |

## Data Model

The main data structure is `Punkt` which represents a currency exchange point with the following key fields:

- **ID**: Unique identifier
- **City**: City location
- **Name**: Exchange point name
- **Address**: Physical address
- **Phones**: Contact information
- **Data**: Currency exchange rates (USD, EUR, RUB, etc.)
- **Workmodes**: Operating hours for each day of the week
- **Map**: Geographic coordinates

## Getting Started

### Prerequisites

- Docker (for containerized deployment)
- Go 1.13+ (for development)

### Running with Docker

```bash
# Pull the latest image
docker pull registry.gitlab.com/kurs-kz/paladin:v4.5

# Run the container
docker run -p 8000:8000 registry.gitlab.com/kurs-kz/paladin:v4.5
```

### Development Setup

1. Clone the repository:
```bash
git clone https://github.com/nvnoskov/paladin.git
cd paladin
```

2. Install dependencies:
```bash
go get -u -t github.com/tinylib/msgp
go get -u -t github.com/dgraph-io/badger/v2
go get -u -t github.com/savsgio/atreugo/v11
go get -u -t github.com/sirupsen/logrus
```

3. Build the application:
```bash
go build -o paladin
```

4. Run the application:
```bash
./paladin --database ./database --listen :8000
```

## Configuration

Paladin accepts the following command-line flags:

- `--database`: Path to the database directory (default: "./database")
- `--listen`: Address and port to listen on (default: ":8000")
- `--verbose`: Enable verbose logging (default: true)
- `--profile`: Enable blocks and mutexes profiling (default: false)

## Docker

The project uses a multi-stage build to create a minimal Docker image:

```dockerfile
FROM scratch
COPY build/ca-certificates.crt /etc/ssl/certs/
COPY build/paladin /go/paladin

WORKDIR /go
VOLUME ["/go/database"]
CMD ["/go/paladin"]
```

## CI/CD Pipeline

The project includes a GitLab CI configuration with the following stages:

1. **Test**: Format checking, static analysis, and unit testing
2. **Build**: Docker image building and publishing

## Performance Considerations

- In-memory cache for frequently accessed data
- BadgerDB for efficient persistent storage
- Automatic garbage collection for database maintenance
- TTL (Time-To-Live) for cached data to ensure freshness

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am 'Add feature'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## License

This project is proprietary software. All rights reserved.

## Contact

For questions or support, please contact the development team.