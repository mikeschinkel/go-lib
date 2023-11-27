

REPO_NAME =

.PHONY: new-repo
new-repo:
    ifndef REPO_NAME
      $(error REPO_NAME is not set)
    endif                                                                     
	git init && \
	echo "git init" && \
	echo ".idea" >> .git/info/exclude  && \
	echo "echo .idea >> .git/info/exclude" && \
	echo ".idea/**" >> .git/info/exclude  && \
	echo "echo .idea/** >> .git/info/exclude" && \
	git config pull.rebase true  && \
	echo "git config pull.rebase true" && \
	git remote add origin "git@github.com:mikeschinkel/$(REPO_NAME)"  && \
	echo "git remote add origin \"git@github.com:mikeschinkel/$(REPO_NAME)"\" && \
	git add .  && \
	echo "git add ." && \
	git commit -m "Initial commit" && \
	echo "git commit -m \"Initial commit\"" && \
	git push --set-upstream origin main && \
	echo "git push --set-upstream origin main"

foo:
#	git pull  && \
#	echo "git pull" && \
#	git branch --set-upstream-to=origin/main main  && \
#	echo "git branch --set-upstream-to=origin/main main" && \
#	echo Done
	# git push