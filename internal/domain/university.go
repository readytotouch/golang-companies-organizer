package domain

type University struct {
	LinkedInProfile LinkedInProfile
}

type Course struct {
	LinkedInProfile   LinkedInProfile
	AlumniCount       int
	DouCurrentCount   int
	DouPastCount      int
	FaangCurrentCount int
	FaangPastCount    int
}
