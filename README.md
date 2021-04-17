# 120vcs
A version control system written in golang. this vc is very similar to git
originaly I wanted to make it accumlative (store patches instead of snapshots)
but doing that proved to be complicated as I would have to rewrite the unix
diff and patch utilities in Go, and then I would have had to add a "base"
for every 20 commits or so which is something that in my oppinion defeats the
purpose of making it accumlative.
