module "ecs_cluster" {
  source  = "HENNGE/ecs/aws"
  version = "1.0.0"

  name = local.name
}