
resource "aws_ecs_cluster" "koronet" {
  name = "koronet"
}

resource "aws_ecs_task_definition" "koronet" {
  family                = "koronet"
  container_definitions = file("files/ecs_task_definition.json")

  cpu    = 1000
  memory = 512

  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.koronet_ecs_task_execution_role.arn
}

resource "aws_ecs_service" "koronet" {
  name            = "koronet"
  cluster         = aws_ecs_cluster.koronet.id
  task_definition = aws_ecs_task_definition.koronet.arn
  launch_type     = "FARGATE"
  desired_count   = 3
  iam_role        = aws_iam_role.koronet_ecs_task_execution_role.arn

  network_configuration {
    subnets          = module.vpc.private_subnets
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.webapp.arn
    container_name   = "webapp"
    container_port   = 8080
  }
}
