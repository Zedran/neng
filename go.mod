module github.com/Zedran/neng

go 1.23.0

require golang.org/x/text v0.24.0

retract v0.9.0 // Retract v0.9.0 due to a post-publishing modification

// Retract versions v0.8.2 through v0.15.4 due to licensing conflict
retract [v0.8.2, v0.15.4]
