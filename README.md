# Download Artifacts from SpeedCurve
## Introduction
[SpeedCurve](www.speedcurve.com) offers services to monitor front end performance through synthetic test and real user monitoring. Test results are nicely shown in the admin portal but they also offers API to download the raw data (as JSON), too. Artifacts such as the thumbnail and HAR files of the page request are available as a hyperlink. The WebPageTest result is also available in the form of XML as well.

Since we run synthetic tests several times a day and manually downloading the thumbnail and WPT XML are tedious. Thus, I have written this script in helping me to pull them down automatically.

## Prerequisite 
### Setup
The script is written in GO and you would need to have GO installed. Refer to the [Getting Started Page](https://golang.org/doc/install) for installation details.

Dependencies of the script can be installed via `go get -d ./...`

Also, you would need to get the API key from SpeedCurve for invoking their API.

### Export the Test Data as JSON
SpeedCurve provides a nice [document for invoking their API](https://api.speedcurve.com/#speedcurve-v1-api). I am using the one that retrieve all tests data for a single URL, a sample in curl as follow,

```
curl "https://api.speedcurve.com/v1/urls/{url_id}?days=30&browser=chrome&region=us-west-1" -u your-api-key:x
```

## Running the script
The script is executed as follows,
`go run download_artifacts.go`

Upon execution, a folder named with the test_id is created and the following files are downloaded within
- Screen captured during the test
- HAR file of the test
- XML export from WebPageTest

Notice that the test_id with the sample_export.json is just a sample. Therefore, the artifacts mentioned above are empty. Besides, you can change the name of the JSON file by modifying the argument of `ioutil.ReadFile` within the main function. 

## License
This project is licensed under the MIT License

## Acknowledgement
SpeedCurve is an essential tools for understanding how your website perform across platform and region. And the idea of writing this script originated from a post on SpeedCurve that a [NPM module](https://www.npmjs.com/package/speedcurve2csv) is developed for exporting the test data to CSV. This triggered my thought of writing this script to download the artifacts.