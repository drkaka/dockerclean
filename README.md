# dockerclean  [![Docker Build][docker-img]][docker] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]
Tools for cleaning up image tags in docker registry.

## Docker container
Using docker container is more flexible to send request.

### List images
```console
$ docker run --rm drkaka/dockerclean -l 127.0.0.1:5000 list
repositories:
	drkaka/alpine
```

### List tags of an image
```console
$ docker run --rm drkaka/dockerclean -l 127.0.0.1:5000 tags --image drkaka/alpine
Tags for drkaka/alpine:
	3.5	3.6	latest
```

### Delete a tag of an image
```console
$ docker run --rm drkaka/dockerclean -l 127.0.0.1:5000 delete --image drkaka/alpine --tag latest
Please run "docker exec -it registry /bin/registry garbage-collect  /etc/docker/registry/config.yml" to free space.
```

If a tag is successfully deleted, the upper information will show to remind users to do [Garbage collection](https://docs.docker.com/registry/garbage-collection/). If some issues occur, like re-pushing to the deleted tag but can't pull back, try to restart the registry.

### Keep certain number of tags
```console
$ docker run --rm drkaka/dockerclean -l 127.0.0.1:5000 keep --image drkaka/alpine --number 5
```

This will delete the older and keep the latest number of tags.

## Command line application
Latest builds can be downloaded [here](https://github.com/drkaka/dockerclean/releases/latest).  

### List images
```console
$ dockerclean --link '127.0.0.1:5000' list
repositories:
	drkaka/alpine
```

### List tags of an image
```console
$ dockerclean --link '127.0.0.1:5000' tags --image drkaka/alpine
Tags for drkaka/alpine:
	3.5	3.6	latest
```

### Delete a tag of an image
```console
$ dockerclean --link '127.0.0.1:5000' delete --image drkaka/alpine --tag latest
```

### Keep certain number of tags
```console
$ dockerclean --link '127.0.0.1:5000' keep --image drkaka/alpine --number 5
```

This will delete the older and keep the latest number of tags.



[docker-img]: https://img.shields.io/docker/build/drkaka/dockerclean.svg
[docker]: https://hub.docker.com/r/drkaka/dockerclean/
[ci-img]: https://travis-ci.org/drkaka/dockerclean.svg?branch=master
[ci]: https://travis-ci.org/drkaka/dockerclean
[cov-img]: https://coveralls.io/repos/github/drkaka/dockerclean/badge.svg?branch=master
[cov]: https://coveralls.io/github/drkaka/dockerclean?branch=master