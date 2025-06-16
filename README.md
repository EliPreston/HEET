# H.E.E.T
## Home Energy Estimator Tool


### Initial Setup:
- `docker compose up -d` or `docker compose up` 
- Go to  http://localhost:5050
- Add postgres server to pgAdmin UI
    - Click "Add New Server"
    - In the General tab:
        - Name: HEET (doesn't actually matter)
    - In the Connection tab:
        - Host name/address must match postgres service from docker compose file: postgres_container
        - Port: 5432
        - Username: postgres
        - Password: admin
    - Press 'Save'
#### This should allow the user to do `docker compose down` and then `docker compose up` and not have to add the postgres server to the pgAdmin UI every time.