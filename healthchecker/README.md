This simple healthchecker thing was supposed to work for Docker. I was trying stuff like:

```
#HEALTHCHECK --timeout=5s --interval=10s --retries=1 CMD ./healthchecker || exit 1
#HEALTHCHECK CMD ./healthchecker || exit 1
```

But it wasn't working, even though running it with `docker exec` was working just fine.

Other useful lines in Dockerfile:

```
WORKDIR /src/healthchecker

run go build -v -o healthchecker
```

```
COPY --from=build /src/healthchecker/healthchecker /app/healthchecker
```
