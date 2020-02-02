# Fork of cbuild

I really liked the original idea of generating `C` source files and headers and build scriptfiles from a single `C module`, thus making the process of compilation almost trivial. I would like to give the project my own spin.

Long term goals to achaive:

1. Port the source from `C` to `Go`, to improve portability, perhaps
even readability and maintainability. `Go` is much easier to cross compile as well.
2. Add more build script backends. [Ninja](https://ninja-build.org/)
for example.
3. Implement a package manager. This would be really neat, but I am not
sure if I will have the time to pull it off. Could be based on [clib](https://clibs.org/)

# Status

Managed to transpile from `C` to `Go` using
[c4go](https://github.com/Konstantin8105/c4go). The output is pretty
messy, but I figured it makes more sense to do the heavy lifting
automatically with a transpiler then cleanup and refactor the code.
