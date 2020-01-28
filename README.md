# Email validation service
[![Build Status](https://travis-ci.com/neville/email-validator.svg?branch=master)](https://travis-ci.com/neville/email-validator)

## Run via Docker 

**Image**  
https://hub.docker.com/r/nevillekb/email-validator

**Command**  
```docker run -t -p 8080:8080:127.0.0.1 -e PORT=8080 nevillekb/email-validator```

## API documentation

**Method**  
HTTP POST 

**URL**  
```<domain>:<port>/email/validate```

**Headers**  
Name - Content-Type  
Value - application/json

**Body**  
```{"email":"xxx@yyy.zzz"}```

**Response**
```
{
  "valid": false,
  "validators": {
    "regexp": {
      "valid": true
    },
    "domain": {
      "valid": false,
      "reason": "INVALID_TLD"
    },
    "smtp": {
      "valid": false,
      "reason": "UNABLE_TO_CONNECT"
    }
  }
}
```

## References

### Project structuring best practises 
- https://github.com/golang-standards/project-layout
- https://github.com/marvincaspar/go-web-app-boilerplate
