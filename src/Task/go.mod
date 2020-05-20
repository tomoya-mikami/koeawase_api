module Task

go 1.12

replace local.packages/voice => ../Domain/Voice

replace local.packages/similarity => ../Domain/Similarity

require (
	local.packages/similarity v0.0.0-00010101000000-000000000000
	local.packages/voice v0.0.0-00010101000000-000000000000
)
