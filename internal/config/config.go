package config

type Config struct {
	LinkedIn struct {
		Email    string
		Password string
	}

	Browser struct {
		Headless  bool
		UserAgent string
	}

	Limits struct {
		MaxConnectionsPerDay int
		MaxMessagesPerDay    int
	}

	Scheduler struct {
		ActiveStartHour int
		ActiveEndHour   int
	}

	Logging struct {
		Level string
	}
}
