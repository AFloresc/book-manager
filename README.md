# 🏗️ Go Hexagonal Architecture: Book API

Este proyecto es una implementación de referencia de la **Arquitectura Hexagonal** (también conocida como Puertos y Adaptadores) utilizando el lenguaje de programación Go.

El objetivo es demostrar cómo separar las reglas de negocio del mundo exterior (bases de datos, APIs, frameworks HTTP), permitiendo que el sistema sea fácil de testear, mantener y evolucionar.

---

## 📂 Estructura del Proyecto

Siguiendo los principios del "Hexágono", el código se organiza de la siguiente manera:

```text
book-api/
├── cmd/
│   └── main.go           # Punto de entrada e Inyección de Dependencias
├── internal/
│   ├── core/             # El corazón de la aplicación (Independiente)
│   │   ├── domain/       # Entidades de negocio (Models)
│   │   ├── ports/        # Interfaces (Contratos de entrada y salida)
│   │   └── services/     # Casos de uso (Lógica de negocio)
│   └── adapters/         # Implementaciones técnicas (Mundo exterior)
│       ├── repository/   # Adaptadores de salida (Persistencia)
│       └── handler/      # Adaptadores de entrada (HTTP/CLI/gRPC)
└── go.mod
```

## 🛠️ Capas de la Arquitectura

 1. Domain: Contiene las estructuras de datos puras. No conoce nada de JSON ni de SQL.

 2. Ports: Define las interfaces. El BookRepository (salida) y el BookService (entrada) actúan como el pegamento del sistema.

 3. Services: Aquí reside la lógica. Utiliza los puertos para comunicarse con el exterior sin saber qué tecnología hay detrás.

 4. Adapters: Implementaciones concretas. En este ejemplo, utilizamos un Memory Repository (podría ser Postgres mañana) y un HTTP Handler.

## 🚀 Instalación y Ejecución

1. **Clona el repositorio:**

   ```bash
   git clone https://github.com/tu-usuario/book-hexagonal-go.git
   cd book-hexagonal-go
   ```

2. **Inicializa el módulo (si no existe):**

   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación:**

    ```bash
    go run cmd/main.go
    ```

4. **Prueba la API (POST):**
La API está configurada para escuchar en el puerto :8080. Puedes probar el endpoint de creación:

    ```bash
    curl -X POST http://localhost:8080/books \ -H "Content-Type: application/json" \ -d '{"title": "El Quijote", "author": "Miguel de Cervantes"}'
    ```
