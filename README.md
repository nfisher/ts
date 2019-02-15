# ts

ts is inspired by the moreutils ts command of the same name. It is used to prepend
timestamps to line output received on stdin. It was developed to coursely profile
commands in CI builds so focuses on the output of time rather than date-time.

Usage
=====

ts is intended to be combined with pipes. The following will timestamp both stderr and
stdout in a single output stream:

```
make 2>&1 | ts
```
