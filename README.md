# elk-logger

    The Golang module for segregating logs for the Elasticsearch

### Configure your environment variables
    export environment=dev // Represents the environment of your application
    export release=1.0.0 // Represents the release of your application
    export elkHost=http://localhost:9200 // Represents the elasticsearch host
    export elkIndex=logger // Represents the name of your index
    export activeLogSegregation=true // Indicates whether logging is active or not, functionality included because there are open source applications that use this module but the executor does not have the interest to send the logs to the elasticsearch

### Import into your project

    Simple message
    Info("Tweet 0")

    Message with arguments 
    type Tweet struct {
		User    string
		Message string
	}

    Info("Tweet 1", Tweet{User: "User 1", Message: "Message 1"})

    To send error messages to elasticsearch, just follow the same basis as the example above using the error method

### Local build

    Run the docker-compose up command at the root of the project


## Releases

### 1.0.0 (19.11.2018)

* Initial version logger module