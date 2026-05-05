# lib-preconditions-golang

On the top of [keytiles/lib-errorhandling-golang](https://github.com/keytiles/lib-errorhandling-golang) brings simple methods helping to validate
parameters / constraints and build `kt_errors.Faults` with less boilerplate.

# How to use?

More details might come later, but basically something like this:

```go
const PACKAGE_NAME = "lib.kt_precond"

func MyFunc(myStrParam string) kt_errors.Fault {
    methodName := "MyFunc()"

    // Checking a parameter - will raise ValidationFault if 'myParam' is empty
    if fb := kt_precond.EnsureStringParamNotEmpty(myStrParam, "myStrParam", PACKAGE_NAME, methodName); fb != nil {
        return faultBuilder.Build()
    }

    ...

    intval, err := strconv.Atoi("12")
    // Checking a state - will raise IllegalStateFault if condition is not met
    if fb := kt_precond.EnsureState(err == nil, "'myStrParam' value '{val}' is not an integer", PACKAGE_NAME, methodName); fb != nil {
        return faultBuilder.WithLabel("val", myStrParam).Build()
    }

    ...

}
```
