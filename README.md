Exercise: Write a simple fizz-buzz REST server. 
---

>The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"". 
>The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."".
>
>Your goal is to implement a web server that will expose a REST API endpoint that:
>- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
>- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

How to
---

Run the server with `go run .`

Run the test suite with `go test ./...`

API
---

I decided to restrict parameters as the program could panic with an OOM error.
* `limit` highest value is 5 000 000
* `int1` and `int2` not equal to 0 to avoid 0 division
* `str1` and `str2` longest possible size is 128

Endpoints are
* `localhost:8080/fizzbuzz`
* `localhost:8080/stats`

Benchmark
---
I used two variation for generating the fizzbuzz and choose the second one for the final result.

First one (`FizzBuzz#Array`) is using an array of strings. Then format it to JSON in order to return it as payload.

`limit` is set to 5 000 000

>cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
>
>13.885s


Second one (`FizzBuzz#Json`) is building the response string already formatted as JSON.
>cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
>
>3.161s

Also I avoided the modulo operator to improve performance.