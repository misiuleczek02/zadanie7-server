# Zadanie 7 — Server (Go + Echo)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=bugs)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=misiuleczek02_zadanie7-server&metric=coverage)](https://sonarcloud.io/summary/new_code?id=misiuleczek02_zadanie7-server)

Aplikacja serwerowa w Go (framework Echo).

## Endpoints

- `GET /api/products` — lista produktów
- `POST /api/payments` — utworzenie płatności
- `GET /api/health` — healthcheck

## Uruchomienie

```
go run ./
```

## Testy

```
go test ./...
```

## Lint (golangci-lint w hookach gita)

Konfiguracja w `.golangci.yml`. Aktywacja hooków:

```
bash scripts/install-hooks.sh
```

(lub `./scripts/install-hooks.ps1` na Windows). Ustawia `core.hooksPath`
na `.githooks/`, gdzie:

- `pre-commit` — `gofmt`, `go vet`, `golangci-lint`, `go test`
- `pre-push` — `golangci-lint`, `go test`

## SonarCloud

- klucz projektu: `misiuleczek02_zadanie7-server`
- workflow: `.github/workflows/sonarcloud.yml`
- konfiguracja: `sonar-project.properties`
