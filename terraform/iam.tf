
resource "aws_iam_role" "koronet_ecs_task_execution_role" {
  name               = "koronet-ecs-execution-role"
  assume_role_policy = file("files/assume_ecs_policy.json")
}

// The built-in AmazonECSTaskExecutionRolePolicy allows:
// 1. Pull from ECR
// 2. Write to CloudWatch
// 3. Read from Secret Manager
resource "aws_iam_policy_attachment" "ecs_task_execution_policy" {
  name       = "ecs-task-execution-policy-attachment"
  roles      = [aws_iam_role.koronet_ecs_task_execution_role.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "koronet_task_role" {
  name               = "koronet-container-role"
  assume_role_policy = file("files/assume_ecs_policy.json")
}