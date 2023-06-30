git add --all && \
git commit -mm && \
git push origin dev && \
git checkout master && \
git pull origin master && \
git merge dev && \
git push origin master && \
git checkout dev