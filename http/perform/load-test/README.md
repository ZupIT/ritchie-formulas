# HTTP Load testing tool

the "http-load" formula named as Ritman is a tiny http load testing client dedicated for performance testing. As of now, ritman is intended for interactive usage, in a near future you'll be able to integrate on pipelines as a simple command call.

## http-load asks for two paremeters:
    - test duration; stamp you desire to ran the tests in seconds
    - max threads; the number of threads you want ritman to spawn for requests. Note, there're no magic numbers, you need to test which numbers firts your infrastructure best.

## First time usage:
##### rit http generate http-config
This command is necessary if you don't have a test template to run the load-test formula

##### rit http perform load-test
You should run this command in the same directory your "ritman-target.json" is, othersiwe, ritman will not know which target you disere to test.
The test is very interactire, it'll notify you once its done, which will pretty much generate a file, named "ritman-test-result-<timestamp>.json".

### Test result file example:
```json
{
	"hits": 314,
	"avgRps": 31,
	"avgMs": 5676,
	"maxMs": 60002,
	"minMs": 687,
	"startedAt": "2020-05-18T11:33:04-03:00",
	"endedAt": "2020-05-18T11:34:13-03:00",
	"histogram": {
		"10": {
			"hits": 22,
			"avgMs": 8748,
			"maxMs": 60000,
			"minMs": 1103,
			"statusCode": {
				"-1": 2,
				"200": 20
			}
		},
		"11": {
			"hits": 19,
			"avgMs": 10043,
			"maxMs": 60000,
			"minMs": 1009,
			"statusCode": {
				"-1": 1,
				"200": 18
			}
		},
		"12": {
			"hits": 29,
			"avgMs": 4799,
			"maxMs": 28349,
			"minMs": 779,
			"statusCode": {
				"200": 29
			}
		},
		"13": {
			"hits": 15,
			"avgMs": 6048,
			"maxMs": 60002,
			"minMs": 724,
			"statusCode": {
				"-1": 1,
				"200": 14
			}
		},
		"14": {
			"hits": 9,
			"avgMs": 4414,
			"maxMs": 12670,
			"minMs": 697,
			"statusCode": {
				"200": 9
			}
		},
		"4": {
			"hits": 100,
			"avgMs": 4788,
			"maxMs": 48020,
			"minMs": 1909,
			"statusCode": {
				"200": 100
			}
		},
		"6": {
			"hits": 23,
			"avgMs": 6903,
			"maxMs": 21718,
			"minMs": 741,
			"statusCode": {
				"200": 23
			}
		},
		"7": {
			"hits": 51,
			"avgMs": 6604,
			"maxMs": 60000,
			"minMs": 687,
			"statusCode": {
				"-1": 1,
				"200": 50
			}
		},
		"8": {
			"hits": 19,
			"avgMs": 2802,
			"maxMs": 8364,
			"minMs": 1061,
			"statusCode": {
				"200": 19
			}
		},
		"9": {
			"hits": 27,
			"avgMs": 3774,
			"maxMs": 23273,
			"minMs": 688,
			"statusCode": {
				"200": 27
			}
		}
	}
}
```
Note:
- The "histogram" field is a series of result where the key is the second that the request was fired to the target endpoint.
- You may notice some "-1" in the "statusCode" field, which means the requests either took a Socket Timeout (maximum wait time is 60 seconds) or your Operation System may exceeded the number of Created Socket in the Socket Stack (65k approx). So make sure to NOT RUN the http service in the same machine you'll run the load testing.

Acronyns:
    - avg - stands for average
    - rps - stands for rate per seconds
    - ms - stands for milliseconds