- cmd
  - releasifier
    + main.go
      -> reads config file and creates config.Config object
      -> creates releasifier object and passes config.Config into it
      -> call releasifier.Start()
- config
- data : can only access libs package inside releasifier and other packages outside.
  + db.go
- etc : all configuration files should be here
- logme : uses private global variable which can be only access by public global Funcs
- lib : everything in lib has to be stateless, no global variables
        they must not depends on any other releasifier packages.
  - crypto
  - utils
- scripts
- web
  - apis
    - auth
    - users
    - bundles
    - releases
  - middlewares : should be stateless, no global variable
  - security
    - jwt
    + security.go
      -> creates a global variable for each package
  + web.go
    -> setup the security package
    -> returns api's handlers

+ releasifier.go
  -> setup logme : internally it creates a private global to its package
  -> setup db : internally it creates a global to its package, e.g. `data.DB`
  -> setup security : internally it creates a global variable for each package

  -> Start()
    -> instantiates web package and assign it to graceful package.
