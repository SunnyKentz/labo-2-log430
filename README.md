## Labo 2

### Explication de l'application:

l'application sur les ports 8080 : magasin, 8090: mere, 8091 : centre logistique

### Comment run :
```
    make run
```

### Comment tester :
```
    make test
```

### Explication du CI
Après avoir fait un push, Github action check le linting du push,<br> execute les testes et si tous passe, le push vers dockerhub