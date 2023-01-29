# git-link

I wanted to manage files in my storages with git.
git-annex is most common choice for doing this stuff but it was too complex and heavy.
git-link is my attempt to replace git-annex.


    $ git init
    $ git link add /mnt/nas/large.mp4

    $ ls 
    large.mp4 -> /mnt/nas/large.mp4
    .gitlinks
    .git

    $ git link fsck large.mp4 # verify checksum of the file, uses XXH64

    $ git link rm large.mp4
