# Appy

Makes your CLI demos 'appy

## Demo

Try this in a fresh terminal

```
# make the git binary be "appy" instead
alias git="./path/to/appy/executable"

# create a working dir for the appy demo. 
# appy will not modify files outside the dir containing .appy
mkdir demo
cd demo

# generate a config that points to the demo script
echo "./path/to/appy/demos/git.yaml" > .appy
```

Now when you type `git` with unrecognized params, you will get a bell character
to warn you that your demo is off script. 
When you type it with the right params, you'll get the scripted behaviour:

Try:
```
$ git init .
$ ls 
$ cat git/never/does/this/readme.md
$ git status
$ git add .
$ git status
$ git commit -m "wip"
```
