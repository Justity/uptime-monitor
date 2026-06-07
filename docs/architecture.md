flowchart TD

User --> Grafana

Scheduler --> Checker

Checker --> PostgreSQL

Checker --> Telegram

Checker --> Metrics

Prometheus --> Metrics

Grafana --> Prometheus