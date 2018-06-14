# dockerclean
Tools for cleaning up image tags in docker registry.

## List Images
```console
./dockerclean --link 'http://127.0.0.1:5000' list
```

## List Tags of an Image
```console
./dockerclean --link 'http://127.0.0.1:5000' tags --image drkaka/alpine
```
