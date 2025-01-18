$date = Get-Date -format "yyyyMMdd"

$counter = $counter + 1
Set TIMESTAMP $date.$counter

Set-Content -Path .\.env -Value TIMESTAMP=$date.$counter
Add-Content -Path .\.env -Value PG_PASS=12345
Add-Content -Path .\.env -Value PG_USER=postgres

docker compose build 

docker compose up -d 