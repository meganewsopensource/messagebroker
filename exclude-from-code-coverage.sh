while read p || [ -n "$p" ]
do
echo -e "$(cat ./cover.out)\n$p:1.1,1.2 1 1" > ./cover.out
done < ./exclude-from-code-coverage.txt
