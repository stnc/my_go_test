# my_go_test

# command list 

## how complete (html output )

```  go test -coverprofile=coverage.out ```

## how complete (html output open to  web browser)

```  go tool cover -html=coverage.out ```

## how complete (percentile)

```  go test -cover . ```

``` go run ./cmd/web ```

##  only single go page

``` go test -timeout 30s -run ^Test_application_handlers$ webapp/cmd/web ```
