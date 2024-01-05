# guild-backend
guild-backend un applicazione che espone delle API per la gestione di una gilda di un videogioco.


## Installation
```bash
git clone https://github.com/fededomm/guild-backend.git

cd guild-backend

go build -o guild-backend

go run . # or ./guild-backend
```

### Docker Compose
utilizziamo docker compose per poter pullare i container necessari al funzionamento dell'applicazione.
```bash
docker-compose build 

docker-compose up
```

## File di configurazione
creare un file allo stesso livello di guild-backend chiamato config.yaml con la seguente struttura:

```yaml
log:
  level: 
  enable_json: 
database:
  host: 
  port: 
  user: 
  password: 
  dbname: 
  sslmode: 
observability:
  enable: 
  serviceName: 
  endpoint: 
```

### log
- level: livello di log 
- enable_json: abilita il log in formato json

### database
specificare i parametri di connessione al database postgres

### observability
specificare le configurazioni per l'invio delle tracce e delle metriche ad un collector opentelemetry (vedi docker-compose.yaml per un esempio di collector).

NB: in endpoint specificare il nome del container specificato all'interno di docker-compose.yaml grazie ad una funzione built-in di docker compose chiamata service discovery.
