# Native Server

A server providing appointment data via JSON (v1) and Protobuf (v2) APIs.

## Endpoints

### Version 1 (JSON)
- `GET /api/v1/appointments`: Get all appointments
- `POST /api/v1/appointments`: Create a new appointment
- `PUT /api/v1/appointments/{id}`: Update an appointment
- `DELETE /api/v1/appointments/{id}`: Delete an appointment

### Version 2 (Protobuf)
- `GET /api/v2/appointments`: Get all appointments

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/harshja1n/native_server.git
    cd native_server
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the server:
    ```sh
    go run main.go
    ```

2. The server runs on port `8000`.

## Examples

### Version 1 (JSON)
- Get all appointments:
    ```sh
    curl http://localhost:8000/api/v1/appointments
    ```

- Create a new appointment:
    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments
    ```

- Update an appointment:
    ```sh
    curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments/1
    ```

- Delete an appointment:
    ```sh
    curl -X DELETE http://localhost:8000/api/v1/appointments/1
    ```

### Version 2 (Protobuf)
- Get all appointments:
    ```sh
    curl http://localhost:8000/api/v2/appointments
    ```

## Custom Logger

This project includes a custom logging middleware for better request tracking and debugging. The middleware logs each request's method, URL, and response status.

## Response Size Comparison

JSON and Protobuf have different characteristics in terms of response size:

- JSON is human-readable but generally larger in size.
- Protobuf is a binary format, more compact, and faster for serialization/deserialization.

To compare response sizes, you can use the following commands:

- JSON response size:
    ```sh
    curl -s -o /dev/null -w "%{size_download}\n" http://localhost:8000/api/v1/appointments
    ```

- Protobuf response size:
    ```sh
    curl -s -o /dev/null -w "%{size_download}\n" http://localhost:8000/api/v2/appointments
    ```

## License

This project is licensed under the MIT License.
