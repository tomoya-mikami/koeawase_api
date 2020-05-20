module src

go 1.12

replace local.packages/task => ./Task

replace local.packages/voice => ./Domain/Voice

replace local.packages/similarity => ./Domain/Similarity

replace local.packages/handler => ./Handler

require (
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/task v0.0.0-00010101000000-000000000000
)
