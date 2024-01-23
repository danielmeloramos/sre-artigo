# Golang, Prometheus, Grafana Example

This is a simple example project demonstrating how to use Golang, Prometheus, and Grafana together. The Golang application exposes metrics, Prometheus scrapes these metrics, and Grafana visualizes the data.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/danielmeloramos/sre-artigo.git
   cd sre-artigo
    ```

2. Build the Golang application:

    ```bash
    make build
    ```

3. Run the entire setup:

    ```bash
    make run
    ```

4. This command starts the Golang application, Prometheus, and Grafana containers.

- Access Grafana at http://localhost:3000 (default credentials: admin/admin).
- Add Prometheus as a data source in Grafana with the URL http://prometheus:9090.
- Create a dashboard in Grafana and visualize the myapp_requests_total metric `sum(increase(myapp_requests_total[1m])) `.

5. Stopping the Setup. To stop and remove all containers, use the following command:

    ```bash
    make clean
    ```

6. Notes

    - The Golang application runs on port 8080;
    - Prometheus is accessible at `http://localhost:9090`;
    - Grafana is accessible at `http://localhost:3000` (default credentials: `admin/admin`).

Feel free to customize the Golang application and the Prometheus/Grafana configurations based on your needs.

```bash
Replace https://github.com/your-username/your-repository.git with the actual URL of your Git repository.
```