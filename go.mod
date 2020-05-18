module github.com/tomoya-paseri/koeawase_api

go 1.13

replace local.packages/voice => ./src/Domain/Voice

replace local.packages/task => ./src/Task

replace local.packages/src => ./src

replace local.packages/handler => ./src/Handler

require (
	cloud.google.com/go/firestore v1.2.0
	github.com/gofiber/fiber v1.9.6
	github.com/google/wire v0.4.0
	github.com/mjibson/go-dsp v0.0.0-20180508042940-11479a337f12
	github.com/youpy/go-wav v0.0.0-20160223082350-b63a9887d320
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.20.0
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/src v0.0.0-00010101000000-000000000000
	local.packages/task v0.0.0-00010101000000-000000000000
	local.packages/voice v0.0.0-00010101000000-000000000000
)
