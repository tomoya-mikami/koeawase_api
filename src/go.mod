module src

go 1.13

replace local.packages/task => ./Task

replace local.packages/voice => ./Domain/Voice

require local.packages/task v0.0.0-00010101000000-000000000000
