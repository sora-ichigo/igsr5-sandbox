FROM alpine

# デフォルトのシークレットの場所から、シークレットを表示
RUN --mount=type=secret,id=mysecret cat /run/secrets/mysecret

# 任意のシークレットの場所から、シークレットを表示
RUN --mount=type=secret,id=mysecret,dst=/foobar cat /foobar
