module "homolog" {
    source = "./config"
    nome_repositorio = "homolog"
    cargoIAM = "homolog"
    ambiante = "homolog"
    imagem="rafaelcarvalhocaetano/ci_golang:34"
}

output "IP_alb" {
  value = module.homolog.IP
}