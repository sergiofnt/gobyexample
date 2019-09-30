#!/usr/bin/env bash
GH="github.com/butuzov/gobyexample"

echo -e "Deploying translations to GitHub"

cd $TRAVIS_BUILD_DIR
# clean up
rm -f public/*

MSG=$(git log -1 --oneline)
# generate new files.
cd public

git config --global user.email "butuzov@made.ua"
git config --global user.name "Oleg Butuzov"
git init
git remote add -f origin "https://$GH"
git checkout -f gh-pages

cd .. && tools/build && cd public

# ls -lah
echo -e "Sources Normalization"
unlink index.html.back
unlink site.css.back

echo -e "Norm Test"
ls -lah

COMMIT=$(git status | grep "nothing to commit, working tree clean")

echo $COMMIT

if [[ -z $COMMIT ]]; then
	echo "Updating Github Pages"
	git add --all
	git commit -m "Travis CI | ${MSG}"
	git push "https://${GITHUB_TOKEN}@${GH}" gh-pages
fi
