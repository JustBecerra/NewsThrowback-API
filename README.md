# NewsThrowback-API
API for processing old news articles from public APIs

flow of endpoints go as follows:
- routes -> handlers (handles HTTP) -> repositories(runs the queries uses models from Models folder)
- db (handles database config and migrations)
- config (handles API config)
sources:
U.S. → Chronicling America
Australia → Trove
Europe → Europeana and/or Gallica