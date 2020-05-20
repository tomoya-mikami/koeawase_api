module Similarity

go 1.12

replace local.packages/voice => ../Voice

require (
	cloud.google.com/go/firestore v1.2.0
	local.packages/voice v0.0.0-00010101000000-000000000000
)
