# GO REST API

The current project is for educational purposes only, and an attempt to get more familiarized to the making of API's/Microservices;

## Tech Stack

* Golang
* Docker
* Postgres

## How to run it

1. **Write a .env file**
   In the repository root, create a file named '.env' with the following variables **PG_PASS** and **PG_USER**
   <br></br>
   Windows Powershell: 
   ```
   Set-Content -Path ./.env -Value 'PG_PASS=1234'
   Add-Content -Path ./.env -Value 'PG_USER=testuser'
   ```
   <br>GNU/Linux:
   ```
   cat << EOF >./.env
   PG_PASS=1234
   PG_USER=testuser
   EOF
   ```
3. **Run docker compose**
    * **[Command to run docker compose]** 
    * **[Optional: Additional instructions for running docker compose]**
