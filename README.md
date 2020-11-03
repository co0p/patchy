![build](https://github.com/co0p/patchy/workflows/build/badge.svg)

patchy
=======

Patchy allows you to get a unified diff of multiple commits on an origin branch for focus on the changes and not the individual commits.

usage: 
    
    $ patchy <remote repository> <branch> <target branch>
    

example
--------

Let's say you have a branch "demo" with 3 commits. 

    * c99a019 (origin/master) 
    | * 8cb9097 (origin/demo) c3    # echo "hi should be \n a new line" > file.txt
    | * 9b69853 c2                  # echo "hi again" > file2.txt 
    | * 6e1c33b c1                  # echo "hi" > file.txt 
    |/  
    * ca8ca2e 
    * 81a6ff4 
    
to get a diff of all changes combined, run `patchy https://github.com/co0p.patchy.git origin/demo origin/master` which will output:

```diff
diff --git a/file.txt b/file.txt
new file mode 100644
index 0000000..0bcc847
--- /dev/null
+++ b/file.txt
@@ -0,0 +1 @@
+ hi should be \n a new line
diff --git a/file2.txt b/file2.txt
new file mode 100644
index 0000000..401adc2
--- /dev/null
+++ b/file2.txt
@@ -0,0 +1 @@
+hi again
```
