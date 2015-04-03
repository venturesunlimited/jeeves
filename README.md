Jeeves
======

An awesome, obedient bot.

To add a new script, implement a function that accepts the incoming message string and outputs a Result interface like:
```
  package scripts
  
  import (
    ...
  )
   
  func NameOfScript(msg string) Result {
      // do some and return JSON
  }
```
