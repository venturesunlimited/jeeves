Jeeves
======

An awesome, obedient bot for Slack.

To get a token, sign into the Slack website and create a new integration that sends a POST request to your bot. Add the token that you get to your ENV and you should be good to go.

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
