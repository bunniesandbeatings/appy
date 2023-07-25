# Appy

Makes your CLI demos 'appy

## What makes it 'appy?

Instead of prototyping a CLI for your design proposal:
> Fake it 'til you make it

Using a [script](./demos/git.yaml) to describe the interactions with a hypothetical CLI, you can demonstrate:
1. How your CLI reponds on stdout to a command
2. The files and directories your CLI creates/deletes/updates (so long as they're plain text)


## Demo

Try this in a fresh terminal

```
# make the git binary be "appy" instead
alias git="./path/to/appy/executable"

# create a working dir for the appy demo. 
mkdir demo
cd demo

# generate a config that points to the demo script
# you can place this in a directory above your demo dir
# if you dont anyone to see it in your directory during
# your demo
echo "./path/to/appy/demos/git.yaml" > .appy
```

Now when you type `git` with unrecognized params, you will get a bell character
to warn you that your demo is off script. 

You can append `APPY_DEBUG=TRUE` to your command if you need more details on what
you got wrong.

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
