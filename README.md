# gorep

gorep is my weekend project to develop an example of clean simple golang application architecture

## The Architecture

- cmd
- databases
- modules
- routes

### cmd

This directory consist main file to run the app. can consist more than one app main file.

### databases

This directory store databases configuration, and return the `db *sql.DB` to use by another packages.

### modules

This directory store modules that used by app . this modules separated by business function.

On this example application modules dir contain another two directory : `user` and `post`.

by default each module containt these folder :

- repository

Repository directory act as a layer that comunicate with data layer. on this dir the application do queries to data layer . can be database or API source

- model

Model directory act as a representative of data model from source, and data model to be delivered on struct form.

- controller

Controller directory act as busines logic of module. Controller can use another module repository.

- api

Api directory act as a module delivery layer. This directory can formed into another way to deliver data to client such as HTML.

### Routes

This directory handle the http route from client and map it into correct handler
