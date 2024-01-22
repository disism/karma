
## Get Version
```shell
curl --location 'localhost:7330/version'
```

## Create User
```shell
curl --location 'localhost:7330/users/create' \
--form 'username="hvturingga"' \
--form 'password="123123"'
```

## Login
```shell
curl --location 'localhost:7330/login' \
--form 'username="hvturingga"' \
--form 'password="123123"'
```


