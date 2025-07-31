
WebApp
Choosed Go 

GH ACTION
el job para deployar ecs esta definido, pero ya que romperia al no haber ECS real, puse un condicional para que no se ejecute
como estamos sobrescribiendo latest, no hace falta definir un nuevo task_definition, sin embargo con mas tiempo y en un ambiente real,
definiria un nuevo task_definition

Terraform
State - Defined a backend with S3 but commented out since can't run the code
State - Mocking aws credentials just to be able to run terraform plan
VPC - Created one NAT per AZ for HA
ECS - Created containers in private subnets and then a LB internet-facing
TaskDefinition - Si bien se asume que en prod se usaria ECR, el ejercicio usa private dockerhub, entonces agrego secret manager