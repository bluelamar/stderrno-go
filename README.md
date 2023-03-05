# stderrno-go : Standardized Errors

This package provides standardized errors to be used instead of system dependent errors from underlying systems.
Errors from the underlying system may not be understood by upper level components.

Inversion of control design means the specifics of an implementation are pushed down.
This is usually done by hiding the implementations of lower level components behind interfaces or other functionality.

Those components that wrap underlying systems can then return known errors mapped from the internal system errors.
This way higher level components can check for specific standardized errors without having to know the internal implementation and its errors.

HTTP errors from an underlying system can be mapped to the standardized errors.
Example:

| HTTP Response Code | Standardized Error |
|--------------------|--------------------|
| 401                | EPERM              |
| 403                | EPERM              |
| 404                | ENOENT             |
| 408                | ETIME              |


Consider a component trying to access an object in AWS S3.
Here are some of the errors that can be returned by S3 as an HTTP 403.

* AccessDenied
* AccountProblem
* AllAccessDisabled
* CrossLocationLoggingProhibited
* InvalidAccessKeyId
* InvalidObjectState
* InvalidPayer
* InvalidSecurity
* NotSignedUp
* RequestTimeTooSkewed
* SignatureDoesNotMatch

The underlying component of your system could return an EPERM wrapped error.
Example:
```
fmt.Errorf("AccountProblem: %w", stderrno.EPERM)
```

## License

Released under the Apache License.

