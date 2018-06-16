# dockerclean  [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]
Tools for cleaning up image tags in docker registry.

## List images
```console
./dockerclean --link '127.0.0.1:5000' list
```

## List tags of an image
```console
./dockerclean --link '127.0.0.1:5000' tags --image drkaka/alpine
```

## Delete a tag of an image
```console
./dockerclean --link '127.0.0.1:5000' delete --image drkaka/alpine --tag latest
```

## Keep certain number of tags
```console
./dockerclean --link '127.0.0.1:5000' keep --image drkaka/alpine --number 5
```
This will delete the older and keep the latest number of tags.

[ci-img]: https://travis-ci.org/drkaka/dockerclean.svg?branch=master
[ci]: https://travis-ci.org/drkaka/dockerclean
[cov-img]: https://coveralls.io/repos/github/drkaka/dockerclean/badge.svg?branch=master
[cov]: https://coveralls.io/github/drkaka/dockerclean?branch=master