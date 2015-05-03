# name

`name`コマンドはランダムに名前を生成します。
中身はDockerのコンテナ名をつけるものとおなじです。

## Install

`go get github.com/ringohub/name`

## Usage

```bash
$ name -h
Usage of name:
  -n=1: Number of generate names
$ name
jolly_kirch
$ name -n 5
sleepy_pare
tender_carson
kickass_bartik
happy_meitner
agitated_pike
```

## References

- [Dockerコンテナのおもしろい名前](http://deeeet.com/writing/2014/07/15/docker-container-name/)
- [docker/pkg/namesgenerator](https://github.com/docker/docker/tree/master/pkg/namesgenerator)
