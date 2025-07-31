
resource "aws_ecs_cluster" "koronet" {
  name = "koronet"
}

resource "aws_ecs_task_definition" "koronet" {
  family = "koronet"
  container_definitions = file("files/ecs_task_definition.json")

  cpu = 1000
  memory = 512

  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  # execution_role_arn       =
}

resource "aws_ecs_service" "koronet" {
  name            = "koronet"
  cluster         = aws_ecs_cluster.koronet.id
  task_definition = aws_ecs_task_definition.koronet.arn
  launch_type     = "FARGATE"
  desired_count   = 3
  # iam_role        = aws_iam_role.foo.arn
  depends_on      = [aws_iam_role_policy.foo]

#   ordered_placement_strategy {
#     type  = "binpack"
#     field = "cpu"
#   }

#   load_balancer {
#     target_group_arn = aws_lb_target_group.foo.arn
#     container_name   = "mongo"
#     container_port   = 8080
#   }

#   placement_constraints {
#     type       = "memberOf"
#     expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
#   }
}
