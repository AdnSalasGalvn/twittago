go build main.go
del main.zip
Compress-Archive -Path "C:\go\src\Gitlab\AdnSalasGalvn\twittago" -DestinationPath "main.zip"