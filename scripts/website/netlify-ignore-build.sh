# This $BRANCH variable comes from the Netlify env
if  [[ $BRANCH == chore/prepare-release-* ]] ;
then
  exit 0
else
  git diff --quiet $COMMIT_REF $CACHED_COMMIT_REF -- website/ specs/ && exit $?
fi
