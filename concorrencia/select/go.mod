module cod3r/select

go 1.21.0

// O caminho relativo deve subir 3 níveis para sair de:
// select -> concorrencia -> cod3r -> GOLANG
// E então entrar em cod3r2
replace cod3r2 => ../../../cod3r2

require cod3r2 v0.0.0-00010101000000-000000000000
