Get-Content .env | ForEach-Object {
  $name, $value = $_.split('=')
  if ([string]::IsNullOrWhiteSpace($name) -or $name.Contains('#')) {
    return
  }
  Set-Content env:\$name $value
}

New-Item languages.txt -ItemType File -Force

$CURRENT_DIR=$pwd
Get-ChildItem docker -Directory -Exclude _base | Foreach-Object {
    $name = $_.Name
    Set-Location docker/$name
    Copy-Item ../_base/* -Destination ./base
    docker build -t "$env:DOCKER_IMAGE_PREFIX$name" -f Dockerfile .
    Set-Location ../..
    Add-Content -Path languages.txt -Value "$name`n" -NoNewline
}
Set-Location $CURRENT_DIR

docker image prune -f