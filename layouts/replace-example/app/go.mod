module layouts/replace-example/app

go 1.23.2

replace github.com/SomeFancyAccount/sender v0.0.0 => ../sender

replace github.com/SomeFancyAccount/receiver v0.0.0 => ../receiver

require (
	github.com/SomeFancyAccount/receiver v0.0.0
	github.com/SomeFancyAccount/sender v0.0.0
)
