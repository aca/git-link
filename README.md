# git-link

I wanted to manage files in my storages with git.
git-annex is most common choice for doing this stuff but it was too complex and heavy.
git-link is my attempt to replace git-annex. It just creates symbolic link to the file, store metadata for it.

[Special remotes](https://git-annex.branchable.com/special_remotes/)? Why not just use [FUSE](https://www.kernel.org/doc/html/latest/filesystems/fuse.html) instead, and mount it if needed.

    $ git init
    $ git link add /mnt/nas/large.mp4

    $ ls 
    large.mp4 -> /mnt/nas/large.mp4

    $ git link fsck large.mp4 # verify checksum of the file, uses XXH64

    $ git link rm large.mp4

    $ git link dump /mnt/nas /mnt
    $ ls
    nas/xxx/large.dump -> /mnt/nas/xxx/large.dump
    nas/xxx/large2.dump -> /mnt/nas/xxx/large2.dump

## Alternatives

- https://github.com/rfjakob/cshatag/
- https://git-annex.branchable.com/
