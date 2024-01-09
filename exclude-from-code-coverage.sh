while read p || [ -n "$p" ]
do
echo -e "$(cat ./coverage.out)\n$p:1.1,1.2 1 1" > ./coverage.out
done < ./exclude-from-code-coverage.txt
